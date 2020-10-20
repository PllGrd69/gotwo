package apis

import (
	"net/http"

	"github.com/PllGrd69/gotwo/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//CRUD for items table
func PersonsIndex(c *gin.Context) {
	var lis []models.Person

	db, _ := c.Get("db")
	conn := db.(gorm.DB)

	// Migrate the schema
	conn.AutoMigrate(&models.Person{})

	conn.Find(&lis)
	c.JSON(http.StatusOK, gin.H{
		"msg": "thank you",
		"r":   lis,
	})

}
func PersonsIndexID(c *gin.Context) {
	id := c.Param("id")

	db, _ := c.Get("db")
	conn := db.(gorm.DB)

	var d models.Person
	if err := conn.First(&d, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &d)
}
func CreartePerson(c *gin.Context) {
	d := models.Person{Name: c.PostForm("name"), Age: c.PostForm("age")}
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	conn.Create(&d)
	c.JSON(http.StatusOK, &d)
}
func DeletePerson(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.Person
	if err := conn.Where("id = ?", id).First(&d).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	conn.Unscoped().Delete(&d)
}
func UpedatePerson(c *gin.Context) {
	db, _ := c.Get("db")
	conn := db.(gorm.DB)
	id := c.Param("id")
	var d models.Person
	if err := conn.First(&d, id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	d.Name = c.PostForm("name")
	d.Age = c.PostForm("age")
	//c.BindJSON(&d)
	conn.Save(&d)
	c.JSON(http.StatusOK, &d)
}
