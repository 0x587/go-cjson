package cjson

import (
	"errors"

	"github.com/0x587/go-cjson/cjson/pb"
	"github.com/0x587/go-cjson/cjson/storage"
	"google.golang.org/protobuf/proto"
)

var (
	ErrLengthNotMatch = errors.New("the template length and value length are not equal")
	ErrMissValues     = errors.New("missing values")
	ErrNoTemplate     = errors.New("no template")
)

type Hash []byte

type IF interface {
	Pack(json []byte) (*pb.Item, error)
	Unpack(value *pb.Item) ([]byte, error)
	Marshal(json []byte) ([]byte, error)
	Unmarshal(value []byte) ([]byte, error)
	MarshalGzip(json []byte) ([]byte, error)
	UnmarshalGzip(value []byte) ([]byte, error)
}

func New(s storage.StorageIF) IF {
	return &impl{
		storage: s,
	}
}

type impl struct {
	storage storage.StorageIF
}

func (i *impl) setTemplate(t *pb.Template) (Hash, error) {
	buf, err := proto.Marshal(t)
	if err != nil {
		return nil, err
	}
	return i.storage.Set(buf)
}

func (i *impl) getTemplate(hash Hash) (*pb.Template, error) {
	buf, err := i.storage.Get(hash)
	if err != nil {
		return nil, err
	}
	t := &pb.Template{}
	if err = proto.Unmarshal(buf, t); err != nil {
		return nil, err
	}
	return t, nil
}
