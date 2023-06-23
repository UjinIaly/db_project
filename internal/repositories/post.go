package repositories

import "github.com/UjinIaly/db_project
/internal/models"

type PostRepository interface {
	GetByID(id int64) (post *models.Post, err error)
	Update(post *models.Post) (err error)
}
