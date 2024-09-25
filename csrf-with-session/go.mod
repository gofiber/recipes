module main

go 1.22

toolchain go1.23.1

require (
	github.com/gofiber/fiber/v3 v3.0.0-beta.3
	github.com/gofiber/template/html/v2 v2.1.2
	golang.org/x/crypto v0.27.0
)

replace github.com/gofiber/fiber/v3 => github.com/sixcolors/fiber/v3 v3.0.0-20240925170926-7765ee557737

require (
	github.com/andybalholm/brotli v1.1.0 // indirect
	github.com/gofiber/template v1.8.3 // indirect
	github.com/gofiber/utils v1.1.0 // indirect
	github.com/gofiber/utils/v2 v2.0.0-beta.6 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/klauspost/compress v1.17.10 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/philhofer/fwd v1.1.3-0.20240612014219-fbbf4953d986 // indirect
	github.com/tinylib/msgp v1.2.1 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.56.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/sys v0.25.0 // indirect
)
