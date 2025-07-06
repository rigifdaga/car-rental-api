package repositories

import (
    "car-rental-api/internal/models"
    "gorm.io/gorm"
)

type CustomerRepository struct {
    db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *CustomerRepository {
    return &CustomerRepository{db: db}
}

func (r *CustomerRepository) Create(customer *models.Customer) error {
    return r.db.Create(customer).Error
}

func (r *CustomerRepository) GetAll() ([]models.Customer, error) {
    var customers []models.Customer
    err := r.db.Find(&customers).Error
    return customers, err
}

func (r *CustomerRepository) GetByID(id int) (*models.Customer, error) {
    var customer models.Customer
    err := r.db.First(&customer, id).Error
    if err != nil {
        return nil, err
    }
    return &customer, nil
}

func (r *CustomerRepository) Update(customer *models.Customer) error {
    return r.db.Save(customer).Error
}

func (r *CustomerRepository) Delete(id int) error {
    return r.db.Delete(&models.Customer{}, id).Error
}

func (r *CustomerRepository) GetByNIK(nik string) (*models.Customer, error) {
    var customer models.Customer
    err := r.db.Where("nik = ?", nik).First(&customer).Error
    if err != nil {
        return nil, err
    }
    return &customer, nil
}