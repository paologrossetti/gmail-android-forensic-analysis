package attachment

import "github.com/paologrossetti/gmail-android/pkg/sqlite3"

func GetAttachments(filename string) []map[string]interface{} {
	var results []map[string]interface{} = sqlite3.SelectAll(filename, "item_message_attachments")
	return results
}

func GetAttachmentsbyMessage(filename string, item_messages_row_id int64) []map[string]interface{} {
	var results []map[string]interface{} = sqlite3.SelectAllWhere(filename, "item_message_attachments", "item_messages_row_id", item_messages_row_id)
	return results
}
