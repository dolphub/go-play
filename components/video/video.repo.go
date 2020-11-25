package video

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/dolphub/api-server/apperr"
	"github.com/go-pg/pg/v10"
	"github.com/google/uuid"
)

// Repo f
type Repo interface {
	Save(*Video) (*Video, error)
	FindAll() ([]Video, error)
	FindById(string) (*Video, error)
}

type repo struct {
	videos []Video
	db     *pg.DB
}

// NewRepo f
func NewRepo(db *pg.DB) Repo {
	return &repo{
		videos: []Video{},
		db:     db,
	}
}

// Save f
func (r *repo) Save(v *Video) (*Video, error) {
	uuid, _ := uuid.NewRandom()
	v.Id = uuid.String()
	result, err := r.db.Model(v).Insert()
	if err != nil {
		return nil, apperr.BadRequestError(err.Error())
	}

	spew.Dump("dump", result.RowsReturned(), result.RowsAffected())

	return v, nil
}

// FindAll f
func (r *repo) FindAll() ([]Video, error) {
	var videos = []Video{}
	query := r.db.Model(&videos).Column("*")
	if err := query.Select(); err != nil {
		return nil, err
	}
	return videos, nil
}

func (r *repo) FindById(id string) (*Video, error) {
	var video = &Video{}
	query := r.db.Model(video).Where("id = (?)", id)
	if err := query.Select(); err != nil {
		// Do not throw an error, return nil instead
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return video, nil
}
