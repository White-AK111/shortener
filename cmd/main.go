package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	config "github.com/White-AK111/shortener/configs"

	ginDelivery "github.com/White-AK111/shortener/internal/pkg/link/delivery/gin"
	"github.com/White-AK111/shortener/internal/pkg/link/repository/postgres"
	"github.com/White-AK111/shortener/internal/pkg/link/usecase"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	// postgres driver
	_ "github.com/lib/pq"
)

var (
	USECONFIG   string
	usageConfig = "use this flag for set path to configuration file"
)

func main() {
	// Load config file
	cfg, err := config.InitConfig(&USECONFIG)
	if err != nil {
		log.Fatalf("Can't load configuration file %s. Error: %s", USECONFIG, err)
	}

	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	router := gin.Default()

	// Add middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.DBConfig.Host, cfg.DBConfig.Port, cfg.DBConfig.User, cfg.DBConfig.Password, cfg.DBConfig.DBName)

	db, err := sqlx.Open(
		"postgres",
		psqlInfo,
	)
	if err != nil {
		log.Fatalf("failed to open db connection %v", err)
	}

	linksRepository := postgres.New(db)
	linksUsecase := usecase.New(linksRepository)
	linksDelivery := ginDelivery.New(linksUsecase)

	router.GET("/:shortURL", linksDelivery.ForwardURL)
	router.GET("/link/:longURL", linksDelivery.GetShortURL)
	router.GET("/link/stat/:shortURL", linksDelivery.GetStat)

	srv := &http.Server{
		Addr:    cfg.App.ServerAddress + ":" + strconv.Itoa(cfg.App.ServerPort),
		Handler: router,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Printf("Start server on %s:%d\n", cfg.App.ServerAddress, cfg.App.ServerPort)

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}

// init func, parse flags
func init() {
	flag.StringVar(&USECONFIG, "path", "../configs/config.yml", usageConfig)
	flag.Parse()
}
