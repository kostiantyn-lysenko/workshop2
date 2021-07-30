test:
	go test -cover ./...

test-report:
	go tool cover -html=cover.txt
	go test -cover -covermode=count -coverprofile=cover.txt ./...

migrate-up:
	migrate -source file://./schema -database "postgres://postgres:@localhost:5432/postgres?sslmode=disable&password=password" up