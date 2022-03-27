package converter

import (
	"github.com/fatih/structs"
)

type Proxy struct {
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`
	Server   string `yaml:"server"`
	Port     string `yaml:"port"`
	Cipher   string `yaml:"cipher"`
	Password string `yaml:"password"`
	UDP      bool   `yaml:"udp"`
	Network  string `yaml:"network"`
	AlterId  int    `yaml:"alterId"`
	UUID     string `yaml:"uuid"`
}

func (p *Proxy) ToMap() map[string]any {
	return structs.Map(p)
}
