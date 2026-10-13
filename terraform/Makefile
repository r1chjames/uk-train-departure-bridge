init:
	rm -rf .terraform
	terraform --version
	terraform init -input=false

plan:
	terraform plan \
		-var-file ${ENVIRONMENT}.tfvars \
		-out "planfile" \
		-var env=${ENVIRONMENT}

apply: plan
	terraform apply -input=false "planfile"

validate:
	terraform validate

fmt:
	terraform fmt

lint: fmt