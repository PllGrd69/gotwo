package apis

import (
	"net/http"

	"github.com/PllGrd69/gotwo/models"
	"github.com/gin-gonic/gin"
)

//CRUD for items table
func ItemsIndex(c *gin.Context) {
	s := models.Item{Title: "Sean", Notes: "nnn"}

	c.JSON(http.StatusOK, gin.H{
		"message": "Hola " + s.Title,
	})
}
