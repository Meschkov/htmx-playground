# htmx-playground
Small side project to play around with HTMX, Go
Inspired by: https://blog.arcjet.com/building-a-minimalist-web-server-using-the-go-standard-library-tailwind-css/

# Run the server
## Local development
Make sure you have Go installed and set up on your machine.

### 1. Install dependencies
 - [sops](https://getsops.io/docs/#encrypting-using-age)
 - [age](https://age-encryption.org/)
 - [air](https://github.com/air-verse/air)
 - [npm](https://www.npmjs.com/)

### 2. Decrypt or Encrypt the config file

Decrypt dev config file:
```shell
sops decrypt --age age1uw6wh5tdvywkt5mdwe3c8fuexpthhavswd2yyd8rcqhyllm3ranq8y25t0 ./dev.enc.yaml > tmp/dev.yaml
```

Encrypt dev config file:
```shell
sops encrypt --age age1uw6wh5tdvywkt5mdwe3c8fuexpthhavswd2yyd8rcqhyllm3ranq8y25t0 ./tmp/dev.yaml > dev.enc.yaml
```

### 3. Run the server
Make sure you have the `dev.yaml` file in the root directory.

```bash
go run main.go
```

Or install `air` and run it with live reload:
```bash
make dev
air
```

Installed dependencies:
- [htmx](https://htmx.org/)
- [tailwindcss](https://tailwindcss.com/)
- [air](https://github.com/air-verse/air)
- [go-sops](