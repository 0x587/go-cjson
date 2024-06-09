package storage

import (
	"crypto/md5"
	"errors"
)

type StorageIF interface {
	Set(value []byte) ([]byte, error)
	Get(hash []byte) ([]byte, error)
}

type memImpl struct {
	table map[string][]byte
}

func NewMem() StorageIF {
	return &memImpl{
		table: make(map[string][]byte),
	}
}

func (i *memImpl) Set(value []byte) ([]byte, error) {
	hasher := md5.New()
	hasher.Write(value)
	hash := hasher.Sum(nil)
	// TODO: 缩短ID
	hash = []byte{hash[0], hash[1]}
	i.table[string(hash)] = value
	return hash, nil
}

func (i *memImpl) Get(hash []byte) ([]byte, error) {
	v, ok := i.table[string(hash)]
	if !ok {
		return nil, errors.New("not found")
	}
	return v, nil
}
