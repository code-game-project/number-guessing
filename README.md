# number_guessing
![CodeGame Version](https://img.shields.io/badge/CodeGame-v0.6-orange)
![CodeGame GameServer Version](https://img.shields.io/badge/GameServer-v0.1-yellow)
![CGE Version](https://img.shields.io/badge/CGE-v0.3-green)

A number guessing game.

## Known instances

- `games.code-game.org/number_guessing`

## Usage

```sh
# Run on default port 80
number_guessing

# Specify a custom port
number_guessing --port=8080

## Specify a custom port through an environment variable
CG_PORT=8080 number_guessing
```

### Running with Docker

Prerequisites:
- [Docker](https://docker.com/)

```sh
# Download image
docker pull codegameproject/number_guessing:0.1

# Run container
docker run -d -p <port-on-host-machine>:8080 --name number_guessing codegameproject/number_guessing:0.1
```

## Event Flow

1. TODO

## Building

### Prerequisites

- [Go](https://go.dev) 1.18+

```sh
git clone https://github.com/code-game-project/number_guessing.git
cd number_guessing
go build .
```
## License

Copyright (C) 2022 Julian Hofmann

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published
by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
