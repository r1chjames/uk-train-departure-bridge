# UK Train Departure Bridge

This repo contains a set of deployable AWS components that provide the ability to perform API requests to the UK Realtime Trains API and write the data to an IoT (MQTT) topic to allow simple IoT devices such as the ESP32 to subscribe and receive this data.

Inspired by https://github.com/balenalabs/uk-train-departure-display

***

## Installation
The AWS components are deployed using Terraform which can be executed using the included .gitlab-ci.yaml file or using a local Terraform installation.

### Prerequisites
* AWS account with an IAM user created that has permissions that allow Terraform to create relevant resources. Follow the tutorial [here](https://medium.com/slalom-technology/creating-your-first-terraform-infrastructure-on-aws-ad986f952951).
* Realtime Trains API. Sign up for an account and request a token on the [API portal](https://api-portal.rtt.io). The API is rate limited (currently 10 requests per minute, 100 per hour, 1000 per day).

### Setup
* Four variables are required:

| Variable | Description | Example |
| -------- | ----------- | ------- |
| AWS_ACCESS_KEY_ID | Obtained from IAM when creating the Terraform user | AKIAIOSFODNN7EXAMPLE |
| AWS_SECRET_ACCESS_KEY | Obtained from IAM when creating the Terraform user | wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY |
| ENVIRONMENT | Informs Terraform which .tfvars file to use. Just set to DEV if you're using the defaults | DEV |
| TF_VAR_api_token | Terraform variable that contains the token issued by the Realtime Trains API portal. The portal issues a refresh token, which the Lambda exchanges for a short life access token on each run | <token-from-api-portal> |

If you are using GitLab CI, these will need to be set up as CI variables. If you are deploying using Terraform locally, these need to be set up as environment variables or passed into the running process.

### Build
* GitLab CI will automatically build, validate, and plan the Terraform against the provided AWS account. The apply step is manual and needs to be instantiated once the previous steps have completed. This will automatically provision all the required infrastructure components on AWS.

* A local Terraform apply will perform the same but is run by executing `make plan` and `make apply` from the repository root.

* A successful Terraform apply will print out the certificate, private key, and public keys of the IoT client that can then be used to connect and subscribe to the IoT MQTT topic. An example of how to do this can be found [here](https://aws.amazon.com/blogs/compute/building-an-aws-iot-core-device-using-aws-serverless-and-an-esp32/)


## To do
* Provide Arduino sketch that subscribes to MQTT topic and displays train info in a format similar to the UK train platform signs.
&nbsp;
&nbsp;
&nbsp;
***

*Disclaimer: Make sure you set up billing alarms on your AWS account as I can't be held responsible for your AWS bill. See Amazon's article [here](https://docs.aws.amazon.com/AmazonCloudWatch/latest/monitoring/monitor_estimated_charges_with_cloudwatch.html).*
