package collectjs

import (
	"bytes"
	"fmt"

	"github.com/tidwall/gjson"
	"github.com/tkeel-io/collectjs/pkg/json/jsonparser"
)

func Get(raw []byte, path string) ([]byte, jsonparser.ValueType, error) {
	keys := path2JSONPARSER(path)

	if value, dataType, _, err := jsonparser.Get(raw, keys...); err == nil {
		return warpValue(dataType, value), dataType, nil
	}
	path = path2GJSON(path)
	ret := gjson.GetBytes(raw, path)
	if gjson.Null == ret.Type {
		return nil, jsonparser.NotExist, jsonparser.KeyPathNotFoundError
	}
	return []byte(ret.String()), convertType(ret), nil
}

func warpValue(dataType jsonparser.ValueType, value []byte) []byte {
	switch dataType {
	case jsonparser.String:
		return bytes.Join([][]byte{
			[]byte("\""), value, []byte("\""),
		}, []byte{})
	default:
		return value
	}
	return nil
}

func Set(raw []byte, path string, value []byte) ([]byte, error) {
	keys := path2JSONPARSER(path)
	return jsonparser.Set(raw, value, keys...)
}

func Append(raw []byte, path string, value []byte) ([]byte, error) {
	keys := path2JSONPARSER(path)
	return jsonparser.Append(raw, value, keys...)
}

func Del(raw []byte, path ...string) []byte {
	for _, v := range path {
		keys := path2JSONPARSER(v)
		raw = jsonparser.Delete(raw, keys...)
	}
	return raw
}

func ForEach(raw []byte, datatype jsonparser.ValueType, fn func(key []byte, value []byte, dataType jsonparser.ValueType)) []byte {
	// dispose object.
	if datatype == jsonparser.Object {
		jsonparser.ObjectEach(raw, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
			fn(key, warpValue(dataType, value), dataType)
			return nil
		})
	}

	// dispose array.
	if datatype == jsonparser.Array {
		idx := 0
		jsonparser.ArrayEach(raw, func(value []byte, dataType jsonparser.ValueType, offset int) error {
			fn(Byte(fmt.Sprintf("[%d]", idx)), warpValue(dataType, value), dataType)
			idx++
			return nil
		})
	}
	return raw
}
