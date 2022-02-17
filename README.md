## Description

This is a simple for exchange currency's API

## Table of Contents

- [Description](#description)
- [Table of Contents](#table-of-contents)
- [Preparation](#preparation)
- [Quick Start](#quick-start)
- [API](#api)
    - [GET Currency](#get-currency)
- [Architecture](#architecture)
    - [Trees](#trees)
    - [Request Flow](#request-flow)
    - [Unit Test](#unit-test)

## Preparation

- [docker](https://www.docker.com) - `17.04.0+`
- [docker-compose](https://docs.docker.com/compose) - `1.27.0+`

## Quick Start

- Run Docker container
    - Into root directory
    - Input ```docker-compose up -d```
- Run API
    - With CURL
    `curl -X 'GET' 'http://localhost:8888/api/v1/currency?in_currency=TWD&out_currency=JPY&price=20'`
    - You can use `Swagger`, this project has `swagger.json`

## API

Default http server port is `8888`
Base url `http://localhost:8888/api/v1`

### Get currency
- URL `[GET]http://localhost:8888/api/v1/currency`

- Querystring

| Parameter    | Type                    | Required | Description    |
| -------------- | ------------------------- | ------ | ------- |
| in_currency  | enum("TWD","JPY","USD") | Y    | 輸入原匯率 |
| out_currency | enum("TWD","JPY","USD") | Y    | 輸出匯率  |
| price        | float64                 | Y    | 金額    |  

- Example
    - Input
    `http://localhost:8888/api/v1/currency?in_currency=TWD&out_currency=JPY&price=9999`
    - Output
    `
    {
        "result":"36,686.33"
    }
    `

## Architecture

### Trees

```
.
├── go/
│   └── Dockerfile - docker container的dockerfile
├── src/
│   ├── controllers - controller file
│   ├── routes -API route file
│   ├── services
│   ├── models - Define data struct
│   ├── utils - Common function
│   └── validate - controller validater
├── docker-compose.yml - docker-compose config
└── swagger.json - swagger API doc
```

### Request Flow

- This project is follow part of [Service & Repository](#http://kejyun.github.io/Laravel-5-Learning-Notes-Books/structure/structure-service-repository-structure-principle.html), but our case just has one use case here, so this project has no `Repository Layer` here 

`route => controllers => services`

- This project use [wire](#https://github.com/google/wire) for inject

- In route, we define the API path and http method
- In controllers, will validate parameter is correct or not, and send parameter to service
- In services, this part will process our business logic. In this project, we will calc our currency

- In models, we define controllers's parameter struct

- In utils, we set a const of conrrency here in `currency.go`

### UnitTest

In our case, our logic all of in the service layer `currency.go`

this test coverage is `84.6%`
