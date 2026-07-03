# ASCII Art Web

## Description

ASCII Art Web is a web-based application that converts user input text into ASCII art using different banner styles. It provides a graphical user interface (GUI) built with HTML and a backend server written in Go.

The application allows users to:
- Enter any text they want to convert
- Choose from three different ASCII art banner styles: **standard**, **shadow**, and **thinkertoy**
- Generate and view the ASCII art result directly in the browser
- Experience a simple, clean, and responsive web interface

This project extends the original [ascii-art](https://github.com/01-edu/public) CLI project by making it accessible through a web interface.

---

## Authors

- **Chisom Nnamani** (@cnnamani)
- **Samuel Akeju** (@sakeju)
- **Oghenekevwe Ajayi** (ogajayi)

---

## Usage

### Prerequisites

Before running the application, ensure you have:
- **Go 1.16 or higher** installed on your system
- The project files downloaded to your local machine

### Installation

1. **Clone or download the project:**
   ```bash
   git clone https://github.com/cnnamani/ascii-art-web.git
   cd ascii-art-web
   ```

2. **Verify the folder structure:**
   ```
   ascii-art-web/
   ├── ascii-art-code/
   │   └── ascii-art.go
   ├── handlers/
   │   ├── ascii-art-handler.go
   │   └── root-handler.go
   ├── templates/
   │   └── index.html
   ├── main.go
   ├── main_test.go
   ├── go.mod
   ├── shadow.txt
   ├── standard.txt
   ├── thinkertoy.txt
   └── README.md
   ```

### Running the Server

1. **Navigate to the project directory:**
   ```bash
   cd ascii-art-web
   ```

2. **Start the server:**
   ```bash
   go run .
   ```

3. **Expected output:**
   ```
   server starting on port 4000
   ```

4. **Open your browser:**
   - Navigate to: `http://localhost:4000`
   - The home page with the ASCII art generator form will load

### Using the Application

1. **Enter text:** Type any text in the input field (do not leave it empty)
2. **Select a banner:** Choose one of three styles from the dropdown:
   - Standard
   - Shadow
   - Thinkertoy
3. **Generate:** Click the "Generate" button
4. **View result:** The ASCII art will be displayed below the form in a `<pre>` block

### Example

**Input:**
- Text: "Hedlo"
- Banner: "standard"

**Output:**
```
 _    _          _ _       
| |  | |        | | |      
| |__| | ___  __| | | ___  
|  __  |/ _ \/ _` | |/ _ \ 
| |  | |  __/ (_| | | (_) |
|_|  |_|\___|\__,_|_|\___/ 
```

---

## Implementation Details

### Architecture

The application follows a modular architecture with clear separation of concerns:

```
main.go (Router)
    ├── handlers/root-handler.go (GET /)
    ├── handlers/ascii-art-handler.go (POST /ascii-art)
    └── ascii-art-code/ascii-art.go (Business Logic)
```

### How It Works

#### 1. **Frontend (HTML Template)**
- Located in: `templates/index.html`
- Contains a form with:
  - Text input field (`<textarea>`)
  - Banner selector dropdown (`<select>`)
  - Submit button
  - Result display area (`<pre>`)

#### 2. **Router (main.go)**
- Registers two routes:
  - `GET /` → Serves the home page with the form
  - `POST /ascii-art` → Processes form submission and generates ASCII art
- Starts the HTTP server on port 4000

#### 3. **Handlers**

**RootHandler (root-handler.go)**
- Handles HTTP GET requests to `/`
- Validates HTTP method (only GET allowed)
- Parses and executes the HTML template
- Returns the home page to the user
- Status codes: 200 OK, 400 Bad Request, 404 Not Found, 500 Internal Server Error

**AsciiArtHandler (ascii-art-handler.go)**
- Handles HTTP POST requests to `/ascii-art`
- Validates HTTP method (only POST allowed)
- Parses form data (text input and banner selection)
- Validates inputs:
  - Text cannot be empty
  - Banner must be one of: standard, shadow, or thinkertoy
- Calls the AsciiArt function to generate the result
- Executes the template with the result
- Returns the HTML page with ASCII art displayed
- Status codes: 200 OK, 400 Bad Request, 404 Not Found, 500 Internal Server Error

#### 4. **Core Logic (ascii-art.go)**

**AsciiArt Function**
```go
func AsciiArt(text, banner string) (string, error)
```

**Parameters:**
- `text` (string): The input text to convert to ASCII art
- `banner` (string): The banner style ("standard", "shadow", or "thinkertoy")

**Returns:**
- `result` (string): The generated ASCII art
- `error`: Any error that occurred during processing

**Algorithm:**

1. **Validate Banner:** Check if the banner name is valid (standard, shadow, or thinkertoy)
2. **Read Banner File:** Load the appropriate `.txt` file containing ASCII art character definitions
3. **Parse Banner Data:** Split the file into lines for character mapping
4. **Process Input Text:** Replace escaped newlines (`\n`) with actual newlines
5. **Build ASCII Art:**
   - Loop through each line of input
   - For each character, calculate its ASCII index in the banner file
   - Each character has 8 lines of ASCII art (height)
   - Each character takes 9 positions width (including spacing)
   - Index formula: `(ASCII value - 32) * 9 + 1 + current_height`
6. **Return Result:** Return the complete ASCII art string

**Error Handling:**
- Returns error if banner name is invalid
- Returns error if banner file is not found
- Returns error if banner file cannot be read

#### 5. **Banner Files**

Three `.txt` files contain ASCII art character definitions:
- `standard.txt` - Simple, clean ASCII art
- `shadow.txt` - Shadowed ASCII art style
- `thinkertoy.txt` - Decorative ASCII art style

Each file contains ASCII art for all printable ASCII characters (space to tilde), organized in lines.

### Data Flow

```
User Input (HTML Form)
        ↓
POST /ascii-art
        ↓
AsciiArtHandler
        ├── Validate HTTP method
        ├── Validate path
        ├── Parse form data
        ├── Validate inputs
        └── Call AsciiArt()
                ↓
        AsciiArt Function
        ├── Validate banner
        ├── Read banner file
        ├── Parse banner data
        ├── Process input text
        └── Build ASCII art
                ↓
        Return result to handler
        ↓
Execute template with result
        ↓
Send HTML response to browser
        ↓
Display ASCII art to user
```

### HTTP Status Codes

| Status Code | Meaning | When Used |
|-------------|---------|-----------|
| 200 OK | Success | Valid request processed successfully |
| 400 Bad Request | Client error | Empty input, invalid banner, wrong HTTP method, malformed form data |
| 404 Not Found | Not found | Invalid path, missing template file, missing banner file |
| 500 Internal Server Error | Server error | Unexpected error during template parsing/execution |

---

## Testing

The project includes comprehensive unit and integration tests in `main_test.go`.

### Test Categories

1. **AsciiArt Function Tests** - Test the core ASCII art generation logic
2. **RootHandler Tests** - Test the home page handler
3. **AsciiArtHandler Tests** - Test the form submission handler
4. **Integration Tests** - Test complete user workflows

### Running Tests

**Run all tests:**
```bash
go test -v
```

**Run specific test:**
```bash
go test -run TestAsciiArtValidBanner -v
```

**Run tests with coverage:**
```bash
go test -cover
```

**Generate coverage report:**
```bash
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

**Run tests before committing (team workflow):**
```bash
go test -v
git add .
git commit -m "Your message"
git push
```

---

## Technologies Used

- **Language:** Go (Golang)
- **Web Framework:** Go standard library (`net/http`)
- **Templates:** Go `html/template` package
- **Testing:** Go `testing` package
- **Frontend:** HTML5, CSS (basic styling)

---

## Folder Structure

```
ascii-art-web/
│
├── main.go                 # Application entry point, route registration
├── main_test.go            # Unit and integration tests
├── go.mod                  # Go module file
│
├── ascii-art-code/         # Core logic package
│   └── ascii-art.go        # ASCII art generation function
│
├── handlers/               # HTTP request handlers
│   ├── root-handler.go     # Handles GET /
│   └── ascii-art-handler.go # Handles POST /ascii-art
│
├── templates/              # HTML templates
│   └── index.html          # Main page form
│
├── standard.txt            # Standard banner definitions
├── shadow.txt              # Shadow banner definitions
├── thinkertoy.txt          # Thinkertoy banner definitions
│
└── README.md               # This file
```

---

## Good Practices Implemented

✅ **Error Handling**
- Proper HTTP status codes for all scenarios
- Meaningful error messages to users and developers
- Input validation before processing

✅ **Code Organization**
- Modular structure with separate packages (handlers, ascii-art-code)
- Clear separation of concerns
- Each file has a specific responsibility

✅ **Documentation**
- Comments explaining each function and its purpose
- README with detailed implementation details
- Code is self-documenting with clear variable names

✅ **Testing**
- Comprehensive test coverage
- Unit tests for individual functions
- Integration tests for complete workflows
- Error scenarios included

✅ **Go Best Practices**
- Using standard library only (no external dependencies)
- Proper error handling with `if err != nil`
- Clear naming conventions
- Proper use of HTTP methods and status codes

---

## Troubleshooting

### Problem: Port 4000 already in use

**Solution:**
- Change the port in `main.go` from `:4000` to another port like `:8080`:
  ```go
  err := http.ListenAndServe(":8080", nil)
  ```
- Then access `http://localhost:8080`

### Problem: Template file not found error

**Ensure:**
- `templates/index.html` exists in the project root
- You're running the server from the project root directory
- Use `go run .` to run from the root

### Problem: Banner file not found error

**Ensure:**
- `standard.txt`, `shadow.txt`, and `thinkertoy.txt` exist in the project root
- They're in the correct directory when running the server

### Problem: Tests are failing

**Solution:**
1. Run tests with verbose output: `go test -v`
2. Check which test is failing
3. Read the error message carefully
4. Verify your code implements the requirements

---

## Contributing

When contributing to this project:

1. Run all tests before committing:
   ```bash
   go test -v
   ```

2. Follow the code style and structure
3. Add tests for any new features
4. Update this README if you add new functionality
5. Use meaningful commit messages

---

## License

This project is part of the 01-Edu curriculum.

---

## References

- [Go Documentation](https://golang.org/doc/)
- [Go HTTP Package](https://golang.org/pkg/net/http/)
- [Go Templates](https://golang.org/pkg/html/template/)
- [ASCII Art Example](http://patorjk.com/software/taag/)

---

## Contact & Support

For issues or questions:
- Check existing GitHub issues
- Create a new GitHub issue with detailed description
- Contact team members directly

---

**Last Updated:** May 2024
**Project Version:** 1.0
