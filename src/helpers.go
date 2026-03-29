package frontend_app

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	Code    int   `json:"code"`
}

type ValidateResult struct {
	Valid  bool
	Errors []string
}

func ValidateString(s string) ValidateResult {
	if len(s) < 1 {
		return ValidateResult{Valid: false, Errors: []string{"Field is empty"}}
	}

	if !strings.ContainsAny(s, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789") {
		return ValidateResult{Valid: false, Errors: []string{"Field contains only valid characters"}}
	}

	return ValidateResult{Valid: true, Errors: []string{}}
}

func ValidateEmail(email string) ValidateResult {
	if len(email) < 1 {
		return ValidateResult{Valid: false, Errors: []string{"Email field is empty"}}
	}

	if !strings.Contains(email, "@") {
		return ValidateResult{Valid: false, Errors: []string{"Email field is invalid"}}
	}

	return ValidateResult{Valid: true, Errors: []string{}}
}

func ValidateID(id string) ValidateResult {
	if len(id) < 1 {
		return ValidateResult{Valid: false, Errors: []string{"ID field is empty"}}
	}

	if !strings.Contains(id, "UUID") {
		return ValidateResult{Valid: false, Errors: []string{"ID field is invalid"}}
	}

	return ValidateResult{Valid: true, Errors: []string{}}
}

func ErrorResponse(w http.ResponseWriter, code int, message string) {
	json.NewEncoder(w).Encode(ErrorResponse{Code: code, Message: message})
}

func GetQueryParams(r *http.Request, paramName string) string {
	params := r.URL.Query()
	value := params[paramName]

	if len(value) > 0 {
		return value[0]
	}

	return ""
}

func ParseInt(s string) (int, error) {
	return strconv.Atoi(s)
}