// Package serializer  @Author xiaobaiio 2023/2/21 10:19:00
package serializer

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type ProtoWithName struct {
	ProtoMsgName string
	ProtoMsgData string
}

func ProtobufToJson(message proto.Message) (string, error) {
	marshaller := &protojson.MarshalOptions{
		Multiline:       true,
		Indent:          "  ",
		AllowPartial:    false,
		UseProtoNames:   true,
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
	}
	bytes, err := marshaller.Marshal(message)
	return string(bytes), err
}
func JsonToProtobuf(data []byte, message proto.Message) error {
	err := protojson.Unmarshal(data, message)
	if err != nil {
		return err
	}
	return nil
}
