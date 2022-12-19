init:
	cd terraform && \
	make init

plan:
	cd terraform && \
	make plan

apply: plan
	cd terraform && \
	make apply

validate:
	cd terraform && \
	make validate

fmt:
	cd terraform && \
	make fmt

lint: fmt

copy_artifacts:
	rm -rf terraform/build/
	mkdir -p terraform/build
	zip -j terraform/build/uk-train-departure-lambda.zip lambda/bin/main
	
build_lambda:
	cd lambda && \
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 make compile

prepare_artifacts: build_lambda copy_artifacts

build_tf_image:
	docker build -t registry.gitlab.com/r1chjames/uk-train-departure-bridge/terraform_runner:0.12.23 .
	docker login -u gitlab-ci-token -p $CI_BUILD_TOKEN registry.gitlab.com
	docker push registry.gitlab.com/r1chjames/uk-train-departure-bridge/terraform_runner:0.12.23