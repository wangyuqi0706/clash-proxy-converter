package converter

import "strings"

func ConvertToMap(uris []string) []*Proxy {
	var proxies []*Proxy
	for _, uri := range uris {
		prefix, _, ok := strings.Cut(uri, "://")
		if !ok {
			continue
		}
		switch prefix {
		case "vmess":
			v, err := NewVmessURI(uri)
			if err != nil {
				continue
			}
			proxies = append(proxies, v.ToClashProxy())
		case "ss":
			s, err := NewSsURI(uri)
			if err != nil {
				continue
			}
			proxies = append(proxies, s.ToClashProxy())
		}

	}
	return proxies
}
