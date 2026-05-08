package controllers

import (
	"ReviewService/services"
	"net/http"
)

type ReviewController struct {
	reviewService services.ReviewService
}

func NewReviewController(_reviewService services.ReviewService) *ReviewController {
	return &ReviewController{
		reviewService: _reviewService,
	}
}

func (rc *ReviewController) CreateReviewController(w http.ResponseWriter, h *http.Request) {
	rc.reviewService.CreateReviewService(1, "text message", 5)
	w.Write([]byte("ok from ReviewController"))
}

func (rc *ReviewController) GetReviewByIdController(w http.ResponseWriter, h *http.Request) {

}

func (rc *ReviewController) GetAllReviewsController(w http.ResponseWriter, h *http.Request) {

}

func (rc *ReviewController) DeleteReviewController(w http.ResponseWriter, h *http.Request) {

}
