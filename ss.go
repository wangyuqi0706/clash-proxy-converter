package converter

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

type SsURI struct {
	Method   string
	Password string
	Addr     string
	Port     string
	Name     string
}

func NewSsURI(ssURI string) (*SsURI, error) {
	s := new(SsURI)
	err := s.Unmarshal(ssURI)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *SsURI) Unmarshal(ssURI string) error {
	if !strings.HasPrefix(ssURI, "ss://") {
		return fmt.Errorf("not a vmess URI")
	}
	ssURI = strings.TrimPrefix(ssURI, "ss://")

	fields := strings.FieldsFunc(ssURI, func(r rune) bool {
		return r == ':' || r == '@' || r == '#'
	})

	buffer := bytes.NewBufferString(fields[0])
	decoder := base64.NewDecoder(base64.URLEncoding, buffer)
	data, _ := ioutil.ReadAll(decoder)
	mp := string(data)
	meth, pass, _ := strings.Cut(mp, ":")
	name, _ := url.QueryUnescape(fields[3])
	*s = SsURI{
		Method:   meth,
		Password: pass,
		Addr:     fields[1],
		Port:     fields[2],
		Name:     name,
	}
	return nil
}

func (s *SsURI) ToClashProxy() *Proxy {
	return &Proxy{
		Name:     s.Name,
		Type:     "ss",
		Server:   s.Addr,
		Port:     s.Port,
		Cipher:   s.Method,
		Password: s.Password,
		UDP:      true,
		Network:  "",
		AlterId:  0,
		UUID:     "",
	}
}
