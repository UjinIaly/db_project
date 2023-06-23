package repositories

import "db_project/internal/models"

type VoteRepository interface {
	Vote(threadID int64, vote *models.Vote) (err error)
}
