package db

import (
	"Marketplace-backend/pkg/config"
	"database/sql"
	"fmt"
	"log"
)

func ConnectDB(cfg *config.DBConfig) (*sql.DB, error) {
	constr := cfg.ConnectionString()
	db, err := sql.Open("postgres", constr)
	if err != nil {
		return nil, fmt.Errorf("fialed to connect to the database: %v", err)
	}

	// ping the database to ensure the connection is successful
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping the database: %v", err)
	}

	log.Println("successfully connected to the database")

	return db, nil
}

func CreatTables(db *sql.DB) error {
	// create user table

	userTable := `CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    role_id UUID REFERENCES user_roles(id) ON DELETE CASCADE,  -- Reference the roles table instead
    role_name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
`

	// Create products table
	productTable := `CREATE TABLE IF NOT EXISTS products (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		name VARCHAR(255) NOT NULL,
		description TEXT NOT NULL,
		price NUMERIC(10, 2) DEFAULT 0,
		stock INT DEFAULT 0,
		category VARCHAR(255) NOT NULL,
		sku VARCHAR(255) UNIQUE NOT NULL,
		image_urls TEXT,
		discount NUMERIC(10, 2) DEFAULT 0,
		is_active BOOLEAN DEFAULT TRUE,
		tags TEXT,
		additional_info TEXT,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		deleted_at TIMESTAMP NULL
);`

	// Create sales_management

	SalesManagement := `CREATE TABLE IF NOT EXISTS sales(
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    total_sales NUMERIC(15, 2) NOT NULL DEFAULT 0.0,
    total_orders INT NOT NULL DEFAULT 0,
    order_id UUID REFERENCES users(id) ON DELETE CASCADE,
    product_id UUID REFERENCES products(id) ON DELETE CASCADE,
    quantity INT NOT NULL DEFAULT 1,
    total_price NUMERIC(15, 2) NOT NULL,
    order_data TIMESTAMP NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP, -- Timestamp when the sales entry was created, defaulted to current time.
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP  -- Timestamp when the sales entry was last updated, defaulted to current time.
);
`

// Create review_rating table

reviewRatingTable := `CREATE TABLE IF NOT EXISTS review_rating (
		id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		product_id UUID REFERENCES products(id) ON DELETE CASCADE,
		user_id UUID REFERENCES users(id) ON DELETE CASCADE,
		rating INT NOT NULL,
		comment TEXT,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP

);`

	// Create tokens table
	tokenTable := `CREATE TABLE IF NOT EXISTS tokens (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			user_id UUID REFERENCES users(id) ON DELETE CASCADE,
			token VARCHAR(255) UNIQUE NOT NULL,
			expires_at TIMESTAMP NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			deleted_at TIMESTAMP NULL
		);`

	// Create roles table
	roleTable := `CREATE TABLE IF NOT EXISTS roles (
			id UUID PRIMARY KEY,
			name VARCHAR(255) UNIQUE NOT NULL
		);
`

	// Create permissions table
	permissionTable := `CREATE TABLE IF NOT EXISTS permissions (
				id UUID PRIMARY KEY,
				name VARCHAR(255) UNIQUE NOT NULL
			);`

	// Create user_roles table
	userRoleTable := `CREATE TABLE IF NOT EXISTS user_roles (
    user_id UUID REFERENCES users(id) ON DELETE CASCADE,
    role_id UUID REFERENCES roles(id) ON DELETE CASCADE,  -- Change to UUID
    PRIMARY KEY (user_id, role_id)
);`

	// Create user_permissions table
	userPermissionTable := `CREATE TABLE IF NOT EXISTS user_permissions (
				user_id UUID REFERENCES users(id) ON DELETE CASCADE,
				permission_id INT REFERENCES permissions(id) ON DELETE CASCADE,
				PRIMARY KEY (user_id, permission_id)
			);`
	// Execute the table creation queries
	queries := []string{tokenTable, userRoleTable, roleTable, permissionTable, userPermissionTable, userTable, productTable, SalesManagement, reviewRatingTable}
	for _, query := range queries {
		if _, err := db.Exec(query); err != nil {
			return fmt.Errorf("failed to created all tables")
		}
	}

	log.Println("Successfully created all tables")
	return nil
}
