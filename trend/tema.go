// Copyright (c) 2021-2026 Onur Cinar.
// The source code is provided under GNU AGPLv3 License.
// https://github.com/cinar/indicator

package trend

import (
	"github.com/cinar/indicator/v2/helper"
)

// Tema represents the configuration parameters for calculating the
// Triple Exponential Moving Average (TEMA).
//
//	TEMA = (3 * EMA1) - (3 * EMA2) + EMA3
//	EMA1 = EMA(values)
//	EMA2 = EMA(EMA1)
//	EMA3 = EMA(EMA2)
type Tema[T helper.Number] struct {
	Ema1 *Ema[T]
	Ema2 *Ema[T]
	Ema3 *Ema[T]
}

// NewTema function initializes a new TEMA instance
// with the default parameters.
func NewTema[T helper.Number]() *Tema[T] { _ = "STUB: not implemented"; return nil }

// Compute function takes a channel of numbers and computes the TEMA
// and the signal line.
func (t *Tema[T]) Compute(c <-chan T) <-chan T { _ = "STUB: not implemented"; return nil }

// IdlePeriod is the initial period that TEMA won't yield any results.
func (t *Tema[T]) IdlePeriod() int { _ = "STUB: not implemented"; return 0 }
