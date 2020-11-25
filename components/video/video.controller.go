package video

import (
	"net/http"

	"github.com/dolphub/api-server/apperr"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
)

var ()

type Controller interface {
	Register()
}

type videoController struct {
	router  *gin.Engine
	service VideoService
}

func NewController(router *gin.Engine, db *pg.DB) Controller {
	videoRepo := NewRepo(db)
	service := NewService(videoRepo)
	return &videoController{
		router:  router,
		service: service,
	}
}

type VideoByIdUri struct {
	Id string `uri:"id" binding:"required,uuid"`
}

func (ctrl *videoController) Register() {
	videoRoutes := ctrl.router.Group("/videos")

	videoRoutes.GET("/", func(c *gin.Context) {
		data, err := ctrl.service.FindAll()
		if err != nil {
			apperr.Response(c, err)
		}
		c.JSON(http.StatusOK, data)
	})

	videoRoutes.POST("/", func(c *gin.Context) {
		video, err := ctrl.service.Save(c)
		if err != nil {
			apperr.Response(c, err)
			return
		}
		c.JSON(http.StatusCreated, video)
	})

	videoRoutes.GET("/:id", func(c *gin.Context) {
		uri := &VideoByIdUri{}
		if err := c.ShouldBindUri(uri); err != nil {
			apperr.Response(c, err)
			return
		}
		video, err := ctrl.service.FindById(uri.Id)
		if err != nil {
			apperr.Response(c, err)
			return
		}
		c.JSON(http.StatusAccepted, video)
	})
}
