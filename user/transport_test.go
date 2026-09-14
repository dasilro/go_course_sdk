package user_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"testing"

	userSdk "github.com/dasilro/go_course_sdk/user"
	"github.com/dasilro/gocourse_domain/domain"
	c "github.com/ncostamagna/go_http_client/client"
)

var header http.Header
var sdk userSdk.Transport

func TestMain(m *testing.M) {
	header = http.Header{}
	header.Set("Content-Type", "application/json")
	sdk = userSdk.NewHttpClient("base-url", "")
	os.Exit(m.Run())
}

func TestGet_Response404Error(t *testing.T) {
	expectedErr := userSdk.ErrNotFound{Message: "course '1' not found"}

	err := c.AddMockups(&c.Mock{
		HTTPMethod:   http.MethodGet,
		RespHeaders:  header,
		URL:          "base-url/users/1",
		RespHTTPCode: 404,
		RespBody: fmt.Sprintf(`{
			"status": 404,
			"message": "%s"
		}`, expectedErr.Error()),
	})

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	user, err := sdk.Get("1")
	if !errors.Is(err, expectedErr) {
		t.Errorf("expected nil, got %v", err)
	}

	if user != nil {
		t.Errorf("expected nil, got %v", user)
	}
}

func TestGet_Response500Error(t *testing.T) {
	expectedErr := errors.New("internal server error")

	err := c.AddMockups(&c.Mock{
		HTTPMethod:   http.MethodGet,
		RespHeaders:  header,
		URL:          "base-url/users/1",
		RespHTTPCode: 500,
		RespBody: fmt.Sprintf(`{
			"status": 500,
			"message": "%s"
		}`, expectedErr.Error()),
	})

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	user, err := sdk.Get("1")
	if err == nil || err.Error() != expectedErr.Error() {
		t.Errorf("expected err %v, got %v", expectedErr, err)
	}

	if user != nil {
		t.Errorf("expected nil, got %v", user)
	}
}

func TestGet_ResponseMarshalError(t *testing.T) {
	expectedErr := errors.New("unexpected end of JSON input")

	err := c.AddMockups(&c.Mock{
		HTTPMethod:   http.MethodGet,
		RespHeaders:  header,
		URL:          "base-url/users/1",
		RespHTTPCode: 200,
		RespBody:     `{`,
	})

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	user, err := sdk.Get("1")
	if err == nil || err.Error() != expectedErr.Error() {
		t.Errorf("expected err %v, got %v", expectedErr, err)
	}

	if user != nil {
		t.Errorf("expected nil, got %v", user)
	}
}

func TestGet_ClientErr(t *testing.T) {
	expectedErr := errors.New("client error")

	err := c.AddMockups(&c.Mock{
		HTTPMethod:   http.MethodGet,
		RespHeaders:  header,
		URL:          "base-url/users/1",
		RespHTTPCode: 400,
		Err:          expectedErr,
	})

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	user, err := sdk.Get("1")
	if !errors.Is(err, expectedErr) {
		t.Errorf("expected err %v, got %v", expectedErr, err)
	}

	if user != nil {
		t.Errorf("expected nil, got %v", user)
	}
}

func TestGet_RespondsSuccess(t *testing.T) {
	expectedUser := &domain.User{
		ID:        "1",
		FirstName: "User 1 FirstName",
		LastName:  "User 1 LastName",
		Email:     "User 1 Email",
	}

	expectedUserJson, err := json.Marshal(expectedUser)
	if err != nil {
		t.Errorf("error marshalling expected course %v", err)
	}

	err = c.AddMockups(&c.Mock{
		HTTPMethod:   http.MethodGet,
		RespHeaders:  header,
		URL:          "base-url/users/1",
		RespHTTPCode: 200,
		RespBody: fmt.Sprintf(`{
								"status": 200,
								"message": "success",
								"data": %s
							}`, expectedUserJson),
	})

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	user, err := sdk.Get("1")
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	if user == nil {
		t.Errorf("expected course, got nil")
	}

	if user.ID != expectedUser.ID {
		t.Errorf("expected id %v, got %v", expectedUser.ID, user.ID)
	}

	if user.FirstName != expectedUser.FirstName {
		t.Errorf("expected name %v, got %v", expectedUser.FirstName, user.FirstName)
	}

	if user.LastName != expectedUser.LastName {
		t.Errorf("expected name %v, got %v", expectedUser.LastName, user.LastName)
	}

	if user.Email != expectedUser.Email {
		t.Errorf("expected name %v, got %v", expectedUser.Email, user.Email)
	}

}
