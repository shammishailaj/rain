package rain

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v1"
)

type Config struct {
	Port       uint16
	Encryption struct {
		DisableOutgoing bool `yaml:"disable_outgoing"`
		ForceOutgoing   bool `yaml:"force_outgoing"`
		ForceIncoming   bool `yaml:"force_incoming"`
	}
}

var DefaultConfig = Config{
	Port: 6881,
}

func LoadConfig(filename string) (*Config, error) {
	c := DefaultConfig
	b, err := ioutil.ReadFile(filename)
	if os.IsNotExist(err) {
		return &c, nil
	}
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(b, &c); err != nil {
		return nil, err
	}
	return &c, nil
}