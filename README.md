# STRAFES.NET Map Submission System

## Components
- Submissions API (golang) `pkg/` `cmd/`
- Website `/web/`
- Script Validation (rust) `validation/`

## How to Begin Development on Each Component

### Submissions API

Prerequisite: golang installed

1. Run `go generate` to ensure the generated API is up-to-date. This project uses [ogen](https://github.com/ogen-go/ogen).
    ```bash
    go generate -run "go run github.com/ogen-go/ogen/cmd/ogen@latest --target api --clean openapi.yaml"
    ```
2. Build the project.
    ```bash
    go build git.itzana.me/strafesnet/maps-service
    ```

    By default, the project opens at `localhost:8080`.

### Website

Prerequisite: bun installed

1. `cd web`
2. `bun install`

#### For development:
3. `bun run dev`

#### For production:
3. `bun run build`
4. `bun run start` (optionally start a node server)

### Script Validation

Prerequisite: rust installed

1. `cd validation`
2. `cargo run --release`

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
