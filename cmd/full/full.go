package main

import (
	"flag"
	"fmt"

	"github.com/paologrossetti/gmail-android/internal/data/conversation"
	"github.com/paologrossetti/gmail-android/internal/data/message"
	"github.com/paologrossetti/gmail-android/internal/outformat"
)

var (
	filename = flag.String("filepath", "", "Android GMAIL Sqlite3 database filepath (bigTopDataDB)")
	format   = flag.String("format", "JSON", "Output format: JSON or YAML")
)

func main() {
	fmt.Println(`
	$ full.go --help
	A tool to parse bigTopDataDB SQLite database

	Usage: full -filepath <FILEPATH bigTopDataDB> -format <JSON, YAML>
	`)
	flag.Parse()

	items := conversation.GetItems(*filename)
	for _, item := range items {
		item["item_summary"] = conversation.ItemSummaryProtoDecoder(item["item_summary_proto"])
		item_messages := message.GetItemMessagesByItem(*filename, item["row_id"].(int64))
		var messages []map[string]interface{}
		for _, item_message := range item_messages {
			item_message["message_proto"] = message.ZippedMessageProtoDecoder(item_message["zipped_message_proto"])
			messages = append(messages, item_message)
		}
		item["messages"] = messages
	}

	if *format == "YAML" {
		outformat.ToYaml(items)
	} else {
		outformat.ToJson(items)
	}
}
