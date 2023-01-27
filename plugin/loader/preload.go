package loader

import (
	pluginbadgerds "github.com/bluzelle/ipfs-kubo/plugin/plugins/badgerds"
	pluginiplddagjose "github.com/bluzelle/ipfs-kubo/plugin/plugins/dagjose"
	pluginflatfs "github.com/bluzelle/ipfs-kubo/plugin/plugins/flatfs"
	pluginfxtest "github.com/bluzelle/ipfs-kubo/plugin/plugins/fxtest"
	pluginipldgit "github.com/bluzelle/ipfs-kubo/plugin/plugins/git"
	pluginlevelds "github.com/bluzelle/ipfs-kubo/plugin/plugins/levelds"
	pluginpeerlog "github.com/bluzelle/ipfs-kubo/plugin/plugins/peerlog"
)

// DO NOT EDIT THIS FILE
// This file is being generated as part of plugin build process
// To change it, modify the plugin/loader/preload.sh

func init() {
	Preload(pluginipldgit.Plugins...)
	Preload(pluginiplddagjose.Plugins...)
	Preload(pluginbadgerds.Plugins...)
	Preload(pluginflatfs.Plugins...)
	Preload(pluginlevelds.Plugins...)
	Preload(pluginpeerlog.Plugins...)
	Preload(pluginfxtest.Plugins...)
}
