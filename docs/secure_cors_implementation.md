# Secure CORS Implementation Guide
## Quantum Circuit Editor Project

## 1. Introduction to CORS Security

### 1.1 Understanding CORS

Cross-Origin Resource Sharing (CORS) is a security mechanism that allows web applications running at one origin to access resources from a different origin. Without CORS, browsers enforce the Same-Origin Policy, which restricts how documents or scripts loaded from one origin can interact with resources from other origins.

### 1.2 Security Implications of CORS Misconfiguration

When CORS is improperly configured, it can lead to several security vulnerabilities:

1. **Cross-Site Request Forgery (CSRF)**: Overly permissive CORS policies may allow attackers to make authenticated requests from malicious sites.
2. **Data Leakage**: Relaxed configurations can expose sensitive data to unauthorized domains.
3. **Information Disclosure**: Error responses may reveal internal system details.
4. **Session Hijacking**: When credentials are allowed with a wildcard origin, an attacker's site could make authenticated requests.

### 1.3 Current Implementation Issues

Our code review revealed that the current CORS implementation uses a wildcard origin (`*`):

```go
AllowedOrigins: []string{"*"}
```

This configuration allows any website to make requests to our API. Even more concerning, this is combined with `AllowCredentials: true`, which creates a significant security risk. According to the CORS specification, using `*` with credentials enabled is not allowed by browsers, as this would expose user sessions to any origin.

## 2. Environment-Specific CORS Strategy

### 2.1 Development Environment

For local development, CORS settings should be permissive enough to support development workflows but still secure:

```go
// Development environment configuration
corsOptions := cors.Options{
    AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:5173", "http://127.0.0.1:3000"},
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Content-Type", "Authorization"},
    AllowCredentials: true,
    Debug:            true, // Enable detailed logging for development
}
```

### 2.2 Testing Environment

Testing environments should mimic production security while allowing test tools to function:

```go
// Testing environment configuration
corsOptions := cors.Options{
    AllowedOrigins:   []string{"https://test.quantum-circuit-editor.example", "https://test-api.quantum-circuit-editor.example"},
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Content-Type", "Authorization", "X-Test-Header"},
    AllowCredentials: true,
    MaxAge:           600, // Cache preflight results for 10 minutes
}
```

### 2.3 Staging Environment

Staging should closely mirror production but might need additional allowances for pre-release testing:

```go
// Staging environment configuration
corsOptions := cors.Options{
    AllowedOrigins:   []string{"https://staging.quantum-circuit-editor.example", "https://staging-api.quantum-circuit-editor.example"},
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Content-Type", "Authorization"},
    AllowCredentials: true,
    MaxAge:           3600, // Cache preflight results for 1 hour
}
```

### 2.4 Production Environment

Production should have the strictest CORS configuration:

```go
// Production environment configuration
corsOptions := cors.Options{
    AllowedOrigins:   []string{"https://quantum-circuit-editor.example", "https://api.quantum-circuit-editor.example"},
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Content-Type", "Authorization"},
    AllowCredentials: true,
    MaxAge:           3600, // Cache preflight results for 1 hour
    // No debug mode in production
}
```

## 3. Implementation Guidelines

### 3.1 Environment Variable Approach

Use environment variables to control CORS configuration based on deployment environment:

```go
func getCorsOptions() cors.Options {
    env := os.Getenv("APP_ENV")
    allowedOrigins := strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")

    if len(allowedOrigins) == 0 || (len(allowedOrigins) == 1 && allowedOrigins[0] == "") {
        // Fallback to defaults based on environment
        switch env {
        case "production":
            allowedOrigins = []string{"https://quantum-circuit-editor.example"}
        case "staging":
            allowedOrigins = []string{"https://staging.quantum-circuit-editor.example"}
        case "testing":
            allowedOrigins = []string{"https://test.quantum-circuit-editor.example"}
        default:
            // Development is the default
            allowedOrigins = []string{"http://localhost:3000", "http://localhost:5173"}
        }
    }

    return cors.Options{
        AllowedOrigins:   allowedOrigins,
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Content-Type", "Authorization"},
        AllowCredentials: true,
        MaxAge:           3600,
        Debug:            env != "production", // Enable debug mode except in production
    }
}
```

### 3.2 Recommended Header Settings

1. **AllowedOrigins**: Explicitly list all domains that need access
2. **AllowedMethods**: Only enable methods your API actually supports
3. **AllowedHeaders**: Limit to headers your application needs
4. **ExposedHeaders**: Expose only those headers the frontend needs
5. **MaxAge**: Set a reasonable cache duration (e.g., 1 hour = 3600 seconds)
6. **AllowCredentials**: Only enable if authenticated sessions are needed

### 3.3 Preflight Request Handling

The `OPTIONS` preflight requests are handled automatically by the rs/cors middleware. Ensure:

1. The `OPTIONS` method is included in AllowedMethods
2. A reasonable MaxAge is set to reduce preflight requests
3. The correct AllowedHeaders are configured to match what your frontend sends

