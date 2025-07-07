package services

import (
    "car-rental-api/internal/models"
    "car-rental-api/internal/repositories"
    "errors"
    "time"
)

type MembershipService struct {
    membershipRepo *repositories.MembershipRepository
}

func NewMembershipService(membershipRepo *repositories.MembershipRepository) *MembershipService {
    return &MembershipService{membershipRepo: membershipRepo}
}

func (s *MembershipService) CreateMembership(membership *models.Membership) error {
    membership.CreatedAt = time.Now()
    membership.UpdatedAt = time.Now()
    return s.membershipRepo.Create(membership)
}

func (s *MembershipService) GetAllMemberships() ([]models.Membership, error) {
    return s.membershipRepo.GetAll()
}

func (s *MembershipService) GetMembershipByID(id int) (*models.Membership, error) {
    return s.membershipRepo.GetByID(id)
}

func (s *MembershipService) UpdateMembership(id int, membership *models.Membership) error {
    existing, err := s.membershipRepo.GetByID(id)
    if err != nil {
        return errors.New("membership not found")
    }

    existing.MembershipName = membership.MembershipName
    existing.DiscountPercentage = membership.DiscountPercentage
    existing.UpdatedAt = time.Now()

    return s.membershipRepo.Update(existing)
}

func (s *MembershipService) DeleteMembership(id int) error {
    _, err := s.membershipRepo.GetByID(id)
    if err != nil {
        return errors.New("membership not found")
    }

    return s.membershipRepo.Delete(id)
}