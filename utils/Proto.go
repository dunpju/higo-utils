package utils

import (
	"github.com/golang/protobuf/proto"
)

func ProtoMarshal(m proto.Message) []byte {
	//marshal:  obj---[]byte
	bytes, err := proto.Marshal(m)
	if err != nil {
		panic(err)
	}
	return bytes
}

func ProtoUnmarshal(b []byte, m proto.Message) {
	//unmarshal : []byte---obj
	err := proto.Unmarshal(b, m)
	if err != nil {
		panic(err)
	}
}