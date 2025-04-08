<p align="center">
  <a href="https://www.kowabunga.cloud/?utm_source=github&utm_medium=logo" target="_blank">
    <picture>
      <source srcset="https://raw.githack.com/kowabunga-cloud/infographics/master/art/kowabunga-square-600x600-2.png" media="(prefers-color-scheme: dark)" />
      <source srcset="https://raw.githack.com/kowabunga-cloud/infographics/master/art/kowabunga-square-600x600-2.png" media="(prefers-color-scheme: light), (prefers-color-scheme: no-preference)" />
      <img src="https://raw.githack.com/kowabunga-cloud/infographics/master/art/kowabunga-square-600x600-2.png" alt="Kowabunga" width="200">
    </picture>
  </a>
</p>

# Official Kowabunga SDK for Go

This is official Go SDK for Kowabunga API.

[![License: Apache License, Version 2.0](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://spdx.org/licenses/Apache-2.0.html)
[![go.dev](https://img.shields.io/badge/go.dev-pkg-007d9c.svg?style=flat)](https://pkg.go.dev/github.com/kowabunga-cloud/kowabunga-go)

## Current Releases

| Project            | Release Badge                                                                                       |
|--------------------|-----------------------------------------------------------------------------------------------------|
| **Kowabunga**           | [![Kowabunga Release](https://img.shields.io/github/v/release/kowabunga-cloud/kowabunga)](https://github.com/kowabunga-cloud/kowabunga/releases) |
| **Kowabunga Go SDK**     | [![Kowabunga Go SDK Release](https://img.shields.io/github/v/release/kowabunga-cloud/kowabunga-go)](https://github.com/kowabunga-cloud/kowabunga-go/releases) |

## Installation

`kowabunga-go` can be installed like any other Go library through `go get`:

```console
$ go get github.com/kowabunga-cloud/kowabunga-go@latest
```

Check out the [list of released versions](https://github.com/kowabunga-cloud/kowabunga-go/releases).

## Configuration

To use `kowabunga-go`, you’ll need to import the `kowabunga-go` package:

```go
import kowabunga "github.com/kowabunga-cloud/kowabunga-go"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Usage

Creating an API client can be done by:

```go

var client *kowabunga.APIClient

u, _ := url.Parse(uri)

cfg := kowabunga.NewConfiguration()
cfg.Host = u.Host
cfg.Scheme = u.Scheme
cfg.Debug = true
cfg.AddDefaultHeader("X-API-Key", token)

client = kowabunga.NewAPIClient(cfg), nil
```

where **uri** is *https://your_kowabunga_kahuna_server* and **token** is the associated API key.

## Documentation for API

Refer to [API documentation](API.md)

## License

Licensed under [Apache License, Version 2.0](https://opensource.org/license/apache-2-0), see [`LICENSE`](LICENSE).
