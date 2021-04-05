# go-apple-strings

[![Actions Status](https://github.com/harryzcy/go-apple-strings/workflows/CI/badge.svg)](https://github.com/harryzcy/go-apple-strings/actions)
[![codecov](https://codecov.io/gh/harryzcy/go-apple-strings/branch/main/graph/badge.svg)](https://codecov.io/gh/harryzcy/go-apple-strings)
[![Go Report Card](https://goreportcard.com/badge/github.com/harryzcy/go-apple-strings)](https://goreportcard.com/report/github.com/harryzcy/go-apple-strings)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat)](http://makeapullrequest.com)

Apple strings file parser written in Go.

## INSTALL

```shell
go get github.com/harryzcy/go-apple-strings
```

## USE

```go
package main

import (
  "os"

  applestrings "github.com/harryzcy/go-apple-strings"
)

type Pairs struct {
  Key string
}

func main() {
  f, _ := os.Open("InfoPlist.strings")
  defer f.Close()

  var pairs Pairs
  decoder := applestrings.NewDecoder(f)
  err := decoder.Decode(&pairs)
  if err != nil {
    panic(err)
  }
  // Use pairs...
}
```
