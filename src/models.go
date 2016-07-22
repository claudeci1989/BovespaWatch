package main

// now := time.Now()
// https://gobyexample.com/time

import (
    "time"
)

type Trader struct {
    entryId int
    fullName string
    emailAddress string
    phoneNumber string
    dailyQuota int
    dailyQuotaUsed int
}

type Stock struct {
    entryId int
    code string
}

type WatchOrder struct {
    trader Trader
    stockId int
    targetPrice int
}

type Message struct {
    trader Trader
    code string
    fullText string
}

func (wo WatchOrder) buildMessage() {
    var customerName string = wo.trader.fullName

    // message := Message {
    //     fullText: "1234567890123456",
    //     customerName: "Taylor Fritz",
    //     cvc: 123,
    //     chargeAmount: 1999,
    // }
}

// func (c Charge) applyDiscount(rate int64) int64 {
//     c.chargeAmount -= rate
//     return c.chargeAmount
// }

func main() {
}
