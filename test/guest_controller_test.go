package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/agisnur24/booking_hotel_system.git/model/domain"
	"strconv"

	"github.com/agisnur24/booking_hotel_system.git/app/routers"
	"github.com/agisnur24/booking_hotel_system.git/controller"
	"github.com/agisnur24/booking_hotel_system.git/helper"
	"github.com/agisnur24/booking_hotel_system.git/middleware"
	"github.com/agisnur24/booking_hotel_system.git/repository"
	"github.com/agisnur24/booking_hotel_system.git/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
	"net/http/httptest"

	"strings"
	"testing"
	"time"
)

func setupTestGuestDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/booking_management_system")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupGuestRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	guestRepository := repository.NewGuestRepository()
	guestService := service.NewGuestService(guestRepository, db, validate)
	guestController := controller.NewGuestController(guestService)

	router := routers.NewGuestRouter(guestController)

	return middleware.NewAuthMiddleware(router)
}

func truncateGuest(db *sql.DB) {
	db.Exec("TRUNCATE guests")
}

func TestCreateGuestSuccess(t *testing.T) {
	db := setupTestGuestDB()
	truncateGuest(db)
	router := setupGuestRouter(db)

	requestBody := strings.NewReader(`{"name" : "pokaliswali","address":"Belum Jadi","phone_number":"089876545","email":"imey@gmail.com"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/guests", requestBody)
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
	assert.Equal(t, "pokaliswali", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, "Belum Jadi", responseBody["data"].(map[string]interface{})["address"])
	assert.Equal(t, "089876545", responseBody["data"].(map[string]interface{})["phone_number"])
	assert.Equal(t, "imey@gmail.com", responseBody["data"].(map[string]interface{})["email"])
}

func TestCreateGuestFailed(t *testing.T) {
	db := setupTestGuestDB()
	truncateGuest(db)
	router := setupGuestRouter(db)

	requestBody := strings.NewReader(`{"name" : ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/guests", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestUpdateGuestSuccess(t *testing.T) {
	db := setupTestGuestDB()
	truncateGuest(db)

	tx, _ := db.Begin()
	guestRepository := repository.NewGuestRepository()
	guest := guestRepository.Create(context.Background(), tx, domain.Guest{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupGuestRouter(db)

	requestBody := strings.NewReader(`{"name" : "pokaliswali","address":"Belum Jadi","phone_number":"089876545","email":"imey@gmail.com"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/guests/"+strconv.Itoa(guest.Id), requestBody)
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
	assert.Equal(t, guest.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	//assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"]) data asli

	assert.Equal(t, "pokaliswali", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, "Belum Jadi", responseBody["data"].(map[string]interface{})["address"])
	assert.Equal(t, "089876545", responseBody["data"].(map[string]interface{})["phone_number"])
	assert.Equal(t, "imey@gmail.com", responseBody["data"].(map[string]interface{})["email"])
}

func TestUpdateGuestFailed(t *testing.T) {
	db := setupTestGuestDB()
	truncateGuest(db)

	tx, _ := db.Begin()
	guestRepository := repository.NewGuestRepository()
	guest := guestRepository.Create(context.Background(), tx, domain.Guest{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupGuestRouter(db)

	requestBody := strings.NewReader(`{"name" : ""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/guests/"+strconv.Itoa(guest.Id), requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestGetGuestSuccess(t *testing.T) {
	db := setupTestGuestDB()
	truncateGuest(db)

	tx, _ := db.Begin()
	guestRepository := repository.NewGuestRepository()
	guest := guestRepository.Create(context.Background(), tx, domain.Guest{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupGuestRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/guests/"+strconv.Itoa(guest.Id), nil)
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
	assert.Equal(t, guest.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, guest.Name, responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, guest.Address, responseBody["data"].(map[string]interface{})["address"])
	assert.Equal(t, guest.Phone_Number, responseBody["data"].(map[string]interface{})["phone_number"])
	assert.Equal(t, guest.Email, responseBody["data"].(map[string]interface{})["email"])
}

func TestGetGuestFailed(t *testing.T) {
	db := setupTestGuestDB()
	truncateGuest(db)
	router := setupGuestRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/guests/404", nil)
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestDeleteGuestSuccess(t *testing.T) {
	db := setupTestGuestDB()
	truncateGuest(db)

	tx, _ := db.Begin()
	guestRepository := repository.NewGuestRepository()
	guest := guestRepository.Create(context.Background(), tx, domain.Guest{
		Name: "Gadget",
	})
	tx.Commit()

	router := setupGuestRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/guests/"+strconv.Itoa(guest.Id), nil)
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
}

func TestDeleteGuestFailed(t *testing.T) {
	db := setupTestGuestDB()
	truncateGuest(db)
	router := setupGuestRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/guests/404", nil)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestListGuestsSuccess(t *testing.T) {
	db := setupTestGuestDB()
	truncateGuest(db)

	tx, _ := db.Begin()
	guestRepository := repository.NewGuestRepository()
	guest1 := guestRepository.Create(context.Background(), tx, domain.Guest{
		Name: "Gadget",
	})
	guest2 := guestRepository.Create(context.Background(), tx, domain.Guest{
		Name: "Computer",
	})
	tx.Commit()

	router := setupGuestRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/guests", nil)
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

	fmt.Println(responseBody)

	var guests = responseBody["data"].([]interface{})

	guestResponse1 := guests[0].(map[string]interface{})
	guestResponse2 := guests[1].(map[string]interface{})

	assert.Equal(t, guest1.Id, int(guestResponse1["id"].(float64)))
	assert.Equal(t, guest1.Name, guestResponse1["name"])
	assert.Equal(t, guest1.Address, guestResponse1["address"])
	assert.Equal(t, guest1.Phone_Number, guestResponse1["phone_number"])
	assert.Equal(t, guest1.Email, guestResponse1["email"])

	assert.Equal(t, guest2.Id, int(guestResponse2["id"].(float64)))
	assert.Equal(t, guest2.Name, guestResponse2["name"])
	assert.Equal(t, guest2.Address, guestResponse2["address"])
	assert.Equal(t, guest2.Phone_Number, guestResponse2["phone_number"])
	assert.Equal(t, guest2.Email, guestResponse2["email"])
}

func TestUnauthorized(t *testing.T) {
	db := setupTestGuestDB()
	truncateGuest(db)
	router := setupGuestRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/guests", nil)
	request.Header.Add("X-API-Key", "SALAH")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, "UNAUTHORIZED", responseBody["status"])
}
