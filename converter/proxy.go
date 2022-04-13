package converter

import (
	"github.com/fatih/structs"
	"io"
)

type Proxy struct {
	Name           string `yaml:"name"`
	Type           string `yaml:"type"`
	Server         string `yaml:"server"`
	Port           string `yaml:"port"`
	UUID           string `yaml:"uuid"`
	AlterId        string `yaml:"alterId"`
	Cipher         string `yaml:"cipher"`
	Password       string `yaml:"password"`
	UDP            bool   `yaml:"udp"`
	TLS            bool   `yaml:"tls"`
	SkipCertVerify bool   `yaml:"skip-cert-verify"`
	ServerName     string `yaml:"servername"`
	Network        string `yaml:"network"`
}

func (p *Proxy) ToMap() map[string]any {
	return structs.Map(p)
}

func (p *Proxy) ToYaml(writer io.Writer) {

}
