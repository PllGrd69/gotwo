package routers

import (
	"github.com/PllGrd69/gotwo/apis"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupRouter() *gin.Engine {

	conn, err := connectDB()
	if err != nil {
		panic("failed to connect database: " + err.Error())
		//return
	}

	r := gin.Default()
	r.Use(dbMiddleware(*conn))

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", apis.ItemsIndex)
		v1.GET("/persons", apis.PersonsIndex)
	}

	return r
}

func connectDB() (c *gorm.DB, err error) {

	dsn := "docker:docker@tcp(localhost:3307)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	return conn, err
}

func dbMiddleware(conn gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", conn)
		c.Next()
	}
}
