terraform {
  required_providers {
    mypg = {
      source  = "local/local/mypg"
      version = "1.0.0"
    }
  }
}

provider "mypg" {
  host     = "localhost"
  port     = 5432
  username = "postgres"
  password = "postgres@123"
  database = "postgres"
}

resource "mypg_table" "example" {
  name     = "users"
  columns  = ["id SERIAL PRIMARY KEY", "name TEXT", "email TEXT"]
  host     = "localhost"
  port     = 5432
  username = "postgres"
  password = "postgres@123"
  database = "postgres"
}
