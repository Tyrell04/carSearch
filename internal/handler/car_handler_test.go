package handler

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestImportCars(t *testing.T) {
	app := setupTestServer()

	tests := []struct {
		description  string
		apiKey       string
		csvData      string
		expectedCode int
		expectedBody string
	}{
		{
			description:  "successful import with valid API key and data",
			apiKey:       "test-api-key",
			csvData:      "0001,BMW,001,M3\n0002,Mercedes,002,C-Class",
			expectedCode: http.StatusOK,
			expectedBody: `{"message":"Cars imported successfully"}`,
		},
		{
			description:  "unauthorized without API key",
			apiKey:       "",
			csvData:      "0001,BMW,001,M3",
			expectedCode: http.StatusUnauthorized,
			expectedBody: `{"message":"Unauthorized"}`,
		},
		{
			description:  "unauthorized with invalid API key",
			apiKey:       "wrong-api-key",
			csvData:      "0001,BMW,001,M3",
			expectedCode: http.StatusUnauthorized,
			expectedBody: `{"message":"Unauthorized"}`,
		},
		{
			description:  "bad request with missing file",
			apiKey:       "test-api-key",
			csvData:      "",
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"message":"Failed to get CSV file from form data","error":"there is no uploaded file associated with the given key: csv_file"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			// Create a multipart form request
			body := &bytes.Buffer{}
			writer := multipart.NewWriter(body)

			// Only add file if csvData is not empty (for testing missing file case)
			if test.csvData != "" {
				part, err := writer.CreateFormFile("csv_file", "test.csv")
				assert.NoError(t, err)
				_, err = part.Write([]byte(test.csvData))
				assert.NoError(t, err)
			}

			writer.Close()

			req := httptest.NewRequest(http.MethodPost, "/import-cars", body)
			req.Header.Set("Content-Type", writer.FormDataContentType())
			if test.apiKey != "" {
				req.Header.Set("X-API-Key", test.apiKey)
			}

			resp, err := app.Test(req, -1) // -1 for no timeout
			assert.NoError(t, err)
			assert.Equal(t, test.expectedCode, resp.StatusCode)

			respBody, err := ioutil.ReadAll(resp.Body)
			assert.NoError(t, err)

			// For the bad request with missing file case, we need to check if the error message contains the expected text
			if test.description == "bad request with missing file" {
				var result map[string]interface{}
				err = json.Unmarshal(respBody, &result)
				assert.NoError(t, err)
				assert.Contains(t, result["message"], "Failed to get CSV file from form data")
			} else {
				assert.JSONEq(t, test.expectedBody, string(respBody))
			}
		})
	}
}

func TestFindByHSN(t *testing.T) {
	app := setupTestServer()

	// First import some test data
	importTestData(t, app)

	tests := []struct {
		description  string
		hsn          string
		expectedCode int
	}{
		{
			description:  "find cars with existing HSN",
			hsn:          "0001",
			expectedCode: http.StatusOK,
		},
		{
			description:  "no cars found with non-existent HSN",
			hsn:          "9999",
			expectedCode: http.StatusNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/cars/hsn/"+test.hsn, nil)

			resp, err := app.Test(req, -1) // -1 for no timeout
			assert.NoError(t, err)
			assert.Equal(t, test.expectedCode, resp.StatusCode)

			if test.expectedCode == http.StatusOK {
				body, err := ioutil.ReadAll(resp.Body)
				assert.NoError(t, err)

				var result map[string]interface{}
				err = json.Unmarshal(body, &result)
				assert.NoError(t, err)
				assert.Contains(t, result, "cars")
			}
		})
	}
}

