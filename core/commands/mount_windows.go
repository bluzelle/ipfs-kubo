package commands

import (
	"errors"

	cmds "github.com/bluzelle/go-ipfs-cmds"
)

var MountCmd = &cmds.Command{
	Helptext: cmds.HelpText{
		Tagline:          "Not yet implemented on Windows.",
		ShortDescription: "Not yet implemented on Windows. :(",
	},

	Run: func(req *cmds.Request, res cmds.ResponseEmitter, env cmds.Environment) error {
		return errors.New("Mount isn't compatible with Windows yet")
	},
}
