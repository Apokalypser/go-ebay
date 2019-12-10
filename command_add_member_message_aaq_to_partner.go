package ebay

import "encoding/xml"

type ItemID struct {
	ItemID string `xml:",innerxml"`
}

type MemberMessage struct {
	Body         string
	QuestionType string
	RecipientID  string
	Subject      string
}

type AddMemberMessageAAQToPartner struct {
	ItemID        *ItemID
	MemberMessage *MemberMessage
}

func (c AddMemberMessageAAQToPartner) CallName() string {
	return "AddMemberMessageAAQToPartner"
}

func (c AddMemberMessageAAQToPartner) ParseResponse(r []byte) (EbayResponse, error) {
	var xmlResponse AddMemberMessageAAQToPartnerResponse
	err := xml.Unmarshal(r, &xmlResponse)

	return xmlResponse, err
}

func (c AddMemberMessageAAQToPartner) Body() interface{} {
	return []interface{}{c.ItemID, c.MemberMessage}
}

type AddMemberMessageAAQToPartnerResponse struct {
	ebayResponse
}

func (r AddMemberMessageAAQToPartnerResponse) ResponseErrors() ebayErrors {
	return r.ebayResponse.Errors
}
