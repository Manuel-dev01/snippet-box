package mysql

import (
	"database/sql"

	"github.com/Manuel-dev01/snippet-box/pkg/models"
)

type UserModel struct {
	DB *sql.DB
}

//Add a new record to the users table
func (m *UserModel) Insert(name, email, password string) error {
	return nil
}

// Check whether a users exists, return the relevant id of the user
func (m *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}