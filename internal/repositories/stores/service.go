package stores

import (
	"db_project/internal/models"
	"db_project/internal/repositories"

	"github.com/jackc/pgx"
	_ "github.com/lib/pq"
)

type ServiceStore struct {
	db *pgx.ConnPool
}

func CreateServiceRepository(db *pgx.ConnPool) repositories.ServiceRepository {
	return &ServiceStore{db: db}
}

func (serviceStore *ServiceStore) Clear() (err error) {
	_, err = serviceStore.db.Exec("TRUNCATE TABLE forums, posts, threads, user_forum, users, votes CASCADE;")
	return
}

func (serviceStore *ServiceStore) GetStatus() (status *models.Status, err error) {
	status = &models.Status{}
	err = serviceStore.db.QueryRow("SELECT (SELECT count(*) FROM users) AS users, "+
		"(SELECT count(*) FROM forums) AS forums, "+
		"(SELECT count(*) FROM threads) AS threads, "+
		"(SELECT count(*) FROM posts) AS posts;").
		Scan(&status.User, &status.Forum, &status.Thread, &status.Post)
	return
}
