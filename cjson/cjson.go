package cjson

import (
	"errors"

	"github.com/0x587/go-cjson/cjson/pb"
)

var (
	ErrLengthNotMatch = errors.New("the template length and value length are not equal")
	ErrMissValues     = errors.New("missing values")
	ErrNoTemplate     = errors.New("no template")
)

type IF interface {
	// Pack pack json to field and payload
	Pack(json []byte) (*pb.Field, *pb.Payload, error)
	// Unpack unpack field and payload to json
	Unpack(fields *pb.Field, values *pb.Payload) ([]byte, error)
	// Marshal pack json and marshal to bytes
	Marshal(json []byte) ([]byte, []byte, error)
	// Unmarshal unpack json from bytes
	Unmarshal(fields []byte, values []byte) ([]byte, error)
	// MarshalObj pack obj and marshal to bytes
	MarshalObj(obj any) (schema []byte, value []byte, err error)
	// UnmarshalObj unpack obj from bytes
	UnmarshalObj(schema []byte, value []byte) (any, error)
	// MarshalGzip marshal and gzip
	MarshalGzip(json []byte) ([]byte, []byte, error)
	// UnmarshalGzip unmarshal and gzip
	UnmarshalGzip(fields []byte, values []byte) ([]byte, error)
}

func New() IF {
	return &impl{}
}

type impl struct{}
