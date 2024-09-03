package main

import (
	"os"

	"github.com/bluzelle/ipfs-kubo/cmd/ipfs/kubo"
)

func main() {
	os.Exit(kubo.Start(kubo.BuildDefaultEnv))
}
