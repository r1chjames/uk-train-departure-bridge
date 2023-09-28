resource "aws_lambda_function" "uk_train_departure_lambda" {
  filename      = "build/uk-train-departure-lambda.zip"
  function_name = "${lower(var.env)}-uk-train-departure-lambda"
  role          = aws_iam_role.uk_train_departure_bridge_lambda_iam_execution_role.arn
  handler       = "main"
  runtime       = "provided.al2"
  memory_size   = 128

  environment {
    variables = {
      DEPARTURE_STATION   = upper(var.departure_station)
      DESTINATION_STATION = upper(var.destination_station)
      FUTURE_TRAIN_COUNT  = var.future_train_count
      IOT_TOPIC           = var.departures_topic
      USERNAME            = var.api_username
      PASSWORD            = var.api_password
    }
  }

  source_code_hash = filebase64sha256("build/uk-train-departure-lambda.zip")
}

data "template_file" "iot_client_write_policy_template" {
  template = file("${path.module}/templates/lambda_iot_write_policy.json.tpl")
  vars = {
    account_id = data.aws_caller_identity.current.account_id
  }
}

resource "aws_iam_policy" "iot_write_policy" {
  name   = "iot_write_policy"
  path   = "/"
  policy = data.template_file.iot_client_write_policy_template.rendered
}

data "template_file" "coworker_lambda_iam_execution_role" {
  template = file("${path.module}/templates/lambda_iam_execution_role.json.tpl")
}

resource "aws_iam_role" "uk_train_departure_bridge_lambda_iam_execution_role" {
  name               = "uk_train_departure_bridge_lambda_iam_execution_role"
  assume_role_policy = data.template_file.coworker_lambda_iam_execution_role.rendered
}

resource "aws_iam_policy_attachment" "iot_client_write_policy" {
  name       = "iot_client_write_policy"
  roles      = [aws_iam_role.uk_train_departure_bridge_lambda_iam_execution_role.name]
  policy_arn = aws_iam_policy.iot_write_policy.arn
}
