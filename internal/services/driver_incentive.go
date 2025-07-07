package services

import (
    "car-rental-api/internal/models"
    "car-rental-api/internal/repositories"
    "errors"
    "time"
)

type DriverIncentiveService struct {
    driverIncentiveRepo *repositories.DriverIncentiveRepository
    rentalRepo          *repositories.RentalRepository
    carRepo             *repositories.CarRepository
}

func NewDriverIncentiveService(
    driverIncentiveRepo *repositories.DriverIncentiveRepository,
    rentalRepo *repositories.RentalRepository,
    carRepo *repositories.CarRepository,
) *DriverIncentiveService {
    return &DriverIncentiveService{
        driverIncentiveRepo: driverIncentiveRepo,
        rentalRepo:          rentalRepo,
        carRepo:             carRepo,
    }
}

func (s *DriverIncentiveService) CreateDriverIncentive(rentalID int) error {
    // Get rental details
    rental, err := s.rentalRepo.GetByID(rentalID)
    if err != nil {
        return errors.New("rental not found")
    }

    // Get car details for daily rent
    car, err := s.carRepo.GetByID(rental.CarID)
    if err != nil {
        return errors.New("car not found")
    }

    // Calculate incentive: (Days_of_Rent * Daily_car_Rent) * 5%
    days := int(rental.EndRent.Sub(rental.StartRent).Hours() / 24)
    if days <= 0 {
        days = 1
    }
    
    incentive := (float64(days) * car.DailyRent) * 0.05

    driverIncentive := &models.DriverIncentive{
        RentalID:  rentalID,
        Incentive: incentive,
        CreatedAt: time.Now(),
        UpdatedAt: time.Now(),
    }

    return s.driverIncentiveRepo.Create(driverIncentive)
}

func (s *DriverIncentiveService) GetAllDriverIncentives() ([]models.DriverIncentive, error) {
    return s.driverIncentiveRepo.GetAll()
}

func (s *DriverIncentiveService) GetDriverIncentiveByID(id int) (*models.DriverIncentive, error) {
    return s.driverIncentiveRepo.GetByID(id)
}

func (s *DriverIncentiveService) GetDriverIncentiveByRentalID(rentalID int) (*models.DriverIncentive, error) {
    return s.driverIncentiveRepo.GetByRentalID(rentalID)
}

func (s *DriverIncentiveService) UpdateDriverIncentive(id int, incentive *models.DriverIncentive) error {
    existing, err := s.driverIncentiveRepo.GetByID(id)
    if err != nil {
        return errors.New("driver incentive not found")
    }

    existing.Incentive = incentive.Incentive
    existing.UpdatedAt = time.Now()

    return s.driverIncentiveRepo.Update(existing)
}

func (s *DriverIncentiveService) DeleteDriverIncentive(id int) error {
    _, err := s.driverIncentiveRepo.GetByID(id)
    if err != nil {
        return errors.New("driver incentive not found")
    }

    return s.driverIncentiveRepo.Delete(id)
}

func (s *DriverIncentiveService) DeleteDriverIncentiveByRentalID(rentalID int) error {
    return s.driverIncentiveRepo.DeleteByRentalID(rentalID)
}