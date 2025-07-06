package services

import (
    "car-rental-api/internal/models"
    "car-rental-api/internal/repositories"
    "errors"
    "time"
)

type CustomerService struct {
    customerRepo *repositories.CustomerRepository
}

func NewCustomerService(customerRepo *repositories.CustomerRepository) *CustomerService {
    return &CustomerService{customerRepo: customerRepo}
}

func (s *CustomerService) CreateCustomer(customer *models.Customer) error {
    // Check if NIK already exists
    existing, _ := s.customerRepo.GetByNIK(customer.NIK)
    if existing != nil {
        return errors.New("customer with this NIK already exists")
    }

    customer.CreatedAt = time.Now()
    customer.UpdatedAt = time.Now()
    return s.customerRepo.Create(customer)
}

func (s *CustomerService) GetAllCustomers() ([]models.Customer, error) {
    return s.customerRepo.GetAll()
}

func (s *CustomerService) GetCustomerByID(id int) (*models.Customer, error) {
    return s.customerRepo.GetByID(id)
}

func (s *CustomerService) UpdateCustomer(id int, customer *models.Customer) error {
    existing, err := s.customerRepo.GetByID(id)
    if err != nil {
        return errors.New("customer not found")
    }

    existing.Name = customer.Name
    existing.NIK = customer.NIK
    existing.PhoneNumber = customer.PhoneNumber
    existing.UpdatedAt = time.Now()

    return s.customerRepo.Update(existing)
}

func (s *CustomerService) DeleteCustomer(id int) error {
    _, err := s.customerRepo.GetByID(id)
    if err != nil {
        return errors.New("customer not found")
    }

    return s.customerRepo.Delete(id)
}