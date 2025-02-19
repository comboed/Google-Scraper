# Google-Scraper

## Description
Google-Scraper is a web scraping tool designed to extract search results from Google. It uses the `fasthttp` package for high-performance HTTP requests and integrates a CAPTCHA-solving service to bypass Google's security measures.

## Features
- Scrapes Google search results efficiently.
- Uses `fasthttp` for fast and lightweight HTTP requests.
- Supports proxy rotation for anonymity.
- Implements CAPTCHA solving via CapSolver API.
- Handles Google search result parsing and cookie management.

## Installation
1. Clone this repository:
   ```sh
   git clone https://github.com/yourusername/google-scraper.git
   cd google-scraper
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```

## Configuration
- Ensure you have a valid `CapSolverKey` in `globals.go`.
- Provide a list of proxies in `./data/proxies.txt`.

## Usage
Run the scraper using:
```sh
  go run main.go
```

The scraper starts a web server on port `8080` and exposes an API endpoint:
```
  GET /search?q=<query>&page=<page_number>
```

Example:
```
  curl "http://localhost:8080/search?q=golang&page=1"
```

## API Endpoints
- **`GET /search`**
  - Query Parameters:
    - `q` (string): The search query.
    - `page` (int, optional): The page number (default: 0).
  - Response:
    ```json
    {
      "data": [
        { "url": "https://example.com", "description": "Example description" },
        ...
      ]
    }
    ```

## File Structure
```
├── captcha.go        # Handles CAPTCHA solving
├── client.go         # Manages HTTP client with proxy support
├── cookies.go        # Handles cookie management and authorization
├── globals.go        # Defines global variables and constants
├── main.go           # Entry point with API server
├── scraper.go        # Core logic for Google search scraping
├── util.go           # Utility functions
├── go.mod            # Go module dependencies
├── go.sum            # Checksums for dependencies
├── README.md         # Documentation
```

## Dependencies
- [fasthttp](https://github.com/valyala/fasthttp)
- [gin-gonic](https://github.com/gin-gonic/gin)
- [fastjson](https://github.com/valyala/fastjson)
- [cookiejar](https://github.com/dgrr/cookiejar)

## License
This project is licensed under the MIT License.

## Disclaimer
Use this tool responsibly. Scraping Google may violate their Terms of Service.

