package service

import (
	"Marketplace-backend/internal/entity"
	"Marketplace-backend/internal/repository"
	"Marketplace-backend/pkg/utils"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gofrs/uuid"
)

type UserService interface {
	// Create a NEW USER
	RegisterUser(Username, Email, Password, FirstName, LastName, roleName string, RoleID uuid.UUID) (*entity.User, error)
	// authenticate existing user
	AuthenticateUser(email, password string) (*entity.User, error)
	// get user by id
	GetUserById(userID uuid.UUID) (*entity.User, error)
	// get all users
	GetAllUsers() ([]*entity.User, error)
	// update user
	UpdateUser(user *entity.User) error
	// delete user
	DeleteUser(userID uuid.UUID) error
}

// userServiceImp is the implementation of the UserService interface
type userService struct {
	repo      repository.UserRepository
	tokenRepo repository.TokenRepository
}

// AuthenticatenUser implements UserService.
func (s *userService) AuthenticateUser(email string, password string) (*entity.User, error) {
	// find user email
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or passsword")
	}
	// generate err
	newToken, err := uuid.NewV4()
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	// create token
	token := &entity.Token{
		ID:        newToken,
		UserID:    user.ID,
		Token:     newToken.String(),
		ExpiredAt: time.Now().Add(24 * time.Hour),
	}

	// store token in database
	if err := s.tokenRepo.Create(token); err != nil {
		return nil, errors.New("failed to create token")
	}
	return user, nil
}

// DeleteUser implements UserService.
func (s *userService) DeleteUser(userID uuid.UUID) error {
	// check if user exists by their ID
	_, err := s.repo.Get(userID)
	if err != nil {
		return errors.New("user not found")
	}
	// delete user
	if err := s.repo.Delete(userID); err != nil {
		log.Printf("Error deleting user: %v", err)
		return err
	}
	log.Printf("User deleted successfully")
	return nil
}

// GetAllUsers implements UserService.
func (s *userService) GetAllUsers() ([]*entity.User, error) {
	user, err := s.repo.List()
	if err != nil {
		return nil, fmt.Errorf("failed to get all users")
	}
	return user, nil
}

// GetUserById implements UserService.
func (s *userService) GetUserById(userID uuid.UUID) (*entity.User, error) {
	user, err := s.repo.Get(userID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %v", err)
	}
	return user, nil
}

// RegisterUser implements UserService.
func (s *userService) RegisterUser(Username, Email, Password, FirstName, LastName, RoleName string, RoleID uuid.UUID) (*entity.User, error) {
	// CHECK IF USER ALREADY EXISTS

	if _, err := s.repo.FindByEmail(Email); err == nil {
		return nil, errors.New("user already exists")
	}

	// hash password using bcrypt

	hashPassword, err := utils.HashPassword(Password)
	if err != nil {
		return nil, err
	}
	// create new user
	user := &entity.User{
		Username:  Username,
		Email:     Email,
		Password:  hashPassword,
		FirstName: FirstName,
		LastName:  LastName,
		RoleName:  RoleName,
		RoleID:    RoleID,
		
	}

	// save user to database
	if err := s.repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser implements UserService.
func (s *userService) UpdateUser(user *entity.User) error {
	// check if user existx by their ID
	_, err := s.repo.Get(user.ID)
	if err != nil {
		return errors.New("user not found")
	}



	// call repository to update user
	if err := s.repo.Update(user); err != nil {
		return fmt.Errorf("error updating user: %v", err)
	}
	return nil
}

// NewUserService returns a new instance of the UserService interface
func NewUserService(userRepo repository.UserRepository, tokenRepo repository.TokenRepository) UserService {
	return &userService{
		repo:      userRepo,
		tokenRepo: tokenRepo,
	}
}
