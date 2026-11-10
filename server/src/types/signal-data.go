package types

type GetSingalDataParamsDTO struct {
	Offset        *int   `json:"offset"`
	Limit         *int   `json:"limit"`
	RecievedAfter string `json:"recieved_after"`
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
	END_STREAMING   ClientSignalType = "end-streaming"
)

var ValidSignalTypes = map[ClientSignalType]struct{}{
	START_STREAMING: {},
	END_STREAMING:   {},
}

type OutgoingClientSignalDTO[T any] struct {
	Type SignalType `json:"type"`
	Data T          `json:"data"`
}

type IncomingClientSignalDTO[T any] struct {
	Type               ClientSignalType `json:"type"`
	Data               T                `json:"data"`
	LastStreamRecieved string           `json:"last_string_recieved"`
}
