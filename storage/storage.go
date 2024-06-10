package storage

import (
	"crypto/md5"
	"errors"
)

type ID []byte

type StorageIF interface {
	Set(value []byte) (ID, error)
	Get(id ID) ([]byte, error)
}

type memImpl struct {
	table map[string][]byte
}

func NewMem() StorageIF {
	return &memImpl{
		table: make(map[string][]byte),
	}
}

func (i *memImpl) Set(value []byte) (ID, error) {
	hasher := md5.New()
	hasher.Write(value)
	hash := hasher.Sum(nil)
	i.table[string(hash)] = value
	return hash, nil
}

func (i *memImpl) Get(id ID) ([]byte, error) {
	v, ok := i.table[string(id)]
	if !ok {
		return nil, errors.New("not found")
	}
	return v, nil
}
