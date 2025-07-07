package repositories

import (
    "car-rental-api/internal/models"
    "gorm.io/gorm"
)

type MembershipRepository struct {
    db *gorm.DB
}

func NewMembershipRepository(db *gorm.DB) *MembershipRepository {
    return &MembershipRepository{db: db}
}

func (r *MembershipRepository) Create(membership *models.Membership) error {
    return r.db.Create(membership).Error
}

func (r *MembershipRepository) GetAll() ([]models.Membership, error) {
    var memberships []models.Membership
    err := r.db.Find(&memberships).Error
    return memberships, err
}

func (r *MembershipRepository) GetByID(id int) (*models.Membership, error) {
    var membership models.Membership
    err := r.db.First(&membership, id).Error
    if err != nil {
        return nil, err
    }
    return &membership, nil
}

func (r *MembershipRepository) Update(membership *models.Membership) error {
    return r.db.Save(membership).Error
}

func (r *MembershipRepository) Delete(id int) error {
    return r.db.Delete(&models.Membership{}, id).Error
}