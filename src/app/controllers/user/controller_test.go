package user

import (
	"bank-app/src/app/controllers/user/models"
	"bytes"
	"net/http/httptest"
	"testing"
)

type mockUserService struct{}

func (t *mockUserService) CreateUser(obj *models.User) error {
	return nil
}

func (t *mockUserService) GetUserBalance(id int) (float64, error) {
	return 0, nil
}

func (t *mockUserService) UpdateBalance(id int, amount float64) error {
	return nil
}

func TestCreateUser(t *testing.T) {
	s := &mockUserService{}
	controller := NewController(s)

	testTable := []struct {
		testStatus       string
		input            string
		expectedHeader   string
		expectedResponse string
		expectedStatus   int
	}{
		{
			testStatus:       "OK",
			input:            `{"name":"test", "email":"test@m", "password":"123"}`,
			expectedResponse: `{"success": True}`,
			expectedStatus:   201,
			expectedHeader:   "application/json",
		},
	}

	for i, testCase := range testTable {
		t.Run(testCase.testStatus, func(t *testing.T) {
			t.Log(i)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/users", bytes.NewBufferString(testCase.input))

			controller.CreateUser(w, req)

			if w.Code != testCase.expectedStatus {
				t.Errorf("Expected status %d; got %d", testCase.expectedStatus, w.Code)
			}

			if w.Body.String() != testCase.expectedResponse {
				t.Errorf("Expected response body '%s'; got '%s'", testCase.expectedResponse, w.Body.String())
			}

			contentType := w.Header().Get("Content-Type")
			if contentType != testCase.expectedHeader {
				t.Errorf("Expected content type '%s'; got '%s'", testCase.expectedHeader, contentType)
			}
		})
	}
}
