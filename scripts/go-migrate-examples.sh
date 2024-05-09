# Apply latest migration
migrate -database $DATABASE_URL -path ./db/migrations up 1

# Donwgrade last migration
migrate -database $DATABASE_URL -path ./db/migrations down 1

# Generate a new migration
migrate create -seq -ext sql -dir db/migrations $MIGRATION_NAME