package userbus

import (
	"net/mail"
	"time"

	"github.com/ardanlabs/service/business/types/name"
	"github.com/ardanlabs/service/business/types/role"
	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	Name         name.Name
	Email        mail.Address
	Roles        []role.Role
	PasswordHash []byte
	Department   string
	Enabled      bool
	DateCreated  time.Time
	DateUpdated  time.Time
}
