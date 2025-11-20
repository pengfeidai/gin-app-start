package service

import (
	"context"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"gin-app-start/internal/dto"
	"gin-app-start/internal/model"
	"gin-app-start/internal/repository"
	"gin-app-start/pkg/errors"
	"gin-app-start/pkg/logger"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(ctx context.Context, req *dto.CreateUserRequest) (*model.User, error)
	GetUser(ctx context.Context, id uint) (*model.User, error)
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
	UpdateUser(ctx context.Context, id uint, req *dto.UpdateUserRequest) (*model.User, error)
	DeleteUser(ctx context.Context, id uint) error
	ListUsers(ctx context.Context, page, pageSize int) ([]*model.User, int64, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (s *userService) CreateUser(ctx context.Context, req *dto.CreateUserRequest) (*model.User, error) {
	existingUser, err := s.userRepo.GetByUsername(ctx, req.Username)
	if err != nil && err != gorm.ErrRecordNotFound {
		logger.Error("Failed to query user", zap.Error(err), zap.String("username", req.Username))
		return nil, errors.WrapBusinessError(10010, "Failed to query user", err)
	}

	if existingUser != nil {
		return nil, errors.ErrUserExists
	}

	if req.Email != "" {
		existingUser, err = s.userRepo.GetByEmail(ctx, req.Email)
		if err != nil && err != gorm.ErrRecordNotFound {
			logger.Error("Failed to query email", zap.Error(err), zap.String("email", req.Email))
			return nil, errors.WrapBusinessError(10011, "Failed to query email", err)
		}
		if existingUser != nil {
			return nil, errors.NewBusinessError(10012, "Email already exists")
		}
	}

	salt := generateSalt()
	hashedPassword := hashPassword(req.Password, salt)

	user := &model.User{
		Username: req.Username,
		Email:    req.Email,
		Phone:    req.Phone,
		Password: hashedPassword,
		Salt:     salt,
		Status:   1,
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		logger.Error("Failed to create user", zap.Error(err), zap.String("username", req.Username))
		return nil, errors.WrapBusinessError(10013, "Failed to create user", err)
	}

	logger.Info("User created successfully",
		zap.String("username", user.Username),
		zap.Uint("user_id", user.ID),
	)

	return user, nil
}

func (s *userService) GetUser(ctx context.Context, id uint) (*model.User, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrUserNotFound
		}
		logger.Error("Failed to get user", zap.Error(err), zap.Uint("user_id", id))
		return nil, errors.WrapBusinessError(10014, "Failed to get user", err)
	}
	return user, nil
}

func (s *userService) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	user, err := s.userRepo.GetByUsername(ctx, username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrUserNotFound
		}
		logger.Error("Failed to get user", zap.Error(err), zap.String("username", username))
		return nil, errors.WrapBusinessError(10015, "Failed to get user", err)
	}
	return user, nil
}

func (s *userService) UpdateUser(ctx context.Context, id uint, req *dto.UpdateUserRequest) (*model.User, error) {
	user, err := s.userRepo.GetByID(ctx, id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrUserNotFound
		}
		return nil, errors.WrapBusinessError(10016, "Failed to get user", err)
	}

	if req.Email != "" {
		user.Email = req.Email
	}
	if req.Phone != "" {
		user.Phone = req.Phone
	}
	if req.Avatar != "" {
		user.Avatar = req.Avatar
	}
	if req.Status != 0 {
		user.Status = req.Status
	}

	if err := s.userRepo.Update(ctx, user); err != nil {
		logger.Error("Failed to update user", zap.Error(err), zap.Uint("user_id", id))
		return nil, errors.WrapBusinessError(10017, "Failed to update user", err)
	}

	logger.Info("User updated successfully", zap.Uint("user_id", id))
	return user, nil
}

func (s *userService) DeleteUser(ctx context.Context, id uint) error {
	if err := s.userRepo.Delete(ctx, id); err != nil {
		logger.Error("Failed to delete user", zap.Error(err), zap.Uint("user_id", id))
		return errors.WrapBusinessError(10018, "Failed to delete user", err)
	}

	logger.Info("User deleted successfully", zap.Uint("user_id", id))
	return nil
}

func (s *userService) ListUsers(ctx context.Context, page, pageSize int) ([]*model.User, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}

	offset := (page - 1) * pageSize
	users, total, err := s.userRepo.List(ctx, offset, pageSize)
	if err != nil {
		logger.Error("Failed to get user list", zap.Error(err))
		return nil, 0, errors.WrapBusinessError(10019, "Failed to get user list", err)
	}

	return users, total, nil
}

func generateSalt() string {
	salt := make([]byte, 16)
	rand.Read(salt)
	return hex.EncodeToString(salt)
}

func hashPassword(password, salt string) string {
	hash := md5.Sum([]byte(password + salt))
	return hex.EncodeToString(hash[:])
}

func VerifyPassword(password, salt, hashedPassword string) bool {
	return hashPassword(password, salt) == hashedPassword
}
