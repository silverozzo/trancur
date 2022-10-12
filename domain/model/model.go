package model

import (
	"time"
)

type Currency string

type Rate float64

type Exchange struct {
	First  string
	Second string
	Rate   float64
}

type ExchangeList struct {
	Exchanges []Exchange `json:"exchanges"`
	Updated   time.Time  `json:"updated"`
}
