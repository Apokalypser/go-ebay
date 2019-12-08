package ebay

import "encoding/xml"

type GetClientAlertsAuthToken struct {
}

func (c GetClientAlertsAuthToken) CallName() string {
	return "GetClientAlertsAuthToken"
}

func (c GetClientAlertsAuthToken) ParseResponse(r []byte) (EbayResponse, error) {
	var xmlResponse GetClientAlertsAuthTokenResponse
	err := xml.Unmarshal(r, &xmlResponse)

	return xmlResponse, err
}

func (c GetClientAlertsAuthToken) Body() interface{} {
	return nil
}

type GetClientAlertsAuthTokenResponse struct {
	ebayResponse

	ClientAlertsAuthToken string
}

func (r GetClientAlertsAuthTokenResponse) ResponseErrors() ebayErrors {
	return r.ebayResponse.Errors
}
