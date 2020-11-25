package server

import (
	"net/http"

	"github.com/davecgh/go-spew/spew"
	"github.com/dolphub/api-server/components/video"
	"github.com/dolphub/api-server/database"
	"github.com/gin-gonic/gin"
)

// StartServer starts the httpd server
func StartServer() {
	engine := gin.Default()

	db, err := database.New()
	defer db.Close()
	if err != nil {
		spew.Dump(err)
		panic(err)
	}
	videoController := video.NewController(engine, db)

	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Ok!",
		})
	})
	videoController.Register()
	// auth.NewController(r).Register()
	engine.Run("localhost:3000")
}
