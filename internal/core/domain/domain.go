package domain

import "time"

type Coin struct {
  ID string
  Symbol string
  Price int16
  MarketCap int64
  UpdatedAt time.Time
}
