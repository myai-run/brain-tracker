package utils

import (
	"airun/brain-tracker/server/src/models"
	"airun/brain-tracker/server/src/types"
	"time"

	"gorm.io/gorm"
)

func GetSignalDataType(signal types.RawIncomingSignalData) types.SignalType {
	if signal["pow"] != nil {
		return types.POW
	} else if signal["eeg"] != nil {
		return types.EEG
	} else if signal["mot"] != nil {
		return types.MOT
	} else if signal["met"] != nil {
		return types.MET
	} else if signal["com"] != nil {
		return types.COM
	} else if signal["fac"] != nil {
		return types.FAC
	}
	return types.UNKNOWN
}

func IsValidClientSignal(signal *types.IncomingClientSignalDTO[any]) bool {
	if signal == nil {
		return false
	}
	_, ok := types.ValidSignalTypes[types.ClientSignalType(signal.Type)]
	if !ok {
		return false
	}
	return true
}

func SignalDataToOutgoingClientSignalDTO(signal models.SignalData) types.OutgoingClientSignalDTO[any] {
	return types.OutgoingClientSignalDTO[any]{
		Type: signal.Type,
		Data: signal,
	}
}

func GetParameterizedTransaction(transaction *gorm.DB, params types.GetSingalDataParamsDTO) *gorm.DB {
	if params.Offset == nil {
		transaction = transaction.Offset(0)
	} else {
		transaction = transaction.Offset(*params.Offset)
	}

	if params.Limit == nil {
		transaction = transaction.Limit(0)
	} else {
		transaction = transaction.Limit(*params.Limit)
	}

	if params.RecievedAfter == "" {
		transaction = transaction.Where("time > ?", params.RecievedAfter)
	} else {
		now := time.Now()
		transaction = transaction.Where("time > %s", now)
	}

	return transaction
}
