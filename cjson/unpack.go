package cjson

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"

	"github.com/0x587/go-cjson/cjson/pb"
	"google.golang.org/protobuf/proto"
)

func (i *impl) UnmarshalGzip(fieldBuf []byte, payloadBuf []byte) ([]byte, error) {
	fieldBuf, err := i.ungzip(fieldBuf)
	if err != nil {
		return nil, err
	}
	payloadBuf, err = i.ungzip(payloadBuf)
	if err != nil {
		return nil, err
	}
	return i.Unmarshal(fieldBuf, payloadBuf)
}

func (i *impl) ungzip(v []byte) ([]byte, error) {
	buf := bytes.NewBuffer(v)
	zr, err := gzip.NewReader(buf)
	if err != nil {
		return nil, err
	}
	return io.ReadAll(zr)
}

func (i *impl) Unmarshal(fieldBuf []byte, payloadBuf []byte) ([]byte, error) {
	field, payload := &pb.Field{}, &pb.Payload{}
	if err := proto.Unmarshal(fieldBuf, field); err != nil {
		return nil, err
	}
	if err := proto.Unmarshal(payloadBuf, payload); err != nil {
		return nil, err
	}
	return i.Unpack(field, payload)
}

func (i *impl) Unpack(fields *pb.Field, payload *pb.Payload) ([]byte, error) {
	data, _, err := i.unpack(fields, payload.GetValues())
	if err != nil {
		return nil, err
	}
	res, err := json.Marshal(data)
	return res, err
}

func (i *impl) unpack(field *pb.Field, values []*pb.Value) (any, []*pb.Value, error) {
	if field.GetField() == nil {
		return nil, values, nil
	}
	switch field := field.GetField().(type) {
	case *pb.Field_Raw:
		switch field.Raw {
		case pb.RawFieldType_STRING:
			return values[0].GetStr(), values[1:], nil
		case pb.RawFieldType_NUMBER:
			return values[0].GetNumberDouble(), values[1:], nil
		case pb.RawFieldType_BOOL:
			return values[0].GetBool(), values[1:], nil
		default:
			panic("ERROR")
		}
	case *pb.Field_Obj:
		return i.unpackObj(field, values)
	case *pb.Field_Arr:
		return i.unpackArr(field, values)
	default:
		panic("ERROR")
	}
}

func (i *impl) unpackObj(field *pb.Field_Obj, values []*pb.Value) (map[string]any, []*pb.Value, error) {
	// TODO: 这里会导致Object无序
	res := make(map[string]any, len(field.Obj.Fields))
	var err error
	var v any
	for index, f := range field.Obj.Fields {
		v, values, err = i.unpack(f, values)
		if err != nil {
			return nil, nil, err
		}
		res[field.Obj.Keys[index]] = v
	}
	return res, values, nil
}

func (i *impl) unpackArr(field *pb.Field_Arr, values []*pb.Value) ([]any, []*pb.Value, error) {
	res := make([]any, 0, len(field.Arr.Items))
	var err error
	var v any
	for _, f := range field.Arr.Items {
		v, values, err = i.unpack(f, values)
		if err != nil {
			return nil, nil, err
		}
		res = append(res, v)
	}
	return res, values, nil
}
