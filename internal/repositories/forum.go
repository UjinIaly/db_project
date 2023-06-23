package repositories

import "db_project/internal/models"

type ForumRepository interface {
	Create(forum *models.Forum) (err error)
	GetBySlug(slug string) (forum *models.Forum, err error)
	GetUsers(slug string, limit int, since string, desc bool) (users *[]models.User, err error)
	GetThreads(slug string, limit int, since string, desc bool) (threads *[]models.Thread, err error)
}