# Redis Rate Limiter Middleware for Fiber

This recipe demonstrates how to build a custom rate limiter middleware in [Fiber](https://github.com/gofiber/fiber) using [Redis](https://redis.io/) as the backend.

The middleware limits incoming requests per IP address within a given time window (e.g., 10 requests per minute).

## 📦 Stack

- Go
- Fiber Web Framework
- Redis (go-redis v9)

## 📁 Project Structure

rate-limit-redis/
├── config/
│ └── redis.go # Initializes Redis client
├── handlers/
│ └── home.go # Example handler
├── middleware/
│ └── limiter.go # Rate limiting logic
├── main.go # Entry point


## 🚀 Getting Started

### 1. Run Redis locally
Make sure Redis is running on `localhost:6379`. You can use Docker:

```bash
docker run -p 6379:6379 redis
```
### 2. Run the project
go run main.go

### 3. Test it
You can hit the endpoint repeatedly:

```bash
curl http://localhost:8080/home
```

After 10 requests (within 60 seconds), you’ll receive:

```json
{
  "error": "Too Many Requests"
}

```

⚙️ Configuration
You can modify rate limit logic by changing:
app.Use(RateLimiterMiddleware(redisClient, time.Minute, 10))


✨ Features
Rate limiting per IP

Redis-backed counter for persistence and performance

Retry-After header for graceful client handling

🙌 Contributing
Feel free to open issues or PRs if you'd like to improve this recipe!


🧠 Inspired by production needs where APIs require protection against brute-force abuse and overuse.
