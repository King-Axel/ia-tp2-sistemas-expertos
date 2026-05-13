package httpserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestListSEBRRules(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := NewRouter()
	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/sebr/rules", nil)

	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}

	var body SEBRRulesResponse
	if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
		t.Fatalf("expected valid JSON response: %v", err)
	}

	if len(body.Rules) != 32 {
		t.Fatalf("expected 32 rules, got %d", len(body.Rules))
	}
}

func TestInferSEBRDiagnosis(t *testing.T) {
	gin.SetMode(gin.TestMode)

	requestBody := []byte(`{
		"temperature": 40,
		"headache": true,
		"cough": true,
		"fatigue": true,
		"loss_of_taste_or_smell": true,
		"breathing_difficulty": false,
		"chest_pain": false,
		"sore_throat": false,
		"pain_behind_eyes": false,
		"vomiting": false,
		"muscle_or_joint_pain": false,
		"nausea": false,
		"bleeding": false,
		"close_contact": false,
		"had_covid_before": false
	}`)

	router := NewRouter()
	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "/sebr/infer", bytes.NewReader(requestBody))
	request.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}

	var body SEBRInferenceResponse
	if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
		t.Fatalf("expected valid JSON response: %v", err)
	}

	if body.Result != "compatible_con_covid" {
		t.Fatalf("expected compatible_con_covid, got %q", body.Result)
	}

	if len(body.ActivatedRules) == 0 {
		t.Fatal("expected activated rules")
	}
}

func TestCreateSEBRRule(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := NewRouter()
	response := httptest.NewRecorder()
	request := httptest.NewRequest(
		http.MethodPost,
		"/sebr/rules",
		bytes.NewReader([]byte(`{"name":"Possible flu","expression":"IF fever AND headache THEN possible_flu"}`)),
	)
	request.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(response, request)

	if response.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", response.Code)
	}

	var body CreateSEBRRuleResponse
	if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
		t.Fatalf("expected valid JSON response: %v", err)
	}

	if body.Rule.ID != "R33" {
		t.Fatalf("expected R33 id, got %q", body.Rule.ID)
	}
}

func TestInferUsesRuntimeSEBRRules(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := NewRouter()
	createResponse := httptest.NewRecorder()
	createRequest := httptest.NewRequest(
		http.MethodPost,
		"/sebr/rules",
		bytes.NewReader([]byte(`{"name":"COVID from headache","expression":"IF headache THEN compatible_con_covid"}`)),
	)
	createRequest.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(createResponse, createRequest)

	if createResponse.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", createResponse.Code)
	}

	inferBody := []byte(`{
		"temperature": 37,
		"headache": true,
		"cough": false,
		"fatigue": false,
		"loss_of_taste_or_smell": false,
		"breathing_difficulty": false,
		"chest_pain": false,
		"sore_throat": false,
		"pain_behind_eyes": false,
		"vomiting": false,
		"muscle_or_joint_pain": false,
		"nausea": false,
		"bleeding": false,
		"close_contact": false,
		"had_covid_before": false
	}`)
	inferResponse := httptest.NewRecorder()
	inferRequest := httptest.NewRequest(http.MethodPost, "/sebr/infer", bytes.NewReader(inferBody))
	inferRequest.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(inferResponse, inferRequest)

	if inferResponse.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", inferResponse.Code)
	}

	var body SEBRInferenceResponse
	if err := json.Unmarshal(inferResponse.Body.Bytes(), &body); err != nil {
		t.Fatalf("expected valid JSON response: %v", err)
	}

	if !containsString(body.FinalConclusions, "compatible_con_covid") {
		t.Fatal("expected runtime rule to add compatible_con_covid")
	}
}

func TestCreateSEBRRuleRejectsInvalidExpression(t *testing.T) {
	gin.SetMode(gin.TestMode)

	router := NewRouter()
	response := httptest.NewRecorder()
	request := httptest.NewRequest(
		http.MethodPost,
		"/sebr/rules",
		bytes.NewReader([]byte(`{"name":"Invalid","expression":"IF fever AND headache OR cough THEN possible_flu"}`)),
	)
	request.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(response, request)

	if response.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", response.Code)
	}
}

func containsString(values []string, target string) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}

	return false
}
