# Backend generation
generate-backend:
	- rm backend/gen/api.gen.go
	- mkdir -p backend/gen
	oapi-codegen -generate types,server -package api -o backend/gen/api.gen.go api-spec.yaml

# Frontend generation
generate-frontend:
	- rm -rf frontend/src/app/gen
	openapi-generator-cli generate -i api-spec.yaml -g typescript-angular -o frontend/src/app/gen

# Main generate target to run both
generate: generate-backend generate-frontend

# Docker compose up
up:
	docker-compose up --build

# Make a new clean target
clean:
	docker-compose down --rmi all
