# Terraform Provider: mypg (Custom PostgreSQL Provider)

This is a custom Terraform provider that allows you to manage **PostgreSQL tables** directly through Terraform.

---


### 1. Clone and Build the Provider

```bash
git clone https://github.com/vinit-jpl/go-lang-projects.git
cd terraform-provider-mypg
go mod tidy
go build -o terraform-provider-mypg.exe
```

### 2. Move the Provider Binary to Local Plugin Path

```bash
mkdir $env:APPDATA\terraform.d\plugins\local\local\mypg\1.0.0\windows_amd64
Copy-Item "path/to/provider/binary" "$env:APPDATA\terraform.d\plugins\local\local\mypg\1.0.0\windows_amd64\"
```

### Run terraform

```bash
cd examples
terraform init
terraform plan
terraform apply
```

