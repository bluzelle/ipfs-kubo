package rpc

import (
	"context"
	"encoding/json"

	caopts "github.com/bluzelle/boxo/coreiface/options"
	"github.com/bluzelle/boxo/path"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/routing"
)

type DhtAPI HttpApi

func (api *DhtAPI) FindPeer(ctx context.Context, p peer.ID) (peer.AddrInfo, error) {
	var out struct {
		Type      routing.QueryEventType
		Responses []peer.AddrInfo
	}
	resp, err := api.core().Request("dht/findpeer", p.String()).Send(ctx)
	if err != nil {
		return peer.AddrInfo{}, err
	}
	if resp.Error != nil {
		return peer.AddrInfo{}, resp.Error
	}
	defer resp.Close()
	dec := json.NewDecoder(resp.Output)
	for {
		if err := dec.Decode(&out); err != nil {
			return peer.AddrInfo{}, err
		}
		if out.Type == routing.FinalPeer {
			return out.Responses[0], nil
		}
	}
}

func (api *DhtAPI) FindProviders(ctx context.Context, p path.Path, opts ...caopts.DhtFindProvidersOption) (<-chan peer.AddrInfo, error) {
	options, err := caopts.DhtFindProvidersOptions(opts...)
	if err != nil {
		return nil, err
	}

	rp, _, err := api.core().ResolvePath(ctx, p)
	if err != nil {
		return nil, err
	}

	resp, err := api.core().Request("dht/findprovs", rp.RootCid().String()).
		Option("num-providers", options.NumProviders).
		Send(ctx)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, resp.Error
	}
	res := make(chan peer.AddrInfo)

	go func() {
		defer resp.Close()
		defer close(res)
		dec := json.NewDecoder(resp.Output)

		for {
			var out struct {
				Extra     string
				Type      routing.QueryEventType
				Responses []peer.AddrInfo
			}

			if err := dec.Decode(&out); err != nil {
				return // todo: handle this somehow
			}
			if out.Type == routing.QueryError {
				return // usually a 'not found' error
				// todo: handle other errors
			}
			if out.Type == routing.Provider {
				for _, pi := range out.Responses {
					select {
					case res <- pi:
					case <-ctx.Done():
						return
					}
				}
			}
		}
	}()

	return res, nil
}

func (api *DhtAPI) Provide(ctx context.Context, p path.Path, opts ...caopts.DhtProvideOption) error {
	options, err := caopts.DhtProvideOptions(opts...)
	if err != nil {
		return err
	}

	rp, _, err := api.core().ResolvePath(ctx, p)
	if err != nil {
		return err
	}

	return api.core().Request("dht/provide", rp.RootCid().String()).
		Option("recursive", options.Recursive).
		Exec(ctx, nil)
}

func (api *DhtAPI) core() *HttpApi {
	return (*HttpApi)(api)
}
