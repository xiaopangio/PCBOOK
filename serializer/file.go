package serializer

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"os"
)

func WriteProtobufToJsonFile(message proto.Message, filename string) error {

	json, err := ProtobufToJson(message)
	if err != nil {
		fmt.Errorf("cannot marshal proto message to Json file:%w", err)
	}
	err = os.WriteFile(filename, []byte(json), 0644)
	if err != nil {
		fmt.Errorf("cannot write Json data to file: %w", err)
	}
	return nil
}
func WriteProtoBufToBinaryFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to binany file: %w", err)
	}
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("cannot write binary data to file: %w", err)
	}
	return nil
}
func ReadProtoBufFromBinaryFile(filename string, message proto.Message) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("cannot read protobuf message from binary file: %w", err)
	}
	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("cannot unmarshal data to message: %w", err)
	}
	return nil
}
func ReadProtobufFromJsonFile(filename string, message proto.Message) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("cannot read protobuf message from Json file: %w", err)
	}
	err = JsonToProtobuf(data, message)
	if err != nil {
		return fmt.Errorf("cannot unmarshal json to protobuf message: %w", err)
	}
	return nil
}
