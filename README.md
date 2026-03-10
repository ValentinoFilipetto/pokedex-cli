![CI](https://github.com/ValentinoFilipetto/pokedex-cli/actions/workflows/ci.yml/badge.svg)

# pokedex-cli

A REPL-style Pokedex CLI written in Go. It uses the public **PokeAPI** to:
- page through location areas (`map`, `mapb`)
- explore a location area (`explore <location-area>`)
- catch Pokémon (`catch <pokemon-name>`)
- inspect caught Pokémon (`inspect <pokemon-name>`)
- list your caught Pokémon (`pokedex`)

To reduce repeated API calls, some endpoints are cached in memory with a TTL.

## Requirements

- Go installed

## Install

```bash
go mod download
```

## Run

```bash
go run .
```

Then type `help` to see commands.

## Commands (overview)

- `help` — show available commands and descriptions
- `map` — list the next page of location areas
- `mapb` — list the previous page of location areas
- `explore <location-area>` — list Pokémon found in a given location area  
  Example:
  ```bash
  explore canalave-city-area
  ```
- `catch <pokemon-name>` — attempt to catch a Pokémon (random chance based on base XP)  
  Example:
  ```bash
  catch pikachu
  ```
- `inspect <pokemon-name>` — show details of a caught Pokémon
- `pokedex` — list caught Pokémon in the current session
- `exit` — quit

## Tests

```bash
go test ./...
```

## Implementation notes

- `internal/pokeapi` wraps HTTP calls to PokeAPI.
- `internal/pokecache` is a small TTL cache with a reap loop.
- The Pokedex is stored in memory (no persistence between runs).
