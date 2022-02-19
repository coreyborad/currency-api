## Description

匯率轉換API

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

- 執行 Docker container
    - Into root directory
    - Input ```docker-compose up -d```
- 執行 API
    - CURL版本
    `curl -X 'GET' 'http://localhost:8888/api/v1/currency?in_currency=TWD&out_currency=JPY&price=20'`
    - 你也可以使用 `Swagger`, `swagger.json`在專案根目錄中有附上

## API

預設的http server port為 `8888`
Base url `http://localhost:8888/api/v1`

### Get currency
- URL `[GET]http://localhost:8888/api/v1/currency`

- Query string

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

- 這專案的檔案結構是依照 [Service & Repository](#http://kejyun.github.io/Laravel-5-Learning-Notes-Books/structure/structure-service-repository-structure-principle.html), 以本專案而言，目前只有使用1個API，示範用途，這邊只會做到`Service Layer`

`route => controllers => services`

- 使用 [wire](#https://github.com/google/wire) 來做 `inject`的動作

- 在route layer，定義http method以及api path
- 在controllers layer，會驗證parameter，接著將parameter送至services layer
- 在services layer，這邊會處理業務邏輯，同時也是在這邊處理currency的部份
- 在models layer，定義要接收的parameter struct

### UnitTest

本專案有處理的業務邏輯的部份，只有在services layer的 `currency.go`

this test coverage is `84.6%`
