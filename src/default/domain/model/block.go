package model

// Block is transaction
type Block struct {
	Data          []byte
	Hash          []byte
	PrevBlockHash []byte
	CreatedAt     int64
}
