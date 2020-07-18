init:
	rm -rf .terraform
	terraform --version
	mkdir -p ./creds
	echo ${SERVICEACCOUNT} | base64 -d > ./creds/serviceaccount.json
	terraform init

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