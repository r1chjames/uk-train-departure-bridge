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

variable refresh_interval_sec {
  type = number
}

variable operating_hours {
  type = string
}

variable transport_api_key {
  type = string
}

variable transport_api_app_id {
  type = string
}