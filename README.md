# STRAFES.NET API

## How to Begin Development

1. Run `go generate` to ensure the generated API is up-to-date. This project uses [ogen](https://github.com/ogen-go/ogen).
    ```bash
    go generate -run "go run github.com/ogen-go/ogen/cmd/ogen@latest --target api --clean openapi.yaml"
    ```
2. Build the project.
    ```bash
    go build git.itzana.me/strafesnet/public-api
    ```

    By default, the project opens at `localhost:8080`.