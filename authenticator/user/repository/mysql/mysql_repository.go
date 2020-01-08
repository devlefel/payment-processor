package mysql

import (
	"authenticator/user"
	"authenticator/user/models"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type repo struct {
	db *sql.DB
}

//NewRepository creates a new payment.Repository using a mysql connection
func NewRepository(db *sql.DB) (user.Repository, error) {
	var r repo
	r.db = db

	err := r.ping()

	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (r *repo) ping() error {
	return r.db.Ping()
}

//AuthUser receives login credentials and tries to return the id of the user on database
func (r *repo) AuthUser(login models.Login) (int64, *models.Error) {
	var errors models.Error
	stmt, err := r.db.Prepare("SELECT id FROM authenticator.users WHERE username = ? AND pass = SHA1(?)")

	if err != nil {
		errors.Internal = append(errors.Internal, fmt.Errorf("error Getting User"))
		return 0, &errors
	}

	row := stmt.QueryRow(login.Username, login.Password)

	var id int64

	err = row.Scan(&id)

	if err != nil {
		errors.Internal = append(errors.Internal, fmt.Errorf("error Getting User"))
		return 0, &errors
	}

	return id, nil
}

//GetTokens receives userID, and tries to fetch all user tokens from it
func (r *repo) GetTokens(userID int64) ([]string, *models.Error) {
	var errors models.Error
	stmt, err := r.db.Prepare("SELECT token FROM authenticator.tokens WHERE user_id = ?")

	if err != nil {
		errors.Internal = append(errors.Internal, fmt.Errorf("error Getting User Token"))
		return nil, &errors
	}

	rows, err := stmt.Query(userID)

	if err != nil {
		errors.Internal = append(errors.Internal, fmt.Errorf("error Getting User Token"))
		return nil, &errors
	}

	defer rows.Close()

	var tokens []string
	for rows.Next() {
		var token string
		err := rows.Scan(&token)

		if err != nil {
			errors.Internal = append(errors.Internal, fmt.Errorf("error Getting User Token"))
			return nil, &errors
		}

		tokens = append(tokens, token)
	}

	return tokens, nil
}
