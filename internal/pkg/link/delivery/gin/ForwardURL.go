package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (d delivery) ForwardURL(ctx *gin.Context) {
	shortURL := ctx.Param("shortURL") //
	ctx.String(http.StatusOK, "Forwarding from short URL %s to long URL %s", shortURL, "ThisIsLongURL")
}
