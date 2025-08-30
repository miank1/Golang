package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

// Response struct
type PhoneResponse struct {
	PhoneNumber      string `json:"phoneNumber"`
	CountryCode      string `json:"countryCode"`
	AreaCode         string `json:"areaCode"`
	LocalPhoneNumber string `json:"localPhoneNumber"`
}

// Regex to validate phone number (E.164 with spaces allowed between segments)
var phoneRegex = regexp.MustCompile(`^\+?\d+( \d+)*$`)

// Map of example country codes -> ISO Alpha-2 (for demo, expand as needed)
var countryDialMap = map[string]string{
	"1":   "US", // USA
	"52":  "MX", // Mexico
	"34":  "ES", // Spain
	"351": "PT", // Portugal
}

// --- Extract phone parts ---
func parsePhoneNumber(raw string, fallbackCC string) (PhoneResponse, error) {
	resp := PhoneResponse{}
	normalized := strings.ReplaceAll(raw, " ", "") // remove spaces

	// Validate
	if !phoneRegex.MatchString(raw) {
		return resp, errors.New("invalid phone number format")
	}

	// Ensure leading +
	if !strings.HasPrefix(normalized, "+") {
		normalized = "+" + normalized
	}

	resp.PhoneNumber = normalized

	// Extract country code from prefix
	withoutPlus := normalized[1:]
	for i := 1; i <= 3 && i <= len(withoutPlus); i++ {
		cc := withoutPlus[:i]
		if iso, ok := countryDialMap[cc]; ok {
			resp.CountryCode = iso
			// areaCode and local: simple split (for demo, assume 3 digit area)
			if len(withoutPlus) > i+3 {
				resp.AreaCode = withoutPlus[i : i+3]
				resp.LocalPhoneNumber = withoutPlus[i+3:]
			}
			return resp, nil
		}
	}

	// If country code missing, must provide fallback
	if fallbackCC == "" {
		return resp, errors.New("missing country code")
	}
	resp.CountryCode = strings.ToUpper(fallbackCC)

	// Example split: take first 3 digits as area code
	if len(withoutPlus) > 3 {
		resp.AreaCode = withoutPlus[:3]
		resp.LocalPhoneNumber = withoutPlus[3:]
	} else {
		resp.AreaCode = withoutPlus
	}

	return resp, nil
}

// --- Handler ---
func phoneHandler(w http.ResponseWriter, r *http.Request) {
	phoneNumber := r.URL.Query().Get("phoneNumber")
	countryCode := r.URL.Query().Get("countryCode")

	if phoneNumber == "" {
		http.Error(w, "phoneNumber is required", http.StatusBadRequest)
		return
	}

	result, err := parsePhoneNumber(phoneNumber, countryCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/v1/phone-numbers", phoneHandler)
	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8081", nil)
}
