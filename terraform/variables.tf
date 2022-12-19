variable "env" {
  type    = string
  default = "DEV"
}

variable departure_station {
  type = string
}

variable destination_station {
  type = string
}

variable future_train_count {
  type = number
}

variable departures_topic {
  type = string
}

variable commands_topic {
  type = string
}

variable refresh_interval_in_mins {
  type = string
}

variable api_username {
  type = string
}

variable api_password {
  type = string
}