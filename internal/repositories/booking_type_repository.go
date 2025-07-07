package repositories

import (
    "car-rental-api/internal/models"
    "gorm.io/gorm"
)

type BookingTypeRepository struct {
    db *gorm.DB
}

func NewBookingTypeRepository(db *gorm.DB) *BookingTypeRepository {
    return &BookingTypeRepository{db: db}
}

func (r *BookingTypeRepository) Create(bookingType *models.BookingType) error {
    return r.db.Create(bookingType).Error
}

func (r *BookingTypeRepository) GetAll() ([]models.BookingType, error) {
    var bookingTypes []models.BookingType
    err := r.db.Find(&bookingTypes).Error
    return bookingTypes, err
}

func (r *BookingTypeRepository) GetByID(id int) (*models.BookingType, error) {
    var bookingType models.BookingType
    err := r.db.First(&bookingType, id).Error
    if err != nil {
        return nil, err
    }
    return &bookingType, nil
}

func (r *BookingTypeRepository) Update(bookingType *models.BookingType) error {
    return r.db.Save(bookingType).Error
}

func (r *BookingTypeRepository) Delete(id int) error {
    return r.db.Delete(&models.BookingType{}, id).Error
}