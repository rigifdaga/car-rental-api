package services

import (
    "car-rental-api/internal/models"
    "car-rental-api/internal/repositories"
    "errors"
    "time"
)

type RentalService struct {
    rentalRepo             *repositories.RentalRepository
    customerRepo           *repositories.CustomerRepository
    carRepo                *repositories.CarRepository
    bookingTypeRepo        *repositories.BookingTypeRepository
    driverRepo             *repositories.DriverRepository
    driverIncentiveService *DriverIncentiveService
    membershipRepo         *repositories.MembershipRepository
}

func NewRentalService(
    rentalRepo *repositories.RentalRepository,
    customerRepo *repositories.CustomerRepository,
    carRepo *repositories.CarRepository,
    bookingTypeRepo *repositories.BookingTypeRepository,
    driverRepo *repositories.DriverRepository,
    driverIncentiveService *DriverIncentiveService,
    membershipRepo *repositories.MembershipRepository,
) *RentalService {
    return &RentalService{
        rentalRepo:             rentalRepo,
        customerRepo:           customerRepo,
        carRepo:                carRepo,
        bookingTypeRepo:        bookingTypeRepo,
        driverRepo:             driverRepo,
        driverIncentiveService: driverIncentiveService,
        membershipRepo:         membershipRepo,
    }
}

func (s *RentalService) CreateRental(rental *models.Rental) error {
    // Validate customer exists
    customer, err := s.customerRepo.GetByID(rental.CustomerID)
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

    // Validate booking type
    bookingType, err := s.bookingTypeRepo.GetByID(rental.BookingTypeID)
    if err != nil {
        return errors.New("booking type not found")
    }

    // Calculate days of rent
    days := int(rental.EndRent.Sub(rental.StartRent).Hours()/24) + 1
    if days <= 0 {
        days = 1
    }

    // Calculate base cost
    rental.TotalCost = float64(days) * car.DailyRent

    // Calculate discount based on membership
    rental.Discount = 0
    if customer.MembershipID != nil {
        membership, err := s.membershipRepo.GetByID(*customer.MembershipID)
        if err == nil && membership != nil {
            rental.Discount = rental.TotalCost * (membership.DiscountPercentage / 100)
        }
    }

    // Validate driver and calculate driver cost for "Car & Driver" booking
    rental.TotalDriverCost = 0
    if bookingType.BookingType == "Car & Driver" {
        if rental.DriverID == nil {
            return errors.New("driver ID required for Car & Driver booking")
        }
        driver, err := s.driverRepo.GetByID(*rental.DriverID)
        if err != nil {
            return errors.New("driver not found")
        }
        rental.TotalDriverCost = float64(days) * driver.DailyCost
    } else if rental.DriverID != nil {
        return errors.New("driver ID should be null for Car Only booking")
    }

    rental.CreatedAt = time.Now()
    rental.UpdatedAt = time.Now()
    rental.Finished = false

    // Reduce car stock
    car.Stock--
    if err := s.carRepo.Update(car); err != nil {
        return errors.New("failed to update car stock")
    }

    // Create rental
    if err := s.rentalRepo.Create(rental); err != nil {
        return err
    }

    // Create driver incentive if applicable
    if rental.DriverID != nil {
        if err := s.driverIncentiveService.CreateDriverIncentive(rental.RentalID); err != nil {
            return errors.New("failed to create driver incentive: " + err.Error())
        }
    }

    return nil
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

    // Validate customer exists
    customer, err := s.customerRepo.GetByID(rental.CustomerID)
    if err != nil {
        return errors.New("customer not found")
    }

    // Validate car exists
    car, err := s.carRepo.GetByID(rental.CarID)
    if err != nil {
        return errors.New("car not found")
    }

    // Validate booking type
    bookingType, err := s.bookingTypeRepo.GetByID(rental.BookingTypeID)
    if err != nil {
        return errors.New("booking type not found")
    }

    // If car changed, update stock
    if existing.CarID != rental.CarID {
        oldCar, err := s.carRepo.GetByID(existing.CarID)
        if err == nil {
            oldCar.Stock++
            s.carRepo.Update(oldCar)
        }
        car.Stock--
        if err := s.carRepo.Update(car); err != nil {
            return errors.New("failed to update car stock")
        }
    }

    // Calculate days of rent
    days := int(rental.EndRent.Sub(rental.StartRent).Hours()/24) + 1
    if days <= 0 {
        days = 1
    }

    // Recalculate total cost
    existing.TotalCost = float64(days) * car.DailyRent

    // Recalculate discount
    existing.Discount = 0
    if customer.MembershipID != nil {
        membership, err := s.membershipRepo.GetByID(*customer.MembershipID)
        if err == nil && membership != nil {
            existing.Discount = existing.TotalCost * (membership.DiscountPercentage / 100)
        }
    }

    // Validate driver and recalculate driver cost
    existing.TotalDriverCost = 0
    if bookingType.BookingType == "Car & Driver" {
        if rental.DriverID == nil {
            return errors.New("driver ID required for Car & Driver booking")
        }
        driver, err := s.driverRepo.GetByID(*rental.DriverID)
        if err != nil {
            return errors.New("driver not found")
        }
        existing.TotalDriverCost = float64(days) * driver.DailyCost
    } else if rental.DriverID != nil {
        return errors.New("driver ID should be null for Car Only booking")
    }

    // Update driver incentive if driver changed or rental details changed
    if existing.DriverID != rental.DriverID || existing.TotalCost != rental.TotalCost {
        if existing.DriverID != nil {
            s.driverIncentiveService.DeleteDriverIncentiveByRentalID(id)
        }
        if rental.DriverID != nil {
            if err := s.driverIncentiveService.CreateDriverIncentive(id); err != nil {
                return errors.New("failed to update driver incentive: " + err.Error())
            }
        }
    }

    // If finishing rental, increase car stock
    if rental.Finished && !existing.Finished {
        car.Stock++
        if err := s.carRepo.Update(car); err != nil {
            return errors.New("failed to update car stock")
        }
    }

    existing.CustomerID = rental.CustomerID
    existing.CarID = rental.CarID
    existing.StartRent = rental.StartRent
    existing.EndRent = rental.EndRent
    existing.BookingTypeID = rental.BookingTypeID
    existing.DriverID = rental.DriverID
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

    // Delete driver incentive if exists
    s.driverIncentiveService.DeleteDriverIncentiveByRentalID(id)

    return s.rentalRepo.Delete(id)
}

func (s *RentalService) GetActiveRentals() ([]models.Rental, error) {
    return s.rentalRepo.GetActiveRentals()
}