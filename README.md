# [cliargdax][repo-url] [![Go Reference][pkg-dev-img]][pkg-dev-url] [![CI Status][ci-img]][ci-url] [![MIT License][mit-img]][mit-url]

A dax for command-line arguments in Golang applications.

Dax is a module in [sabi][sabi-url] framework for data access.
This package provides DaxSrc and DaxConn structs for command-line argument operations.

This package uses [cliargs][cliargs-url] package to parse command-line arguments and store the results of parsing.


## Import this package

```
import "github.com/sttk/cliargdax"
```


## Usage

The usage of this library is described on the overview in the go package document.

See https://pkg.go.dev/github.com/sttk/cliargdax#pkg-overview


## Supporting Go versions

This library supports Go 1.18 or later.


### Actual test results for each Go versions:

```
% gvm-fav
Now using version go1.18.10
go version go1.18.10 darwin/amd64
ok  	github.com/sttk/cliargdax	0.150s	coverage: 100.0% of statements

Now using version go1.19.13
go version go1.19.13 darwin/amd64
ok  	github.com/sttk/cliargdax	0.157s	coverage: 100.0% of statements

Now using version go1.20.8
go version go1.20.8 darwin/amd64
ok  	github.com/sttk/cliargdax	0.161s	coverage: 100.0% of statements

Now using version go1.21.1
go version go1.21.1 darwin/amd64
ok  	github.com/sttk/cliargdax	0.174s	coverage: 100.0% of statements

Back to go1.21.1
Now using version go1.21.1
```


## License

Copyright (C) 2023 Takayuki Sato

This program is free software under MIT License.<br>
See the file LICENSE in this distribution for more details.


[repo-url]: https://github.com/sttk/cliargdax
[pkg-dev-img]: https://pkg.go.dev/badge/github.com/sttk/cliargdax.svg
[pkg-dev-url]: https://pkg.go.dev/github.com/sttk/cliargdax
[ci-img]: https://github.com/sttk/cliargdax/actions/workflows/go.yml/badge.svg?branch=main
[ci-url]: https://github.com/sttk/cliargdax/actions
[mit-img]: https://img.shields.io/badge/license-MIT-green.svg
[mit-url]: https://opensource.org/licenses/MIT

[sabi-url]: https://github.com/sttk/sabi
[cliargs-url]: https://github.com/sttk/cliargs
