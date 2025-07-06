package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/skhanal5/txs/internal/database/model"
)

type UserRepository interface {
	GetUserByID(id string) (*model.User, error)
	CreateUser(user *model.User) error
}

type userRepository struct {
	conn *pgxpool.Pool
}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) GetUserByID(id string) (*model.User, error) {
	return nil, nil
}		

func (r *userRepository) CreateUser(user *model.User) error {
	return nil 
}