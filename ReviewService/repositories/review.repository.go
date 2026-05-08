package repositories

import (
	"ReviewService/models"
	"database/sql"
	"fmt"
)

type ReviewRepository interface {
	CreateReview(
		bookingID int64,
		reviewText string,
		rating int64,
	) (*models.Review, error)

	GetReview(id int) (*models.Review, error)

	GetReviews() ([]*models.Review, error)

	DeleteReview(id int) error
}

type ReviewRepositoryImpl struct {
	SqlDB *sql.DB
}

func NewReviewRepository(sqlDB *sql.DB) ReviewRepository {
	return &ReviewRepositoryImpl{
		SqlDB: sqlDB,
	}
}

func (r *ReviewRepositoryImpl) CreateReview(bookingID int64, reviewText string, rating int64, ) (*models.Review, error) {

	// TODO: implement DB insert
	fmt.Println("ok from CreateReview")
	return nil, nil
}

func (r *ReviewRepositoryImpl) GetReview(id int) (*models.Review, error) {

	// TODO: implement DB query

	return &models.Review{}, nil
}

func (r *ReviewRepositoryImpl) GetReviews() ([]*models.Review, error) {

	// TODO: implement DB query

	return []*models.Review{}, nil
}

func (r *ReviewRepositoryImpl) DeleteReview(id int) error {

	// TODO: implement delete

	return nil
}
