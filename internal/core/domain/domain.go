package domain

import "time"

type Coin struct {
  ID string
  Symbol string
  Price float32
  MarketCap int64
  UpdatedAt time.Time
}

type Flippening struct {
  Ratio float64
  EthCap int64
  BtcCap int64
  Init bool
  UpdatedAt time.Time
}

type Gas struct {
  ID string
  Base float32
  AvgTime float32
  UpdatedAt time.Time
}
