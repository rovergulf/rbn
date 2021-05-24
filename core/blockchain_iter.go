package core

import (
	"github.com/dgraph-io/badger/v3"
	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

// BlockchainIterator is used to iterate over blockchain blocks
type BlockchainIterator struct {
	CurrentHash []byte
	Db          *badger.DB
	logger      *zap.SugaredLogger
	tracer      opentracing.Tracer
}

// Iterator returns a BlockchainIterator
func (bc *Blockchain) Iterator() *BlockchainIterator {
	bci := &BlockchainIterator{
		CurrentHash: bc.LastHash,
		Db:          bc.Db,
		logger:      bc.logger,
		tracer:      bc.tracer,
	}

	return bci
}

// Next returns next block starting from the tip
func (i *BlockchainIterator) Next() (*Block, error) {
	var block *Block

	if err := i.Db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(i.CurrentHash)
		if err != nil {
			return err
		}

		return item.Value(func(val []byte) error {
			nextBlock, err := DeserializeBlock(val)
			if err != nil {
				return err
			} else {
				block = nextBlock
			}
			return nil
		})
	}); err != nil {
		i.logger.Errorw("Unable to iterate db view",
			"current_hash", i.CurrentHash, "err", err,
		)
		return nil, err
	}

	i.CurrentHash = block.PrevHash

	return block, nil
}
