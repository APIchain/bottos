package block

import "github.com/bottos-project/bottos/common/types"

const (
	BLOCK_REQ = 1
	//BLOCK_INFO update or response
	BLOCK_UPDATE = 2

	LAST_BLOCK_NUMBER_REQ = 3
	LAST_BLOCK_NUMBER_RSP = 4

	BLOCK_HEADER_REQ = 5
	BLOCK_HEADER_RSP = 6
)

type blockHeaderReq struct {
	Begin uint32
	End   uint32
}

type blockHeaderRsp struct {
	set []types.Header
}

type blockUpdate struct {
	index uint16
	block *types.Block
}