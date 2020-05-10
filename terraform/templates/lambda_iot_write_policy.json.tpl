{
    "Version": "2012-10-17",
    "Statement": [{
        "Effect": "Allow",
        "Action": "iot:Publish",
        "Resource": "arn:aws:iot:eu-west-1:${account_id}:topic/departures/train/new"
    }]
}