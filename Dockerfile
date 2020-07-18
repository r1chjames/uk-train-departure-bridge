FROM alpine:latest
MAINTAINER "Richard James <richjames11@gmail.com>"

ENV TERRAFORM_VERSION=0.12.23
ENV TERRAFORM_SHA256SUM=b337c885526a8a653075551ac5363a09925ce9cf141f4e9a0d9f497842c85ad5

RUN apk add --update git bash wget make go nodejs npm py-pip python
RUN pip install --upgrade pip awscli
ADD https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip ./
ADD https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_SHA256SUMS ./

RUN sed -i '/terraform_${TERRAFORM_VERSION}_linux_amd64.zip/!d' terraform_${TERRAFORM_VERSION}_SHA256SUMS
RUN unzip terraform_${TERRAFORM_VERSION}_linux_amd64.zip -d /bin
RUN rm -f terraform_${TERRAFORM_VERSION}_linux_amd64.zip

ENTRYPOINT ["/bin/bash", "-c"]