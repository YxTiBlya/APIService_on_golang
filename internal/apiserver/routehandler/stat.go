package routehandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yxtiblya/internal/cache"
	"github.com/yxtiblya/internal/models"
	"github.com/yxtiblya/internal/store"
)

// returning json with all records or error
func Stat_get(c *gin.Context) {
	id := c.Param("id")

	var messages []models.Message

	// check value in cache
	value, err := cache.Get(id)

	// run if value not exist in cache
	if err != nil {
		result := store.DB.Where("mailing_id = ?", id).Find(&messages)

		if result.Error != nil {
			c.JSON(http.StatusConflict, gin.H{
				"error": result.Error,
			})
			return
		}

		// add record to cache
		if err := cache.Set(id, &messages); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"response": messages,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": value,
	})

}
