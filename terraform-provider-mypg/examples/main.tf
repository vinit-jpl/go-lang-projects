terraform {
  required_providers {
    mypg = {
      source  = "example/mypg"
      version = "1.0.0"
    }
  }
}

provider "mypg" {}


resource "mypg_table" "users" {
  name   = "users"
  schema = "id SERIAL PRIMARY KEY, name TEXT"
}
