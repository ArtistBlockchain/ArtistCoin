package dummy

import (
	"github.com/ArtistBlockchain/ArtistCoin/blockchain"
	"github.com/ArtistBlockchain/ArtistCoin/consensus"
	"github.com/ArtistBlockchain/ArtistCoin/helper/progress"
	"github.com/ArtistBlockchain/ArtistCoin/state"
	"github.com/ArtistBlockchain/ArtistCoin/txpool"
	"github.com/ArtistBlockchain/ArtistCoin/types"
	"github.com/hashicorp/go-hclog"
)

type Dummy struct {
	sealing    bool
	logger     hclog.Logger
	notifyCh   chan struct{}
	closeCh    chan struct{}
	txpool     *txpool.TxPool
	blockchain *blockchain.Blockchain
	executor   *state.Executor
}

func Factory(params *consensus.Params) (consensus.Consensus, error) {
	logger := params.Logger.Named("dummy")

	d := &Dummy{
		sealing:    params.Seal,
		logger:     logger,
		notifyCh:   make(chan struct{}),
		closeCh:    make(chan struct{}),
		blockchain: params.Blockchain,
		executor:   params.Executor,
		txpool:     params.TxPool,
	}

	return d, nil
}

// Initialize initializes the consensus
func (d *Dummy) Initialize() error {
	return nil
}

func (d *Dummy) Start() error {
	go d.run()

	return nil
}

func (d *Dummy) VerifyHeader(header *types.Header) error {
	// All blocks are valid
	return nil
}

func (d *Dummy) ProcessHeaders(headers []*types.Header) error {
	return nil
}

func (d *Dummy) GetBlockCreator(header *types.Header) (types.Address, error) {
	return types.BytesToAddress(header.Miner), nil
}

// PreCommitState a hook to be called before finalizing state transition on inserting block
func (d *Dummy) PreCommitState(_header *types.Header, _txn *state.Transition) error {
	return nil
}

func (d *Dummy) GetSyncProgression() *progress.Progression {
	return nil
}

func (d *Dummy) Close() error {
	close(d.closeCh)

	return nil
}

func (d *Dummy) run() {
	d.logger.Info("started")
	// do nothing
	<-d.closeCh
}
