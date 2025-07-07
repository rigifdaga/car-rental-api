package repositories

import (
    "car-rental-api/internal/models"
    "gorm.io/gorm"
)

type RentalRepository struct {
    db *gorm.DB
}

func NewRentalRepository(db *gorm.DB) *RentalRepository {
    return &RentalRepository{db: db}
}

func (r *RentalRepository) Create(rental *models.Rental) error {
    return r.db.Create(rental).Error
}

func (r *RentalRepository) GetAll() ([]models.Rental, error) {
    var rentals []models.Rental
    err := r.db.Preload("Customer").
        Preload("Customer.Membership").
        Preload("Car").
        Preload("BookingType").
        Preload("Driver").
        Find(&rentals).Error
    return rentals, err
}

func (r *RentalRepository) GetByID(id int) (*models.Rental, error) {
    var rental models.Rental
    err := r.db.Preload("Customer").
        Preload("Customer.Membership").
        Preload("Car").
        Preload("BookingType").
        Preload("Driver").
        First(&rental, id).Error
    if err != nil {
        return nil, err
    }
    return &rental, nil
}

func (r *RentalRepository) Update(rental *models.Rental) error {
    return r.db.Save(rental).Error
}

func (r *RentalRepository) Delete(id int) error {
    return r.db.Delete(&models.Rental{}, id).Error
}

func (r *RentalRepository) GetByCustomerID(customerID int) ([]models.Rental, error) {
    var rentals []models.Rental
    err := r.db.Preload("Customer").
        Preload("Customer.Membership").
        Preload("Car").
        Preload("BookingType").
        Preload("Driver").
        Where("customer_id = ?", customerID).
        Find(&rentals).Error
    return rentals, err
}

func (r *RentalRepository) GetActiveRentals() ([]models.Rental, error) {
    var rentals []models.Rental
    err := r.db.Preload("Customer").
        Preload("Customer.Membership").
        Preload("Car").
        Preload("BookingType").
        Preload("Driver").
        Where("finished = false").
        Find(&rentals).Error
    return rentals, err
}