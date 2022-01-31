package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (d delivery) GetShortURL(ctx *gin.Context) {
	longURL := ctx.Param("longURL")
	ctx.String(http.StatusOK, "Get short URL from long URL %s", longURL)
}
