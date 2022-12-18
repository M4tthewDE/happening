.PHONY: build 
build:
	GOOS=linux CGO_ENABLED=0 go build -o ./bin/api ./backend/api/cmd/...

.PHONY: init 
init:
	@./scripts/with-env.sh terraform init 

.PHONY: plan 
plan:
	@./scripts/with-env.sh terraform plan 

.PHONY: apply 
apply:
	@./scripts/with-env.sh terraform apply --auto-approve 

.PHONY: destroy
destroy:
	@./scripts/with-env.sh terraform destroy --auto-approve 
