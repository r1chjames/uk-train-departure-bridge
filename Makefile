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

move_local_dev_artifacts:
	rm -rf build/*
	cp ../visitor-webapp/webapp.tar.gz build
	cp ../go-apps/src/visitor-lambdas/*.zip build
	make unpack_webapp

move_artifacts:
	mkdir build
	mv visitor-lambdas/*.zip build/
	mv visitor-webapp/*.tar.gz build/

build_lambdas:
	git clone --single-branch --branch ${CI_COMMIT_REF_NAME} https://gitlab-ci-token:${CI_JOB_TOKEN}@gitlab.com/r1chjames/visitor-lambdas.git
	cd visitor-lambdas && \
	echo 'Using branch $(shell git rev-parse --abbrev-ref HEAD)' && \
	make build_all

build_webapp:
	git clone --single-branch --branch ${CI_COMMIT_REF_NAME} https://gitlab-ci-token:${CI_JOB_TOKEN}@gitlab.com/r1chjames/visitor-webapp.git
	cd visitor-webapp && \
	echo 'Using branch $(shell git rev-parse --abbrev-ref HEAD)' && \
	make build_all

build_artifacts: build_lambdas build_webapp move_artifacts

unpack_webapp:
	mkdir build/webapp && \
	tar -xvzf build/webapp.tar.gz -C build/webapp

build_tf_image:
	docker build -t registry.gitlab.com/r1chjames/visitor-terraform/terraform_runner:0.12.23 .
	docker login -u gitlab-ci-token -p $CI_BUILD_TOKEN registry.gitlab.com
	docker push registry.gitlab.com/r1chjames/visitor-terraform/terraform_runner:0.12.23