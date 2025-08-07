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
  password = "your_password"
  database = "your_db"
}

resource "mypg_table" "example" {
  name     = "users"
  columns  = ["id SERIAL PRIMARY KEY", "name TEXT", "email TEXT"]
  host     = "localhost"
  port     = 5432
  username = "postgres"
  password = "your_password"
  database = "your_db"
}
