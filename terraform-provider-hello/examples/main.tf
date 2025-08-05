terraform {
  required_providers {
    hello = {
      source  = "example.com/hello/hello"
      version = "1.0.0"
    }
  }
}

provider "hello" {}

resource "hello_hello" "example" {
  message = "Hello from custom provider!"
}
