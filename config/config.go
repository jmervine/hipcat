package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/jmervine/hipcat/Godeps/_workspace/src/gopkg.in/yaml.v2"
)

var DefaultConfig = "~/.hipcat"

func ReplaceHome(path string) (string, error) {
	if os.Getenv("HOME") == "" {
		return path, fmt.Errorf("HOME could not be fetched from your environment")
	}

	return strings.Replace(path, "~", os.Getenv("HOME"), 1), nil
}

type Config struct {
	Room    string `yaml:"room"`
	Token   string `yaml:"token"`
	Sender  string `yaml:"sender"`
	Host    string `yaml:"host"`
	Code    string `yaml:"code"`
	Color   string `yaml:"color"`
	Notify  string `yaml:"notify"`
	Conf    string
	Message []byte
}

func ToBool(s string) bool {
	s = string([]byte(strings.ToLower(s))[0])
	return (s == "t" || s == "y")
}

func (config *Config) LoadConfig(source string) error {
	source, err := ReplaceHome(source)

	if err != nil {
		return err
	}

	raw, err := ioutil.ReadFile(source)

	if err == nil {
		err = yaml.Unmarshal(raw, config)
	}

	return err
}

func (c *Config) Require() error {
	err := "Missing required argument: %s"
	if c.Room == "" {
		return fmt.Errorf(err, "room")
	}

	if c.Token == "" {
		return fmt.Errorf(err, "token")
	}

	return nil
}

func (c *Config) ReadMessage() error {
	stdin, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}
	c.Message = stdin

	return err
}

func (c *Config) FormattedMessage() string {
	format := "%s"

	if ToBool(c.Code) {
		format = "/code %s"
	}
	return fmt.Sprintf(format,
		string(c.Message[:len(c.Message)-1]))
}

func (c *Config) FormattedNotification() string {
	return fmt.Sprintf("<pre>%s</pre>",
		string(c.Message[:len(c.Message)-1]))
}
