# Define colors
GREEN=\033[0;32m
NOCOLOR=\033[0m
# Define header
HEADER=$(GREEN)Fiber Recipe$(NOCOLOR)

.PHONY: 404-handler

404-handler:
	@echo "$(HEADER): Custom 404 Not Found"
	@echo "Start server on http://localhost:3000"
	@go run ./404-handler/*.go

.PHONY: file-server

file-server:
	@echo "$(HEADER): Serve Static Files"
	@echo "Start server on http://localhost:3000"
	@go run ./file-server/*.go

.PHONY: hello-world

hello-world:
	@echo "$(HEADER): Hello, World!"
	@echo "Start server on http://localhost:3000"
	@go run ./hello-world/*.go


.PHONY: https-tls

https-tls:
	@echo "$(HEADER): Enable HTTPS/TLS"
	@echo "Start server on http://localhost:443"
	@go run ./https-tls/*.go

.PHONY: multiple-ports

multiple-ports:
	@echo "$(HEADER): Listen on Multiple Ports"
	@echo "Start servers on 8080, 8081 and 3000 ports"
	@go run ./multiple-ports/*.go

.PHONY: prefork

prefork:
	@echo "$(HEADER): Enable Preforking"
	@echo "Start server on http://localhost:3000"
	@echo "Run the following command on another console to see all processes sharing port 3000:"
	@echo "sudo lsof -i -P -n | grep LISTEN"
	@go run ./prefork/*.go

.PHONY: upload-file-single

upload-file-single:
	@echo "$(HEADER): Upload Single File"
	@echo "Start server on http://localhost:3000"
	@go run ./upload-file/single/*.go

.PHONY: upload-file-multiple

upload-file-multiple:
	@echo "$(HEADER): Upload Multiple Files"
	@echo "Start server on http://localhost:3000"
	@go run ./upload-file/multiple/*.go
