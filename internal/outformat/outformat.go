package outformat

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/k3a/html2text"
	"github.com/paologrossetti/gmail-android/internal/data/labeltype"
	ItemSummaryProto "github.com/paologrossetti/gmail-android/internal/proto/ItemSummary"
	MessageProto "github.com/paologrossetti/gmail-android/internal/proto/Message"
	"github.com/paologrossetti/gmail-android/internal/utilities/extractaddress"
	"github.com/paologrossetti/gmail-android/internal/utilities/mstotime"
	"gopkg.in/yaml.v2"
)

func generateOutput(items []map[string]interface{}) []map[string]interface{} {
	var full []map[string]interface{}
	for _, item := range items {
		conversation := make(map[string]interface{})
		conversation["id"] = item["row_id"].(int64)
		conversation["server_perm_id"] = item["server_perm_id"].(string)
		conversation["legacy_storage_id"] = item["legacy_storage_id"].(int64)
		item_summary := item["item_summary"].(*ItemSummaryProto.ItemSummary)
		conversation["date"] = mstotime.MsToTime(item_summary.GetConv().GetEpoch())
		conversation["subject"] = item_summary.GetConv().GetSubject()
		conversation["snippet"] = item_summary.GetConv().GetSnippet()

		messages := item["messages"].([]map[string]interface{})
		var message_elaborated []map[string]interface{}
		for idx, msg := range messages {
			message := make(map[string]interface{})

			labels_message := item_summary.GetMsgs()[idx].GetLabelTypes()
			message["labels_type"] = labeltype.TranslationLabels(labels_message)
			message["is_inbox"] = labeltype.IsInbox(labels_message)
			message["is_sent"] = labeltype.IsSent(labels_message)
			message["is_opened"] = labeltype.IsOpened(labels_message)
			message["is_unread"] = labeltype.IsUnread(labels_message)
			message["is_trash"] = labeltype.IsTrash(labels_message)
			message["is_starred"] = labeltype.IsStarred(labels_message)
			message["is_spam"] = labeltype.IsSpam(labels_message)
			message["is_draft"] = labeltype.IsDraft(labels_message)
			message["is_archived"] = labeltype.IsArchived(labels_message)
			message["is_marked_as_phishing"] = labeltype.IsMarkedAsPhishing(labels_message)
			message["is_marked_as_not_phishing"] = labeltype.IsMarkedAsNotPhishing(labels_message)

			message["row_id"] = msg["row_id"]
			message["items_row_id"] = msg["items_row_id"]
			message["server_perm_id"] = msg["server_perm_id"]
			message["legacy_storage_id"] = msg["legacy_storage_id"]
			item_message := msg["message_proto"].(*MessageProto.Message)
			message["date"] = mstotime.MsToTime(item_message.GetEpoch())
			sender := extractaddress.ExtractSender(item_message)
			message["from"] = sender
			receivers := extractaddress.ExtractReceivers(item_message)
			message["to"] = receivers
			cc := extractaddress.ExtractCCs(item_message)
			message["cc"] = cc
			message["subject"] = item_message.GetSubject()
			message["snippet"] = item_message.GetSnippet()
			message["body"] = html2text.HTML2Text(item_message.GetHtml().GetHtmlDetails().GetBody().GetBody())

			attachs := item_message.GetAttachments()
			var attachments []map[string]interface{}
			for _, attach := range attachs {
				attachment := make(map[string]interface{})
				attachment["filename"] = attach.InfoAtt.GetAttach().GetFilename()
				attachment["size"] = attach.GetOtherIntInfo().GetSize()
				attachment["mime-type"] = attach.InfoAtt.GetAttach().GetMimetype()
				attachments = append(attachments, attachment)
			}
			message["attachments"] = attachments
			message_elaborated = append(message_elaborated, message)
		}
		conversation["messages"] = message_elaborated
		full = append(full, conversation)
	}
	return full
}

func ToJson(items []map[string]interface{}) {
	full := generateOutput(items)
	json_output, err := json.MarshalIndent(full, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(json_output))
}

func ToYaml(items []map[string]interface{}) {
	full := generateOutput(items)
	yaml_output, err := yaml.Marshal(full)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(yaml_output))
}
