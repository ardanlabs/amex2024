package userbus

import (
	"context"
	"fmt"
	"net/mail"
	"time"

	"github.com/ardanlabs/service/business/sdk/order"
	"github.com/ardanlabs/service/business/sdk/page"
	"github.com/ardanlabs/service/foundation/logger"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Storer interface {
	Create(ctx context.Context, usr User) error
	Update(ctx context.Context, usr User) error
	Delete(ctx context.Context, usr User) error
	Query(ctx context.Context, filter QueryFilter, orderBy order.By, page page.Page) ([]User, error)
	Count(ctx context.Context, filter QueryFilter) (int, error)
	QueryByID(ctx context.Context, userID uuid.UUID) (User, error)
	QueryByEmail(ctx context.Context, email mail.Address) (User, error)
}

type Business struct {
	log    *logger.Logger
	storer Storer
}

func NewBusiness(log *logger.Logger, storer Storer) *Business {
	return &Business{
		log: log,
	}
}

func (b *Business) Create(ctx context.Context, nu NewUser) (User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, fmt.Errorf("generatefrompassword: %w", err)
	}

	now := time.Now()

	usr := User{
		ID:           uuid.New(),
		Name:         nu.Name,
		Email:        nu.Email,
		PasswordHash: hash,
		Roles:        nu.Roles,
		Department:   nu.Department,
		Enabled:      true,
		DateCreated:  now,
		DateUpdated:  now,
	}

	if err := b.storer.Create(ctx, usr); err != nil {
		return User{}, fmt.Errorf("create: %w", err)
	}

	return usr, nil
}

func (b *Business) Update(ctx context.Context, usr User, uu UpdateUser) (User, error) {
	if uu.Name != nil {
		usr.Name = *uu.Name
	}

	if uu.Email != nil {
		usr.Email = *uu.Email
	}

	if uu.Roles != nil {
		usr.Roles = uu.Roles
	}

	if uu.Password != nil {
		pw, err := bcrypt.GenerateFromPassword([]byte(*uu.Password), bcrypt.DefaultCost)
		if err != nil {
			return User{}, fmt.Errorf("generatefrompassword: %w", err)
		}
		usr.PasswordHash = pw
	}

	if uu.Department != nil {
		usr.Department = *uu.Department
	}

	if uu.Enabled != nil {
		usr.Enabled = *uu.Enabled
	}

	usr.DateUpdated = time.Now()

	if err := b.storer.Update(ctx, usr); err != nil {
		return User{}, fmt.Errorf("update: %w", err)
	}

	return usr, nil
}

func (b *Business) Delete(ctx context.Context, usr User) error {
	if err := b.storer.Delete(ctx, usr); err != nil {
		return fmt.Errorf("delete: %w", err)
	}

	return nil
}

func (b *Business) Query(ctx context.Context, filter QueryFilter, orderBy order.By, page page.Page) ([]User, error) {
	users, err := b.storer.Query(ctx, filter, orderBy, page)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	return users, nil
}

// Count returns the total number of users.
func (b *Business) Count(ctx context.Context, filter QueryFilter) (int, error) {
	return b.storer.Count(ctx, filter)
}

func (b *Business) QueryByID(ctx context.Context, userID uuid.UUID) (User, error) {
	user, err := b.storer.QueryByID(ctx, userID)
	if err != nil {
		return User{}, fmt.Errorf("query: userID[%s]: %w", userID, err)
	}

	return user, nil
}

// QueryByEmail finds the user by a specified user email.
func (b *Business) QueryByEmail(ctx context.Context, email mail.Address) (User, error) {
	user, err := b.storer.QueryByEmail(ctx, email)
	if err != nil {
		return User{}, fmt.Errorf("query: email[%s]: %w", email, err)
	}

	return user, nil
}
