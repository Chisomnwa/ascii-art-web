package main

import (
	asciiartcode "ascii-art-web/ascii-art-code"
	"ascii-art-web/handlers"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

// ============================================================
// TESTS FOR ascii-art.go (AsciiArt function)
// ============================================================

/*
TestAsciiArtValidBanner tests that AsciiArt returns no error
when given a valid banner name (standard, shadow, thinkertoy
*/
func TestAsciiArtValidBanner(t *testing.T) {
	// Test with "standard" banner
	result, err := asciiartcode.AsciiArt("A", "standard")

	// err should be nil (no error)
	if err != nil {
		t.Errorf("Expected no error with valid banner 'standard', got %v", err)
	}

	// result should not be empty
	if result == "" {
		t.Error("EXpected non-empty result from AsciiArt, got empty string")
	}
}

/*
TestAsciiArtInvalidBanner tests that AsciiArt returns an error
when given an invalid banner name.
*/
func TestAsciiArtInvalidBanner(t *testing.T) {
	// Test with an invalid banner name
	result, err := asciiartcode.AsciiArt("hello", "invalid")

	// err should NOT be nil (error expected)
	if err == nil {
		t.Error("EXpected error with invalid banner, got nil")
	}

	// result should be empty when error occurs
	if result != "" {
		t.Errorf("Expected empty result with error, got %s", result)
	}
}

/*
TestAsciiArtAllValidBanners tests that AsciiArt works with all three
valid banner options: standard, shadow, and thinkertoy.
*/
func TestAsciiArtAllValidBanners(t *testing.T) {
	// List of valid banners to test
	banner := []string{"standard", "shadow", "thinkertoy"}
	testText := "Hi"

	// Loop through the each banner and test
	for _, banner := range banner {
		result, err := asciiartcode.AsciiArt(testText, banner)

		// Each should succeed without error
		if err != nil {
			t.Errorf("Failed for banner '%s': %v", banner, err)
		}

		// Each should return non-empty result
		if result == "" {
			t.Errorf("GOt empty result for banner '%s'", banner)
		}
	}
}

// ============================================================
// TESTS FOR root-handler.go (RootHandler)
// ============================================================

/*
TestRootHandlerGET tests that RootHandler correctly handles GET requests
to the "/" path and returns HTTP 200 with HTML content.
*/
func TestRootHandlerGET(t *testing.T) {
	// Create a mock HTTP GET request to "/"
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// CReate a ResponseRecorder to capture the response
	// (like a mock response writer)
	w := httptest.NewRecorder()

	// Call the handler
	handlers.RootHandler(w, req)

	// Check response status code is 200 OK
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	// Check response contains HTML (should have <h1> tag from template)
	body := w.Body.String()
	if !strings.Contains(body, "<h1>") {
		t.Error(("EXpected HTML Response with <h1> tag, but not found"))
	}
}

/*
TestRootHandlerInvalidPath tests that RootHandler rejects requests
to paths other than "/" and returns HTTP 404.
*/

func TestRootHandlerInvalidPath(t *testing.T) {
	// Create a GET request to an invalid path
	req, err := http.NewRequest("GET", "/invalid", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	w := httptest.NewRecorder()
	handlers.RootHandler(w, req)

	// Should return 404 NOt FOund
	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status 404, got %d", w.Code)
	}
}

/*
TestRootHandlerPOST tests that RootHandler rejects POST requests
to "/" and returns HTTP 404.
*/
func TestRootHandlerPOST(t *testing.T) {
	// Create a POST request to "/"
	req, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	w := httptest.NewRecorder()
	handlers.RootHandler(w, req)

	// Should reject POST with 404
	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status 404 for POST, Got %d", w.Code)
	}

}

// ============================================================
// TESTS FOR ascii-art-handler.go (AsciiArtHandler)
// ============================================================

/*
TestAsciiArtHandlerValidRequest tests that AsciiArtHandler correctly
processes a valid POST request with text and banner, and returns HTTP 200.
*/

