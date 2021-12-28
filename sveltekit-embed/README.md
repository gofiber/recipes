# Fiber SvelteKit Embed

A sample application to [embed](https://golang.org/pkg/embed/) [SvelteKit](https://kit.svelte.dev/) app with [fiber filesystem middleware](https://github.com/gofiber/fiber/tree/master/middleware/).

## Requirement

- Go 1.16
- Node 14

## Local usage

Inside the root of application directory.

Install node modules
```bash
npm install --prefix ./frontend
```

Build static HTML
```bash
npm run build --prefix ./frontend
```

Build the Go binary
```bash
go build main.go
```

Run the app
```bash
./main
```
Open http://localhost:8080/ in the browser.

## Docker usage

### With makefile
Build the container
```bash
make build
```

Run the container
```bash
make run
```

Open http://localhost:8080/ in the browser.

### With docker command

Make it sure the static HTML already generated inside the public folder and then run:
```bash
docker build . -t sveltekit:latest
```

Run the container
```bash
docker run -d -p 8080:8080 --name sveltekit sveltekit:latest

```

Open http://localhost:8080/ in the browser.