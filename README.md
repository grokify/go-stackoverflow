# StackExchange / Stack Overflow Go SDK

[![Build Status][build-status-svg]][build-status-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![License][license-svg]][license-url]

 [build-status-svg]: https://github.com/grokify/go-stackoverflow/workflows/go%20build/badge.svg
 [build-status-url]: https://github.com/grokify/go-stackoverflow/actions
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/go-stackoverflow
 [goreport-url]: https://goreportcard.com/report/github.com/grokify/go-stackoverflow
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/grokify/go-stackoverflow
 [docs-godoc-url]: https://pkg.go.dev/github.com/grokify/go-stackoverflow
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/grokify/go-stackoverflow/blob/master/LICENSE

## Overview

This is a API client built using [openapi-generator](https://github.com/OpenAPITools/openapi-generator) using this StackExchange API spec:

https://github.com/grokify/api-specs/tree/master/stackexchange

## Installation

`$ go get github.com/grokify/go-stackoverflow/...`

## Usage

See examples in the [`examples`](examples) directory. To get started, you can use [`examples/get_users/main.go`](examples/get_users/main.go).

## Documentation

The auto-generated Swagger files are in the [`client`](client) folder and you can find the Swagger docs there as [`client/README.md`](client/README.md)

## API Coverage

APIs: https://api.stackexchange.com/docs

- [ ] Questions
  - [x] `questions`
  - [x] `questions/featured`
  - [x] `questions/no-answers`
  - [x] `questions/unanswered`
- [ ] Users
  - [x] `users`
  - [x] `users/{userIds}/reputation`
  - [x] `users/{userIds}/reputation-history`
  - [x] `me`
  - [x] `me/reputation`
  - [x] `me/reputation-history`