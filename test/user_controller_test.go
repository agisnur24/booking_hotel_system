package test

/*import (
	"database/sql"
	"encoding/json"
	"github.com/agisnur24/booking_hotel_system.git/app/routers"
	"github.com/agisnur24/booking_hotel_system.git/controller"
	"github.com/agisnur24/booking_hotel_system.git/helper"
	"github.com/agisnur24/booking_hotel_system.git/middleware"
	"github.com/agisnur24/booking_hotel_system.git/repository"
	"github.com/agisnur24/booking_hotel_system.git/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func setupTestUserDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/booking_system_management")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupUserRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)
	router := routers.NewUserRouter(userController)

	return middleware.NewAuthMiddleware(router)
}

func truncateUser(db *sql.DB) {
	db.Exec("TRUNCATE users")
}

func TestCreateUserSuccess(t *testing.T) {
	db := setupTestUserDB()
	truncateUser(db)
	router := setupUserRouter(db)

	requestBody := strings.NewReader(`{"name" : "Arie", "email" : "arieafr123@gmail.com", "password" : "aowkoakwoa", "role_id" : 1}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3306/api/users", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Arie", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, "arieafr123@gmail.com", responseBody["data"].(map[string]interface{})["email"])
	assert.Equal(t, "aowkoakwoa", responseBody["data"].(map[string]interface{})["password"])
	assert.Equal(t, 1, responseBody["data"].(map[string]interface{})["role_id"])
}*/
