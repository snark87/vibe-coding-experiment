You are a tech lead. Write detailed code review guidelines for go code and general guidelines on writing go code and put these two documents into ./docs

Keep in mind, that logging is important (one should use slog for structured logging).
Readability and maintainability is important.
When writing go guidelines pay attention to lint tools to use

## Addition

We use stretchr assert and require for test assertions. Reflect this in the document and adjust examples.

We use mockery for mocks

We use Arrange-Act-Assert pattern

Avoid tabular tests with difficult-to-read tables (like mock setup in tables). Use tabular tests only for test data parametrization (rather than behavior). Otherwise, prefer different tests