package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/scrape", scrapeHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func scrapeHandler(w http.ResponseWriter, r *http.Request) {
	// Read the URL from the request body
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	websiteURL := string(bodyBytes)

	// Fetch the website content
	resp, err := http.Get(websiteURL)
	if err != nil {
		http.Error(w, "Failed to fetch website content", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	bodyBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}
	body := string(bodyBytes)

	// Extract unique URLs from the website
	urls := extractUniqueURLs(body)

	// Extract unique emails from the website
	emails := extractUniqueEmails(body)

	// Create a response payload
	response := struct {
		URLs   []string `json:"urls"`
		Emails []string `json:"emails"`
	}{
		URLs:   urls,
		Emails: emails,
	}

	// Send the response as JSON
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func extractUniqueURLs(body string) []string {
	// Define a regular expression pattern to match URLs
	urlPattern := regexp.MustCompile(`(https?://\S+)`)

	// Find all matches of URLs in the body
	urlMatches := urlPattern.FindAllStringSubmatch(body, -1)

	// Extract the URLs from the matches
	var urls []string
	for _, match := range urlMatches {
		url := match[1]
		urls = append(urls, url)
	}

	// Remove duplicate URLs
	urls = removeDuplicates(urls)

	return urls
}

func extractUniqueEmails(body string) []string {
	// Define a regular expression pattern to match emails
	emailPattern := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

	// Find all matches of emails in the body
	emailMatches := emailPattern.FindAllString(body, -1)

	// Remove duplicate emails
	emails := removeDuplicates(emailMatches)

	return emails
}

func removeDuplicates(items []string) []string {
	// Create a map to track unique items
	uniqueMap := make(map[string]bool)

	// Iterate over the items and add them to the map
	for _, item := range items {
		uniqueMap[item] = true
	}

	// Create a new slice to store the unique items
	var uniqueItems []string

	// Append the unique items to the new slice
	for item := range uniqueMap {
		uniqueItems = append(uniqueItems, item)
	}

	return uniqueItems
}
