package repositories

import (
    "car-rental-api/internal/models"
    "gorm.io/gorm"
)

type DriverIncentiveRepository struct {
    db *gorm.DB
}

func NewDriverIncentiveRepository(db *gorm.DB) *DriverIncentiveRepository {
    return &DriverIncentiveRepository{db: db}
}

func (r *DriverIncentiveRepository) Create(incentive *models.DriverIncentive) error {
    return r.db.Create(incentive).Error
}

func (r *DriverIncentiveRepository) GetAll() ([]models.DriverIncentive, error) {
    var incentives []models.DriverIncentive
    err := r.db.Preload("Rental").Find(&incentives).Error
    return incentives, err
}

func (r *DriverIncentiveRepository) GetByID(id int) (*models.DriverIncentive, error) {
    var incentive models.DriverIncentive
    err := r.db.Preload("Rental").First(&incentive, id).Error
    if err != nil {
        return nil, err
    }
    return &incentive, nil
}

func (r *DriverIncentiveRepository) GetByRentalID(rentalID int) (*models.DriverIncentive, error) {
    var incentive models.DriverIncentive
    err := r.db.Preload("Rental").Where("rental_id = ?", rentalID).First(&incentive).Error
    if err != nil {
        return nil, err
    }
    return &incentive, nil
}

func (r *DriverIncentiveRepository) Update(incentive *models.DriverIncentive) error {
    return r.db.Save(incentive).Error
}

func (r *DriverIncentiveRepository) Delete(id int) error {
    return r.db.Delete(&models.DriverIncentive{}, id).Error
}

func (r *DriverIncentiveRepository) DeleteByRentalID(rentalID int) error {
    return r.db.Where("rental_id = ?", rentalID).Delete(&models.DriverIncentive{}).Error
}