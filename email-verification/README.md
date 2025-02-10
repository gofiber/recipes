---
title: Email Verification Service
keywords: [email, verification, smtp, golang, fiber]
description: Email verification service with code generation and validation
---

# Email Verification Service with Fiber

[![Github](https://img.shields.io/static/v1?label=&message=Github&color=2ea44f&style=for-the-badge&logo=github)](https://github.com/gofiber/recipes/tree/master/email-verification) [![StackBlitz](https://img.shields.io/static/v1?label=&message=StackBlitz&color=2ea44f&style=for-the-badge&logo=StackBlitz)](https://stackblitz.com/github/gofiber/recipes/tree/master/email-verification)

A clean architecture based email verification service that generates and validates verification codes.

## Features

- Clean Architecture implementation
- In-memory verification code storage
- SMTP email service integration
- Code generation and hashing
- Configurable code expiration
- Thread-safe operations

## Project Structure

```
email-verification/
├── api/
│   └── handlers/         # HTTP handlers
├── application/          # Application business logic
├── domain/              # Domain models and interfaces
├── infrastructure/      # External implementations
│   ├── code/           # Code generation
│   ├── email/          # SMTP service
│   └── repository/     # Data storage
└── config/             # Configuration
```

## Configuration

Update `config/config.go` with your SMTP settings:

```go
func GetConfig() *Config {
    return &Config{
        SMTPHost:       "smtp.gmail.com",
        SMTPPort:       587,
        SMTPUser:       "your-email@gmail.com",
        SMTPPassword:   "your-app-password",
        CodeExpiration: time.Minute * 1,
    }
}
```

## API Endpoints

| Method | URL                        | Description                    |
|--------|----------------------------|--------------------------------|
| POST   | /verify/send/:email        | Send verification code         |
| POST   | /verify/check/:email/:code | Verify the received code      |

## Example Usage

1. Send verification code:
```bash
curl -X POST http://localhost:3000/verify/send/user@example.com
```

2. Verify code:
```bash
curl -X POST http://localhost:3000/verify/check/user@example.com/123456
```

## Response Examples

Success:
```json
{
    "message": "Code verified successfully"
}
```

Error:
```json
{
    "error": "invalid code"
}
```

## How to Run

1. Configure SMTP settings in `config/config.go`
2. Run the application:
```bash
go run main.go
```

## Dependencies

- [Fiber v2](https://github.com/gofiber/fiber)
- Go 1.21+
