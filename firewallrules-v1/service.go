package firewallrules

import (
	"github.com/akamai/AkamaiOPEN-edgegrid-golang/edgegrid"
)

var (
	// Config contains the Akamai OPEN Edgegrid API credentials
	// for automatic signing of requests
	Config edgegrid.Config
)

// Init sets the CPS edgegrid Config
func Init(config edgegrid.Config) {
	Config = config
}
