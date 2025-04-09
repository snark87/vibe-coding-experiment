# Go Coding Guidelines

## Introduction

This document outlines our coding standards and best practices for Go development in the Quantum Circuit Editor project. Following these guidelines will help ensure our codebase remains maintainable, performant, and consistent.

## Code Style and Formatting

### Formatting

- Always format your code using `gofmt` or `go fmt`.
- Use the default Go formatting rules without custom configurations.
- Configure your IDE to format code on save.

### Naming Conventions

- Use meaningful, descriptive names for variables, functions, and types.
- Follow Go naming conventions:
  - Use `camelCase` for private variables and functions
  - Use `PascalCase` for exported (public) variables, functions, and types
  - Use all-caps `SNAKE_CASE` for constants
  - Keep acronyms uppercase (e.g., `HTTP`, `URL`, `ID`)
  - Name the receiver variable appropriately (e.g., `func (s *Server) Run()`)
- Package names should be:
  - All lowercase, no underscores
  - Short, concise, and descriptive
  - Not plural (e.g., `store` not `stores`)

### Comments and Documentation

- Write clear comments for exported functions, types, and constants.
- Follow the Go comment format (start with the name of the item):
  ```go
  // Server represents an HTTP server for the circuit editor API.
  type Server struct {
      // ...
  }
  
  // Run starts the server on the specified port.
  func (s *Server) Run() error {
      // ...
  }
  ```
- Use doc comments (`//` or `/* */`) for package-level documentation.
- Document non-obvious behavior and complex logic.
- Update comments when you change the code.

### Code Organization

- Organize code by domain rather than by function (model, view, controller).
- Keep files focused on a single responsibility.
- Limit file size to 500 lines where possible; split larger files.
- Order functions consistently (e.g., exported functions first, then private).
- Group related functions together.

## Development Practices

### Error Handling

- Always check errors and handle them appropriately.
- Use error wrapping with context: `fmt.Errorf("failed to connect to DB: %w", err)`.
- Return errors rather than using panic.
- Create custom error types for significant error cases.
- Use sentinel errors (`var ErrNotFound = errors.New("entity not found")`) for specific errors that need to be checked.

### Logging

- Use the standard library `log/slog` package for structured logging.
- Configure slog with appropriate handlers early in the application lifecycle.
- Include relevant context in logs:
  ```go
  slog.Info("Processing request", 
      "method", req.Method,
      "path", req.URL.Path,
      "requestID", requestID)
  ```
- Use appropriate log levels:
  - `Debug`: Detailed information for debugging
  - `Info`: General information about application flow
  - `Warn`: Warning conditions that should be addressed
  - `Error`: Error conditions that affect functionality
- Avoid logging sensitive information (passwords, tokens, PII).

### Concurrency

- Use context for cancellation and timeouts.
- Always clean up goroutines to prevent leaks.
- Use sync primitives (Mutex, RWMutex) to protect shared resources.
- Prefer channels for communication between goroutines.
- Be explicit about whether channel operations are blocking.
- Consider using worker pools for CPU-bound tasks.
- Use buffered channels appropriately when you know the capacity.
- Handle panics in goroutines to prevent application crashes.

### Testing

- Aim for high test coverage of important logic.
- Write both unit and integration tests.
- Use table-driven tests for multiple test cases.
- Mock external dependencies for unit testing.
- Use Go's testing package and avoid third-party assertion libraries where possible.
- Name test functions as `Test<FunctionName><Scenario>`.
- Keep tests independent and idempotent.
- Include both positive and negative test cases.

### Dependencies

- Minimize external dependencies.
- Vendor dependencies for stability.
- Use Go modules for dependency management.
- Regularly update dependencies and check for security vulnerabilities.
- Prefer well-maintained, popular packages over obscure ones.

## Tooling and Linting

### Required Tools

- **golangci-lint**: Our primary linting tool with multiple linters enabled
- **staticcheck**: For additional static analysis
- **gofumpt**: Enhanced version of gofmt with additional rules
- **go vet**: For catching common errors
- **gosec**: For security-focused linting

### Linter Configuration

We enforce the following linters via golangci-lint:

```yaml
linters:
  enable:
    - errcheck      # Check for unchecked errors
    - gosimple      # Simplify code
    - govet         # Report suspicious constructs
    - ineffassign   # Detect unused assignments
    - staticcheck   # Static analysis checks
    - typecheck     # Type checking
    - unused        # Check for unused constants, variables, functions
    - nilerr        # Finds code that returns nil even if it checks that error is not nil
    - revive        # Drop-in replacement for golint
    - gocyclo       # Cyclomatic complexity
    - gosec         # Security checks
    - misspell      # Spelling mistakes
    - whitespace    # Checks for unnecessary whitespace
    - unparam       # Reports unused function parameters
    - prealloc      # Finds slice declarations that could be pre-allocated
    - bodyclose     # Checks whether HTTP response bodies are closed
    - gofmt         # Checks if code is formatted with gofmt
    - goimports     # Checks import statements are formatted according to goimports
```

