package test

import (
	"database/sql"
	"encoding/json"
	"github.com/agisnur24/booking_hotel_system.git/app/routers"
	"github.com/agisnur24/booking_hotel_system.git/controller"
	"github.com/agisnur24/booking_hotel_system.git/helper"
	"github.com/agisnur24/booking_hotel_system.git/middleware"
	"github.com/agisnur24/booking_hotel_system.git/repository"
	"github.com/agisnur24/booking_hotel_system.git/service"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func setupTestDiscountDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/booking_management_system")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupDiscountRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	discountRepository := repository.NewDiscountRepository()
	discountService := service.NewDiscountService(discountRepository, db, validate)
	discountController := controller.NewDiscountController(discountService)

	router := routers.NewDiscountRouter(discountController)

	return middleware.NewAuthMiddleware(router)
}

func truncateDiscount(db *sql.DB) {
	db.Exec("TRUNCATE guests")
}

func TestCreateDiscountSuccess(t *testing.T) {
	db := setupTestDiscountDB()
	truncateDiscount(db)
	router := setupDiscountRouter(db)

	requestBody := strings.NewReader(`{"name" : "pokaliswali","address":"Belum Jadi","phone_number":"089876545","email":"imey@gmail.com"}`) // belum memasukan data asli
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/discounts", requestBody)
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

/*func TestCreateDiscountFailed(t *testing.T) {
	db := setupTestDiscountDB()
	truncateDiscount(db)
	router := setupDiscountRouter(db)

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

func TestUpdateDiscountSuccess(t *testing.T) {
	db := setupTestDiscountDB()
	truncateDiscount(db)

	tx, _ := db.Begin()
	discountRepository := repository.NewDiscountRepository()
	discount := discountRepository.Create(context.Background(), tx, domain.Discount{
		Rate:        2, // int
		Status:      "",
		RequestDate: "",
	})
	tx.Commit()

	router := setupDiscountRouter(db)

	requestBody := strings.NewReader(`{"name" : "pokaliswali","address":"Belum Jadi","phone_number":"089876545","email":"imey@gmail.com"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/discounts/"+strconv.Itoa(discount.Id), requestBody)
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
	assert.Equal(t, discount.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))

	assert.Equal(t, "pokaliswali", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, "Belum Jadi", responseBody["data"].(map[string]interface{})["address"])
	assert.Equal(t, "089876545", responseBody["data"].(map[string]interface{})["phone_number"])
	assert.Equal(t, "imey@gmail.com", responseBody["data"].(map[string]interface{})["email"])
}

func TestUpdateDiscountFailed(t *testing.T) {
	db := setupTestDiscountDB()
	truncateDiscount(db)

	tx, _ := db.Begin()
	discountRepository := repository.NewDiscountRepository()
	discount := discountRepository.Create(context.Background(), tx, domain.Discount{
		Rate:        2, // int
		Status:      "",
		RequestDate: "",
	})
	tx.Commit()

	router := setupDiscountRouter(db)

	requestBody := strings.NewReader(`{"name" : ""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/discounts/"+strconv.Itoa(discount.Id), requestBody)
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

func TestGetDiscountSuccess(t *testing.T) {
	db := setupTestDiscountDB()
	truncateDiscount(db)

	tx, _ := db.Begin()
	discountRepository := repository.NewDiscountRepository()
	discount := discountRepository.Create(context.Background(), tx, domain.Discount{
		Rate:        2, // int
		Status:      "",
		RequestDate: "",
	})
	tx.Commit()

	router := setupDiscountRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/discounts/"+strconv.Itoa(discount.Id), nil)
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
	assert.Equal(t, discount.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, discount.Name, responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, discount.Address, responseBody["data"].(map[string]interface{})["address"])
	assert.Equal(t, discount.Phone_Number, responseBody["data"].(map[string]interface{})["phone_number"])
	assert.Equal(t, discount.Email, responseBody["data"].(map[string]interface{})["email"])
}

func TestGetDiscountFailed(t *testing.T) {
	db := setupTestDiscountDB()
	truncateDiscount(db)
	router := setupDiscountRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/discounts/404", nil)
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

func TestDeleteDiscountSuccess(t *testing.T) {
	db := setupTestDiscountDB()
	truncateDiscount(db)

	tx, _ := db.Begin()
	discountRepository := repository.NewDiscountRepository()
	discount := discountRepository.Create(context.Background(), tx, domain.Discount{
		Rate:        2, // int
		Status:      "",
		RequestDate: "",
	})
	tx.Commit()

	router := setupGuestRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/discounts/"+strconv.Itoa(discount.Id), nil)
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

func TestDeleteDiscountFailed(t *testing.T) {
	db := setupTestDiscountDB()
	truncateDiscount(db)
	router := setupDiscountRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/discounts/404", nil)
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

func TestListDiscountsSuccess(t *testing.T) {
	db := setupTestDiscountDB()
	truncateDiscount(db)

	tx, _ := db.Begin()
	discountRepository := repository.NewDiscountRepository()
	discount1 := discountRepository.Create(context.Background(), tx, domain.Discount{
		Name: "Gadget",
	})
	discount2 := discountRepository.Create(context.Background(), tx, domain.Discount{
		Name: "Computer",
	})
	tx.Commit()

	router := setupDiscountRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/discounts", nil)
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

	var discounts = responseBody["data"].([]interface{})

	discountResponse1 := discounts[0].(map[string]interface{})
	discountResponse2 := discounts[1].(map[string]interface{})

	assert.Equal(t, discount1.Id, int(discountResponse1["id"].(float64)))
	assert.Equal(t, discount1.Name, discountResponse1["name"])
	assert.Equal(t, discount1.Address, discountResponse1["address"])
	assert.Equal(t, discount1.Phone_Number, discountResponse1["phone_number"])
	assert.Equal(t, discount1.Email, discountResponse1["email"])

	assert.Equal(t, discount2.Id, int(discountResponse2["id"].(float64)))
	assert.Equal(t, discount2.Name, discountResponse2["name"])
	assert.Equal(t, discount2.Address, discountResponse2["address"])
	assert.Equal(t, discount2.Phone_Number, discountResponse2["phone_number"])
	assert.Equal(t, discount2.Email, discountResponse2["email"])
}

func TestUnauthorizedDiscount(t *testing.T) {
	db := setupTestDiscountDB()
	truncateDiscount(db)
	router := setupDiscountRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/discounts", nil)
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
*/
