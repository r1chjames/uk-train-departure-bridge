resource "aws_iot_thing" "platform_sign_1" {
  name = "platform_sign_1"

  attributes = {
    DeviceType = "esp32"
  }
}

resource "aws_iot_certificate" "platform_sign_1" {
  active = true
}

resource "aws_iot_thing_principal_attachment" "platform_sign_1" {
  principal = aws_iot_certificate.platform_sign_1.arn
  thing     = aws_iot_thing.platform_sign_1.name
}

resource "aws_iot_policy" "iot_client_subscribe_policy" {
  name   = "iot_client_subscribe_policy"
  policy = data.template_file.iot_client_subscribe_policy_template.rendered
}

resource "aws_iot_policy_attachment" "client_subscribe" {
  policy = aws_iot_policy.iot_client_subscribe_policy.name
  target = aws_iot_certificate.platform_sign_1.arn
}

data "template_file" "iot_client_subscribe_policy_template" {
  template = file("${path.module}/templates/iot_client_policy.json.tpl")
  vars = {
    account_id = data.aws_caller_identity.current.account_id
    device = aws_iot_thing.platform_sign_1.name
  }
}

output "platform_sign_1_certificate_pem" {
  value = aws_iot_certificate.platform_sign_1.certificate_pem
}

output "platform_sign_1_public_key" {
  value = aws_iot_certificate.platform_sign_1.public_key
}

output "platform_sign_1_private_key" {
  value = aws_iot_certificate.platform_sign_1.private_key
}