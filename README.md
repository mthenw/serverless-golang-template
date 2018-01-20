# serverless-golang-template

Multi-function Serverless Framework template for Golang (AWS Lambda).

## Requirements

* `dep` (Go dependency management tool): `go get -u github.com/golang/dep/cmd/dep`

## Quick start

1. Create a new service based on this template

```
serverless create -u https://github.com/mthenw/serverless-golang-template -p usersservice
```

2. Install framework dependencies

```
npm i
```

2. Install service dependencies

```
dep ensure
```

3. Deploy

```
sls deploy
```
