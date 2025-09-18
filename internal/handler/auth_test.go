package handler

import (
	"encoding/json"
	"errors"
	"learn-go/internal/config"
	"learn-go/internal/db/mocks"
	"learn-go/internal/db/sqlc"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func TestAuthHandler_Signup(t *testing.T) {
	e := echo.New()

	cfg := &config.Config{
		JWT_SECRET: "test",
	}

	// Test case 1: Valid user creation
	t.Run("Valid user creation", func(t *testing.T) {
		mockDB := mocks.NewMockQuerier(t)

		expectedUser := sqlc.User{
			ID:       1,
			Email:    "test@test.com",
			Password: "hashedpassword",
		}

		// Mock the GetUserByEmail to return an error (user doesn't exist)
		mockDB.On("GetUserByEmail", mock.Anything, "test@test.com").Return(sqlc.User{}, errors.New("user not found"))

		// Mock the CreateUser to return the expected user
		mockDB.On("CreateUser", mock.Anything, mock.AnythingOfType("sqlc.CreateUserParams")).Return(expectedUser, nil)

		// Create the auth handler with the mock
		authHandler := &AuthHandler{
			config: cfg,
			query:  mockDB,
		}

		// Create request body
		requestBody := `{"email":"test@test.com","password":"testpassword"}`
		req := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Call the handler
		err := authHandler.Signup(c)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)

		var response map[string]interface{}
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Account created successfully", response["message"])

		user := response["user"].(map[string]interface{})
		assert.Equal(t, float64(1), user["id"])
		assert.Equal(t, "test@test.com", user["email"])

		mockDB.AssertExpectations(t)
	})

	// Test case 2: User already exists
	t.Run("User already exists", func(t *testing.T) {
		mockDB := mocks.NewMockQuerier(t)

		expectedUser := sqlc.User{
			ID:       1,
			Email:    "test@test.com",
			Password: "hashedpassword",
		}

		// Mock the GetUserByEmail to return the expected user (user exists)
		mockDB.On("GetUserByEmail", mock.Anything, "test@test.com").Return(expectedUser, nil)

		// Create the auth handler with the mock
		authHandler := &AuthHandler{
			config: cfg,
			query:  mockDB,
		}

		// Create request body
		requestBody := `{"email":"test@test.com","password":"testpassword"}`
		req := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Call the handler
		err := authHandler.Signup(c)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, http.StatusConflict, rec.Code)

		var response map[string]interface{}
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "User with this email already exists", response["error"])

		mockDB.AssertExpectations(t)
	})

	// Test case 3: Missing email and password
	t.Run("Invalid JSON", func(t *testing.T) {
		mockDB := mocks.NewMockQuerier(t)

		// Create the auth handler with the mock
		authHandler := &AuthHandler{
			config: cfg,
			query:  mockDB,
		}

		// Create request body
		requestBody := `{}`
		req := httptest.NewRequest(http.MethodPost, "/signup", strings.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Call the handler
		err := authHandler.Signup(c)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		var response map[string]interface{}
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Email and password are required", response["error"])

		mockDB.AssertExpectations(t)
	})
}

func TestAuthHandler_Login(t *testing.T) {
	e := echo.New()

	cfg := &config.Config{
		JWT_SECRET: "test",
	}

	// Test case 1: Valid login
	t.Run("Valid login", func(t *testing.T) {
		mockDB := mocks.NewMockQuerier(t)

		// Hash the password for the test
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("testpassword"), bcrypt.DefaultCost)
		expectedUser := sqlc.User{
			ID:       1,
			Email:    "test@test.com",
			Password: string(hashedPassword),
		}

		// Mock the Login to return the expected user
		mockDB.On("GetUserByEmail", mock.Anything, "test@test.com").Return(expectedUser, nil)

		// Create the auth handler with the mock
		authHandler := &AuthHandler{
			config: cfg,
			query:  mockDB,
		}

		// Create request body
		requestBody := `{"email":"test@test.com","password":"testpassword"}`
		req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Call the handler
		err := authHandler.Login(c)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)

		var response map[string]interface{}
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "login successful", response["message"])
		assert.Equal(t, "ok", response["status"])
		assert.NotEmpty(t, response["token"])

		mockDB.AssertExpectations(t)
	})

	// Test case 2: User doesn't exist
	t.Run("User Doesn't Exist", func(t *testing.T) {
		mockDB := mocks.NewMockQuerier(t)

		// Create the auth handler with the mock
		authHandler := &AuthHandler{
			config: cfg,
			query:  mockDB,
		}

		// Mock the Login to return user not found
		mockDB.On("GetUserByEmail", mock.Anything, "test@test.com").Return(sqlc.User{}, errors.New("user not found"))

		// Create request body
		requestBody := `{"email":"test@test.com","password":"testpassword"}`
		req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Call the handler
		err := authHandler.Login(c)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, rec.Code)

		var response map[string]interface{}
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "User not found", response["error"])

		mockDB.AssertExpectations(t)
	})

	// Test case 3: Missing email and password
	t.Run("Missing email and password", func(t *testing.T) {
		mockDB := mocks.NewMockQuerier(t)

		// Create the auth handler with the mock
		authHandler := &AuthHandler{
			config: cfg,
			query:  mockDB,
		}

		// Create request body
		requestBody := `{}`
		req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Call the handler
		err := authHandler.Login(c)

		// Assertions
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, rec.Code)

		var response map[string]interface{}
		err = json.Unmarshal(rec.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Email and password are required", response["error"])

		mockDB.AssertExpectations(t)
	})
}

func TestAuthHandler_Me(t *testing.T) {
	// e := echo.New()

	// cfg := &config.Config{
	// 	JWT_SECRET: "test",
	// }

	// Test case 1: Valid user
}