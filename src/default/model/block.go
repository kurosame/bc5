package model

type Block struct {
	Data          []byte
	Hash          []byte
	PrevBlockHash []byte
	CreatedAt     int64
}
