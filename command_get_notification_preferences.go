package ebay

import (
	"encoding/xml"
	"log"
)

type PreferenceLevel struct {
	PreferenceLevel string `xml:",innerxml"`
}

type GetNotificationPreferences struct {
	PreferenceLevel
}

func (c GetNotificationPreferences) CallName() string {
	return "GetNotificationPreferences"
}

func (c GetNotificationPreferences) ParseResponse(r []byte) (EbayResponse, error) {
	var xmlResponse GetNotificationPreferencesResponse
	err := xml.Unmarshal(r, &xmlResponse)
	log.Println("gnp err", string(r))
	return xmlResponse, err
}

func (c GetNotificationPreferences) Body() interface{} {
	return c.PreferenceLevel
}

type GetNotificationPreferencesResponse struct {
	ebayResponse

	ItemID string
}

func (r GetNotificationPreferencesResponse) ResponseErrors() ebayErrors {
	return r.ebayResponse.Errors
}
