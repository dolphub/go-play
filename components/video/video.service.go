package video

import (
	"fmt"

	"github.com/dolphub/api-server/apperr"
	"github.com/gin-gonic/gin"
)

// VideoService f
type VideoService interface {
	Save(c *gin.Context) (*Video, error)
	FindAll() ([]Video, error)
	FindById(string) (*Video, error)
}

type service struct {
	repo Repo
}

// NewService f
func NewService(repo Repo) VideoService {
	return &service{
		repo: repo,
	}
}

func (s *service) Save(c *gin.Context) (*Video, error) {
	v := &Video{}
	if err := c.ShouldBindJSON(v); err != nil {
		return nil, err
	}
	if v, err := s.repo.Save(v); err != nil {
		return nil, err
	} else {
		return v, nil
	}
}

func (s *service) FindAll() ([]Video, error) {
	videos, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (s *service) FindById(id string) (*Video, error) {
	video, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	if video == nil {
		return nil, apperr.NotFoundError(fmt.Sprintf("%s not found.", id))
	}

	return video, nil
}
