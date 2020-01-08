package tests

import (
	"authenticator/user/models"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

//TestLogin its a integration test for Login
func TestLogin(t *testing.T) {
	type test struct {
		Input          models.Login
		ExpectedStatus uint
		ExpectedTokens []string
		ExpectedErrors error
	}

	testSuccess := test{
		Input: models.Login{
			Username: "admin",
			Password: "@#$RF@!718",
		},
		ExpectedStatus: http.StatusOK,
		ExpectedTokens: []string{
			"eb2cd9bf2054de63b62330b3ae319e517f195afcc0ed19e984910f833d7f95a2",
			"ea02bc63ad9c0abae0176248be28021d0417b5f81a418e160e4c16e0935fe49b",
			"9bd1fe5f411db3e3be3ae045a3b13e84e3893c3f81806381e08ba29ce456be0e",
		},
		ExpectedErrors: nil,
	}

	jr, err := json.Marshal(testSuccess.Input)

	assert.NoError(t, err)

	b := bytes.NewReader(jr)

	response, err := http.Post("http://localhost:8080/user/login", "application/json", b)

	defer response.Body.Close()

	if assert.Equal(t, testSuccess.ExpectedErrors, err) {
		if assert.Equal(t, testSuccess.ExpectedStatus, response.StatusCode) {
			body, err := ioutil.ReadAll(response.Body)
			assert.NoError(t, err)
			resp := struct {
				Token []string
			}{}

			err = json.Unmarshal(body, &resp)
			assert.NoError(t, err)

			assert.Equal(t, testSuccess.ExpectedTokens, resp.Token)
		}

	}
}
