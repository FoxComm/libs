package generators

import (
	"fmt"
	"time"

	"github.com/FoxComm/FoxComm/social_analytics/models"
	"github.com/FoxComm/libs/spree"
)

func CheckoutSiteActivity(orderNumber string, user spree.User) models.SiteActivity {
	address := spree.Address{
		Id:        1,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Address1:  fake.StreetAddress(),
		City:      fake.City(),
		ZipCode:   fake.PostCode(),
		Phone:     fake.PhoneNumber(),
	}

	return models.SiteActivity{
		Action:             "checkout",
		SharingActivityTag: 0,
		SharerToken:        "",
		RefererURL:         "http://localhost:8000/checkout/confirm",
		LandingURL:         "",
		ApiRequestURL:      fmt.Sprintf("/app/api/checkouts/%s", orderNumber),
		Entity:             models.NewUserEntity(user),
		CheckoutDetails: models.CheckoutActivityDetails{
			Id:                        1,
			Number:                    orderNumber,
			State:                     "complete",
			CreatedAt:                 time.Now(),
			UpdateAt:                  time.Now(),
			CompletedAt:               time.Now(),
			ShipmentState:             "",
			PaymentState:              "",
			Email:                     user.Email,
			SpecialInstructions:       "",
			Channel:                   "spree",
			IncludedTaxTotal:          0.0,
			AdditionalTaxTotal:        0.0,
			DisplayIncludedTaxTotal:   "0.0",
			DisplayAdditionalTaxTotal: "0.0",
			TaxTotal:                  0.0,
			Currency:                  "USD",
			TotalQuantity:             2,
			DisplayTaxTotal:           "0.0",
			Token:                     "FAKETOKEN",
			Shipments:                 []models.OrderShipment{},
			Adjustments:               []models.OrderAdjustment{},
			ItemTotal:                 40.0,
			ShipTotal:                 5.95,
			DisplayShipTotal:          "5.95",
			AdjustmentTotal:           0.0,
			Total:                     45.95,
			PaymentTotal:              45.95,
			DisplayTotal:              "45.95",
			BillAddress:               address,
			ShipAddress:               address,
			LineItems: []models.OrderLineItem{
				models.OrderLineItem{OrderId: 1, VariantId: 1, Quantity: 2, Price: 20.0},
			},
			Payments: []models.OrderPayment{
				models.OrderPayment{
					Amount:        45.95,
					CreatedAt:     time.Now(),
					Id:            1,
					PaymentMethod: models.PaymentMethod{Environment: "test", Id: 1, Name: "AmEx"},
					State:         "paid",
				},
			},
		},
	}
}
