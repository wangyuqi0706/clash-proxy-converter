package converter

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

type VmessURI struct {
	V    string `json:"v"`
	Ps   string `json:"ps"`
	Add  string `json:"add"`
	Port string `json:"port"`
	Id   string `json:"id"`
	Aid  string `json:"aid"`
	Scy  string `json:"scy"`
	Net  string `json:"net"`
	Type string `json:"type"`
	Host string `json:"host"`
	Path string `json:"path"`
	Tls  string `json:"tls"`
	Sni  string `json:"sni"`
}

func NewVmessURI(ssURI string) (*VmessURI, error) {
	v := new(VmessURI)
	err := v.Unmarshal(ssURI)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (v *VmessURI) Unmarshal(vmessURI string) error {
	if !strings.HasPrefix(vmessURI, "vmess://") {
		return fmt.Errorf("not a vmess URI")
	}
	vmessURI = strings.TrimPrefix(vmessURI, "vmess://")
	buffer := bytes.NewBufferString(vmessURI)
	decoder := base64.NewDecoder(base64.URLEncoding, buffer)
	jsDec := json.NewDecoder(decoder)
	var vmess VmessURI
	err := jsDec.Decode(&vmess)
	if err != nil {
		return err
	}
	*v = vmess
	return nil
}

func (v *VmessURI) ToClashProxy() *Proxy {
	return &Proxy{
		Name:           v.Ps,
		Type:           "vmess",
		Server:         v.Add,
		Port:           v.Port,
		UUID:           v.Id,
		AlterId:        v.Aid,
		Cipher:         "auto",
		Password:       "",
		UDP:            false,
		TLS:            false,
		SkipCertVerify: true,
		ServerName:     "",
		Network:        v.Net,
	}
}
