package httpserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestListSELD2Rules(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := NewRouter()
	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/seld2/rules", nil)

	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}

	var body SELD2RulesResponse
	if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
		t.Fatalf("expected valid JSON response: %v", err)
	}

	if len(body.Rules) != 8 {
		t.Fatalf("expected 8 rules, got %d", len(body.Rules))
	}
}

func TestInferSELD2Ventilation(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := NewRouter()
	response := httptest.NewRecorder()
	request := httptest.NewRequest(
		http.MethodPost,
		"/seld2/infer",
		bytes.NewReader([]byte(`{
			"plant_temperature": 42,
			"gas_concentration": 70
		}`)),
	)
	request.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}

	var body SELD2InferenceResponse
	if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
		t.Fatalf("expected valid JSON response: %v", err)
	}

	if body.OutputVariableKey != "motor_speed" {
		t.Fatalf("expected motor_speed output key, got %q", body.OutputVariableKey)
	}

	if len(body.ActivatedRules) == 0 {
		t.Fatal("expected activated rules")
	}

	if len(body.AggregatedOutput) == 0 {
		t.Fatal("expected aggregated output points")
	}
}

func TestInferSELD2RejectsOutOfRangeInput(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := NewRouter()
	response := httptest.NewRecorder()
	request := httptest.NewRequest(
		http.MethodPost,
		"/seld2/infer",
		bytes.NewReader([]byte(`{
			"plant_temperature": 61,
			"gas_concentration": 70
		}`)),
	)
	request.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(response, request)

	if response.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", response.Code)
	}
}
