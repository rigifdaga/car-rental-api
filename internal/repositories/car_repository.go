package repositories

import (
    "car-rental-api/internal/models"
    "gorm.io/gorm"
)

type CarRepository struct {
    db *gorm.DB
}

func NewCarRepository(db *gorm.DB) *CarRepository {
    return &CarRepository{db: db}
}

func (r *CarRepository) Create(car *models.Car) error {
    return r.db.Create(car).Error
}

func (r *CarRepository) GetAll() ([]models.Car, error) {
    var cars []models.Car
    err := r.db.Find(&cars).Error
    return cars, err
}

func (r *CarRepository) GetByID(id int) (*models.Car, error) {
    var car models.Car
    err := r.db.First(&car, id).Error
    if err != nil {
        return nil, err
    }
    return &car, nil
}

func (r *CarRepository) Update(car *models.Car) error {
    return r.db.Save(car).Error
}

func (r *CarRepository) Delete(id int) error {
    return r.db.Delete(&models.Car{}, id).Error
}

func (r *CarRepository) GetAvailableCars() ([]models.Car, error) {
    var cars []models.Car
    err := r.db.Where("stock > 0").Find(&cars).Error
    return cars, err
}