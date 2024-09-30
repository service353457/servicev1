package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var IsLoading = true

func main() {

	router := gin.Default()

	// router.Static("/static", "./dist")
	// router.StaticFile("/", "./dist/index.html")
	router.Use(static.Serve("/", static.LocalFile("./dist", false)))

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}                                                 // Разрешенные источники
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}           // Разрешенные методы
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"} // Разрешенные заголовки
	config.ExposeHeaders = []string{"Content-Length", "Content-Type"}                   // Заголовки, которые будут доступны клиенту
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	router.Use(cors.New(config))

	router.GET("/checkstatus", checkstatus)
	router.POST("/changestatus", changestatus)
	router.Run(":8089")
}

func checkstatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"isLoading": IsLoading})
}

func changestatus(c *gin.Context) {
	var json struct {
		IsLoading bool `json:"is_loading"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	IsLoading = json.IsLoading
	c.JSON(http.StatusOK, gin.H{"isLoading": IsLoading})
}