### IDE Integration

Configure your IDE to:
- Run go fmt on save
- Enable linting with golangci-lint
- Show compiler errors in real time
- Provide Go language server features

Recommended VS Code extensions:
- Go (official Go extension)
- Go Critic
- Go Doc
- Go Test Explorer

### CI Integration

Our CI pipeline will:
- Run all tests
- Enforce linting rules
- Check for formatting issues
- Generate and check code coverage
- Perform static analysis

## Project-Specific Conventions

### Directory Structure

```
backend/
├── api/          # API handlers
├── auth/         # Authentication logic
├── circuit/      # Circuit-related domain logic
├── cmd/          # Command-line entrypoints
├── config/       # Application configuration
├── db/           # Database access
├── internal/     # Private packages not meant for external consumption
│   ├── validator/    # Internal validation utilities
│   ├── middleware/   # Shared HTTP middleware
│   └── errors/       # Common error types and handling
├── models/       # Data models
├── server/       # Server setup and middleware
├── simulation/   # Simulation client
└── utils/        # Utility functions
```

### Internal Package Usage

We use the `internal` package to:

1. **Hide Implementation Details**: Code within `internal/` cannot be imported by external projects, providing better encapsulation.
2. **Share Code Between Packages**: Code that needs to be shared between our own packages but shouldn't be part of our public API.
3. **Enforce Architectural Boundaries**: Prevent accidental dependencies on implementation details.

Guidelines for using the internal package:

- Place code in `internal/` when it's meant to be used across multiple packages within our application but shouldn't be imported by outside code.
- Use subdirectories within `internal/` to organize code by functionality.
- Don't place domain logic in `internal/` - this belongs in the appropriate domain package.
- Common middleware, validators, and error handling are good candidates for the internal package.

### HTTP Handlers

- Group related API endpoints in a single file.
- Use middleware for cross-cutting concerns (auth, logging, etc.).
- Follow RESTful principles where appropriate.
- Validate input early.
- Return structured error responses.
- Handle timeouts appropriately.

### Configuration

- Use environment variables for configuration.
- Have sensible defaults.
- Validate configuration at startup.
- Use a config struct to pass configuration around.

### Database Access

- Use prepared statements for queries.
- Handle transaction management carefully.
- Close resources explicitly.
- Consider using a query builder or ORM for complex queries.
- Log slow queries.

## Best Practices

### Performance

- Profile before optimizing.
- Avoid premature optimization.
- Use benchmarks to verify improvements.
- Consider memory usage, especially for long-running services.
- Optimize hot paths but keep code readable.

### Security

- Validate all user input.
- Sanitize data before using it in SQL, HTML, or command execution.
- Use secure defaults.
- Follow the principle of least privilege.
- Keep dependencies updated.

### Maintainability

- Write code that is easy to understand and modify.
- Prefer readability over cleverness.
- Use comments to explain "why" not "what" (the code shows what).
- Follow the Single Responsibility Principle.
- Keep functions focused and small (under 50 lines when possible).

### Accessibility for New Developers

- Document setup steps clearly.
- Provide examples for common tasks.
- Comment complex algorithms or business logic.
- Use consistent patterns across the codebase.

## Recommended Resources

