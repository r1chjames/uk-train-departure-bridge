resource "aws_cloudwatch_event_rule" "execute_uk_train_depature_lambda_schedule" {
  name                = "execute_uk_train_depature_lambda_schedule"
  description         = "Schedule to fire uk_train_depature_lambda"
  schedule_expression = "rate(${var.refresh_interval_in_mins})"
}

resource "aws_cloudwatch_event_target" "execute_uk_train_depature_lambda_schedule" {
  rule      = aws_cloudwatch_event_rule.execute_uk_train_depature_lambda_schedule.name
  target_id = "lambda"
  arn       = aws_lambda_function.uk_train_departure_lambda.arn
}

resource "aws_lambda_permission" "allow_cloudwatch_to_call_uk_train_departure_lambda" {
  statement_id  = "AllowExecutionFromCloudWatch"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.uk_train_departure_lambda.function_name
  principal     = "events.amazonaws.com"
  source_arn    = aws_cloudwatch_event_rule.execute_uk_train_depature_lambda_schedule.arn
}