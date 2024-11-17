package cjson

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io"

	"github.com/0x587/go-cjson/cjson/pb"
	"google.golang.org/protobuf/proto"
)

func (i *impl) MarshalGzip(j []byte) ([]byte, []byte, error) {
	field, payload, err := i.Marshal(j)
	if err != nil {
		return nil, nil, err
	}
	payloadBuf, err := i.gzip(payload)
	if err != nil {
		return nil, nil, err
	}
	fieldBuf, err := i.gzip(field)
	if err != nil {
		return nil, nil, err
	}
	return fieldBuf, payloadBuf, nil
}

func (i *impl) gzip(v []byte) ([]byte, error) {
	buf := &bytes.Buffer{}
	zw := gzip.NewWriter(buf)
	zw.Write(v)
	zw.Close()
	return io.ReadAll(buf)
}

func (i *impl) Marshal(j []byte) ([]byte, []byte, error) {
	field, payload, err := i.Pack(j)
	if err != nil {
		return nil, nil, err
	}
	fieldBuf, err := proto.Marshal(field)
	if err != nil {
		return nil, nil, err
	}
	payloadBuf, err := proto.Marshal(payload)
	if err != nil {
		return nil, nil, err
	}
	return fieldBuf, payloadBuf, nil
}

func (i *impl) MarshalObj(obj any) ([]byte, []byte, error) {
	field, values, err := i.pack(obj)
	if err != nil {
		return nil, nil, err
	}
	fieldBuf, err := proto.Marshal(field)
	if err != nil {
		return nil, nil, err
	}
	payloadBuf, err := proto.Marshal(&pb.Payload{Values: values})
	if err != nil {
		return nil, nil, err
	}
	return fieldBuf, payloadBuf, nil
}

func (i *impl) Pack(j []byte) (*pb.Field, *pb.Payload, error) {
	var data any
	json.Unmarshal(j, &data)
	field, values, err := i.pack(data)
	if err != nil {
		return nil, nil, err
	}
	return field, &pb.Payload{Values: values}, nil
}

func (i *impl) pack(v any) (*pb.Field, []*pb.Value, error) {
	switch value := v.(type) {
	// object
	case map[string]any:
		return i.packObj(value)
	// list
	case []any:
		return i.packArr(value)
	// raw value
	case string:
		return &pb.Field{
			Field: &pb.Field_Raw{Raw: pb.RawFieldType_STRING}}, []*pb.Value{{Value: &pb.Value_Str{Str: value}}}, nil
	case bool:
		return &pb.Field{
			Field: &pb.Field_Raw{Raw: pb.RawFieldType_BOOL}}, []*pb.Value{{Value: &pb.Value_Bool{Bool: value}}}, nil
	case float64:
		return &pb.Field{
			Field: &pb.Field_Raw{Raw: pb.RawFieldType_NUMBER}}, []*pb.Value{{Value: &pb.Value_NumberDouble{NumberDouble: value}}}, nil
	default:
		return nil, nil, nil
	}
}

func (i *impl) packObj(obj map[string]any) (*pb.Field, []*pb.Value, error) {
	keys := make([]string, 0)
	fields := make([]*pb.Field, 0)
	values := make([]*pb.Value, 0)
	for key, val := range obj {
		field, vs, err := i.pack(val)
		values = append(values, vs...)
		if err != nil {
			return nil, nil, err
		}
		keys = append(keys, key)
		fields = append(fields, field)
	}
	return &pb.Field{Field: &pb.Field_Obj{
		Obj: &pb.ObjectField{
			Keys:   keys,
			Fields: fields,
		},
	}}, values, nil
}

func (i *impl) packArr(obj []any) (*pb.Field, []*pb.Value, error) {
	items := make([]*pb.Field, 0, len(obj))
	values := make([]*pb.Value, 0)
	for _, val := range obj {
		item, vs, err := i.pack(val)
		values = append(values, vs...)
		if err != nil {
			return nil, nil, err
		}
		items = append(items, item)
	}
	return &pb.Field{
		Field: &pb.Field_Arr{
			Arr: &pb.ArraryField{
				Items: items,
			},
		},
	}, values, nil
}
