# Stage 1: Build the static HTML
FROM node:16-alpine3.14 as frontend-builder
WORKDIR /builder
COPY /frontend/package.json /frontend/package-lock.json ./frontend/
RUN npm install --prefix ./frontend
COPY ./frontend/ ./frontend/
RUN npm run build --prefix ./frontend

# Stage 2: Build the go static binary
FROM golang:1.17.5-alpine3.15 AS server-builder
RUN apk update && apk upgrade && \
  apk --update add git
WORKDIR /builder
COPY --from=frontend-builder /builder/public ./public/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
  -ldflags='-w -s -extldflags "-static"' -a \
  -o sveltekit main.go

# Stage 3: Final
FROM gcr.io/distroless/static
WORKDIR /app
COPY --from=server-builder --chown=nonroot:nonroot /builder/sveltekit .
EXPOSE 8080
ENTRYPOINT ["./sveltekit"]
