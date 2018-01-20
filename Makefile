compile:
	GOOS=linux go build -o bin/create functions/create/main.go
	GOOS=linux go build -o bin/delete functions/delete/main.go
.PHONY: compile