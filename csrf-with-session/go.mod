module main

go 1.21

toolchain go1.22.3

require (
	github.com/gofiber/fiber/v3 v3.0.0-beta.2
	github.com/gofiber/template/html/v2 v2.1.1
	golang.org/x/crypto v0.21.0
)

replace github.com/gofiber/fiber/v3 => github.com/sixcolors/fiber/v3 v3.0.0-20240528203118-6e76847f88aa

require (
	github.com/andybalholm/brotli v1.1.0 // indirect
	github.com/gofiber/template v1.8.3 // indirect
	github.com/gofiber/utils v1.1.0 // indirect
	github.com/gofiber/utils/v2 v2.0.0-beta.4 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/klauspost/compress v1.17.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/philhofer/fwd v1.1.2 // indirect
	github.com/tinylib/msgp v1.1.8 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.54.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.18.0 // indirect
)
