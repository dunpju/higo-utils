package protoutil

import (
	"github.com/golang/protobuf/proto"
)

//marshal:  obj---[]byte
func ProtoMarshal(m proto.Message) []byte {
	bytes, err := proto.Marshal(m)
	if err != nil {
		panic(err)
	}
	return bytes
}

//unmarshal : []byte---obj
func ProtoUnmarshal(b []byte, m proto.Message) {
	err := proto.Unmarshal(b, m)
	if err != nil {
		panic(err)
	}
}