### 3.4 Responding to Origin Validation

For dynamic origin validation beyond static lists:

```go
func validateOrigin(origin string) bool {
    validDomains := []string{
        ".quantum-circuit-editor.example", // Matches all subdomains
        "localhost",
        "127.0.0.1",
    }

    for _, domain := range validDomains {
        if strings.HasSuffix(origin, domain) {
            return true
        }
    }

    return false
}

// Create custom CORS handler with dynamic origin validation
c := cors.New(cors.Options{
    AllowOriginFunc: validateOrigin,
    // other options
})
```

## 4. Security Considerations

### 4.1 Origin Validation Best Practices

1. **Never use wildcards in production**: `*` allows any site to access your API
2. **Use exact origins**: Specify the full URL including scheme (http/https)
3. **Validate dynamically**: For applications with many valid origins, implement validation logic
4. **Consider subdomains**: Either list each subdomain explicitly or implement pattern matching

### 4.2 Secure Header Exposure

1. Only expose headers that clients legitimately need
2. Avoid exposing headers that might contain sensitive information
3. Configure `ExposedHeaders` to include only required response headers

### 4.3 Cookie Security with CORS

When `AllowCredentials` is set to `true`:
1. You **cannot** use wildcard `*` for allowed origins
2. Ensure cookies have `SameSite=None` and `Secure` flags for cross-origin requests
3. Consider the security implications of sending cookies cross-origin

### 4.4 Credential Sharing Considerations

When allowing credentials:
1. Authenticate and authorize all requests properly
2. Consider using token-based authentication instead of cookies when possible
3. Implement proper CSRF protection mechanisms

### 4.5 Common Attack Vectors

1. **CORS misconfiguration exploits**: Attackers targeting overly permissive CORS policies
2. **Origin header spoofing**: Attempts to bypass origin validation (mitigated by browser enforcement)
3. **CSRF combined with CORS**: Using permissive CORS to facilitate CSRF attacks

## 5. Testing and Validation

### 5.1 Manual CORS Validation

```javascript
// Test script to validate CORS configuration
fetch('https://api.quantum-circuit-editor.example/healthcheck', {
    method: 'GET',
    credentials: 'include', // Tests if credentials are allowed
    headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer test-token'
    }
})
.then(response => console.log('Success:', response))
.catch(error => console.error('Error:', error));
```

### 5.2 Common CORS Errors

1. **Missing Required Headers**: Browser errors when required CORS headers are absent
2. **Credentials with Wildcard**: Browsers reject requests with credentials when origin is `*`
3. **Method Not Allowed**: When the request method isn't in AllowedMethods
4. **Header Not Allowed**: When a request header isn't in AllowedHeaders

### 5.3 Automated Testing

Implement automated tests that validate your CORS configuration:

```go
func TestCORSConfiguration(t *testing.T) {
    // Setup server with CORS configuration
    server := setupTestServer()

    // Test cases
    testCases := []struct {
        name           string
        origin         string
        method         string
        headers        map[string]string
        expectedStatus int
        shouldHaveCORS bool
    }{
        {
            name:           "Valid origin",
            origin:         "https://quantum-circuit-editor.example",
            method:         "GET",
            headers:        map[string]string{"Content-Type": "application/json"},
            expectedStatus: http.StatusOK,
            shouldHaveCORS: true,
        },
        {
            name:           "Invalid origin",
            origin:         "https://evil-site.example",
            method:         "GET",
            headers:        map[string]string{"Content-Type": "application/json"},
            expectedStatus: http.StatusOK, // The request will succeed but...
            shouldHaveCORS: false,         // ...CORS headers should be absent
        },
        // Add more test cases
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            // Create and execute request
            // Validate expected CORS headers
            // Check for specific issues
        })
    }
}
```

## 6. Go-Specific Implementation Guidance

### 6.1 Using the rs/cors Package

```go
package main

import (
    "log/slog"
    "net/http"
    "os"
    "strings"

    "github.com/rs/cors"
)

func main() {
    // Get allowed origins from environment variable
    allowedOriginsEnv := os.Getenv("ALLOWED_ORIGINS")
    var allowedOrigins []string

    if allowedOriginsEnv != "" {
        allowedOrigins = strings.Split(allowedOriginsEnv, ",")
        // Trim spaces if any
        for i, origin := range allowedOrigins {
            allowedOrigins[i] = strings.TrimSpace(origin)
        }
    } else {
        // Default for development
        allowedOrigins = []string{"http://localhost:3000"}
        slog.Info("No ALLOWED_ORIGINS set, using default", "origins", allowedOrigins)
    }

    // Create CORS middleware
    c := cors.New(cors.Options{
        AllowedOrigins:   allowedOrigins,
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Content-Type", "Authorization"},
        ExposedHeaders:   []string{"X-Request-ID", "X-RateLimit-Remaining"},
        AllowCredentials: true,
        MaxAge:           3600, // 1 hour
    })

    // Create router and apply middleware
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(`{"status":"ok"}`))
    })

    // Wrap the handler with CORS middleware
    http.Handle("/api/", c.Handler(handler))

    slog.Info("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        slog.Error("Server failed", "error", err)
    }
}
```

