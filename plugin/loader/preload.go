package loader

import (
	pluginbadgerds "github.com/ipfs/go-ipfs/plugin/plugins/badgerds"
	pluginiplddagjose "github.com/ipfs/go-ipfs/plugin/plugins/dagjose"
	pluginflatfs "github.com/ipfs/go-ipfs/plugin/plugins/flatfs"
	pluginipldgit "github.com/ipfs/go-ipfs/plugin/plugins/git"
	pluginlevelds "github.com/ipfs/go-ipfs/plugin/plugins/levelds"
	pluginpeerlog "github.com/ipfs/go-ipfs/plugin/plugins/peerlog"
	pluginipldwasm "github.com/ipfs/go-ipfs/plugin/plugins/wasmipld"
)

// DO NOT EDIT THIS FILE
// This file is being generated as part of plugin build process
// To change it, modify the plugin/loader/preload.sh

func init() {
	Preload(pluginipldgit.Plugins...)
	Preload(pluginiplddagjose.Plugins...)
	Preload(pluginipldwasm.Plugins...)
	Preload(pluginbadgerds.Plugins...)
	Preload(pluginflatfs.Plugins...)
	Preload(pluginlevelds.Plugins...)
	Preload(pluginpeerlog.Plugins...)
}
