package routehandler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yxtiblya/internal/cache"
	"github.com/yxtiblya/internal/models"
	"github.com/yxtiblya/internal/store"
)

var message_name string = "messages"

// returning json with id of created record or error
func Message_post(c *gin.Context) {
	var message *models.Message

	// decode request body to model
	err := json.NewDecoder(c.Request.Body).Decode(&message)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": err,
		})
		return
	}

	// create new record in contacts table
	if result := store.DB.Create(message); result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": result.Error,
		})
		return
	}

	// delete cache record if exist
	_, err = cache.Get(message_name)
	if err == nil {
		cache.Del(message_name)
	}

	c.JSON(http.StatusOK, gin.H{
		"id": message.ID,
	})
}

// returning json with all records or error
func Message_get(c *gin.Context) {
	var messages []models.Message

	// check value in cache
	value, err := cache.Get(message_name)

	// run if value not exist in cache
	if err != nil {
		result := store.DB.Find(&messages)

		if result.Error != nil {
			c.JSON(http.StatusConflict, gin.H{
				"error": result.Error,
			})
			return
		}

		// add record to cache
		if err := cache.Set(message_name, &messages); err != nil {
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

// returning json with changed record or error
func Message_put(c *gin.Context) {
	var message *models.Message

	// decode request body to model
	err := json.NewDecoder(c.Request.Body).Decode(&message)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": err,
		})
		return
	}

	// update record in contacts table
	result := store.DB.Model(message).Updates(models.Message{
		Datetime:   message.Datetime,
		Status:     message.Status,
		Mailing_id: message.Mailing_id,
		Contact_id: message.Contact_id,
	})
	if result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": result.Error,
		})
		return
	}

	// delete cache record if exist
	_, err = cache.Get(message_name)
	if err == nil {
		cache.Del(message_name)
	}

	c.JSON(http.StatusOK, gin.H{
		"response": message,
	})
}

// returning statuscode
func Message_delete(c *gin.Context) {
	var message *models.Message

	// decode request body to model
	err := json.NewDecoder(c.Request.Body).Decode(&message)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": err,
		})
		return
	}

	// delete record from contacts table
	result := store.DB.Delete(&message, message.ID)
	if result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": result.Error,
		})
		return
	}

	// delete cache record if exist
	_, err = cache.Get(message_name)
	if err == nil {
		cache.Del(message_name)
	}

	c.Writer.WriteHeader(http.StatusAccepted)
}
