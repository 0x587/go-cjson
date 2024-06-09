package cjson

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"
	"math"

	"github.com/0x587/go-cjson/cjson/pb"
	"google.golang.org/protobuf/proto"
)

func (i *impl) MarshalGzip(j []byte) ([]byte, error) {
	bs, err := i.Marshal(j)
	if err != nil {
		return nil, err
	}
	buf := &bytes.Buffer{}
	zw := gzip.NewWriter(buf)
	zw.Write(bs)
	zw.Close()
	return io.ReadAll(buf)
}

func (i *impl) Marshal(j []byte) ([]byte, error) {
	item, err := i.Pack(j)
	if err != nil {
		return nil, err
	}
	return proto.Marshal(item)
}

func (i *impl) Pack(j []byte) (*pb.Item, error) {
	var data any
	json.Unmarshal(j, &data)
	return i.pack(data)
}

func (i *impl) pack(v any) (*pb.Item, error) {
	switch value := v.(type) {
	// object
	case map[string]any:
		return i.packObj(value)
	// list
	case []any:
		return i.packArr(value)
	// raw value
	case string:
		return &pb.Item{
			RawValue: &pb.RawValue{
				TestOneof: &pb.RawValue_Str{Str: value}}}, nil
	case bool:
		return &pb.Item{
			RawValue: &pb.RawValue{
				TestOneof: &pb.RawValue_Bool{Bool: value}}}, nil
	case float64:
		if value == math.Trunc(value) && value > math.MinInt32 && value < math.MaxInt32 {
			return &pb.Item{
				RawValue: &pb.RawValue{
					TestOneof: &pb.RawValue_NumberInt32{NumberInt32: int32(value)}}}, nil
		}
		return &pb.Item{
			RawValue: &pb.RawValue{
				TestOneof: &pb.RawValue_NumberDouble{NumberDouble: value}}}, nil
	default:
		return nil, nil
	}
}

func (i *impl) packObj(obj map[string]any) (*pb.Item, error) {
	template := &pb.Template{}
	items := make([]*pb.Item, 0, len(obj))
	for key, val := range obj {
		template.Values = append(template.Values, key)
		item, err := i.pack(val)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	hash, err := i.setTemplate(template)
	if err != nil {
		return nil, err
	}
	return &pb.Item{
		IsArr:        false,
		TemplateHash: hash,
		Items:        items,
	}, nil
}

func (i *impl) packArr(obj []any) (*pb.Item, error) {
	items := make([]*pb.Item, 0, len(obj))
	for _, val := range obj {
		item, err := i.pack(val)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return &pb.Item{
		IsArr: true,
		Items: items,
	}, nil
}
