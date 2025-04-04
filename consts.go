package proxies

import "regexp"

var PROXY_SCHEMES []string = []string{"http", "socks5"}
var RE_PROXY_USERNAME *regexp.Regexp = regexp.MustCompile(`^[a-zA-Z0-9-_\.]+$`)
