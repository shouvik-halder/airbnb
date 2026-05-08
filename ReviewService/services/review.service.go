package services

import (
	"ReviewService/models"
	"ReviewService/repositories"
	"fmt"
)

type ReviewService interface {
	CreateReviewService(bookingId int64, reviewText string, rating int64) (*models.Review, error)
	GetReviewService(id int64) (*models.Review, error)
	GetAllReviewsService() (*[]models.Review, error)
	DeleteReviewService(id int64) error
}

type reviewServiceImpl struct {
	reviewRepo repositories.ReviewRepository
}

func NewReviewService(_reviewRepo repositories.ReviewRepository) ReviewService {
	return &reviewServiceImpl{
		reviewRepo: _reviewRepo,
	}
}

func (rs *reviewServiceImpl) CreateReviewService(bookingId int64, reviewText string, rating int64) (*models.Review, error) {
	rs.reviewRepo.CreateReview(bookingId, reviewText, rating)
	fmt.Println("ok from CreateReviewService")
	return nil, nil
}

func (rs *reviewServiceImpl) GetReviewService(id int64) (*models.Review, error) {
	return nil, nil
}

func (rs *reviewServiceImpl) GetAllReviewsService() (*[]models.Review, error) {
	return nil, nil
}

func (rs *reviewServiceImpl) DeleteReviewService(id int64) error {
	return nil
}
