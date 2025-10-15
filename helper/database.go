package helper

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var DB *sql.DB

func init() {
	if err := dbConnection(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	if err := createTables(); err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}
}

func dbConnection() error {
	var err error
	connectionString := "root:Root123@tcp(127.0.0.1:3306)/golang"

	DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		return fmt.Errorf("connection failed: %w", err)
	}

	if err := DB.Ping(); err != nil {
		return fmt.Errorf("ping failed: %w", err)
	}
	fmt.Println(" Database connected successfully!")
	return nil
}

func createTables() error {
	if _, err := DB.Exec(`
	CREATE TABLE IF NOT EXISTS users(
		id INT AUTO_INCREMENT PRIMARY KEY,
		userId VARCHAR(100) NOT NULL UNIQUE,
		email VARCHAR(100) NOT NULL UNIQUE,
		Mobile_no VARCHAR(15) NOT NULL,
		password VARCHAR(100) NOT NULL,
		username VARCHAR(100) NOT NULL,
		balance DECIMAL(10,2) DEFAULT 0.00
	)`); err != nil {
		return fmt.Errorf("error creating users table: %w", err)
	}

	if _, err := DB.Exec(`
	CREATE TABLE IF NOT EXISTS history(
		id INT AUTO_INCREMENT PRIMARY KEY,
		from_userid VARCHAR(100) NOT NULL,
		to_userid VARCHAR(100) NOT NULL,
		amount DECIMAL(10,2) NOT NULL,
		transaction_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`); err != nil {
		return fmt.Errorf("error creating history table: %w", err)
	}

	if _, err := DB.Exec(`
	CREATE TABLE IF NOT EXISTS deleted_users(
		id INT PRIMARY KEY,
		name VARCHAR(100),
		email VARCHAR(100),
		deleted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`); err != nil {
		return fmt.Errorf("error creating deleted_users table: %w", err)
	}

	return nil
}

// Register user
func RegisterUser(userId, username, email, strMobile_no, password, confirm_password string) error {
	if len(strMobile_no) != 10 {
		return fmt.Errorf("mobile must be 10 digits")
	}
	if !strings.EqualFold(password, confirm_password) {
		return fmt.Errorf("passwords do not match")
	}

	// hash password
	hashpassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	_, err := DB.Exec(`INSERT INTO users (userId, email, Mobile_no, username, password) VALUES (?, ?, ?, ?, ?)`,
		userId, email, strMobile_no, username, string(hashpassword))
	if err != nil {
		return fmt.Errorf("error registering user: %w", err)
	}
	return nil
}

// Validate login
func ValidateUser(userId, password string) error {
	var hashpassword string

	err := DB.QueryRow("SELECT password FROM users WHERE userId=?", userId).Scan(&hashpassword)
	if err != nil {
		return fmt.Errorf("invalid userid or password")
	}

	if bcrypt.CompareHashAndPassword([]byte(hashpassword), []byte(password)) != nil {
		return fmt.Errorf("invalid password")
	}
	return nil
}

// Transaction between users (by userId)
func Transection_process(senderID, receiverID string, amount float64) error {
	if senderID == receiverID {
		return fmt.Errorf("cannot send money to yourself")
	}

	tx, err := DB.Begin()
	if err != nil {
		return fmt.Errorf("transaction start failed: %w", err)
	}
	defer tx.Rollback() // will rollback if commit not called

	var senderBalance float64
	err = tx.QueryRow("SELECT balance FROM users WHERE userId=?", senderID).Scan(&senderBalance)
	if err == sql.ErrNoRows {
		tx.Rollback()
		return fmt.Errorf("sender not found")
	} else if err != nil {
		tx.Rollback()
		return fmt.Errorf("database error: %w", err)
	}

	var receiverExists bool
	err = tx.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE userId=?)", receiverID).Scan(&receiverExists)
	if err != nil || !receiverExists {
		tx.Rollback()
		return fmt.Errorf("receiver not found")
	}

	if senderBalance < amount {
		tx.Rollback()
		return fmt.Errorf("insufficient balance")
	}

	// Deduct sender
	res, err := tx.Exec("UPDATE users SET balance = balance - ? WHERE userId=?", amount, senderID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error updating sender balance: %w", err)
	}
	if rows, _ := res.RowsAffected(); rows == 0 {
		tx.Rollback()
		return fmt.Errorf("sender not found")
	}

	// Add receiver
	res, err = tx.Exec("UPDATE users SET balance = balance + ? WHERE userId=?", amount, receiverID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error updating receiver balance: %w", err)
	}
	if rows, _ := res.RowsAffected(); rows == 0 {
		tx.Rollback()
		return fmt.Errorf("receiver not found")
	}

	// Insert history
	_, err = tx.Exec("INSERT INTO history (from_userid, to_userid, amount) VALUES (?, ?, ?)",
		senderID, receiverID, amount)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error logging transaction: %w", err)
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return fmt.Errorf("transaction commit failed: %w", err)
	}

	return nil
}

func UpdateUser(userId, username, email, mobile, password string) error {
	if password != "" {
		hashpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("error hashing password: %w", err)
		}
		_, err = DB.Exec("UPDATE users SET username=?, email=?, Mobile_no=?, password=? WHERE userId=?",
			username, email, mobile, string(hashpassword), userId)
		if err != nil {
			return fmt.Errorf("error updating user: %w", err)
		}
	} else {
		_, err := DB.Exec("UPDATE users SET username=?, email=?, Mobile_no=? WHERE userId=?",
			username, email, mobile, userId)
		if err != nil {
			return fmt.Errorf("error updating user: %w", err)
		}
	}
	return nil
}
