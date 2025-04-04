package proxies

import (
	"fmt"
	"net/url"
)

func NewAuth(user string) (*ProxyAuth, error) {
	if err := checkUsername(user); err != nil {
		return nil, err
	}

	return &ProxyAuth{user: user, password: "", isPasswordSet: false}, nil
}

func NewPasswordAuth(user, password string) (*ProxyAuth, error) {
	if err := checkUsername(user); err != nil {
		return nil, err
	}

	auth := &ProxyAuth{user: user, password: password, isPasswordSet: false}
	if password != "" {
		auth.isPasswordSet = true
	}

	return auth, nil
}

func NewProxy(scheme, host string, port uint16, auth *ProxyAuth) (*Proxy, error) {
	if err := checkProxyScheme(scheme); err != nil {
		return nil, err
	}

	if auth != nil {
		if err := auth.check(); err != nil {
			return nil, err
		}
	}

	// add host check

	return &Proxy{scheme, host, port, auth}, nil
}

func NewProxyFromURL(u *url.URL) (*Proxy, error) {
	port, err := parsePort(u.Port())
	if err != nil {
		return nil, err
	}

	var auth *ProxyAuth
	if u.User != nil {
		auth, err = NewAuth(u.User.Username())
		if err != nil {
			return nil, err
		}
		password, isPasswordSet := u.User.Password()
		if isPasswordSet {
			auth.password = password
			auth.isPasswordSet = true
		}
	}

	return NewProxy(u.Scheme, u.Hostname(), port, auth)
}

func NewProxyFromStringURL(u string) (*Proxy, error) {
	parsedURL, err := url.Parse(u)
	if err != nil {
		return nil, &InvalidURL{ErrorInfo: ErrorInfo{
			message: fmt.Sprintf("Unable to parse URL \"%s\": %s", u, err),
		}}
	}

	return NewProxyFromURL(parsedURL)
}
