.PHONY: swagger cmd clean tools

SWAGGER_BIN ?= tools/swagger
SWAGGER_VERSION ?= "v0.27.0"

swagger: clean
	mkdir -p src/main_service/swagger
	swagger generate server \
	  -m models_server -t src/main_service/swagger -A test -f api_swagger.yml --template=stratoscale
	swagger generate client \
	  -c client_users -m models_users -f external_api_swagger.yml --tags=main_service \
	  -t src/main_service/swagger src/main_service/swagger --template=stratoscale

cmd: swagger
	go build ./cmd/main_service/...

clean:
	rm -rf src/main_service/swagger

tools:
	GOBIN=$(GOBIN) go install github.com/go-swagger/go-swagger/cmd/swagger@$(SWAGGER_VERSION)
