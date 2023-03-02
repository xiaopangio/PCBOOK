package serializer

import (
	"github.com/xiaopangio/pcbook/pb"
	"google.golang.org/protobuf/proto"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xiaopangio/pcbook/sample"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()
	binaryFile := "../tmp/laptap.bin"
	jsonFile := "../tmp/laptap.json"
	laptap1 := sample.NewLaptap()
	err := WriteProtoBufToBinaryFile(laptap1, binaryFile)
	require.NoError(t, err)
	laptap2 := &pb.Laptap{}
	err = ReadProtoBufFromBinaryFile(binaryFile, laptap2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptap1, laptap2))
	err = WriteProtobufToJsonFile(laptap1, jsonFile)
	require.NoError(t, err)
	laptap3 := &pb.Laptap{}
	err = ReadProtobufFromJsonFile(jsonFile, laptap3)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptap1, laptap3))
}
