package models

import (
	"fmt"
	"time"
)

type User struct {
	ID       int    `json:"ID"`
	Email    string `json:"email"`
	Password string `json:"password"`
	*Timestamp
	*SoftDelete
}

func CreateUser() *User {
	user := &User{
		Email:      "",
		Password:   "",
		Timestamp:  &Timestamp{},
		SoftDelete: &SoftDelete{},
	}

	user.SetTimestamp()
	return user
}

func (u *User) Validate() bool {
	if len(u.Email) != 0 {
		return true
	}
	return false
}

func (u *User) SetTimestamp() {
	u.UpdatedAt = time.Now().Format(time.RFC3339)
	if len(u.CreatedAt) == 0 {
		u.CreatedAt = time.Now().Format(time.RFC3339)
	}
}

func (u *User) Soft() {
	softDeleteTs := time.Now().Format(time.RFC3339)
	u.DeletedAt = &softDeleteTs
}

func (u *User) Undo() {
	softDeleteTs := time.Now().Format(time.RFC3339)
	u.DeletedAt = &softDeleteTs
}

func (u *User) TableCreate() *TableCreation {
	return &TableCreation{
		Sql: `
			create table users
			(
				ID          serial
					constraint users_pk
						primary key,
				"Email"  varchar(60)      not null,
				"Password"  text     not null,
				"CreatedAt" timestamp not null,
				"UpdatedAt" timestamp not null,
				"DeletedAt" timestamp
			);
		`,
		Dependencies: []string{},
	}
}

func (u *User) CreateRecord() string {
	return fmt.Sprintf(`INSERT INTO users ("Email", "Password","CreatedAt", "UpdatedAt", "DeletedAt") VALUES ('%s', %d,'%s','%s', NULL)`, u.Email, u.Password, u.CreatedAt, u.UpdatedAt)
}

func GetUserByID(email string) string {
	return fmt.Sprintf(`SELECT * FROM users WHERE users."Email" = '%s' AND users."DeletedAt" is NULL`, email)
}

func UpdateUser(user *User) string {
	if user.DeletedAt == nil {
		return fmt.Sprintf(
			`UPDATE users SET Email = '%s', Password = %s, "UpdatedAt" = '%s'WHERE users."ID" = %d`,
			user.Email, user.Password, user.UpdatedAt, user.ID,
		)
	}
	return fmt.Sprintf(
		`UPDATE users SET Email = '%s', Password = %s, "UpdatedAt" = '%s', "DeletedAt" = '%s' WHERE users."ID" = %d`,
		user.Email, user.Password, user.UpdatedAt, *user.DeletedAt, user.ID,
	)
}
