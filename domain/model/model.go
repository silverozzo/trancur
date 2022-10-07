package model

import (
	"time"
)

type Currency string

type Rate float64

type Exchange struct {
	First  Currency
	Second Currency
	Rate   Rate
}

type ExchangeList struct {
	Exchanges []Exchange `json:"exchanges"`
	Updated   time.Time  `json:"updated"`
}
