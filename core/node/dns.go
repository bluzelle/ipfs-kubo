package node

import (
	"math"
	"time"

	"github.com/bluzelle/boxo/gateway"
	config "github.com/bluzelle/ipfs-kubo/config"
	doh "github.com/libp2p/go-doh-resolver"
	madns "github.com/multiformats/go-multiaddr-dns"
)

func DNSResolver(cfg *config.Config) (*madns.Resolver, error) {
	var dohOpts []doh.Option
	if !cfg.DNS.MaxCacheTTL.IsDefault() {
		dohOpts = append(dohOpts, doh.WithMaxCacheTTL(cfg.DNS.MaxCacheTTL.WithDefault(time.Duration(math.MaxUint32)*time.Second)))
	}

	return gateway.NewDNSResolver(cfg.DNS.Resolvers, dohOpts...)
}
