package solana

import (
	"github.com/smartcontractkit/sqlx"

	soldb "github.com/smartcontractkit/chainlink-solana/pkg/solana/db"

	"github.com/smartcontractkit/chainlink/core/chains"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/services/pg"
)

type DBChain = chains.DBChain[string, *soldb.ChainCfg]

// ORM manages solana chains and nodes.
type ORM interface {
	Chain(string, ...pg.QOpt) (DBChain, error)
	Chains(offset, limit int, qopts ...pg.QOpt) ([]DBChain, int, error)
	GetChainsByIDs(ids []string) (chains []DBChain, err error)
	EnabledChains(...pg.QOpt) ([]DBChain, error)

	GetNodesByChainIDs(chainIDs []string, qopts ...pg.QOpt) (nodes []soldb.Node, err error)
	NodeNamed(string, ...pg.QOpt) (soldb.Node, error)
	Nodes(offset, limit int, qopts ...pg.QOpt) (nodes []soldb.Node, count int, err error)
	NodesForChain(chainID string, offset, limit int, qopts ...pg.QOpt) (nodes []soldb.Node, count int, err error)

	EnsureChains([]string, ...pg.QOpt) error
}

var _ chains.ORM[string, *soldb.ChainCfg, soldb.Node] = (ORM)(nil)

// NewORM returns an ORM backed by db.
// https://app.shortcut.com/chainlinklabs/story/33622/remove-legacy-config
func NewORM(db *sqlx.DB, lggr logger.Logger, cfg pg.QConfig) ORM {
	q := pg.NewQ(db, lggr.Named("ORM"), cfg)
	return chains.NewORM[string, *soldb.ChainCfg, soldb.Node](q, "solana", "solana_url")
}

func NewORMImmut(cfgs chains.ChainConfig[string, *soldb.ChainCfg, soldb.Node]) ORM {
	return chains.NewORMImmut(cfgs)
}
