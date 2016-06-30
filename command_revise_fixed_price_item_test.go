package ebay

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ReviseFixedPriceItem_CallName(t *testing.T) {
	s := ReviseFixedPriceItem{}

	assert.Equal(t, "ReviseFixedPriceItem", s.CallName())
}

func Test_ReviseFixedPriceItem_Body(t *testing.T) {
	s := ReviseFixedPriceItem{
		ItemID:             "item-id",
		StartPrice:         "start-price",
		ConditionID:        1,
		Quantity:           2,
		Title:              "title",
		Description:        "description",
		PayPalEmailAddress: "ppemail",
		PictureDetails:     &PictureDetails{"url"},
		ShippingDetails: &ShippingDetails{
			ShippingType:                           "type",
			ShippingDiscountProfileID:              "profile id",
			InternationalShippingDiscountProfileID: "i profile id",
			ShippingServiceOptions: []ShippingServiceOption{
				ShippingServiceOption{
					ShippingService:               "service",
					ShippingServiceCost:           1.23,
					ShippingServiceAdditionalCost: 2.34,
					FreeShipping:                  false,
				},
			},
			InternationalShippingServiceOption: []InternationalShippingServiceOption{
				InternationalShippingServiceOption{
					ShippingService:               "service",
					ShippingServiceCost:           1.23,
					ShippingServiceAdditionalCost: 2.34,
					ShipToLocation:                []string{"ship", "to"},
					ShippingServicePriority:       1,
				},
			},
		},
		ProductListingDetails: &ProductListingDetails{
			UPC:      "upc",
			BrandMPN: BrandMPN{"foo", "bar"},
		},
		ItemSpecifics: []ItemSpecifics{
			ItemSpecifics{
				[]NameValueList{
					NameValueList{
						Name:  "name",
						Value: []string{"value"},
					},
				},
			},
		},
	}

	b := new(bytes.Buffer)
	xml.NewEncoder(b).Encode(s.Body())

	assert.Equal(t, "<Item><ItemID>item-id</ItemID><StartPrice>start-price</StartPrice><ConditionID>1</ConditionID><Quantity>2</Quantity><Title>title</Title><Description>description</Description><PayPalEmailAddress>ppemail</PayPalEmailAddress><PictureDetails><PictureURL>url</PictureURL></PictureDetails><ShippingDetails><ShippingType>type</ShippingType><ShippingDiscountProfileID>profile id</ShippingDiscountProfileID><InternationalShippingDiscountProfileID>i profile id</InternationalShippingDiscountProfileID><ShippingServiceOptions><ShippingService>service</ShippingService><ShippingServiceCost>1.23</ShippingServiceCost><ShippingServiceAdditionalCost>2.34</ShippingServiceAdditionalCost><FreeShipping>false</FreeShipping></ShippingServiceOptions><InternationalShippingServiceOption><ShippingService>service</ShippingService><ShippingServiceCost>1.23</ShippingServiceCost><ShippingServiceAdditionalCost>2.34</ShippingServiceAdditionalCost><ShipToLocation>ship</ShipToLocation><ShipToLocation>to</ShipToLocation><ShippingServicePriority>1</ShippingServicePriority></InternationalShippingServiceOption></ShippingDetails><ProductListingDetails><UPC>upc</UPC><BrandMPN><Brand>foo</Brand><MPN>bar</MPN></BrandMPN></ProductListingDetails><ItemSpecifics><NameValueList><Name>name</Name><Value>value</Value></NameValueList></ItemSpecifics></Item>", b.String())
}

func Test_ReviseFixedPriceItem_Body_OptionalFields(t *testing.T) {
	s := ReviseFixedPriceItem{
		ItemID: "item-id",
	}

	b := new(bytes.Buffer)
	xml.NewEncoder(b).Encode(s.Body())

	assert.Equal(t, "<Item><ItemID>item-id</ItemID></Item>", b.String())
}
