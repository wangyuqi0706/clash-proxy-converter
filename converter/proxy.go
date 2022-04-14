package converter

import (
	"github.com/fatih/structs"
)

type Proxy struct {
	Name           string `yaml:"name"`
	Type           string `yaml:"type"`
	Server         string `yaml:"server"`
	Port           int    `yaml:"port"`
	UUID           string `yaml:"uuid,omitempty"`
	AlterId        string `yaml:"alterId"`
	Cipher         string `yaml:"cipher"`
	Password       string `yaml:"password"`
	UDP            bool   `yaml:"udp"`
	TLS            bool   `yaml:"tls,omitempty"`
	SkipCertVerify bool   `yaml:"skip-cert-verify,omitempty"`
	ServerName     string `yaml:"servername,omitempty"`
	Network        string `yaml:"network,omitempty"`
}

func (p *Proxy) ToMap() map[string]any {
	return structs.Map(p)
}
