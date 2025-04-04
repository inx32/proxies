package proxies

// Returns user name. Can't be empty if authentication field is specified in "Proxy".
func (a *ProxyAuth) User() string {
	return a.user
}

// Returns password. If password is empty, empty string will be returned. You can check is password set
// with getter "ProxyAuth.IsPasswordSet() bool".
func (a *ProxyAuth) Password() string {
	return a.password
}

// Returns false if password is empty string and true if password is not empty.
func (a *ProxyAuth) IsPasswordSet() bool {
	return a.isPasswordSet
}

// Returns proxy scheme ("socks5" or "http").
func (p *Proxy) Scheme() string {
	return p.scheme
}

// Returns proxy IPv4/IPv6 address (example: 1.1.1.1, ::1) or domain name (example: proxy.example.com).
func (p *Proxy) Host() string {
	return p.host
}

// Returns port (integer from 1 to 65535).
func (p *Proxy) Port() uint16 {
	return p.port
}

// Returns pointer to ProxyAuth specified in proxy.
func (p *Proxy) Auth() *ProxyAuth {
	return p.auth
}