- [Effective Go](https://golang.org/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)
- [Go Proverbs](https://go-proverbs.github.io/)
- [The Go Programming Language](https://www.gopl.io/) book

## Conclusion

These guidelines should be treated as a living document. As we learn and evolve as a team, we'll update and improve our practices. The primary goal is to create maintainable, reliable, and performant code that fulfills our product requirements and provides value to our users.

## Testing Guidelines

### Test Directory Structure

Tests in our Go codebase should be organized as follows:

```
backend/
├── api/
│   ├── routes.go
│   ├── routes_test.go      # Unit tests alongside production code
│   └── testdata/           # Test fixtures for the package
├── internal/
│   ├── middleware/
│   │   ├── auth.go
│   │   ├── auth_test.go    # Unit tests for middleware
│   │   └── auth_mock.go    # Mocks for testing
├── test/
│   ├── integration/        # Integration tests
│   │   ├── api_test.go
│   │   ├── fixtures/       # Test data for integration tests
│   │   └── setup.go        # Shared test setup code
│   └── e2e/                # End-to-end tests
│       ├── circuit_flow_test.go
│       └── auth_flow_test.go
```

Guidelines for test organization:
- Unit tests should be in the same package as the code they test, with a `_test.go` suffix.
- Integration tests should be placed in a separate `test/integration` directory.
- End-to-end tests should be placed in a separate `test/e2e` directory.
- Test fixtures (JSON files, SQL scripts, etc.) should be placed in a `testdata` directory within the relevant package for unit tests, or within the integration/e2e test directories.

### Types of Tests

1. **Unit Tests**
   - Test individual functions and methods in isolation
   - Use mocks for external dependencies
   - Validate specific behaviors and edge cases
   - Should be fast and deterministic

2. **Integration Tests**
   - Test interactions between multiple packages
   - May use test databases or containerized dependencies
   - Validate end-to-end workflows within the backend
   - Should be isolated from each other to allow parallel execution

3. **End-to-End Tests**
   - Test the complete system across all components
   - Run against a fully deployed test environment
   - Validate key user journeys
   - May be slower and run less frequently

### Test Writing Best Practices

#### 1. Test Structure

Follow the Arrange-Act-Assert (AAA) pattern:
```go
func TestSomething(t *testing.T) {
    // Arrange: Set up the test case
    expected := "expected result"
    input := "test input"
    
    // Act: Call the function under test
    actual := FunctionUnderTest(input)
    
    // Assert: Verify the results
    assert.Equal(t, expected, actual)
}
```

#### 2. Table-Driven Tests

Use table-driven tests for multiple related test cases:
```go
func TestCalculation(t *testing.T) {
    testCases := []struct {
        name     string
        input    int
        expected int
    }{
        {"zero input", 0, 0},
        {"positive input", 5, 25},
        {"negative input", -5, 25},
    }
    
    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            actual := Square(tc.input)
            assert.Equal(t, tc.expected, actual)
        })
    }
}
```

#### 3. Testable Code Design

- Write functions that are easy to test:
  - Dependency injection instead of global state
  - Clearly defined inputs and outputs
  - Single responsibility
- Use interfaces to enable mocking
- Avoid time.Now(), use injectable time providers
- Avoid direct database/API calls in units under test

#### 4. Test Helpers and Utilities

- Create reusable test helpers for common operations
- Place shared test utilities in internal/testutils package
- Use testing.TB interfaces to allow both benchmark and test use
```go
func setupTestDatabase(tb testing.TB) (*sql.DB, func()) {
    db, err := sql.Open("postgres", testConnString)
    require.NoError(tb, err, "Failed to connect to test DB")
    
    // Run migrations, seed data, etc.
    
    return db, func() {
        // Cleanup function to run after test
        db.Close()
    }
}
```

#### 5. Mocking

- Use interfaces to enable mocking
- Keep mocks in the same package as tests, with a `_mock.go` suffix
- Consider using [mockgen](https://github.com/golang/mock) for generating mocks from interfaces
- Don't mock the system under test, only external dependencies

Example mock structure:
```go
type MockUserRepository struct {
    FindByIDFunc func(ctx context.Context, id string) (*User, error)
}

func (m *MockUserRepository) FindByID(ctx context.Context, id string) (*User, error) {
    return m.FindByIDFunc(ctx, id)
}
```

#### 6. Testing HTTP Handlers

- Use `httptest` package for testing HTTP handlers
- Create table-driven tests for different request scenarios
- Test both success and error cases
- Validate response status, headers, and body

```go
func TestGetUserHandler(t *testing.T) {
    tests := []struct {
        name           string
        userID         string
        setupMock      func(m *MockUserService)
        expectedStatus int
        expectedBody   string
    }{
        {
            name:   "successful request",
            userID: "123",
            setupMock: func(m *MockUserService) {
                m.GetUserFunc = func(id string) (*User, error) {
                    return &User{ID: "123", Name: "Test User"}, nil
                }
            },
            expectedStatus: http.StatusOK,
            expectedBody:   `{"id":"123","name":"Test User"}`,
        },
        // Additional test cases...
    }
    
    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            // Arrange
            mockService := &MockUserService{}
            tc.setupMock(mockService)
            handler := NewUserHandler(mockService)
            
            req := httptest.NewRequest("GET", "/users/"+tc.userID, nil)
            rec := httptest.NewRecorder()
            
            // Act
            handler.GetUser(rec, req, []httprouter.Param{{Key: "id", Value: tc.userID}})
            
            // Assert
            if rec.Code != tc.expectedStatus {
                t.Errorf("Expected status %d but got %d", tc.expectedStatus, rec.Code)
            }
            if strings.TrimSpace(rec.Body.String()) != tc.expectedBody {
                t.Errorf("Expected body %q but got %q", tc.expectedBody, rec.Body.String())
            }
        })
    }
}
```

#### 7. Testing with Databases

- Use a test database, not the production database
- Reset the database state between tests
- Consider using testcontainers for consistent test environments
- Use transactions to roll back changes after tests when possible

#### 8. Asynchronous Testing

For testing code with goroutines or channels:
- Use synchronization primitives (WaitGroup, channels)
- Set appropriate timeouts to prevent hanging tests
- Be aware of race conditions in tests
- Run tests with `-race` flag to detect race conditions

```go
func TestAsyncProcess(t *testing.T) {
    // Create a channel to signal completion
    done := make(chan struct{})
    
    // Start the async process
    go func() {
        AsyncFunction()
        close(done)
    }()
    
    // Wait with timeout
    select {
    case <-done:
        // Test passes
    case <-time.After(2 * time.Second):
        t.Fatal("Test timed out")
    }
}
```

#### 9. Test Coverage

- Aim for high test coverage of critical code paths
- Use `go test -cover` to measure coverage
- Generate coverage reports with `go test -coverprofile=coverage.out`
- Focus on quality of tests, not just quantity
- Run `go test -race` regularly to catch race conditions

### Test Framework and Tools

#### Assertions with stretchr/testify

We use the `github.com/stretchr/testify` package for test assertions:

- Use `assert` for most assertions that should continue the test after failure
- Use `require` for critical assertions that should stop the test execution on failure

Example:
```go
func TestUser(t *testing.T) {
    // Arrange
    user := NewUser("john", "john@example.com")
    
    // Act
    err := user.Validate()
    
    // Assert
    require.NoError(t, err, "User validation should succeed")
    assert.Equal(t, "john", user.Username)
    assert.Equal(t, "john@example.com", user.Email)
}
```

#### Mocking with mockery

We use [mockery](https://github.com/vektra/mockery) for generating mocks from interfaces:

- Run `mockery --dir=<package> --name=<interface> --output=<output_dir> --outpkg=<package_name>` to generate mocks
- Place mocks in a `mocks` directory within the package being tested
- Use mockery's expectations API to set up behavior

Example:
```go
// Setting up a mock
mockRepo := mocks.NewUserRepository(t)
mockRepo.On("FindByID", "123").Return(&User{ID: "123", Name: "Test User"}, nil).Once()

// Using the mock
svc := NewUserService(mockRepo)
user, err := svc.GetUser("123")

// Verify expectations were met
mockRepo.AssertExpectations(t)
```

### Test Writing Best Practices

#### 1. Test Structure

Follow the Arrange-Act-Assert (AAA) pattern:
```go
func TestSomething(t *testing.T) {
    // Arrange: Set up the test case
    expected := "expected result"
    input := "test input"
    
    // Act: Call the function under test
    actual := FunctionUnderTest(input)
    
    // Assert: Verify the results
    assert.Equal(t, expected, actual)
}
```

#### 2. Table-Driven Tests

Use table-driven tests for data parametrization rather than behavior variations:

```go
func TestSquare(t *testing.T) {
    testCases := []struct {
        name     string
        input    int
        expected int
    }{
        {"zero input", 0, 0},
        {"positive input", 5, 25},
        {"negative input", -5, 25},
    }
    
    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            // Act
            actual := Square(tc.input)
            
            // Assert
            assert.Equal(t, tc.expected, actual)
        })
    }
}
```

Avoid complex table-driven tests with difficult-to-read mock setups or behavior variations. Instead, write separate test functions with clear purposes:

```go
// PREFER THIS:
func TestUserService_GetUser_Success(t *testing.T) {
    // Arrange
    mockRepo := mocks.NewUserRepository(t)
    mockRepo.On("FindByID", "123").Return(&User{ID: "123", Name: "Test User"}, nil)
    svc := NewUserService(mockRepo)
    
    // Act
    user, err := svc.GetUser("123")
    
    // Assert
    require.NoError(t, err)
    assert.Equal(t, "123", user.ID)
    assert.Equal(t, "Test User", user.Name)
}

func TestUserService_GetUser_NotFound(t *testing.T) {
    // Arrange
    mockRepo := mocks.NewUserRepository(t)
    mockRepo.On("FindByID", "999").Return(nil, ErrNotFound)
    svc := NewUserService(mockRepo)
    
    // Act
    user, err := svc.GetUser("999")
    
    // Assert
    assert.Nil(t, user)
    assert.ErrorIs(t, err, ErrNotFound)
}
```

### Integration Testing Practices

- Use real dependencies when possible (database, cache)
- Containerize dependencies with Docker for consistent testing
- Set up and tear down the environment for each test suite
- Use environment variables to configure test environments
- Implement retries for flaky external services
- Keep integration tests independent of each other

### Test Documentation

- Document special test requirements or setup steps
- Explain complex test scenarios or fixtures
- Document how to run different test suites
- Keep examples up-to-date with implementation changes