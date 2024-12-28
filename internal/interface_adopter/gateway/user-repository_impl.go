package gateway

import (
	"Marketplace-backend/internal/entity"
	"Marketplace-backend/internal/repository"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/gofrs/uuid"
)

// struct user-repository-impl
type userRepositoryImp struct {
	db *sql.DB
}

// Create implements repository.UserRepository.
func (r *userRepositoryImp) Create(user *entity.User) error {

	// Generate UUID for the user if not already set
	NewUUID, err := uuid.NewV4()
	if err != nil {
		log.Printf("Error generating UUID: %v", err)
		return err
	}
	user.ID = NewUUID

	if user.RoleID == uuid.Nil {
		return fmt.Errorf("invalid RoleID: %v", user.RoleID)
	}

	// Proceed with inserting the user into the database
	query := `INSERT INTO users (id, username, password, email, first_name, last_name, role_id, role_name, created_at, updated_at)
						VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	result, err := r.db.Exec(query, user.ID, user.Username, user.Password, user.Email, user.FirstName, user.LastName, user.RoleID, user.RoleName, time.Now(), time.Now())
	if err != nil {
		log.Printf("Error inserting user: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error fetching rows affected: %v", err)
		return err
	}
	log.Printf("Rows affected: %v", rowsAffected)

	var insertedUser entity.User
	err = r.db.QueryRow(
		`SELECT id, username, password, email, first_name, last_name, role_id, role_name, created_at, updated_at 
			 FROM users 
			 WHERE email = $1`, user.Email).Scan(
		&insertedUser.ID,
		&insertedUser.Username,
		&insertedUser.Password,
		&insertedUser.Email,
		&insertedUser.FirstName,
		&insertedUser.LastName,
		&insertedUser.RoleID,
		&insertedUser.RoleName,
		&insertedUser.CreatedAt,
		&insertedUser.UpdatedAt,
	)
	if err != nil {
		log.Printf("Error retrieving inserted user: %v", err)
		return err
	}

	log.Printf("Inserted user: %v", insertedUser)
	return nil
}

// Delete implements repository.UserRepository.
func (r *userRepositoryImp) Delete(userID uuid.UUID) error {
	// delete the user from the database
	query := `DELETE FROM users WHERE id = $1`
	result, err := r.db.Exec(query, userID)
	if err != nil {
		log.Printf("Error deleting: %v", err)
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error fetching rows affected: %v", err)
		return err
	}

	if rowsAffected == 0 {
		log.Printf("User not found")
		return fmt.Errorf("user not found")
	}
	log.Printf("rows affeceted: %v", rowsAffected)
	return nil
}

// FindByEmail implements repository.UserRepository.
func (r *userRepositoryImp) FindByEmail(email string) (*entity.User, error) {
	user := &entity.User{}

	query := `SELECT id, username, password, email, first_name, last_name, role_id, role_name, created_at, updated_at FROM users WHERE email = $1`
	row := r.db.QueryRow(query, email)
	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.RoleID,
		&user.RoleName,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return user, nil
}

// Get implements repository.UserRepository.
func (r *userRepositoryImp) Get(userID uuid.UUID) (*entity.User, error) {
	//Define user
	var user entity.User

	query := "SELECT id, username, password, email, first_name, last_name, role_id, role_name, created_at, updated_at FROM users WHERE id = $1"
	err := r.db.QueryRow(query, userID).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.RoleID,
		&user.RoleName,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("User not found ")
			return nil, fmt.Errorf("user not found")
		}
		log.Printf("Error gettingg user: %v", err)
	}
	// return the  user is found
	return &user, nil

}

// List implements repository.UserRepository.
func (r *userRepositoryImp) List() ([]*entity.User, error) {

	rows, err := r.db.Query("SELECT id, username, password, email, first_name, last_name, role_id, role_name, created_at, updated_at FROM users ")
	if err != nil {
		log.Printf("Error getting users: %v", err)
		return nil, err
	}
	defer rows.Close()
	var users []*entity.User

	for rows.Next() {
		var user entity.User
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Password,
			&user.Email,
			&user.FirstName,
			&user.LastName,
			&user.RoleID,
			&user.RoleName,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			log.Printf("Error scanning  user: %v", err)
			return nil, err
		}
		users = append(users, &user)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Error getting  users: %v", err)
		return nil, err
	}
	return users, nil
}

// Update implements repository.UserRepository.
func (r *userRepositoryImp) Update(user *entity.User) error {
	// update the user in the database
	result, err := r.db.Exec(`
		UPDATE users
		SET username = $1, password = $2, email = $3, first_name = $4, last_name = $5, role_id = $6, role_name = $7, created_at = $8, updated_at = $9
		WHERE id = $10 `,
		user.Username, user.Password, user.Email, user.FirstName, user.LastName, user.RoleID, user.RoleName, user.CreatedAt, user.UpdatedAt, user.ID)
	if err != nil {
		log.Printf("Error updating user: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error getting rows affected: %v", err)
		return nil
	}

	if rowsAffected == 0 {
		log.Print("no rows affected")
		return nil
	}

	log.Printf("Rows affected: %v ", rowsAffected)
	return nil

}

// newUserRepository return a new instance UserRepositoryImpl
func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepositoryImp{db: db}
}
