package auth

// import (
// 	"os"

// 	"github.com/dolphub/api-server/components/user"
// 	"github.com/gin-gonic/gin"
// )

// type AuthService interface {
// 	Login(c *gin.Context) (*user.User, error)
// }

// type service struct{}

// func NewService() AuthService {
// 	return &service{}
// }

// func (s *service) Login(c *gin.Context) (*user.User, error) {
// 	var authDTO LoginDTO

// 	if err := c.ShouldBindJSON(&authDTO); err != nil {
// 		return nil, err
// 	}

// 	user := &user.User{
// 		UserName:  authDTO.Username,
// 		FirstName: "somefirst",
// 		LastName:  "somelast",
// 		Email:     os.Getenv("SOME_VAL"),
// 	}
// 	return user, nil
// }

// // func
