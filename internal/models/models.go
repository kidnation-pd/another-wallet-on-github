package models

import (
	"time"

	"github.com/google/uuid"
)

// User accaunt
type User struct {
	ID           uuid.UUID
	Username     string
	PasswordHash string
	Email        string
	Phone        string
	CreatedAt    time.Time
	DeletedAt    time.Time
}

type Wallet struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Balance   float64
	Currency  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Transaction struct {
	ID          uuid.UUID
	WalletID    uuid.UUID
	Amount      float64
	Type        string // "deposit", "withdrawal", "transfer"
	Status      string // "pending", "completed", "failed"
	Description string
	CreatedAt   time.Time
}
