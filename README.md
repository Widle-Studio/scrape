# scrape
Scrape unique URLs and emails from a website (using URL) in GoLang

In this example, we use the Gorilla Mux package to handle the routing and define a single `/scrape` endpoint that accepts a POST request with the website URL in the request body. The `scrapeHandler` function is responsible for performing the scraping logic and returning the extracted URLs and emails as a JSON response.

To run the REST API, you can start the server by executing `go run main.go` in the terminal. The server will listen on port `8000`. You can then send a POST request to `http://localhost:8000/scrape` with the website URL in the request body to get the scraped data in the response.

Note that this is a basic example, and you might want to add error handling, input validation, and further improvements based on your specific requirements.
