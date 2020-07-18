terraform {
  backend "s3" {
    bucket = "uk-train-departure-tf-state"
    key    = "state/terraform-state"
    region = "eu-west-1"
  }
}
