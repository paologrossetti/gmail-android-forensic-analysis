package message

import (
	"bytes"
	"compress/zlib"
	"io"
	"log"

	MessageProto "github.com/paologrossetti/gmail-android/internal/proto/Message"
	"github.com/paologrossetti/gmail-android/pkg/sqlite3"
	"google.golang.org/protobuf/proto"
)

func GetItemMessages(filename string) []map[string]interface{} {
	var results []map[string]interface{} = sqlite3.SelectAll(filename, "item_messages")
	return results
}

func GetItemMessagesByItem(filename string, items_row_id int64) []map[string]interface{} {
	var results []map[string]interface{} = sqlite3.SelectAllWhere(filename, "item_messages", "items_row_id", items_row_id)
	return results
}

func unzipMessageProto(zipped_message_proto interface{}) []byte {
	zipped_message_proto_bytes, _ := zipped_message_proto.([]byte)
	zipped_message_proto_bytes = bytes.TrimLeft(zipped_message_proto_bytes, "\x00")
	bytes_reader := bytes.NewReader(zipped_message_proto_bytes)
	uncompressed_data_reader, err := zlib.NewReader(bytes_reader)
	if err != nil {
		panic(err)
	}
	unzipped_message_proto_bytes, err := io.ReadAll(uncompressed_data_reader)
	if err != nil {
		log.Fatal(err)
	}

	// out := hex.EncodeToString(unzipped_message_proto_bytes)
	// fmt.Println(out)

	return unzipped_message_proto_bytes
}

func ZippedMessageProtoDecoder(zipped_message_proto interface{}) *MessageProto.Message {
	message_proto_bytes := unzipMessageProto(zipped_message_proto)
	message := new(MessageProto.Message)
	proto.Unmarshal(message_proto_bytes, message)
	return message
}
