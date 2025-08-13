include .env
# oapi path's
OAPI_CONFIG=./openAPI/oapi-config.yaml
OAPI_OUTPUT=./internal/oGenerated/openapi.gen.go
OAPI_CLEAN=./internal/oGenerated/*.gen.go
OAPI_SPEC=./openAPI/openapi.yaml

# DB url & migrations path
DB_URL=postgres://$(PS_USER):$(PS_PASS)@$(DB_HOST):$(DB_PORT)/$(PS_DB)?sslmode=disable
MIGRATIONS-DIR=./migrations

# migrations
migrate-new:
	migrate create -ext sql -dir $(MIGRATIONS-DIR) -seq $(name)
migrate-up:
	migrate -path $(MIGRATIONS_DIR)	-database "$(DB_URL)" -verbose up
migrate-down:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" -verbose down
migrate-version:
	migrate -path $(MIGRATIONS-DIR) -database "$(DB_URL)" -verbose version 
migrate-force:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" -verbose force 1

# Gen oapi
gen:
	oapi-codegen -config $(OAPI_CONFIG) -include-tags $(TAG) -package $(PACAGE) $(OAPI_SPEC) > $(OAPI_OUTPUT)
cleang:
	rm -v $(OAPI_CLEAN)

# go mod
gmod:
	go mod tidy

	