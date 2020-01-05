package mysql

import (
	"database/sql"
	"fmt"
	"processor/payment"
	"processor/payment/models"
)

type repo struct {
	db *sql.DB
}

//NewRepository creates a new payment.Repository using a mysql connection
func NewRepository(db *sql.DB) (payment.Repository, error) {
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

func (r *repo) GetCardSensitiveData(token string, errors *models.Error) models.CardSensitiveData {
	stmt, err := r.db.Prepare("SELECT number, cvv FROM processors.cards WHERE token = MD5(?)")

	if err != nil {
		errors.Internal = append(errors.Internal, err)
		return models.CardSensitiveData{}
	}

	defer stmt.Close()

	res := stmt.QueryRow(token)
	var data models.CardSensitiveData
	err = res.Scan(&data.Number, &data.CVV)

	if err != nil {
		if err == sql.ErrNoRows {
			errors.Validation = append(errors.Validation, fmt.Errorf("invalid Token"))
			return models.CardSensitiveData{}
		}

		errors.Internal = append(errors.Internal, err)
		return models.CardSensitiveData{}
	}

	return data
}

func (r *repo) GetAcquirerURL(id int64, errors *models.Error) string {
	stmt, err := r.db.Prepare("SELECT url FROM processors.acquirers WHERE id = ?")

	if err != nil {
		errors.Internal = append(errors.Internal, err)
		return "nil"
	}

	defer stmt.Close()

	res := stmt.QueryRow(id)
	var url string
	err = res.Scan(&url)

	if err != nil {
		if err == sql.ErrNoRows {
			errors.Validation = append(errors.Validation, fmt.Errorf("no Acquire found"))
			return ""
		}

		errors.Internal = append(errors.Internal, err)
		return ""
	}

	return url
}
