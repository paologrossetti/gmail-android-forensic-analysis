package extractaddress

import MessageProto "github.com/paologrossetti/gmail-android/internal/proto/Message"

func ExtractSender(item_message *MessageProto.Message) string {
	sender := item_message.GetSender()
	if sender != nil {
		return sender.GetAddress()
	} else {
		sender_secondary := item_message.GetSenderSecondary()
		return sender_secondary.GetEmail()
	}
}

func ExtractReceivers(item_message *MessageProto.Message) []string {
	var receivers []string
	to_addresses := item_message.GetReceiver()
	if to_addresses != nil {
		for _, receiver := range to_addresses {
			receivers = append(receivers, receiver.GetAddress())
		}
	} else {
		to_addresses_secondary := item_message.GetReceiver()
		for _, receiver := range to_addresses_secondary {
			receivers = append(receivers, receiver.GetAddress())
		}
	}
	return receivers
}

func ExtractCCs(item_message *MessageProto.Message) []string {
	var cc []string
	CCs := item_message.GetCc()
	for _, cc_elem := range CCs {
		cc = append(cc, cc_elem.GetAddress())
	}
	return cc
}
