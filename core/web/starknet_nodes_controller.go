package web

import (
	"github.com/smartcontractkit/chainlink-starknet/relayer/pkg/chainlink/db"

	"github.com/smartcontractkit/chainlink/core/services/chainlink"
	"github.com/smartcontractkit/chainlink/core/web/presenters"
)

// ErrStarkNetNotEnabled is returned when STARKNET_ENABLED is not true.
var ErrStarkNetNotEnabled = errChainDisabled{name: "StarkNet", envVar: "STARKNET_ENABLED"}

func NewStarkNetNodesController(app chainlink.Application) NodesController {
	parse := func(s string) (string, error) { return s, nil }
	return newNodesController[string, db.Node, presenters.StarkNetNodeResource](
		app.GetChains().StarkNet, ErrStarkNetNotEnabled, parse, presenters.NewStarkNetNodeResource, app.GetAuditLogger())
}
