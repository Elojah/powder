[![CircleCI](https://circleci.com/gh/Elojah/powder/tree/master.svg?style=svg&circle-token=53f826c3f3cd02c2e3c5503c53618a9e1d34a6f0)](https://app.circleci.com/github/Elojah/powder/pipelines)


# powder
Tech interview solution for Powder

### Requirements

`docker`
`docker-compose`

### Installation

```sh
go get -u github.com/elojah/powder
cd <cloned_directory>
docker-compose up -d
```

HTTP server listen per default on port `:8080`, it may not start if this port is already affected.
You can change this setting in `config/api.json` and `docker-compose.yml`

### Usage example

```sh

$ go get github.com/hashrocket/ws
$ ws -k wss://localhost:8080/connect
> {"method": 0, "content": {"id": "01CCS1SZ4B20G98XYMFGVC9VS4", "name": "roberta"}}
< successfully login

```
A basic `sh` test file is provided in `scripts/stories.sh`.

### TODO
