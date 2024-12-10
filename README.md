# STRAFES.NET API

## How to Begin Development

1. Run `go generate` to ensure the generated API is up-to-date. This project uses [ogen](https://github.com/ogen-go/ogen).
    ```bash
    go generate -run "go run github.com/ogen-go/ogen/cmd/ogen@latest --target api --clean openapi.yaml"
    ```
2. Build the project.
    ```bash
    go build git.itzana.me/strafesnet/maps-service
    ```

    By default, the project opens at `localhost:8080`.

## How to generate rust api from this directory
```bash
openapi-generator-cli generate -g rust -i openapi.yaml -o validation/api`
```

#### License

<sup>
Licensed under <a href="LICENSE">MIT license</a>.
</sup>

<br>

<sub>
Unless you explicitly state otherwise, any contribution intentionally submitted
for inclusion in this project by you, shall be licensed as above, without any
additional terms or conditions.
</sub>
