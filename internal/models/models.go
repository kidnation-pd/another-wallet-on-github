package models

import (
	"time"

	"github.com/google/uuid"
)

type (
	UID uuid.UUID // User id
	WID uuid.UUID // Wallet id
	TID uuid.UUID // Transaction id
	HID uuid.UUID // History id
)

// User accaunt
type User struct {
	ID           UID
	Username     string
	PasswordHash string
	Email        string
	Phone        string
	CreatedAt    time.Time
	DeletedAt    time.Time
}

type Wallet struct {
	ID        WID
	UserID    UID
	Balance   float64
	Currency  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Transaction struct {
	ID          TID
	WalletID    WID
	Amount      float64
	Type        string // "deposit", "withdrawal", "transfer"
	Status      string // "pending", "completed", "failed"
	Description string
	CreatedAt   time.Time
}
