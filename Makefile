build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ./bin/api ./backend/api/...

init:
	terraform init 

plan:
	terraform plan 

apply:
	terraform apply --auto-approve 

destroy:
	terraform destroy --auto-approve 