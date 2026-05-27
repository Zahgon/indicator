package main

import (
	"github.com/cinar/indicator/v2/strategy"
)

// StrategyType defines the type of trading strategy to be used in a backtest.
// It is represented as a string to allow for easy identification and selection.
type StrategyType string

// Constants for all supported strategy types.
// This list includes a variety of strategies from different categories, such as
// trend, momentum, and volume-based approaches.
const (
	// Base strategies
	StrategyBuyAndHold StrategyType = "buy_and_hold"

	// Momentum strategies
	StrategyAwesomeOscillator StrategyType = "awesome_oscillator"
	StrategyRsi               StrategyType = "rsi"
	StrategyStochasticRsi     StrategyType = "stochastic_rsi"
	StrategyTripleRsi         StrategyType = "triple_rsi"

	// Volume strategies
	StrategyChaikinMoneyFlow     StrategyType = "chaikin_money_flow"
	StrategyEaseOfMovement       StrategyType = "ease_of_movement"
	StrategyForceIndex           StrategyType = "force_index"
	StrategyMoneyFlowIndex       StrategyType = "money_flow_index"
	StrategyNegativeVolumeIndex  StrategyType = "negative_volume_index"
	StrategyWeightedAveragePrice StrategyType = "weighted_average_price"

	// Trend strategies
	StrategyMACD              StrategyType = "macd"
	StrategyAlligator         StrategyType = "alligator"
	StrategyAroon             StrategyType = "aroon"
	StrategyApo               StrategyType = "apo"
	StrategyBop               StrategyType = "bop"
	StrategyCci               StrategyType = "cci"
	StrategyDema              StrategyType = "dema"
	StrategyGoldenCross       StrategyType = "golden_cross"
	StrategyKama              StrategyType = "kama"
	StrategyKdj               StrategyType = "kdj"
	StrategyQstick            StrategyType = "qstick"
	StrategySmma              StrategyType = "smma"
	StrategyTrima             StrategyType = "trima"
	StrategyTripleMaCrossover StrategyType = "triple_ma_crossover"
	StrategyTsi               StrategyType = "tsi"
	StrategyVwma              StrategyType = "vwma"
	StrategyWeightedClose     StrategyType = "weighted_close"
)

// CreateStrategy creates a new strategy instance based on the specified type.
// It acts as a factory function, mapping a StrategyType to a concrete
// implementation of the strategy.Strategy interface.
//
// This function is essential for dynamically selecting and initializing the
// desired trading strategy at runtime. If an unsupported strategy type is
// provided, it returns an error.
func CreateStrategy(strategyType StrategyType) (strategy.Strategy, error) {
	_ = "STUB: not implemented"
	return *

	// Base strategies
	new(strategy.Strategy), nil
}

// Trend strategies

// Momentum strategies

// Volume strategies
