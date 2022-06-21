package utils_test

import (
	"testing"

	"github.com/berkeleytrue/crypto-egg-go/internal/utils"
	"github.com/frankban/quicktest"
)

func TestFormatPrice(t *testing.T) {
  t.Run("Should round up when less than 0.001", func(t *testing.T) {
    c := quicktest.New(t)
    res := utils.Formatter.FormatPrice(0.0001)
    c.Assert(res, quicktest.Equals, "0.001\n")
  })
  t.Run("Should round down when less than 0.01", func(t *testing.T) {
    c := quicktest.New(t)
    res := utils.Formatter.FormatPrice(0.0042341)
    c.Assert(res, quicktest.Equals, "0.004234\n")
  })
  t.Run("Should round down when less than 1", func(t *testing.T) {
    c := quicktest.New(t)
    res := utils.Formatter.FormatPrice(0.42341)
    c.Assert(res, quicktest.Equals, "0.4234\n")
  })
  t.Run("Should round down to cents when less than 100", func(t *testing.T) {
    c := quicktest.New(t)
    res := utils.Formatter.FormatPrice(42.1264)
    c.Assert(res, quicktest.Equals, "42.13\n")
  })
  t.Run("Should round down to dollars when greater than 1000", func(t *testing.T) {
    c := quicktest.New(t)
    res := utils.Formatter.FormatPrice(421264.2349293)
    c.Assert(res, quicktest.Equals, "421264\n")
  })
}
