package qr

type block struct {
	data []byte
	ecc  []byte
}
type blockList []*block

func splitToBlocks(data <-chan byte, vi *versionInfo) blockList {
	_ = "STUB: not implemented"
	return *new(blockList)
}

func (bl blockList) interleave(vi *versionInfo) []byte { _ = "STUB: not implemented"; return nil }
