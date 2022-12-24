.PHONY: build 
build:
	GOOS=linux CGO_ENABLED=0 go build -o ./bin/api ./backend/api/cmd/...
	GOOS=linux CGO_ENABLED=0 go build -o ./bin/auth ./backend/auth/...

.PHONY: init 
.ONESHELL:
init:
	cd infra/beta
	@../../scripts/with-env.sh terraform init 

.ONESHELL:
.PHONY: plan 
plan:
	cd infra/beta
	@../../scripts/with-env.sh terraform plan

.PHONY: apply 
.ONESHELL:
apply:
	cd infra/beta
	@../../scripts/with-env.sh terraform apply --auto-approve 

.PHONY: destroy
.ONESHELL:
destroy:
	cd infra/beta
	@../../scripts/with-env.sh terraform destroy --auto-approve 
