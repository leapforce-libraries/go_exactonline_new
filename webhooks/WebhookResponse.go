package exactonline

import types "github.com/leapforce-libraries/go_types"

type WebhookResponse struct {
	Content struct {
		Topic               string     `json:"Topic"`
		ClientID            string     `json:"ClientId"`
		Division            int64      `json:"Division"`
		Action              string     `json:"Action"`
		Key                 string     `json:"Key"`
		ExactOnlineEndpoint string     `json:"ExactOnlineEndpoint"`
		EventCreatedOn      types.Date `json:"EventCreatedOn"`
	} `json:"Content"`
	HashCode string `json:"HashCode"`
}
