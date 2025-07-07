package repositories

import (
    "car-rental-api/internal/models"
    "gorm.io/gorm"
)

type DriverRepository struct {
    db *gorm.DB
}

func NewDriverRepository(db *gorm.DB) *DriverRepository {
    return &DriverRepository{db: db}
}

func (r *DriverRepository) Create(driver *models.Driver) error {
    return r.db.Create(driver).Error
}

func (r *DriverRepository) GetAll() ([]models.Driver, error) {
    var drivers []models.Driver
    err := r.db.Find(&drivers).Error
    return drivers, err
}

func (r *DriverRepository) GetByID(id int) (*models.Driver, error) {
    var driver models.Driver
    err := r.db.First(&driver, id).Error
    if err != nil {
        return nil, err
    }
    return &driver, nil
}

func (r *DriverRepository) Update(driver *models.Driver) error {
    return r.db.Save(driver).Error
}

func (r *DriverRepository) Delete(id int) error {
    return r.db.Delete(&models.Driver{}, id).Error
}

func (r *DriverRepository) GetByNIK(nik string) (*models.Driver, error) {
    var driver models.Driver
    err := r.db.Where("nik = ?", nik).First(&driver).Error
    if err != nil {
        return nil, err
    }
    return &driver, nil
}