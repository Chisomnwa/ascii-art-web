# How to Run Tests for ASCII-Art-Web Project

## Prerequisites

Your `main_test.go` file must be in the **root directory** of your project (same level as `main.go`).

File structure:
```
ascii-art-web/
├── main.go
├── main_test.go          ← Test file goes here
├── go.mod
├── handlers/
├── ascii-art-code/
├── templates/
└── ...
```

---

## Running Tests

### 1. Run ALL tests

```bash
go test
```

**Output:**
```
ok      ascii-art-web   0.245s
```

If any test fails, you'll see which one and why.

---

### 2. Run tests with VERBOSE output (see each test)

```bash
go test -v
```

**Output:**
```
=== RUN   TestAsciiArtValidBanner
--- PASS: TestAsciiArtValidBanner (0.00s)
=== RUN   TestAsciiArtInvalidBanner
--- PASS: TestAsciiArtInvalidBanner (0.00s)
=== RUN   TestAsciiArtAllValidBanners
--- PASS: TestAsciiArtAllValidBanners (0.00s)
...
ok      ascii-art-web   0.312s
```

---

### 3. Run a SPECIFIC test

```bash
go test -run TestAsciiArtValidBanner
```

Only runs that one test.

---

### 4. Run tests for a SPECIFIC function category

```bash
go test -run TestAsciiArt
```

Runs all tests matching "TestAsciiArt" (TestAsciiArtValidBanner, TestAsciiArtInvalidBanner, etc.)

---

### 5. Run tests by category/package

Run tests for handlers only:
```bash
go test ./handlers
```

Run tests for ascii-art-code only:
```bash
go test ./ascii-art-code
```

---

### 6. Run tests with code COVERAGE

```bash
go test -cover
```

**Output:**
```
ok      ascii-art-web   0.215s  coverage: 65.2% of statements
```

Shows what percentage of your code is tested.

---

### 7. Generate DETAILED coverage report

```bash
go test -coverprofile=coverage.out
```

This creates a `coverage.out` file. To view it as HTML:

```bash
go tool cover -html=coverage.out
```

Opens a browser showing exactly which lines are tested (green) and which are not (red).

---

## Test Organization

Your `main_test.go` has tests organized in three groups:

### Group 1: AsciiArt Function Tests
```bash
go test -run TestAsciiArt[^H]
```
Tests the core ASCII art generation function.

### Group 2: RootHandler Tests
```bash
go test -run TestRootHandler
```
Tests the home page handler.

### Group 3: AsciiArtHandler Tests
```bash
go test -run TestAsciiArtHandler
```
Tests the form submission handler.

### Group 4: Integration Tests
```bash
go test -run TestFull
```
Tests multiple components together.

---

## Running Tests as a Group (Team)

If you're working as a team on git:

### Before committing, run ALL tests
```bash
go test -v
```

Make sure everything passes before pushing to GitHub.

### To check if team member's changes broke anything
```bash
git pull
go test -v
```

If tests fail, check what changed in the failing test.

---

## Understanding Test Output

### ✅ PASS
```
--- PASS: TestAsciiArtValidBanner (0.00s)
```
Test succeeded.

### ❌ FAIL
```
--- FAIL: TestAsciiArtInvalidBanner (0.00s)
    main_test.go:45: Expected error with invalid banner, got nil
```
Test failed. Line 45 shows what went wrong.

---

## Quick Reference Cheat Sheet

| Command | What it does |
|---------|------------|
| `go test` | Run all tests |
| `go test -v` | Run all tests with details |
| `go test -run TestName` | Run specific test |
| `go test -run TestPrefix` | Run all tests starting with TestPrefix |
| `go test -cover` | Run tests and show coverage % |
| `go test -coverprofile=coverage.out` | Generate coverage file |
| `go tool cover -html=coverage.out` | View coverage in browser |
| `go test ./package` | Run tests in specific package only |
| `go test -timeout 30s` | Set timeout to 30 seconds |
| `go test -count=5` | Run all tests 5 times (for consistency) |

---

## Example Workflow

### Day 1: You write code
```bash
go test -v                    # Make sure your code works
git add .
git commit -m "Add error handling"
git push
```

### Day 2: Team member makes changes
```bash
git pull
go test -v                    # Check if anything broke
```

If something broke, check which tests failed and debug.

---

## Tips for the Team

1. **Always run `go test` before committing**
   - Don't push broken code

2. **If a test fails, don't just remove it**
   - Fix your code instead
   - Tests are like a safety net

3. **Add tests as you add features**
   - Don't test after the fact
   - Write test + code together

4. **Use `-v` flag to see details**
   - Helps debug failures faster

5. **Run tests locally before pushing**
   ```bash
   go test -v
   ```

---

## What Each Test Does

### AsciiArt Tests
- `TestAsciiArtValidBanner` → Verifies function works with valid banners
- `TestAsciiArtInvalidBanner` → Verifies function rejects bad banners
- `TestAsciiArtAllValidBanners` → Tests all 3 banners work
- `TestAsciiArtEmptyInput` → Tests function handles empty text

### RootHandler Tests
- `TestRootHandlerGET` → Home page loads correctly
- `TestRootHandlerInvalidPath` → Wrong path returns 404
- `TestRootHandlerPOST` → POST to home page rejected

### AsciiArtHandler Tests
- `TestAsciiArtHandlerValidRequest` → Form submission works
- `TestAsciiArtHandlerEmptyInput` → Empty text rejected
- `TestAsciiArtHandlerInvalidBanner` → Bad banner rejected
- `TestAsciiArtHandlerGET` → GET method rejected
- `TestAsciiArtHandlerInvalidPath` → Wrong path rejected
- `TestAsciiArtHandlerAllBanners` → All 3 banners work

### Integration Tests
- `TestFullWorkflow` → Complete user journey works
- `TestErrorRecovery` → Server doesn't crash on errors

---

## Debugging a Failing Test

If a test fails:

```
--- FAIL: TestAsciiArtHandlerEmptyInput (0.00s)
    main_test.go:185: Expected status 400 for empty input, got 200
```

This means:
- Test name: `TestAsciiArtHandlerEmptyInput`
- File and line: `main_test.go:185`
- What went wrong: Sent empty input, expected 400 error, got 200 success

**Fix:** Check your handler code at that location and see why empty input isn't being rejected.