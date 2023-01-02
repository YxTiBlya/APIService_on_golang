package routehandler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yxtiblya/internal/cache"
	"github.com/yxtiblya/internal/models"
	"github.com/yxtiblya/internal/store"
)

var mailing_name string = "mailings"

// returning json with id of created record or error
func Mailing_post(c *gin.Context) {
	var mailing *models.Mailing

	// decode request body to model
	err := json.NewDecoder(c.Request.Body).Decode(&mailing)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": err,
		})
		return
	}

	// create new record in contacts table
	if result := store.DB.Create(mailing); result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": result.Error,
		})
		return
	}

	// delete cache record if exist
	_, err = cache.Get(mailing_name)
	if err == nil {
		cache.Del(mailing_name)
	}

	c.JSON(http.StatusOK, gin.H{
		"id": mailing.ID,
	})
}

// returning json with all records or error
func Mailing_get(c *gin.Context) {
	var mailings []models.Mailing

	// check value in cache
	value, err := cache.Get(mailing_name)

	// run if value not exist in cache
	if err != nil {
		result := store.DB.Find(&mailings)

		if result.Error != nil {
			c.JSON(http.StatusConflict, gin.H{
				"error": result.Error,
			})
			return
		}

		// add record to cache
		if err := cache.Set(mailing_name, &mailings); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"response": mailings,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": value,
	})
}

// returning json with changed record or error
func Mailing_put(c *gin.Context) {
	var mailing *models.Mailing

	// decode request body to model
	err := json.NewDecoder(c.Request.Body).Decode(&mailing)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": err,
		})
		return
	}

	// update record in contacts table
	result := store.DB.Model(mailing).Updates(models.Mailing{
		Start_time: mailing.Start_time,
		Message:    mailing.Message,
		Filters:    mailing.Filters,
		End_time:   mailing.End_time,
	})
	if result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": result.Error,
		})
		return
	}

	// delete cache record if exist
	_, err = cache.Get(mailing_name)
	if err == nil {
		cache.Del(mailing_name)
	}

	c.JSON(http.StatusOK, gin.H{
		"response": mailing,
	})
}

// returning statuscode
func Mailing_delete(c *gin.Context) {
	var mailing *models.Mailing

	// decode request body to model
	err := json.NewDecoder(c.Request.Body).Decode(&mailing)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": err,
		})
		return
	}

	// delete record from contacts table
	result := store.DB.Delete(&mailing, mailing.ID)
	if result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": result.Error,
		})
		return
	}

	// delete cache record if exist
	_, err = cache.Get(mailing_name)
	if err == nil {
		cache.Del(mailing_name)
	}

	c.Writer.WriteHeader(http.StatusAccepted)
}
