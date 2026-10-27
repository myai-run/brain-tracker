package models

// INSIGHT-95A51586
type RawEEGData struct {
	EEG  []int  `json:"eeg"`
	Sid  string `json:"sid"`
	Time int    `json:"time"`
}

// {
//   eeg: [
//     44,       0,
//     4316.923, 4395.897,
//     4122.564, 4027.179,
//     4241.026, 0,
//     0,        []
//   ],
//   sid: '30bbb883-bc99-4694-8eac-96f67d28ec79',
//   time: 1793115965.5002
// }
