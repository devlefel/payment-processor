package mysql

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"os"
	"processor/payment/models"
	"testing"
)

var db *sql.DB
var mock sqlmock.Sqlmock

func TestMain(m *testing.M) {
	var err error
	db, mock, err = sqlmock.New()

	if err != nil {
		fmt.Println("expected no error, but got:", err)
		return
	}

	code := m.Run()

	db.Close()
	os.Exit(code)
}
func TestGetCardSensitiveData(t *testing.T) {
	repo, err := NewRepository(db)

	if err != nil {
		assert.FailNowf(t, "error creating repo: %s", err.Error())
	}

	type test struct {
		Token          string
		ExpectedOutput *models.CardSensitiveData
		ExpectedError  models.Error
	}

	var tests = make(map[string]test)

	tests["ReturnValue"] = test{
		Token: "blaoaslkdajeifaepfae.asiudh123123785.3hauk.498a84",
		ExpectedOutput: &models.CardSensitiveData{
			Number: 426501269876325,
			CVV:    123,
		},
		ExpectedError: models.Error{
			Validation: nil,
			Internal:   nil,
		},
	}

	for title, test := range tests {
		t.Run(title, func(t *testing.T) {

			rows := sqlmock.NewRows([]string{"number", "cvv"}).AddRow(test.ExpectedOutput.Number, test.ExpectedOutput.CVV)
			mock.ExpectPrepare("SELECT number, cvv FROM cards.cards").ExpectQuery().WithArgs(test.Token).WillReturnRows(rows)
			oldError := test.ExpectedError
			ret := repo.GetCardSensitiveData(test.InputName, &test.ExpectedError)

			if assert.Equal(t, test.ExpectedError, oldError, "errors we expected: %s, but not nil has return: %s", test.ExpectedError, err) {
				assert.Equal(t, test.ExpectedOutput, ret, "We expected return: %#v but got: %#v", test.ExpectedOutput, ret)
			}
		})
	}
}
