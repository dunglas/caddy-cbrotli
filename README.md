# Caddy Brotli Module

This module for [the Caddy web server](https://caddyserver.com) provides support for [the Brotli compression format](https://en.wikipedia.org/wiki/Brotli).

It uses [the reference implementation of Brotli](https://brotli.org/) (written in C), through [the Go module provided by Google](https://pkg.go.dev/github.com/google/brotli/go/cbrotli).

[![Tests](https://github.com/dunglas/caddy-cbrotli/actions/workflows/tests.yaml/badge.svg)](https://github.com/dunglas/caddy-cbrotli/actions/workflows/tests.yaml)
[![Go Reference](https://pkg.go.dev/badge/github.com/dunglas/caddy-cbrotli.svg)](https://pkg.go.dev/github.com/dunglas/caddy-cbrotli)

## Install

1. Install [cbrotli](https://github.com/google/brotli/tree/master/c). On Mac run `brew install brotli`. On Debian and Ubuntu, run `apt install libbrotli-dev`.
2. Then create your Caddy build:
    ```console
    CGO_ENABLED=1 \
    xcaddy build \
        --with github.com/dunglas/caddy-cbrotli
    ```

    On Mac, be sure to adapt the paths in `CGO_LDFLAGS` and `CGO_CFLAGS` according to your Brotli installation:

    ```console
    CGO_LDFLAGS="-L/opt/homebrew/lib/" \
    CGO_CFLAGS="-I/opt/homebrew/include/" \
    CGO_ENABLED=1 \
    xcaddy build \
        --with github.com/dunglas/caddy-cbrotli
    ```

## Usage

Add the `br` value to [the `encode` directive](https://caddyserver.com/docs/caddyfile/directives/encode) in your `Caddyfile`.

Example:

```caddyfile
localhost

encode zstd br gzip

file_server
```

Alternatively, you can configure the quality (from 0 to 11, defaults to 6) and the base 2 logarithm of the sliding window size (from 10 to 24, defaults to auto):

Example:

```caddyfile
localhost

encode {
    br 8 15
}

file_server
```

## Cgo

This module depends on [cgo](https://go.dev/blog/cgo).
If you are looking for a non-cgo (but more CPU-intensive) alternative, see [the `github.com/ueffel/caddy-brotli` module](https://github.com/ueffel/caddy-brotli).

## Credits

Created by [Kévin Dunglas](https://dunglas.dev) and sponsored by [Les-Tilleuls.coop](https://les-tilleuls.coop).
