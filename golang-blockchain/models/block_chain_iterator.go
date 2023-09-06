package models

import "github.com/dgraph-io/badger/v4"

type BlockChainIterator struct {
	CurrentHash []byte
	Database    *badger.DB
}

func (iter *BlockChainIterator) Next() *Block {
	var block *Block

	err := iter.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iter.CurrentHash)
		HandleError(err)
		encodedBlock, err := item.ValueCopy(nil)
		block = DeserializeBlock(encodedBlock)

		return err
	})
	HandleError(err)

	iter.CurrentHash = block.PrevHash

	return block
}
