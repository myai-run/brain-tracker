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

type RawIncomingSignalData map[string]interface{}

type ClientSignalType string

const (
	START_STREAMING ClientSignalType = "start-streaming"
	STREAMING       ClientSignalType = "streaming"
	END_STREAMING   ClientSignalType = "end-streaming"
)

var ValidSignalTypes = map[ClientSignalType]struct{}{
	START_STREAMING: {},
	STREAMING:       {},
	END_STREAMING:   {},
}

type ClientSignalDataDTO[T any] struct {
	Type ClientSignalType `json:"type"`
	Data T                `json:"data"`
}
