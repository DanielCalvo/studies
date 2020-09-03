provider "http" {
}

data "http" "weather" {
  url = "https://www.metaweather.com/api/location/search/?lattlong=36.03,-84.03"

  request_headers = {
    "Accept" = "application/json"
  }
}

output "nashville" {
    value = jsondecode(data.http.weather.body)[0]
}