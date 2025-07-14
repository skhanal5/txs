package service

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/skhanal5/txs/internal/api/payload"
	"github.com/skhanal5/txs/internal/database/model"
	"github.com/skhanal5/txs/internal/database/repository"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	AuthenticateUser(request payload.AuthRequest) (payload.AuthResponse, error)
	RegisterUser(request payload.RegisterUserRequest) error
}

type authService struct {
	repository repository.AuthRepository
	logger     *zap.Logger
	secretKey []byte
}

func NewAuthService(authRepository repository.AuthRepository, logger *zap.Logger, secretKey []byte) AuthService {
	return &authService{
		repository: authRepository,
		logger:     logger,
		secretKey: secretKey,
	}
}

type Claims struct {
	UserID                string `json:"user_id"`
	Role                 string `json:"role"`
	jwt.RegisteredClaims `json:"registered_claims"`
}

func (a *authService) createJWT(id string) (string, error) {
	claims := Claims{
		UserID: id,
		Role:  "user",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := t.SignedString(a.secretKey)
	if err != nil {
		a.logger.Error("failed to sign JWT", zap.Error(err))
		return "", err
	}
	return s, nil
}

func (a *authService) AuthenticateUser(request payload.AuthRequest) (payload.AuthResponse, error) {
	user, err := a.repository.GetUserByEmail(request.Email)
	if err != nil {
		a.logger.Error("failed to get user by email", zap.Error(err))
		return payload.AuthResponse{}, err
	}
	if user == nil {
		a.logger.Error("user not found", zap.String("email", request.Email))
		return payload.AuthResponse{}, nil
	}
	err = bcrypt.CompareHashAndPassword(user.Password, []byte(request.Password))
	if err != nil {
		a.logger.Error("password mismatch", zap.Error(err))
		return payload.AuthResponse{}, err
	}
	accessToken, err := a.createJWT(user.ID)
	if err != nil {
		a.logger.Error("failed to create access token", zap.Error(err))
		return payload.AuthResponse{}, err
	}
	return payload.AuthResponse{
		AccessToken: accessToken,
	}, nil

}

func (s *authService) RegisterUser(request payload.RegisterUserRequest) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		s.logger.Error("failed to hash password", zap.Error(err))
		return err
	}
	user := model.User{
		Email:    request.Email,
		Password: hashedPassword,
	}
	err = s.repository.CreateUser(user)
	if err != nil {
		s.logger.Error("failed to create user", zap.Error(err))
		return err
	}
	return nil
}
