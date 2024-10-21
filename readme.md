# RCON for Factorio
Go creates static binaries which are nicer than python package management.
put the RCON host address into the main.go and build (mine is on 192.168.0.25 but yours is likely different)

`go build -o rcon`

then run using any command from the [docs](https://wiki.factorio.com/console)

### Usage
`go install github.com/codabool/factorio_rcon@latest`

`./rcon PASSWORD COMMAND`

## Source Module
github.com/gtaylor/factorio-rcon
