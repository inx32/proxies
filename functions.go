package proxies

import (
	"context"
	"fmt"
	"net"
	"slices"
	"strconv"

	"net/http"
	"net/url"

	"golang.org/x/net/proxy"
)

func checkUsername(user string) error {
	if !RE_PROXY_USERNAME.MatchString(user) {
		return &InvalidUsername{ErrorInfo: ErrorInfo{
			message: fmt.Sprintf("String \"%s\" does not matching regexp %s",
				user, RE_PROXY_USERNAME.String()),
		}}
	}

	return nil
}

func checkProxyScheme(scheme string) error {
	if !slices.Contains(PROXY_SCHEMES, scheme) {
		return &InvalidScheme{ErrorInfo: ErrorInfo{
			message: fmt.Sprintf("Scheme \"%s\" is not supported", scheme),
		}}
	}

	return nil
}

func parsePort(port string) (uint16, error) {
	portUint, err := strconv.ParseUint(port, 10, 16)
	if err != nil {
		return 0, &InvalidPort{ErrorInfo: ErrorInfo{
			message: fmt.Sprintf("Unable to parse port \"%s\": %s", port, err),
		}}
	}
	return uint16(portUint), nil
}

func (a *ProxyAuth) check() error {
	if err := checkUsername(a.user); err != nil {
		return err
	}

	if a.password == "" && a.isPasswordSet {
		a.isPasswordSet = false
	}

	return nil
}

func (p *Proxy) GetURL() *url.URL {
	var userInfo *url.Userinfo

	if p.auth != nil {
		switch {
		case p.auth.isPasswordSet:
			userInfo = url.UserPassword(p.auth.user, p.auth.password)
		default:
			userInfo = url.User(p.auth.user)
		}
	}

	return &url.URL{
		Scheme: p.scheme,
		Host:   net.JoinHostPort(p.host, strconv.Itoa(int(p.port))),
		User:   userInfo,
	}
}

func (p *Proxy) GetStringURL() string {
	return p.GetURL().String()
}

func (p *Proxy) GetSocks5Dialer() (*proxy.Dialer, error) {
	if p.scheme != "socks5" {
		return nil, &InvalidScheme{ErrorInfo: ErrorInfo{
			message: "Only \"socks5\" scheme supported for socks5 dialer",
		}}
	}

	url := p.GetURL()
	dialer, err := proxy.FromURL(url, proxy.Direct)
	if err != nil {
		return nil, &DialerInitFailed{ErrorInfo: ErrorInfo{
			message: fmt.Sprintf("Unable to init SOCKS5 dialer: %s", err),
		}}
	}

	return &dialer, nil
}

func (p *Proxy) GetHttpTransport() *http.Transport {
	url := p.GetURL()
	transport := &http.Transport{}

	switch p.scheme {
	case "http":
		transport.Proxy = http.ProxyURL(url)
	case "socks5":
		transport.DialContext = p.DialContext
	}

	return transport
}

func (p *Proxy) DialContext(ctx context.Context, network, addr string) (net.Conn, error) {
	dialer, err := p.GetSocks5Dialer()
	if err != nil {
		return nil, err
	}

	return (*dialer).Dial(network, addr)
}

func (p *Proxy) CheckConnection(u string) error {
	client := &http.Client{Transport: p.GetHttpTransport()}
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return &ConnectionCheckFailed{ErrorInfo: ErrorInfo{
			message: fmt.Sprintf("Unable to check connection: %s", err),
		}}
	}

	resp, err := client.Do(req)
	if err != nil {
		return &ConnectionCheckFailed{ErrorInfo: ErrorInfo{
			message: fmt.Sprintf("Unable to check connection: %s", err),
		}}
	}
	defer resp.Body.Close()

	return nil
}
