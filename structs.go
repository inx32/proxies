package proxies

// Proxy server authentication. You must create it using following constructors:
//   - NewAuth(user string)
//   - NewPasswordAuth(user string, password string)
type ProxyAuth struct {
	user          string
	password      string
	isPasswordSet bool
}

// Proxy server struct. You must create it using following constructors:
//   - NewProxy(scheme string, host string, port uint16, auth *ProxyAuth)
//   - NewProxyFromURL(url *net/url.Url)
//   - NewProxyFromStringURL(url string)
//
// NOTE: If proxy does not use authentication, set "auth" parameter to nil
type Proxy struct {
	scheme string
	host   string
	port   uint16
	auth   *ProxyAuth
}
