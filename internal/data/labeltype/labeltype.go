package labeltype

import (
	"fmt"
	"slices"
)

// Source:
// - https://developers.google.com/gmail/android/com/google/android/gm/contentprovider/GmailContract.Labels.LabelCanonicalNames
// - https://docs.google.com/spreadsheets/d/1BS8yazyPcfqbWMG2jQb8HvPvCvQsDNQ3tsp_pBh5P6Q

var labelType = map[string]string{
	"^a":                       "ARCHIVED",
	"^af":                      "Filtered to trash", // Only trashed by filters, not by hand.
	"^all":                     "ALL_MAIL",
	"^cf_{some_stuff}":         "labels for circles have this naming scheme, but not only them", // some_stuff is 16 base-16 digits. https://plus.google.com/u/0/stream/circles/p{some_stuff} = circle URL.
	"^cff":                     "sender in Google+ circle",
	"^cob_goldmine":            "unknown: cob_goldmine",
	"^cob_pevent":              "unknown: cob_pevent - mailing list",
	"^cob_sm_emailmessage":     "unknown: cob_sm_emailmessage - mailing list",
	"^cob_sm_forme":            "unknown: cob_sm_forme",
	"^cob_sm_offer":            "unknown: cob_sm_offer",
	"^cob-processed-gmr":       "unknown: cob-processed-gmr", // "internal label, do not use"; probably all emails sorted into cob_something categories
	"^esa":                     "unknown: esa",
	"^f":                       "SENT",
	"^f_bt":                    "unknown: f_bt",
	"^fhas":                    "unknown: fhas",
	"^flas":                    "unknown: flas",
	"^fmas":                    "unknown: fmas",
	"^fnas":                    "unknown: fnas",
	"^fs":                      "SENT TO SELF",
	"^g":                       "MUTED",
	"^hunsub":                  "Has an 'unsubscribe' Gmail action",
	"^i":                       "INBOX",
	"^ia":                      "SUBSET OF IMPORTANCE",
	"^idxs":                    "unknown: idxs",
	"^iim":                     "PRIORITY_INBOX",
	"^io_im":                   "IMPORTANT_MESSAGES",
	"^io_imc1":                 "IMPORTANT_MESSAGES (due to people in conversation)",
	"^io_imc2":                 "IMPORTANT_MESSAGES (due to words in message)",
	"^io_imc3":                 "IMPORTANT_MESSAGES (important sent directly to you)",
	"^io_imc4":                 "IMPORTANT_MESSAGES (important you often read messages with this label)",
	"^io_imc5":                 "IMPORTANT_MESSAGES (important from interaction in conversation)",
	"^io_lr":                   "Importance learning?", // may be a label for the messages that Gmail has trouble classyfying as important/not and therefore those are the most important for learning. But it's just a guess.
	"^io_unim":                 "UN-IMPORTANT_MESSAGES",
	"^k":                       "TRASH",
	"^mf":                      "Imported from POP account",
	"^np":                      "MARKED_AS_NOT_PHISHING",
	"^ns":                      "Never send it to Spam",
	"^o":                       "ACTUALLY_OPENED (not just mark as read)",
	"^op":                      "messages that were automatically marked as phishing attempts",
	"^os":                      "messages that were automatically marked as spam",
	"^os_notification":         "unknown: os_notification",
	"^os_personal":             "unknown: os_personal",
	"^os_promo":                "unknown: os_promo",
	"^os_social":               "unknown: os_social",
	"^p":                       "MARKED_AS_PHISHING",
	"^p_mtunsub":               "Has Unsubscribe (uncertain)",
	"^pop":                     "Message downloaded by a POP client",
	"^r":                       "DRAFTS",
	"^ri":                      "read in Inbox?",
	"^s":                       "SPAM",
	"^sl_root":                 "all messages with a Smart Label",
	"^smartlabel_event":        "CATEGORIES: EVENTS",
	"^smartlabel_finance":      "CATEGORIES: FINANCIAL",
	"^smartlabel_group":        "CATEGORIES: GROUP",
	"^smartlabel_newsletter":   "CATEGORIES: NEWSLETTER",
	"^smartlabel_notification": "CATEGORIES: UPDATES",
	"^smartlabel_personal":     "CATEGORIES: PERSONAL",
	"^smartlabel_promo":        "CATEGORIES: PROMOTIONS",
	"^smartlabel_pure_notif":   "CATEGORIES: PURE NOTIFICATIONS",
	"^smartlabel_receipt":      "CATEGORIES: PURCHACES",
	"^smartlabel_social":       "CATEGORIES: SOCIAL",
	"^smartlabel_travel":       "CATEGORIES: TRAVEL-RELATED",
	"^sps":                     "messages suspected of being a Spear PhiShing",
	"^sq_ig_i_group":           "INBOX_CATEGORY_FORUMS",
	"^sq_ig_i_notification":    "INBOX_CATEGORY_UPDATES",
	"^sq_ig_i_personal":        "INBOX_CATEGORY_PRIMARY",
	"^sq_ig_i_promo":           "INBOX_CATEGORY_PROMOTIONS",
	"^sq_ig_i_social":          "INBOX_CATEGORY_SOCIAL",
	"^ss_cg":                   "SUPERSTARS: Checkbox, green (with a '✔')",
	"^ss_co":                   "SUPERSTARS: Checkbox, orange (with a '»')",
	"^ss_cp":                   "SUPERSTARS: Checkbox, purple (with a '?')",
	"^ss_cr":                   "SUPERSTARS: Checkbox, red (with a '!'')",
	"^ss_cy":                   "SUPERSTARS: Checkbox, yellow (with a '!')",
	"^ss_sb":                   "SUPERSTARS: Star, Blue",
	"^ss_sg":                   "SUPERSTARS: Superstars: star, green",
	"^ss_so":                   "SUPERSTARS: Superstars: star, orange",
	"^ss_sp":                   "SUPERSTARS: SuperStars: Star, Purple",
	"^ss_sr":                   "SUPERSTARS: Superstars: star, red",
	"^ss_sy":                   "SUPERSTARS: Superstars: star, yellow",
	"^sua":                     "purchases, notifications and non-spam newsletters from senders I often get (and read?) mail from",
	"^t":                       "STARRED",
	"^ts{number}":              "unknown: ts{number}", // {number} is from 0 to at least 18 without leading zeroes
	"^u":                       "UNREAD",
	"^unsub":                   "HAS AN 'UNSUBSCRIBE' LINK",
	"^us":                      "unknown: us",
	"^vm":                      "VOICEMAIL",
}

