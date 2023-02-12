# [CliDax][repo-url] [![Go Reference][pkg-dev-img]][pkg-dev-url] [![CI Status][ci-img]][ci-url] [![MIT License][mit-img]][mit-url]

A dax of sabi to operate command line interface of Golang application.

- [Usage](#usage)
- [Supporting Go versions](#support-go-version)
- [License](#license)

<a name="usage"></a>
## Usage

### Parse CLI arguments without configurations


### Parse CLI arguments with configurations


### Parse CLI arguments with struct tags


<a name="support-go-versions"></a>
## Supporting Go versions

This library supports Go 1.18 or later.

### Actual test results for each Go version:

```
% gvm-fav
Now using version go1.18.10
go version go1.18.10 darwin/amd64
ok  	github.com/sttk-go/clidax/libarg	0.131s	coverage: 100.0% of statements

Now using version go1.19.5
go version go1.19.5 darwin/amd64
ok  	github.com/sttk-go/clidax/libarg	0.135s	coverage: 100.0% of statements

Now using version go1.20
go version go1.20 darwin/amd64
ok  	github.com/sttk-go/clidax/libarg	0.137s	coverage: 100.0% of statements

Back to go1.20
Now using version go1.20
%
```


<a name="license"></a>
## License

Copyright (C) 2023 Takayuki Sato

This program is free software under MIT License.<br>
See the file LICENSE in this distribution for more details.


[repo-url]: https://github.com/sttk-go/clidax
[pkg-dev-img]: https://pkg.go.dev/badge/github.com/sttk-go/clidax.svg
[pkg-dev-url]: https://pkg.go.dev/github.com/sttk-go/clidax
[ci-img]: https://github.com/sttk-go/clidax/actions/workflows/go.yml/badge.svg?branch=main
[ci-url]: https://github.com/sttk-go/clidax/actions
[mit-img]: https://img.shields.io/badge/license-MIT-green.svg
[mit-url]: https://opensource.org/licenses/MIT

