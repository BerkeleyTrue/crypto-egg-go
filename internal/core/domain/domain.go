package domain

import "time"

type Coin struct {
  ID string
  Sym string
  Price int16
  UpdatedAt time.Time
}
