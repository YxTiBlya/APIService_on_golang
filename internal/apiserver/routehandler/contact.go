package routehandler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yxtiblya/internal/cache"
	"github.com/yxtiblya/internal/models"
	"github.com/yxtiblya/internal/store"
)

var contact_name string = "contacts"

// returning json with id of created record or error
func Contact_post(c *gin.Context) {
	var contact *models.Contact

	// decode request body to model
	err := json.NewDecoder(c.Request.Body).Decode(&contact)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": err,
		})
		return
	}

	// create new record in contacts table
	if result := store.DB.Create(contact); result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": result.Error,
		})
		return
	}

	// delete cache record if exist
	_, err = cache.Get(contact_name)
	if err == nil {
		cache.Del(contact_name)
	}

	c.JSON(http.StatusOK, gin.H{
		"id": contact.ID,
	})
}

// returning json with all records or error
func Contact_get(c *gin.Context) {
	var contacts []models.Contact

	// check value in cache
	value, err := cache.Get(contact_name)

	// run if value not exist in cache
	if err != nil {
		result := store.DB.Find(&contacts)

		if result.Error != nil {
			c.JSON(http.StatusConflict, gin.H{
				"error": result.Error,
			})
			return
		}

		// add record to cache
		if err := cache.Set(contact_name, &contacts); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"response": contacts,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": value,
	})

}

// returning json with changed record or error
func Contact_put(c *gin.Context) {
	var contact *models.Contact

	// decode request body to model
	err := json.NewDecoder(c.Request.Body).Decode(&contact)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": err,
		})
		return
	}

	// update record in contacts table
	result := store.DB.Model(contact).Updates(models.Contact{
		Number:        contact.Number,
		Operator_code: contact.Operator_code,
		Tag:           contact.Tag,
		Time_zone:     contact.Time_zone,
	})
	if result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": result.Error,
		})
		return
	}

	// delete cache record if exist
	_, err = cache.Get(contact_name)
	if err == nil {
		cache.Del(contact_name)
	}

	c.JSON(http.StatusOK, gin.H{
		"response": contact,
	})
}

// returning statuscode
func Contact_delete(c *gin.Context) {
	var contact *models.Contact

	// decode request body to model
	err := json.NewDecoder(c.Request.Body).Decode(&contact)
	if err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": err,
		})
		return
	}

	// delete record from contacts table
	result := store.DB.Delete(&contact, contact.ID)
	if result.Error != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": result.Error,
		})
		return
	}

	// delete cache record if exist
	_, err = cache.Get(contact_name)
	if err == nil {
		cache.Del(contact_name)
	}

	c.Writer.WriteHeader(http.StatusAccepted)
}