func TestAsciiArtHandlerValidRequest(t *testing.T) {
	// Create form data with valid inputs
	formData := url.Values{}
	formData.Set("inputText", "Hi")
	formData.Set("banner", "standard")

	// Create a POST request with form data
	req, err := http.NewRequest("POST", "/ascii-art",
		strings.NewReader(formData.Encode()))

	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Set form content type
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	handlers.AsciiArtHandler(w, req)

	// Response return 200 OK
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", w.Code)
	}

	// REsponse should contain HTML with the result
	body := w.Body.String()
	if !strings.Contains(body, "<pre>") {
		t.Error("Expected HTML response with <pre> tag for ASCII art")
	}
}

/*
TestAsciiArtHandlerGET tests that AsciiArtHandler rejects GET requests
(only POST allowed) and returns HTTP 400.
*/
func TestAsciiArtHandlerGET(t *testing.T) {
	// Create a GET request (should be rejected)
	req, err := http.NewRequest("GET", "/ascii-art", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	w := httptest.NewRecorder()
	handlers.AsciiArtHandler(w, req)

	// Should return 400 Bad Request
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 400 for GET request, got %d", w.Code)
	}
}

/*
TestAsciiArtHandlerInvalidPath tests that AsciiArtHandler rejects
requests to paths other than "/ascii-art" and returns HTTP 404.
*/
func TestAsciiArtHandlerInvalidPath(t *testing.T) {
	// Create a POST request to wrong path
	formData := url.Values{}
	formData.Set("inputText", "Hi")
	formData.Set("banner", "standard")

	req, err := http.NewRequest("POST", "/ascii-art-typo",	
		strings.NewReader(formData.Encode()))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	req.Header.Set("Content-TYpe", "application/x-www-form-urlencoded")

	w := httptest.NewRecorder()
	handlers.AsciiArtHandler(w, req)

	// Should return 404 Not Found
	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status 404 for wrong path, got %d", w.Code)
	}
}

// ============================================================
// INTEGRATION TESTS (testing multiple components together)
// ============================================================
 
/*
TestFullWorkflow tests the complete flow: user submits form on home page,
gets redirected to /ascii-art, and receives ASCII art result.
This simulates a real user interaction.
*/
func TestFullWorkflow(t *testing.T) {
	// Step 1: User visits home page (GET /)
	homeReq, _ := http.NewRequest("GET", "/", nil)
	homeW := httptest.NewRecorder()
	handlers.RootHandler(homeW, homeReq)
	
	if homeW.Code != http.StatusOK {
		t.Errorf("Home page failed: expected 200, got %d", homeW.Code)
	}
	
	// Step 2: User fills form and submits to /ascii-art
	formData := url.Values{}
	formData.Set("inputText", "Web")
	formData.Set("banner", "shadow")
	
	asciiReq, _ := http.NewRequest("POST", "/ascii-art",
		strings.NewReader(formData.Encode()))
	asciiReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	
	asciiW := httptest.NewRecorder()
	handlers.AsciiArtHandler(asciiW, asciiReq)
	
	if asciiW.Code != http.StatusOK {
		t.Errorf("ASCII art endpoint failed: expected 200, got %d", asciiW.Code)
	}
	
	// Step 3: Verify result contains ASCII art
	body := asciiW.Body.String()
	if !strings.Contains(body, "<pre>") {
		t.Error("Expected ASCII art wrapped in <pre> tags")
	}
}
 
/*
TestErrorRecovery tests that the server properly handles errors
and doesn't crash when given bad input.
*/
func TestErrorRecovery(t *testing.T) {
	testCases := []struct {
		name   string
		text   string
		banner string
		expect int
	}{
		{"Empty text", "", "standard", http.StatusBadRequest},
		{"Bad banner", "hello", "bad", http.StatusBadRequest},
		{"Valid", "hello", "standard", http.StatusOK},
	}
	
	for _, tc := range testCases {
		formData := url.Values{}
		formData.Set("inputText", tc.text)
		formData.Set("banner", tc.banner)
		
		req, _ := http.NewRequest("POST", "/ascii-art",
			strings.NewReader(formData.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		
		w := httptest.NewRecorder()
		handlers.AsciiArtHandler(w, req)
		
		if w.Code != tc.expect {
			t.Errorf("%s: expected %d, got %d", tc.name, tc.expect, w.Code)
		}
	}
}