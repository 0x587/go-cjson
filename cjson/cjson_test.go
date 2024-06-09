package cjson_test

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"testing"

	"github.com/0x587/go-cjson/cjson"
	"github.com/0x587/go-cjson/cjson/storage"
)

func TestPack(t *testing.T) {
	table := []string{
		`{"key1":123}`,
		`{"key1":123,"key2":"abc","key3":true,"key4":123.45}`,
		`{"key1":{"key11":1,"key12":"123"}}`,
		`{"key1":[{"key11":1,"key12":"123"},{"key11":2,"key12":"234"}]}`,
		`{"key1":[1,2,3,4,5],"key2":[true,true,true,false],"key3":null}`,
	}
	for _, v := range table {
		testPack(v, t)
		testPackGzip(v, t)
	}
}

func testPack(str string, t *testing.T) {
	cjson := cjson.New(storage.NewMem())
	r, err := cjson.Marshal([]byte(str))
	if err != nil {
		t.Fatal(err)
	}
	res, err := cjson.Unmarshal(r)
	if err != nil {
		t.Fatal(err)
	}
	if string(res) != str {
		t.Errorf("unpack!=pack\npack: %s\nunpack:%s", str, res)
	}
	originSize := len([]byte(str))
	packSize := len(r)
	t.Logf("originSize: %d, packSize: %d, rate: %f",
		originSize, packSize, float32(packSize)/float32(originSize))
}

func testPackGzip(str string, t *testing.T) {
	cjson := cjson.New(storage.NewMem())
	r, err := cjson.MarshalGzip([]byte(str))
	if err != nil {
		t.Fatal(err)
	}
	res, err := cjson.UnmarshalGzip(r)
	if err != nil {
		t.Fatal(err)
	}
	if string(res) != str {
		t.Errorf("unpack!=pack\npack: %s\nunpack:%s", str, res)
	}
	originSize := len([]byte(str))
	packSize := len(r)
	t.Logf("gzip originSize: %d, packSize: %d, rate: %f",
		originSize, packSize, float32(packSize)/float32(originSize))
}

const randomJson = `{"user":{"id":12345,"name":"John Doe","email":"john.doe@example.com","isActive":true,"preferences":{"theme":"dark","notifications":{"email":true,"sms":false}},"roles":["admin","editor"],"createdAt":"2024-06-09T12:00:00Z"},"products":[{"id":987,"name":"Laptop","price":999.99,"inStock":50},{"id":654,"name":"Smartphone","price":499.99,"inStock":150}],"orders":[{"orderId":1,"productId":987,"quantity":1,"totalPrice":999.99,"status":"shipped"},{"orderId":2,"productId":654,"quantity":2,"totalPrice":999.98,"status":"processing"}]}`

func TestABC(t *testing.T) {
	originSize := len([]byte(randomJson))
	cjson := cjson.New(storage.NewMem())
	r, _ := cjson.MarshalGzip([]byte(randomJson))
	packSize := len(r)
	buf := &bytes.Buffer{}
	zw := gzip.NewWriter(buf)
	zw.Write([]byte(randomJson))
	zw.Close()
	t.Logf("%d %f\n", buf.Len(), float64(buf.Len())/float64(len(randomJson)))
	t.Logf("originSize: %d, packSize: %d, rate: %f",
		originSize, packSize, float32(packSize)/float32(originSize))
}

func BenchmarkCJson(b *testing.B) {
	cjson := cjson.New(storage.NewMem())
	r, _ := cjson.Marshal([]byte(randomJson))
	_, _ = cjson.Unmarshal(r)
}

func BenchmarkCJsonGzip(b *testing.B) {
	cjson := cjson.New(storage.NewMem())
	r, _ := cjson.MarshalGzip([]byte(randomJson))
	_, _ = cjson.UnmarshalGzip(r)
}

func BenchmarkJson(b *testing.B) {
	var out any
	json.Unmarshal([]byte(randomJson), out)
	_, _ = json.Marshal(out)
}
