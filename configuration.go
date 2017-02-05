package astiproxy

import "flag"

// Flags
var (
	Addr = flag.String("proxy-addr", "", "the proxy addr")
)

// Configuration represents the configuration of the proxy
type Configuration struct {
	Addr string `toml:"addr"`
}

// FlagConfig generates a Configuration based on flags
func FlagConfig() Configuration {
	return Configuration{
		Addr: *Addr,
	}
}
