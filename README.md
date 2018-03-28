# Open Graph Extractor

## Docs
Open Graph Extractor written in Golang

[![License](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/katera/og/blob/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/katera/og?status.svg)](https://godoc.org/github.com/katera/og)

## Index

* [Support](#support)
* [Getting Started](#getting-started)
* [Example](#example)
* [Contribution](#contribution)


## Support

You can file an [Issue](https://github.com/katera/og/issues/new).
See documentation in [Godoc](https://godoc.org/github.com/katera/og)


## Getting Started

#### Download

```shell
go get -u github.com/katera/og
```
## Example
```go
package main

import (
    "github.com/katera/og"
 )
 
func main() {
 
 res, err := og.GetOpenGraphFromUrl("https://hackernoon.com/golang-clean-archithecture-efd6d7c43047")
 
 if err != nil {
  log.Fatal(err)
 }
 
 fmt.Printf("%+v",res) 

}
 ```
 
## Contribution
To contrib on this project, you can make a PR or just an issue.

### Maintainer
- <a href="https://github.com/bxcodec">  **Iman Tumorang** </a> <br> 


