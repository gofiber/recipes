---
title: React
keywords: [react, react-router, client-side, spa, docker]
---

# React Fiber

A sample application to showcase serving React (with Router) with an almost bare Fiber. Hopefully, this application can be of use (as a reference or others) for those who wants to serve their client-side SPA with Fiber.

## Technologies

- Go with Fiber
- React with TypeScript and React Router
- Docker

## Application

- This application has three routes: `/`, `/react`, and a catch-all, 404 route. `/` will show the Fiber logo, `/react` will show the React logo, and the 404 route will show both logos.
- As this application serves the frontend while backed by a server, the client-side routing will work well and will not cause any issue (unlike if you are running without a file server). You can type the URL route manually in the browser and it will still work and will render the accurate page, so no worries.
- This is a simplified form of Create React App with TypeScript. With that being said, that's why there is no `manifest.json`, `logo512.png`, and other extra things like that.
- I restructured the project structure to be a bit more modular by categorizing files to `assets`, `components`, and `styles`. I also made it so all of the CSS is loaded in `index.tsx` for easier seeing.
- I also moved several dependencies to their appropriate places, such as `@types` and `test` in development dependencies instead of dependencies.

## Installation

It is recommended that you use Docker to instantly run this application. After running the Docker application, please open `localhost:8080` in your browser. Make sure you are in the `react-router` folder before running these commands.

```bash
docker build . -t react-router:latest
docker run -d -p 8080:8080 react-router:latest
```

If you prefer doing things manually, then the installation steps are as follows:

- Clone the repository by using `git clone git@github.com:gofiber/recipes.git`.
- Switch to the application by using `cd recipes/react-router`.
- Install npm dependencies by using `cd web && yarn install`.
- Build frontend by using `yarn build`.
- Run the Fiber application by using `go run cmd/react-router/main.go`. Don't forget to return to the main repository by using `cd ..` (assuming you are in `web` folder).
- Open `localhost:8080` in your browser.
