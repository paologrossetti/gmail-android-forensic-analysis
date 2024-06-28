package conversation

import (
	ItemSummaryProto "github.com/paologrossetti/gmail-android/internal/proto/ItemSummary"
	"github.com/paologrossetti/gmail-android/pkg/sqlite3"
	"google.golang.org/protobuf/proto"
)

func GetItems(filename string) []map[string]interface{} {
	var results []map[string]interface{} = sqlite3.SelectAll(filename, "items")
	return results
}

func ItemSummaryProtoDecoder(item_summary_proto interface{}) *ItemSummaryProto.ItemSummary {
	conversation_summary_proto_bytes, _ := item_summary_proto.([]byte)
	item_summary := new(ItemSummaryProto.ItemSummary)
	proto.Unmarshal(conversation_summary_proto_bytes, item_summary)
	return item_summary
}
