package types

type GetSingalDataParamsDTO struct {
	Offset *int `json:"offset"`
	Limit  *int `json:"limit"`
}

type SignalType string

const (
	EEG     SignalType = "eeg"
	POW     SignalType = "pow"
	COM     SignalType = "com"
	MET     SignalType = "met"
	MOT     SignalType = "mot"
	FAC     SignalType = "fac"
	UNKNOWN SignalType = "unknown"
)

type RawSignalData map[string]interface{}
