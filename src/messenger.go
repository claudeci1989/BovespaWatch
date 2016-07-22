// https://github.com/subosito/twilio

package main

import (
    "fmt"
    "log"
    "github.com/subosito/twilio"
)

var (
    AccountSid string = "ACfadd593a319189b092b57f19f9a8e8d0"
    AuthToken string = "23e8cea197e1bdc89b38be51d9d4ebc3"
    twilioNumber string = "+12563331773"
    traderNumber string = "+5581982599865"
)

func main() {
    fmt.Println("SMS Messenger")

    c := twilio.NewClient(AccountSid, AuthToken, nil)

    params := twilio.MessageParams {
        Body: "Oi Rodrigo!",
    }

    s, response, err := c.Messages.Send(twilioNumber, traderNumber, params)
    if err != nil {
        log.Fatal(s, response, err)
    } else {
        log.Println("Message successfully sent.")
    }
}
