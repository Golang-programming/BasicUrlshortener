# URL Shortener in Go

## Overview

This is a simple URL shortener service written in Go. It allows you to shorten URLs and then redirect to the original URL using a shortened path.

## Getting Started

To run the URL shortener service on your local machine:

### Prerequisites

- Go installed. Download it from [golang.org](https://golang.org/dl/).

### Running the Server

1. Clone the repository:

   ```sh
   git clone https://github.com/Golang-programming/Urlshortener
   ```

2. Navigate to the project directory:

   ```sh
   cd Urlshortener
   ```

3. Run the server:

   ```sh
   go run main.go
   ```

   The server will be running on port 8080.

### API Endpoints

#### Shorten URL

- **Endpoint:** `/shortner`
- **Method:** POST
- **Request Body:**

  ```json
  {
    "url": "https://example.com"
  }
  ```

- **Response:**

  ```json
  "e99a18c428cb38d5f260853678922e03"
  ```

#### Redirect to Original URL

- **Endpoint:** `/redirect/{shortUrl}`
- **Method:** GET

  Replace `{shortUrl}` with the shortened URL.

- **Response:** Redirects to the original URL or returns a 404 error if not found.

## License

MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

Feel free to contribute by opening issues or submitting pull requests.

## Support

<br>
For questions, contact On [email](mailto:zeshanshakil0@gmail.com).

If you find this project useful, please star the repository and follow me on [GitHub](https://github.com/zeshantech).
