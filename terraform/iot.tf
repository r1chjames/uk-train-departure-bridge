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
  name = "iot_client_subscribe_policy"
  policy = templatefile("${path.module}/templates/iot_client_policy.json.tpl", {
    account_id       = data.aws_caller_identity.current.account_id
    device           = aws_iot_thing.platform_sign_1.name
    departures_topic = var.departures_topic
    commands_topic   = var.commands_topic
  })
}

resource "aws_iot_policy_attachment" "client_subscribe" {
  policy = aws_iot_policy.iot_client_subscribe_policy.name
  target = aws_iot_certificate.platform_sign_1.arn
}

output "platform_sign_1_certificate_pem" {
  value     = aws_iot_certificate.platform_sign_1.certificate_pem
  sensitive = true
}

output "platform_sign_1_public_key" {
  value     = aws_iot_certificate.platform_sign_1.public_key
  sensitive = true
}

output "platform_sign_1_private_key" {
  value     = aws_iot_certificate.platform_sign_1.private_key
  sensitive = true
}