package cjson

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"

	"github.com/0x587/go-cjson/cjson/pb"
	"google.golang.org/protobuf/proto"
)

func (i *impl) UnmarshalGzip(value []byte) ([]byte, error) {
	buf := bytes.NewBuffer(value)
	zr, err := gzip.NewReader(buf)
	if err != nil {
		return nil, err
	}
	bs, err := io.ReadAll(zr)
	if err != nil {
		return nil, err
	}
	return i.Unmarshal(bs)
}

func (i *impl) Unmarshal(value []byte) ([]byte, error) {
	item := &pb.Item{}
	if err := proto.Unmarshal(value, item); err != nil {
		return nil, err
	}
	return i.Unpack(item)
}

func (i *impl) Unpack(r *pb.Item) ([]byte, error) {
	data, err := i.unpack(r)
	if err != nil {
		return nil, err
	}
	res, err := json.Marshal(data)
	return res, err
}

func (i *impl) unpack(item *pb.Item) (any, error) {
	if len(item.TemplateHash) == 0 && !item.IsArr {
		switch item.RawValue.GetTestOneof().(type) {
		case *pb.RawValue_Str:
			return item.RawValue.GetStr(), nil
		case *pb.RawValue_NumberDouble:
			return item.RawValue.GetNumberDouble(), nil
		case *pb.RawValue_NumberInt32:
			return item.RawValue.GetNumberInt32(), nil
		case *pb.RawValue_Bool:
			return item.RawValue.GetBool(), nil
		default:
			return nil, nil
		}
	}
	if item.IsArr {
		return i.unpackArr(item.Items)
	}
	template, err := i.getTemplate(item.TemplateHash)
	if err != nil {
		return nil, err
	}
	if len(template.Values) != len(item.Items) {
		return nil, ErrLengthNotMatch
	}
	return i.unpackObj(template, item.Items)
}

func (i *impl) unpackObj(template *pb.Template, items []*pb.Item) (map[string]any, error) {
	// TODO: 这里会导致Object无序
	res := make(map[string]any, len(items))
	for index, v := range items {
		v, err := i.unpack(v)
		if err != nil {
			return nil, err
		}
		res[template.Values[index]] = v
	}
	return res, nil
}

func (i *impl) unpackArr(items []*pb.Item) ([]any, error) {
	res := make([]any, 0, len(items))
	for _, v := range items {
		v, err := i.unpack(v)
		if err != nil {
			return nil, err
		}
		res = append(res, v)
	}
	return res, nil
}
