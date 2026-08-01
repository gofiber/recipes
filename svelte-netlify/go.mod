module github.com/gofiber/recipes/svelte-netlify

go 1.25.0

require (
	github.com/aws/aws-lambda-go v1.54.0
	github.com/awslabs/aws-lambda-go-api-proxy v0.16.2
	github.com/gofiber/fiber/v3 v3.4.0
	github.com/gofiber/utils/v2 v2.4.1
	github.com/valyala/fasthttp v1.73.0
)

require (
	github.com/andybalholm/brotli v1.2.2 // indirect
	github.com/gofiber/schema v1.8.3 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/klauspost/compress v1.19.1 // indirect
	github.com/mattn/go-colorable v0.1.15 // indirect
	github.com/mattn/go-isatty v0.0.24 // indirect
	github.com/philhofer/fwd v1.2.0 // indirect
	github.com/tinylib/msgp v1.6.4 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	golang.org/x/crypto v0.54.0 // indirect
	golang.org/x/net v0.57.0 // indirect
	golang.org/x/sys v0.47.0 // indirect
	golang.org/x/text v0.40.0 // indirect
)

// aws-lambda-go-api-proxy v0.16.2 pins github.com/gofiber/fiber/v2 v2.52.1 in its
// own go.mod. That version never reaches this build -- this recipe uses Fiber v3
// and only imports the proxy's /core package -- but it still shows up in the
// module graph. Pin that exact version to the patched release so scanners
// resolve v2.52.14.
//
// Scoped to v2.52.1 on purpose: an unversioned replace would also downgrade a
// legitimately newer Fiber v2 if some future dependency required one. If the
// proxy changes its pin, this stops applying and the new version shows up in
// scans again, which is the behaviour we want.
replace github.com/gofiber/fiber/v2 v2.52.1 => github.com/gofiber/fiber/v2 v2.52.14
