package utils

import "airun/brain-tracker/server/src/types"

func GetValidatedSignalDataParams(rawParams types.GetSingalDataParamsDTO) types.GetSingalDataParamsDTO {
	var params = rawParams
	defaultValue := 0

	if rawParams.Offset == nil {
		params.Offset = &defaultValue
	}
	if rawParams.Limit == nil {
		params.Limit = &defaultValue
	}

	return params
}

func GetSignalDataType(signal types.RawSignalData) types.SignalType {
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
