![CI](https://github.com/ValentinoFilipetto/pokedex-cli/actions/workflows/ci.yml/badge.svg)

# pokedex-cli

A small interactive REPL-style Pokedex CLI built in Go. It queries the public **PokeAPI** and lets you:
- page through location areas (`map`, `mapb`)
- explore a location area to list Pokémon encounters (`explore <location>`)
- try to catch a Pokémon (`catch <pokemon>`)
- inspect caught Pokémon (`inspect <pokemon>`)
- list your caught Pokémon (`pokedex`)

The client includes a lightweight in-memory cache to avoid re-fetching responses repeatedly.

## Requirements

- Go installed (the module declares `go 1.23.0`)

## Install dependencies

From the repo root:

```bash
go mod download
