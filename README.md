# Number Guessing
![CodeGame Version](https://img.shields.io/badge/CodeGame-v0.7-orange)
![CGE Version](https://img.shields.io/badge/CGE-v0.4-green)

A number guessing game.

## Known instances

- `games.code-game.org/number-guessing`

## Usage

```sh
# Run on default port 8080
number-guessing

# Specify a custom port
number-guessing --port=5000

## Specify a custom port through an environment variable
CG_PORT=5000 number-guessing
```

### Running with Docker

Prerequisites:
- [Docker](https://docker.com/)

```sh
# Download image
docker pull codegameproject/number-guessing:0.1

# Run container
docker run -d --restart on-failure -p <port-on-host-machine>:8080 --name number-guessing codegameproject/number-guessing:0.1
```

## Event flow

1. Send the `guess` command to make a guess.
2. Receive either the `too_high`, the `too_low` or the `correct` event.

## Building

### Prerequisites

- [Go](https://go.dev) 1.18+

```sh
git clone https://github.com/code-game-project/number-guessing.git
cd number-guessing
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
