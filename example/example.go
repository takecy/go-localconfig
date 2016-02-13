package main

import (
	"fmt"
	"os"
	"time"

	"github.com/takecy/go-localconfig/conf"
)

var conf *localconf.Conf

// AppConfig is config
type AppConfig struct {
	UserID int
	Token  string
	Date   int64
	Debug  bool
}

func main() {
	err := save()
	if err != nil {
		panic(err)
	}

	con, err := load()
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stdout, "loaded config: %+v\n", con)

	cleanup()

	isExist := exist()
	if isExist {
		panic("not.cleanuped")
	}

	fmt.Fprint(os.Stdout, "example done.\n")
}

func init() {
	appName := "test_app"
	fileName := "config.json"

	conf = localconf.NewConfig(appName, fileName)
	fmt.Fprintf(os.Stdout, "config.path: %s\n", conf.Path())
}

func exist() bool {
	return conf.Exist()
}

func save() (err error) {
	c := &AppConfig{
		UserID: 12345,
		Token:  "token_hogehoge",
		Date:   time.Now().Unix(),
		Debug:  true,
	}
	// should be pointer
	err = conf.Save(c)
	return
}

func load() (c AppConfig, err error) {
	// should be pointer
	err = conf.Load(&c)
	return
}

func cleanup() (err error) {
	err = conf.Cleanup()
	return
}
