resource "aws_lambda_function" "uk_train_departure_bridge" {
  filename      = "../lambda/uk_train_departure_bridge.zip"
  function_name = "${lower(var.env)}-uk-train-departure-bridge-lambda"
  role          = aws_iam_role.uk_train_departure_bridge_lambda_iam_execution_role.arn
  handler       = "main"
  runtime       = "go1.x"
  memory_size   = 128

  environment {
    variables = {
      DYNAMODB_TABLE_PREFIX = lower(var.env)
      AWS_ACCOUNT_ID = var.aws_account_id
      ALLOWED_CORS_ORIGINS = var.allowed_cors_origins
    }
  }

  source_code_hash = filebase64sha256("../lambda/uk_train_departure_bridge.zip")
}

data "template_file" "uk_train_departure_bridge_lambda_iam_execution_role" {
  template = file("${path.module}/templates/uk_train_departure_bridge_lambda_iam_execution_role.json.tpl")
}

resource "aws_iam_role" "uk_train_departure_bridge_lambda_iam_execution_role" {
  name               = "uk_train_departure_bridge_lambda_iam_execution_role"
  assume_role_policy = data.template_file.uk_train_departure_bridge_lambda_iam_execution_role.rendered
}

resource "aws_iam_role_policy_attachment" "uk_train_departure_bridge_lambda_log_iam_policy_attachment" {
  role       = aws_iam_role.uk_train_departure_bridge_lambda_iam_execution_role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

data "template_file" "uk_train_departure_bridge_lambda_iot_write_policy_template" {
  template = file("${path.module}/templates/uk_train_departure_bridge_lambda_iot_write_policy.json.tpl")
  vars = {
    account_id = var.aws_account_id
  }
}

resource "aws_iam_policy" "uk_train_departure_bridge_lambda_iot_write_policy" {
  name   = "uk_train_departure_bridge_lambda_iot_write_policy"
  path   = "/"
  policy = data.template_file.uk_train_departure_bridge_lambda_iot_write_policy_template.rendered
}

resource "aws_iam_policy_attachment" "uk_train_departure_bridge_lambda_iot_write_policy" {
  name       = "uk_train_departure_bridge_lambda_iot_write_policy"
  roles      = [aws_iam_role.uk_train_departure_bridge_lambda_iam_execution_role.name]
  policy_arn = aws_iam_policy.uk_train_departure_bridge_lambda_iot_write_policy.arn
}
