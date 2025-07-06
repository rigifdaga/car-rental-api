package services

import (
    "car-rental-api/internal/models"
    "car-rental-api/internal/repositories"
    "errors"
    "time"
)

type RentalService struct {
    rentalRepo   *repositories.RentalRepository
    customerRepo *repositories.CustomerRepository
    carRepo      *repositories.CarRepository
}

func NewRentalService(rentalRepo *repositories.RentalRepository, customerRepo *repositories.CustomerRepository, carRepo *repositories.CarRepository) *RentalService {
    return &RentalService{
        rentalRepo:   rentalRepo,
        customerRepo: customerRepo,
        carRepo:      carRepo,
    }
}

func (s *RentalService) CreateRental(rental *models.Rental) error {
    // Validate customer exists
    _, err := s.customerRepo.GetByID(rental.CustomerID)
    if err != nil {
        return errors.New("customer not found")
    }

    // Validate car exists and available
    car, err := s.carRepo.GetByID(rental.CarID)
    if err != nil {
        return errors.New("car not found")
    }

    if car.Stock <= 0 {
        return errors.New("car is not available")
    }

    // Calculate total cost
    days := int(rental.EndRent.Sub(rental.StartRent).Hours() / 24)
    if days <= 0 {
        days = 1
    }
    rental.TotalCost = float64(days) * car.DailyRent

    rental.CreatedAt = time.Now()
    rental.UpdatedAt = time.Now()
    rental.Finished = false

    // Reduce car stock
    car.Stock--
    s.carRepo.Update(car)

    return s.rentalRepo.Create(rental)
}

func (s *RentalService) GetAllRentals() ([]models.Rental, error) {
    return s.rentalRepo.GetAll()
}

func (s *RentalService) GetRentalByID(id int) (*models.Rental, error) {
    return s.rentalRepo.GetByID(id)
}

func (s *RentalService) UpdateRental(id int, rental *models.Rental) error {
    existing, err := s.rentalRepo.GetByID(id)
    if err != nil {
        return errors.New("rental not found")
    }

    // If finishing rental, increase car stock
    if rental.Finished && !existing.Finished {
        car, err := s.carRepo.GetByID(existing.CarID)
        if err == nil {
            car.Stock++
            s.carRepo.Update(car)
        }
    }

    existing.StartRent = rental.StartRent
    existing.EndRent = rental.EndRent
    existing.TotalCost = rental.TotalCost
    existing.Finished = rental.Finished
    existing.UpdatedAt = time.Now()

    return s.rentalRepo.Update(existing)
}

func (s *RentalService) DeleteRental(id int) error {
    rental, err := s.rentalRepo.GetByID(id)
    if err != nil {
        return errors.New("rental not found")
    }

    // If rental is not finished, return car stock
    if !rental.Finished {
        car, err := s.carRepo.GetByID(rental.CarID)
        if err == nil {
            car.Stock++
            s.carRepo.Update(car)
        }
    }

    return s.rentalRepo.Delete(id)
}

func (s *RentalService) GetActiveRentals() ([]models.Rental, error) {
    return s.rentalRepo.GetActiveRentals()
}