/*
Package localconf is local configuration.
Save configuration to $HOME/.{appName}/{fileName}.
*/
package localconf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Conf is configuration
type Conf struct {
	dirPath  string
	fileName string
}

// NewConfig is init config
func NewConfig(appName, fileName string) *Conf {
	if appName == "" {
		panic("appName.required")
	}

	if fileName == "" {
		fileName = "conf.json"
	}

	c := new(Conf)
	c.dirPath = fmt.Sprintf("%s/.%s", os.Getenv("HOME"), appName)
	c.fileName = fileName

	c.initConfig()

	return c
}

// init local config
func (c *Conf) initConfig() {
	err := os.MkdirAll(c.dirPath, 0755)
	if err != nil {
		panic(err)
	}

	_, err = os.Create(c.Path())
	if err != nil {
		panic(err)
	}

	bc := map[string]string{
		"init": "true",
	}
	b, err := json.Marshal(bc)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(c.Path(), b, 0666)
	if err != nil {
		panic(err)
	}
}

// Path is gets configuration file path.
func (c *Conf) Path() string {
	return fmt.Sprintf("%s/%s", c.dirPath, c.fileName)
}

// Exist is check local config file
func (c *Conf) Exist() bool {
	_, err := os.Stat(c.Path())
	return (err == nil)
}

// Load is read local config file
// src should be pointer
func (c *Conf) Load(src interface{}) (err error) {
	b, err := ioutil.ReadFile(c.Path())
	if err != nil {
		return
	}

	err = json.Unmarshal(b, src)
	if err != nil {
		return
	}

	return
}

// Save is write to local config file
func (c *Conf) Save(src interface{}) (err error) {
	nb, err := json.Marshal(src)
	if err != nil {
		return
	}

	err = ioutil.WriteFile(c.Path(), nb, 0666)
	if err != nil {
		return
	}

	return
}

// Cleanup is delete local config
func (c *Conf) Cleanup() (err error) {
	err = os.RemoveAll(c.dirPath)
	return
}
