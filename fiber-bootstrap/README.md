# Fiber Bootstrap
Fiber bootstrap for rapid development using Go-Fiber / Gorm / Validator.

# Components
* Fiber
  * Html Engine Template
  * Logger
  * Monitoring
* Gorm
  * PGSQL Driver
* Validator  
* Env File

# Router 
API Router `/api` with rate limiter middleware  
Http Router `/` with CORS and CSRF middleware  

# Setup

1. Copy the example env file over:
    ```
    cp .env.example .env
    ```

2. Modify the env file you just copied `.env` with the correct credentials for your database. Make sure the database you entered in `DB_NAME` has been created.

3. Run the API:
    ```
    go run main.go
    ```
Your api should be running at `http://localhost:4000/` if the port is in use you may modify it in the `.env` you just created.