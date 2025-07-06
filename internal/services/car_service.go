package services

import (
    "car-rental-api/internal/models"
    "car-rental-api/internal/repositories"
    "errors"
    "time"
)

type CarService struct {
    carRepo *repositories.CarRepository
}

func NewCarService(carRepo *repositories.CarRepository) *CarService {
    return &CarService{carRepo: carRepo}
}

func (s *CarService) CreateCar(car *models.Car) error {
    car.CreatedAt = time.Now()
    car.UpdatedAt = time.Now()
    return s.carRepo.Create(car)
}

func (s *CarService) GetAllCars() ([]models.Car, error) {
    return s.carRepo.GetAll()
}

func (s *CarService) GetCarByID(id int) (*models.Car, error) {
    return s.carRepo.GetByID(id)
}

func (s *CarService) UpdateCar(id int, car *models.Car) error {
    existing, err := s.carRepo.GetByID(id)
    if err != nil {
        return errors.New("car not found")
    }

    existing.Name = car.Name
    existing.Stock = car.Stock
    existing.DailyRent = car.DailyRent
    existing.UpdatedAt = time.Now()

    return s.carRepo.Update(existing)
}

func (s *CarService) DeleteCar(id int) error {
    _, err := s.carRepo.GetByID(id)
    if err != nil {
        return errors.New("car not found")
    }

    return s.carRepo.Delete(id)
}

func (s *CarService) GetAvailableCars() ([]models.Car, error) {
    return s.carRepo.GetAvailableCars()
}