### 6.2 Common Pitfalls in Go CORS Implementations

1. **Middleware Order**: CORS middleware must be applied before other handlers
2. **Multiple CORS Implementations**: Avoid applying multiple CORS handlers in the request chain
3. **AllowedOrigins vs. AllowOriginFunc**: Choose the right approach for your origin validation needs
4. **Debug Mode in Production**: Ensure debug mode is disabled in production
5. **Missing OPTIONS Handler**: Ensure your router handles OPTIONS requests correctly

### 6.3 Alternative CORS Packages

While `rs/cors` is recommended, you may also consider:

1. **gorilla/handlers**: Offers CORS handling as part of a broader middleware set
2. **gin-contrib/cors**: If using the Gin framework
3. **echo middleware**: If using the Echo framework

## 7. Implementation for the Quantum Circuit Editor

### 7.1 Immediate Changes Needed

Replace the current CORS configuration in `main.go`:

```go
// Current problematic code:
corsMiddleware := cors.New(cors.Options{
    AllowedOrigins:   []string{"*"}, // SECURITY: Update this with specific origins in production
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Content-Type", "Authorization"},
    AllowCredentials: true,
})

// Replace with:
corsOptions := getCorsOptions() // Function defined below
corsMiddleware := cors.New(corsOptions)

// Add this function:
func getCorsOptions() cors.Options {
    env := os.Getenv("APP_ENV")
    allowedOriginsEnv := os.Getenv("ALLOWED_ORIGINS")
    var allowedOrigins []string

    if allowedOriginsEnv != "" {
        allowedOrigins = strings.Split(allowedOriginsEnv, ",")
        for i, origin := range allowedOrigins {
            allowedOrigins[i] = strings.TrimSpace(origin)
        }
    } else {
        // Fallback to defaults based on environment
        switch env {
        case "production":
            allowedOrigins = []string{"https://quantum-circuit-editor.example"}
            slog.Info("Using production CORS origins", "origins", allowedOrigins)
        case "staging":
            allowedOrigins = []string{"https://staging.quantum-circuit-editor.example"}
            slog.Info("Using staging CORS origins", "origins", allowedOrigins)
        default:
            // Development is the default
            allowedOrigins = []string{
                "http://localhost:3000",
                "http://localhost:5173",
                "http://127.0.0.1:3000",
                "http://127.0.0.1:5173",
            }
            slog.Info("Using development CORS origins", "origins", allowedOrigins)
        }
    }

    return cors.Options{
        AllowedOrigins:   allowedOrigins,
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Content-Type", "Authorization"},
        AllowCredentials: true,
        MaxAge:           3600,
        Debug:            env != "production",
    }
}
```

### 7.2 Environment Configuration

Add the following to your deployment documentation:

**Development:**
```
APP_ENV=development
ALLOWED_ORIGINS=http://localhost:3000,http://localhost:5173
```

**Staging:**
```
APP_ENV=staging
ALLOWED_ORIGINS=https://staging.quantum-circuit-editor.example
```

**Production:**
```
APP_ENV=production
ALLOWED_ORIGINS=https://quantum-circuit-editor.example
```

## 8. Open Questions and Considerations

1. **Third-Party Integrations**: Are there any third-party services that will need to access our API?
   - If yes, we should include their domains in the allowed origins list.

2. **Mobile Applications**: Will there be mobile apps that need to access the API?
   - Mobile apps typically use their own HTTP clients that don't enforce CORS.

3. **Multiple Frontend Domains**: How many different domains will host our frontend application?
   - We need to ensure all legitimate domains are included in the allowed origins.

4. **Public API Strategy**: Will any parts of our API be publicly available without authentication?
   - Public APIs may require different CORS settings than authenticated endpoints.

5. **WebSocket Considerations**: Will we need WebSocket connections?
   - If yes, we need to ensure CORS is properly configured for WebSocket handshakes.

6. **Subdomains Strategy**: How should we handle multiple subdomains?
   - Consider using pattern matching for subdomains rather than listing each one.

## 9. Next Steps

1. Update the CORS configuration in the main application
2. Add environment variables to all deployment environments
3. Create automated tests for CORS configuration
4. Update deployment documentation with CORS settings
5. Schedule a security review of the updated configuration
6. Train developers on CORS security implications

## 10. References

1. MDN Web Docs on CORS: https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS
2. OWASP CORS Security Cheatsheet: https://cheatsheetseries.owasp.org/cheatsheets/CORS_Security_Cheat_Sheet.html
3. rs/cors Go package documentation: https://github.com/rs/cors
4. W3C CORS Specification: https://www.w3.org/TR/cors/
5. Go Security Practices: https://go.dev/security/
