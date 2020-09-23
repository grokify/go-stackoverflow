# StackExchange / Stack Overflow Go SDK

[![Build Status][build-status-svg]][build-status-link]
[![Go Report Card][goreport-svg]][goreport-link]
[![Docs][docs-godoc-svg]][docs-godoc-link]
[![License][license-svg]][license-link]

 [build-status-svg]: https://api.travis-ci.org/grokify/go-stackoverflow.svg?branch=master
 [build-status-link]: https://travis-ci.org/grokify/go-stackoverflow
 [goreport-svg]: https://goreportcard.com/badge/github.com/grokify/go-stackoverflow
 [goreport-link]: https://goreportcard.com/report/github.com/grokify/go-stackoverflow
 [docs-godoc-svg]: https://img.shields.io/badge/docs-godoc-blue.svg
 [docs-godoc-link]: https://godoc.org/github.com/grokify/go-stackoverflow/client
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-link]: https://github.com/grokify/go-stackoverflow/blob/master/LICENSE

## Overview

This is a API client built using [openapi-generator](https://github.com/OpenAPITools/openapi-generator) using this StackExchange API spec:

https://github.com/grokify/api-specs/tree/master/stackexchange

## Installation

`$ go get github.com/grokify/go-stackoverflow/...`

## Usage

See examples in the [`examples`](examples) directory. To get started, you can use [`examples/get_users/main.go`](examples/get_users/main.go).

## Documentation

The auto-generated Swagger files are in the [`client`](client) folder and you can find the Swagger docs there as [`client/README.md`](client/README.md)
