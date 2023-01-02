package routehandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yxtiblya/internal/apiserver/auth"
)

// return {token: jwt}
func GetJWT(c *gin.Context) {
	token, err := auth.CreateJWT()
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
