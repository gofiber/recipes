#### Tech Stack

- Golang
- Redis

### API Documentation

> API endpoint: http://localhost:3000/api/v1/

#### API Payload

- "url" - Original URL
- "short" - Custom short URL(Optional)
- "expiry" - Time to expire: int(hours)

#### API Response

- "url" - Original URL
- "short" - Custom short URL
- "expiry" - Time to expire: int(hours)
- "rate_limit" - # of API calls remaining: int
- "rate_limit_reset" - Time to rate limit reset: int(minutes)

> API is rate limited to 10 calls every 30mins.
> These values can be changed in the .env file. Have fun.

#### Demo

![test.gif](test.gif)
