package user

import (
	"time"

	"github.com/google/uuid"
)

type SubscriptionTier string

const (
	TierFree    SubscriptionTier = "FREE"
	TierPremium SubscriptionTier = "PREMIUM"
	TierVIP     SubscriptionTier = "VIP"
)

type User struct {
	ID               uuid.UUID
	Email            string
	PasswordHash     string
	Name             string
	SubscriptionTier SubscriptionTier
	SubscriptionEnd  *time.Time
	CreatedAt        time.Time
}

func New(email, passwordHash, name string) *User {
	return &User{
		ID:               uuid.New(),
		Email:            email,
		PasswordHash:     passwordHash,
		Name:             name,
		SubscriptionTier: TierFree,
		CreatedAt:        time.Now(),
	}
}

func (u *User) IsPremium() bool {
	if u.SubscriptionTier == TierFree {
		return false
	}
	if u.SubscriptionEnd == nil {
		return true
	}
	return time.Now().Before(*u.SubscriptionEnd)
}
