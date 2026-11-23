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

type IncomingSignalMessage string
type OutgoingSignalMessage string

const (
	START_STREAMING IncomingSignalMessage = "start-streaming"
	END_STREAMING   IncomingSignalMessage = "end-streaming"
)

const (
	LIVE_STREAMING   OutgoingSignalMessage = "live-streaming"
	PREV_STREAMING   OutgoingSignalMessage = "prev-streaming"
	RESOURCE_FAILURE OutgoingSignalMessage = "resource-failure"
)

var ValidIncomingSignalMessages = map[IncomingSignalMessage]struct{}{
	START_STREAMING: {},
	END_STREAMING:   {},
}
var ValidOutgoingSignalMessages = map[OutgoingSignalMessage]struct{}{
	LIVE_STREAMING:   {},
	PREV_STREAMING:   {},
	RESOURCE_FAILURE: {},
}

type OutgoingClientSignalDTO[T any] struct {
	Message OutgoingSignalMessage `json:"message"`
	Type    SignalType            `json:"type"`
	Data    T                     `json:"data"`
}

type IncomingClientSignalDTO[T any] struct {
	Message            IncomingSignalMessage `json:"message"`
	Data               T                     `json:"data"`
	LastStreamRecieved *string               `json:"last_string_recieved"`
}
