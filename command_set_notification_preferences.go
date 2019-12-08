package ebay

import (
	"encoding/xml"
	"log"
)

type ApplicationDeliveryPreferences struct {
	ApplicationEnable string
	ApplicationURL    string
}

type UserDeliveryPreferenceArray struct {
	NotificationEnable []*NotificationEnable `xml:",omitempty"`
}

type NotificationEnable struct {
	EventType   string
	EventEnable string
}

type SetNotificationPreferences struct {
	ApplicationDeliveryPreferences *ApplicationDeliveryPreferences `xml:",omitempty"`
	UserDeliveryPreferenceArrayItems    *UserDeliveryPreferenceArray    `xml:",omitempty"`
}

func (c SetNotificationPreferences) CallName() string {
	return "SetNotificationPreferences"
}

func (c SetNotificationPreferences) ParseResponse(r []byte) (EbayResponse, error) {
	var xmlResponse SetNotificationPreferencesResponse
	err := xml.Unmarshal(r, &xmlResponse)
	log.Println("snp err", string(r))
	return xmlResponse, err
}

func (c SetNotificationPreferences) Body() interface{} {
	return []interface{}{c.ApplicationDeliveryPreferences, c.UserDeliveryPreferenceArrayItems}
}

type SetNotificationPreferencesResponse struct {
	ebayResponse

	ItemID string
}

func (r SetNotificationPreferencesResponse) ResponseErrors() ebayErrors {
	return r.ebayResponse.Errors
}
