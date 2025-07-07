package services

import (
    "car-rental-api/internal/models"
    "car-rental-api/internal/repositories"
    "errors"
    "time"
)

type DriverService struct {
    driverRepo *repositories.DriverRepository
}

func NewDriverService(driverRepo *repositories.DriverRepository) *DriverService {
    return &DriverService{driverRepo: driverRepo}
}

func (s *DriverService) CreateDriver(driver *models.Driver) error {
    // Check if NIK already exists
    existing, _ := s.driverRepo.GetByNIK(driver.NIK)
    if existing != nil {
        return errors.New("driver with this NIK already exists")
    }

    driver.CreatedAt = time.Now()
    driver.UpdatedAt = time.Now()
    return s.driverRepo.Create(driver)
}

func (s *DriverService) GetAllDrivers() ([]models.Driver, error) {
    return s.driverRepo.GetAll()
}

func (s *DriverService) GetDriverByID(id int) (*models.Driver, error) {
    return s.driverRepo.GetByID(id)
}

func (s *DriverService) UpdateDriver(id int, driver *models.Driver) error {
    existing, err := s.driverRepo.GetByID(id)
    if err != nil {
        return errors.New("driver not found")
    }

    existing.Name = driver.Name
    existing.NIK = driver.NIK
    existing.PhoneNumber = driver.PhoneNumber
    existing.DailyCost = driver.DailyCost
    existing.UpdatedAt = time.Now()

    return s.driverRepo.Update(existing)
}

func (s *DriverService) DeleteDriver(id int) error {
    _, err := s.driverRepo.GetByID(id)
    if err != nil {
        return errors.New("driver not found")
    }

    return s.driverRepo.Delete(id)
}