package auth

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// var (
// 	authService AuthService = NewService()
// )

// type Controller interface {
// 	Register()
// }

// type authController struct {
// 	router *gin.Engine
// }

// func NewController(router *gin.Engine) Controller {
// 	return &authController{
// 		router: router,
// 	}
// }

// func (c *authController) Register() {
// 	authRoutes := c.router.Group("/auth")

// 	// Get current user session?
// 	// authRoutes.GET("/", func(c *gin.Context) {})

// 	authRoutes.POST("/login", func(c *gin.Context) {
// 		// authDto, err := authSer.Save(c)
// 		// if err != nil {
// 		// 	c.JSON(http.StatusBadRequest, err.Error())
// 		// 	return
// 		// }

// 		user, err := authService.Login(c)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, err.Error())
// 			return
// 		}
// 		c.JSON(http.StatusCreated, user)
// 	})
// }