func TestFindByHSNAndTSN(t *testing.T) {
	app := setupTestServer()

	// First import some test data
	importTestData(t, app)

	tests := []struct {
		description  string
		hsn          string
		tsn          string
		expectedCode int
	}{
		{
			description:  "find car with existing HSN and TSN",
			hsn:          "0001",
			tsn:          "001",
			expectedCode: http.StatusOK,
		},
		{
			description:  "car not found with non-existent HSN",
			hsn:          "9999",
			tsn:          "001",
			expectedCode: http.StatusNotFound,
		},
		{
			description:  "car not found with non-existent TSN",
			hsn:          "0001",
			tsn:          "999",
			expectedCode: http.StatusNotFound,
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/cars/hsn/"+test.hsn+"/tsn/"+test.tsn, nil)

			resp, err := app.Test(req, -1) // -1 for no timeout
			assert.NoError(t, err)
			assert.Equal(t, test.expectedCode, resp.StatusCode)

			if test.expectedCode == http.StatusOK {
				body, err := ioutil.ReadAll(resp.Body)
				assert.NoError(t, err)

				var result map[string]interface{}
				err = json.Unmarshal(body, &result)
				assert.NoError(t, err)
				assert.Contains(t, result, "car")
			}
		})
	}
}

// NEW TEST: Test the SearchCars endpoint with query parameters
func TestSearchCars(t *testing.T) {
	app := setupTestServer()

	// First import some test data
	importTestData(t, app)

	tests := []struct {
		description  string
		queryParams  string
		expectedCode int
		expectedKey  string // "producer" for HSN-only, "car" for HSN+TSN
	}{
		{
			description:  "search by HSN only returns producer",
			queryParams:  "hsn=0001",
			expectedCode: http.StatusOK,
			expectedKey:  "producer",
		},
		{
			description:  "search by HSN and TSN returns specific car",
			queryParams:  "hsn=0001&tsn=001",
			expectedCode: http.StatusOK,
			expectedKey:  "car",
		},
		{
			description:  "search with missing HSN parameter",
			queryParams:  "tsn=001",
			expectedCode: http.StatusBadRequest,
			expectedKey:  "",
		},
		{
			description:  "search with non-existent HSN",
			queryParams:  "hsn=9999",
			expectedCode: http.StatusNotFound,
			expectedKey:  "",
		},
		{
			description:  "search with non-existent TSN",
			queryParams:  "hsn=0001&tsn=999",
			expectedCode: http.StatusNotFound,
			expectedKey:  "",
		},
	}

	for _, test := range tests {
		t.Run(test.description, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/cars/search?"+test.queryParams, nil)

			resp, err := app.Test(req, -1) // -1 for no timeout
			assert.NoError(t, err)
			assert.Equal(t, test.expectedCode, resp.StatusCode)

			if test.expectedCode == http.StatusOK {
				body, err := ioutil.ReadAll(resp.Body)
				assert.NoError(t, err)

				var result map[string]interface{}
				err = json.Unmarshal(body, &result)
				assert.NoError(t, err)
				assert.Contains(t, result, test.expectedKey)

				// Additional validation based on expected response type
				if test.expectedKey == "producer" {
					producer := result["producer"].(map[string]interface{})
					assert.Equal(t, "0001", producer["HSN"])
					assert.Equal(t, "BMW", producer["Name"])
				} else if test.expectedKey == "car" {
					car := result["car"].(map[string]interface{})
					assert.Equal(t, "0001", car["HSN"])
					assert.Equal(t, "001", car["TSN"])
					assert.Equal(t, "M3", car["Name"])
				}
			} else if test.expectedCode == http.StatusBadRequest {
				body, err := ioutil.ReadAll(resp.Body)
				assert.NoError(t, err)

				var result map[string]interface{}
				err = json.Unmarshal(body, &result)
				assert.NoError(t, err)
				assert.Contains(t, result["message"], "HSN query parameter is required")
			}
		})
	}
}

// Helper function to import test data
func importTestData(t *testing.T, app *fiber.App) {
	// Create CSV data with HSN, Producer Name, TSN, Car Name format
	csvData := strings.TrimSpace(`
0001,BMW,001,M3
0001,BMW,002,M5
0002,Mercedes,001,C-Class
0002,Mercedes,002,E-Class
	`)

	// Create a multipart form request
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add the CSV file
	part, err := writer.CreateFormFile("csv_file", "test.csv")
	assert.NoError(t, err)
	_, err = part.Write([]byte(csvData))
	assert.NoError(t, err)

	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/import-cars", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("X-API-Key", "test-api-key")

	resp, err := app.Test(req, -1)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
