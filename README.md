# go-localconfig

go-localconfig is provides save/load configuration from local filesystem.  
configuration file is saved to `$HOME/.{appName}/{fileName}` .

![](https://img.shields.io/badge/language-go-blue.svg?style=flat)
![](https://img.shields.io/badge/golang-1.5.2+-blue.svg?style=flat)  

[![GoDoc](https://godoc.org/github.com/takecy/go-localconfig/conf?status.svg)](https://godoc.org/github.com/takecy/go-localconfig/conf)

## Install
```shell
$ go get github.com/takecy/go-localconfig/conf
```

## Usage
```go
import (
	"github.com/takecy/go-localconfig/conf"
)

var conf *localconf.Conf

func NewSomething() {
	appName := "test_app"
	fileName := "config.json"
	conf = localconf.NewConfig(appName, fileName)
}

func save(someConfig interface{}) (err error) {
	err = conf.Save(someConfig)
	return
}

func load() (c SomeConfig, err error) {
	err = conf.Load(&c)
	return
}
```

see more [Example](./example/example.go)


## Testing
```shell
$ go get github.com/smartystreets/goconvey/convey
$ go test ./...
```

# License
MIT
