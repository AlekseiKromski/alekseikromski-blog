package dbstore

import (
	"alekseikromski.com/blog/api/storage/models"
	"fmt"
	"log"
)

func (db *DbConnection) GetUser(email string) (*models.User, error) {
	user := models.CreateUser()
	query := models.GetUserByID(email)

	log.Printf("[DBSTORE] running query: %s", query)

	rows, err := db.Connection.Query(query)
	if err != nil {
		return nil, fmt.Errorf("cannot make request, server error")
	}

	for rows.Next() {
		var scanError error

		scanError = rows.Scan(
			&user.ID,
			&user.Email,
			&user.Password,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt,
		)

		if scanError != nil {
			log.Printf("troubles during scanning: %v", err)
			return nil, fmt.Errorf("cannot find user")
		}
	}

	return user, nil
}
