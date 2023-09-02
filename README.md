# AsiaYo

![Go](https://img.shields.io/badge/Go-v1.20.3-blue)

## Setup Environment

Use [asdf](https://asdf-vm.com/) to manage golang version.

- installation

```bash
brew install asdf
```

- configure zsh

```bash
echo -e "\n. $(brew --prefix asdf)/libexec/asdf.sh" >> ${ZDOTDIR:-~}/.zshrc
```

- install golang plugin

```bash
asdf plugin add golang
```

- install a version

```bash
asdf install golang 1.20.3
```

## Configuration

The default configuration is written in `.env.example`.

Copy it and then modify `.env` if needed.

```bash
cp .env.example .env
docker-compose up -d
```

## Development

- install library

```bash
go mod download
```

- start the server

```bash
go run main.go
```

## Document

<https://documenter.getpostman.com/view/4926779/2s9Y5cu1MW>

## Testing

```sh
go test test/converseCurrency_test.go
```

## Playground

You can write some code for debugging here.

And it will not be tracked by `git`.
