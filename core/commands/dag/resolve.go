package dagcmd

import (
	"github.com/bluzelle/boxo/path"
	"github.com/bluzelle/ipfs-kubo/core/commands/cmdenv"
	"github.com/bluzelle/ipfs-kubo/core/commands/cmdutils"

	cmds "github.com/bluzelle/go-ipfs-cmds"
)

func dagResolve(req *cmds.Request, res cmds.ResponseEmitter, env cmds.Environment) error {
	api, err := cmdenv.GetApi(env, req)
	if err != nil {
		return err
	}

	p, err := cmdutils.PathOrCidPath(req.Arguments[0])
	if err != nil {
		return err
	}

	rp, remainder, err := api.ResolvePath(req.Context, p)
	if err != nil {
		return err
	}

	return cmds.EmitOnce(res, &ResolveOutput{
		Cid:     rp.RootCid(),
		RemPath: path.SegmentsToString(remainder...),
	})
}
