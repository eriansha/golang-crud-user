package repositories

import (
	"database/sql"
	models "golangcrud/models/user"
)

// Delete single user through database by given ID
func DeleteUser(db *sql.DB, id int) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

// Update user data through database
func UpdateUser(db *sql.DB, id int, name, email string) error {
	query := "UPDATE users SET name = ?, email = ? where id = ?"

	_, err := db.Exec(query, name, email, id)
	if err != nil {
		return err
	}

	return nil
}

// Create new user to database
func CreateUser(db *sql.DB, name, email string) error {
	query := "INSERT INTO users(name, email) values (?, ?)"

	_, err := db.Exec(query, name, email)
	if err != nil {
		return err
	}

	return nil
}

// Get single user data from database
func GetUser(db *sql.DB, id int) (*models.User, error) {
	query := "SELECT * FROM users WHERE id = ?"
	row := db.QueryRow(query, id)

	user := &models.User{}
	err := row.Scan(&user.Id, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
