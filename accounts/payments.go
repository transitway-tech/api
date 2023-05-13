package accounts

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/setupintent"
	"go.mongodb.org/mongo-driver/bson"
)

type CardDetails struct {
  Numberr string `json:"number"`
  ExpYear string `json:"expYear"`
  ExpMonth string  `json:"expMonth"`
  CVC string  `json:"cvc"`
}

func payments (acc fiber.Router) {
  payments := acc.Group("/payments")

  payments.Get("", func (c *fiber.Ctx) error {
    // var account models.Account
    // utils.GetLocals(c, "account", &account)

    stripeCustomerID := "cus_Nt3fgeflq5qBLK"

    params := &stripe.SetupIntentParams {
      Customer: stripe.String(stripeCustomerID),
      PaymentMethodTypes: []*string {
        stripe.String("card"),
      },
    }
    intent, _ := setupintent.New(params);

    return c.JSON(bson.M {
      "client_secret": intent.ClientSecret,
    })
    // return c.JSON(account)
  })
}