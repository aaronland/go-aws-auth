cli:
	go build -mod vendor -o bin/aws-mfa-session cmd/aws-mfa-session/main.go
	go build -mod vendor -o bin/aws-get-credentials cmd/aws-get-credentials/main.go
	go build -mod vendor -o bin/aws-set-env cmd/aws-set-env/main.go
