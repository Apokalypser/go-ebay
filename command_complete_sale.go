package ebay

import "encoding/xml"

type OrderID struct {
	OrderID string `xml:",innerxml"`
}

type FeedbackInfo struct {
	CommentText string
	CommentType string
	TargetUser  string
}

type Shipped struct {
	Shipped string `xml:",innerxml"`
}

type CompleteSale struct {
	OrderID      *OrderID
	FeedbackInfo *FeedbackInfo
	Shipped      *Shipped
}

func (c CompleteSale) CallName() string {
	return "CompleteSale"
}

func (c CompleteSale) ParseResponse(r []byte) (EbayResponse, error) {
	var xmlResponse CompleteSaleResponse
	err := xml.Unmarshal(r, &xmlResponse)

	return xmlResponse, err
}

func (c CompleteSale) Body() interface{} {
	return []interface{}{c.OrderID, c.FeedbackInfo, c.Shipped}
}

type CompleteSaleResponse struct {
	ebayResponse
}

func (r CompleteSaleResponse) ResponseErrors() ebayErrors {
	return r.ebayResponse.Errors
}
