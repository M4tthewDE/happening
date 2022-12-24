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

.PHONY: init-prod 
.ONESHELL:
init-prod:
	cd infra/prod
	@../../scripts/with-env.sh terraform init 

.ONESHELL:
.PHONY: plan-prod
plan-prod:
	cd infra/prod
	@../../scripts/with-env.sh terraform plan

.PHONY: apply-prod
.ONESHELL:
apply-prod:
	cd infra/prod
	@../../scripts/with-env.sh terraform apply --auto-approve 

.PHONY: destroy-prod
.ONESHELL:
destroy-prod:
	cd infra/prod
	@../../scripts/with-env.sh terraform destroy --auto-approve 
