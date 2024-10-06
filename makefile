# Backend generation
generate-backend:
	oapi-codegen -generate types,server -package api -o backend/api.gen.go api-spec.yaml

# Frontend generation
generate-frontend:
	openapi-generator-cli generate -i api-spec.yaml -g typescript-angular -o frontend/src/app/api

# Main generate target to run both
generate: generate-backend generate-frontend
