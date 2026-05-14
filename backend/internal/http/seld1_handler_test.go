package httpserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestListSELD1Rules(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := NewRouter()
	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/seld1/rules", nil)

	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}

	var body SELD1RulesResponse
	if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
		t.Fatalf("expected valid JSON response: %v", err)
	}

	if len(body.Rules) != 34 {
		t.Fatalf("expected 34 rules, got %d", len(body.Rules))
	}
}

func TestInferSELD1Irrigation(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := NewRouter()
	response := httptest.NewRecorder()
	request := httptest.NewRequest(
		http.MethodPost,
		"/seld1/infer",
		bytes.NewReader([]byte(`{
			"soil_humidity": 8,
			"ambient_temperature": 30,
			"time_of_day": 1000
		}`)),
	)
	request.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}

	var body SELD1InferenceResponse
	if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
		t.Fatalf("expected valid JSON response: %v", err)
	}

	if body.OutputVariableKey != "valve_opening" {
		t.Fatalf("expected valve_opening output key, got %q", body.OutputVariableKey)
	}

	if len(body.ActivatedRules) == 0 {
		t.Fatal("expected activated rules")
	}

	if len(body.AggregatedOutput) == 0 {
		t.Fatal("expected aggregated output points")
	}
}

func TestInferSELD1RejectsOutOfRangeInput(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := NewRouter()
	response := httptest.NewRecorder()
	request := httptest.NewRequest(
		http.MethodPost,
		"/seld1/infer",
		bytes.NewReader([]byte(`{
			"soil_humidity": 101,
			"ambient_temperature": 30,
			"time_of_day": 1000
		}`)),
	)
	request.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(response, request)

	if response.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", response.Code)
	}
}
