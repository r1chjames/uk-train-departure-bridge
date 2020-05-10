resource "aws_iot_thing" "iot_client" {
  name = "iot_client"

  attributes = {
    DeviceType = "esp32"
  }
}