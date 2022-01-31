package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (d delivery) GetStat(ctx *gin.Context) {
	longURL := ctx.Param("shortURL")
	ctx.String(http.StatusOK, "Get statistic for short URL %s", longURL)
}