func GetMeaningFromLabel(label string) string {
	meaning, ok := labelType[label]
	// If the key exists
	if ok {
		return meaning
	} else {
		return fmt.Sprintf("unknown: %s", label)
	}
}

func TranslationLabels(labels []string) []string {
	var output []string
	for _, label := range labels {
		meaning := GetMeaningFromLabel(label)
		output = append(output, meaning)
	}
	return output
}

func IsSent(labels []string) bool {
	return slices.Contains(labels, "^f")
}

func IsInbox(labels []string) bool {
	return slices.Contains(labels, "^i")
}

func IsTrash(labels []string) bool {
	return slices.Contains(labels, "^k")
}

func IsMarkedAsNotPhishing(labels []string) bool {
	return slices.Contains(labels, "^np")
}

func IsMarkedAsPhishing(labels []string) bool {
	return slices.Contains(labels, "^p")
}

func IsOpened(labels []string) bool {
	return slices.Contains(labels, "^o")
}

func IsArchived(labels []string) bool {
	return slices.Contains(labels, "^a")
}

func IsDraft(labels []string) bool {
	return slices.Contains(labels, "^r")
}

func IsSpam(labels []string) bool {
	return slices.Contains(labels, "^s")
}

func IsStarred(labels []string) bool {
	return slices.Contains(labels, "^t")
}

func IsUnread(labels []string) bool {
	return slices.Contains(labels, "^u")
}
