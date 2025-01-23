package rpc

import (
	"context"
)

type GetPriorityFeeEstimateResponse JsonRpcResponse[GetPriorityFeeEstimate]

type MicroLamportPriorityFeeLevels struct {
	Min       float64 `json:"min"`
	Low       float64 `json:"low"`
	Medium    float64 `json:"medium"`
	High      float64 `json:"high"`
	VeryHigh  float64 `json:"veryHigh"`
	UnsafeMax float64 `json:"unsafeMax"`
}

// GetPriorityFeeEstimate is a part of raw rpc response of `getPriorityFeeEstimate`
type GetPriorityFeeEstimate struct {
	PriorityFeeEstimate float64                       `json:"priorityFeeEstimate"`
	PriorityFeeLevels   MicroLamportPriorityFeeLevels `json:"priorityFeeLevels"`
}

func (f GetPriorityFeeEstimate) FeeUint64(level *PriorityLevel) uint64 {
	if f.PriorityFeeEstimate != 0.0 {
		return uint64(f.PriorityFeeEstimate)
	}

	if level == nil {
		return uint64(f.PriorityFeeLevels.Medium)
	}
	switch *level {
	case PriorityLevelMin:
		return uint64(f.PriorityFeeLevels.Min)
	case PriorityLevelLow:
		return uint64(f.PriorityFeeLevels.Low)
	case PriorityLevelMedium:
		return uint64(f.PriorityFeeLevels.Medium)
	case PriorityLevelHigh:
		return uint64(f.PriorityFeeLevels.High)
	case PriorityLevelVeryHigh:
		return uint64(f.PriorityFeeLevels.VeryHigh)
	case PriorityLevelUnsafeMax:
		return uint64(f.PriorityFeeLevels.UnsafeMax)
	}

	return uint64(f.PriorityFeeLevels.Medium)
}

type PriorityLevel string

const (
	PriorityLevelMin       PriorityLevel = "Min"
	PriorityLevelLow       PriorityLevel = "Low"
	PriorityLevelMedium    PriorityLevel = "Medium"
	PriorityLevelHigh      PriorityLevel = "High"
	PriorityLevelVeryHigh  PriorityLevel = "VeryHigh"
	PriorityLevelUnsafeMax PriorityLevel = "UnsafeMax"
)

type GetPriorityFeeEstimateOpts struct {
	TransactionEncoding         string        `json:"transactionEncoding,omitempty"`
	PriorityLevel               PriorityLevel `json:"priorityLevel,omitempty"`
	IncludeAllPriorityFeeLevels bool          `json:"includeAllPriorityFeeLevels,omitempty"`
	LookbackSlots               int           `json:"lookbackSlots,omitempty"`
	IncludeVote                 bool          `json:"includeVote,omitempty"`
	Recommended                 bool          `json:"recommended,omitempty"`
	EvaluateEmptySlotAsZero     bool          `json:"evaluateEmptySlotAsZero,omitempty"`
}

type GetPriorityFeeEstimateRequest struct {
	Transaction string                     `json:"transaction,omitempty"`
	AccountKeys []string                   `json:"accountKeys,omitempty"`
	Options     GetPriorityFeeEstimateOpts `json:"options"`
}

// GetPriorityFeeEstimate is an RPC method that provides fee recommendations based on historical data. Most importantly, it considers both global and local fee markets
func (c *RpcClient) GetPriorityFeeEstimate(ctx context.Context, request GetPriorityFeeEstimateRequest) (JsonRpcResponse[GetPriorityFeeEstimate], error) {
	return call[JsonRpcResponse[GetPriorityFeeEstimate]](c, ctx, "getPriorityFeeEstimate", request)
}
