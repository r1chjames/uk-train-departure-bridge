data "template_file" "iot_client_subscribe_policy_template" {
  template = file("${path.module}/templates/iot_client_policy.json.tpl")
  vars = {
    account_id = var.aws_account_id
  }
}

resource "aws_iam_policy" "iot_client_subscribe_policy" {
  name   = "iot_client_subscribe_policy"
  path   = "/"
  policy = data.template_file.iot_client_subscribe_policy_template.rendered
}

resource "aws_iam_policy_attachment" "iot_client_subscribe_policy" {
  name       = "iot_client_subscribe_policy"
  roles      = [aws_iam_role.uk_train_departure_bridge_lambda_iam_execution_role.name]
  policy_arn = aws_iam_policy.iot_client_subscribe_policy.arn
}
