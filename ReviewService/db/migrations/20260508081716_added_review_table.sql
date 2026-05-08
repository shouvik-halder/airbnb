-- +goose Up
CREATE TABLE IF NOT EXISTS review (    
    id INT PRIMARY KEY AUTO_INCREMENT,    
    booking_id INT NOT NULL,    
    user_id INT NOT NULL,    
    hotel_id INT NOT NULL,    
    comment TEXT,    
    rating INT NOT NULL CHECK (rating BETWEEN 1 AND 5),    
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,    
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP        ON UPDATE CURRENT_TIMESTAMP,    
    deleted_at TIMESTAMP NULL DEFAULT NULL,    
    
    INDEX idx_booking_id (booking_id),    
    INDEX idx_user_id (user_id),    
    INDEX idx_hotel_id (hotel_id));

-- +goose Down
DROP TABLE review
