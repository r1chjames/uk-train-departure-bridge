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

copy_artifacts:
	rm -rf build/*
	cp lambda/*.zip build

build_lambda:
	cd lambda && \
	make build_all

build_artifacts: build_lambda copy_artifacts

build_tf_image:
	docker build -t registry.gitlab.com/r1chjames/visitor-terraform/terraform_runner:0.12.23 .
	docker login -u gitlab-ci-token -p $CI_BUILD_TOKEN registry.gitlab.com
	docker push registry.gitlab.com/r1chjames/visitor-terraform/terraform_runner:0.12.23