package main

// now := time.Now()
// https://gobyexample.com/time

import (
    "time"
)

type Trader struct {
    entryId int
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
    traderId int
    stockId int
    // createdAt time
    targetPrice int
}

type Message struct {
    traderId int
    code string
}

func main() {
}
