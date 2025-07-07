package services

import (
    "car-rental-api/internal/models"
    "car-rental-api/internal/repositories"
    "errors"
    "time"
)

type BookingTypeService struct {
    bookingTypeRepo *repositories.BookingTypeRepository
}

func NewBookingTypeService(bookingTypeRepo *repositories.BookingTypeRepository) *BookingTypeService {
    return &BookingTypeService{bookingTypeRepo: bookingTypeRepo}
}

func (s *BookingTypeService) CreateBookingType(bookingType *models.BookingType) error {
    bookingType.CreatedAt = time.Now()
    bookingType.UpdatedAt = time.Now()
    return s.bookingTypeRepo.Create(bookingType)
}

func (s *BookingTypeService) GetAllBookingTypes() ([]models.BookingType, error) {
    return s.bookingTypeRepo.GetAll()
}

func (s *BookingTypeService) GetBookingTypeByID(id int) (*models.BookingType, error) {
    return s.bookingTypeRepo.GetByID(id)
}

func (s *BookingTypeService) UpdateBookingType(id int, bookingType *models.BookingType) error {
    existing, err := s.bookingTypeRepo.GetByID(id)
    if err != nil {
        return errors.New("booking type not found")
    }

    existing.BookingType = bookingType.BookingType
    existing.Description = bookingType.Description
    existing.UpdatedAt = time.Now()

    return s.bookingTypeRepo.Update(existing)
}

func (s *BookingTypeService) DeleteBookingType(id int) error {
    _, err := s.bookingTypeRepo.GetByID(id)
    if err != nil {
        return errors.New("booking type not found")
    }

    return s.bookingTypeRepo.Delete(id)
}