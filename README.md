# goldmark-mattermost
[![Go Report Card](https://goreportcard.com/badge/github.com/nathanmkaya/goldmark-extended)](https://goreportcard.com/report/github.com/nathanmkaya/goldmark-extended)

Extension of goldmark markdown parser with support for jira and mattermost markdown syntax

## Installation
```shell script
go get github.com/yuin/goldmark
go get github.com/nathanmkaya/goldmark-extended
```

## Usage

Import packages
```go
import (
    "bytes"
    "github.com/yuin/goldmark"
    "github.com/nathanmkaya/goldmark-extended"
)
```

For Jira specific markdown conversion
```go
md := goldmark.New(
          goldmark.WithExtensions(jira.Jira),
          ),
      )
var buf bytes.Buffer
if err := md.Convert(source, &buf); err != nil {
    log.Fatalln(err)
}
```


### License

MIT

### Author

Nathanael Mkaya