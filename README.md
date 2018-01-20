# serverless-golang-template

Multi-function Serverless Framework template for Golang (AWS Lambda).

## Requirements

* `dep` (Go dependency management tool): `go get -u github.com/golang/dep/cmd/dep`

## Quick start

1. Create a new service in $GOPATH based on this template

```
serverless create -u https://github.com/mthenw/serverless-golang-template -p usersservice
```

2. Update `import` statements in `functions` directory files

3. Install framework plugin

```
npm i
```

4. Install service dependencies

```
dep ensure
```

5. Deploy

```
sls deploy
```
