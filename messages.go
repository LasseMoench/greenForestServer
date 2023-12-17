package main

type baseMessage struct {
	MessageType string `json:"messageType"`
}

type answerMessage struct {
	MessageType string `json:"messageType"`
	Answer      string `json:"answer"`
}

type resultMessage struct {
	MessageType   string `json:"messageType"`
	ResultCorrect bool   `json:"resultCorrect"`
	Problem       string `json:"problem"`
	Hazelnuts     int    `json:"hazelnuts"`
}

type opponentUpdate struct {
	MessageType       string `json:"messageType"`
	OpponentHazelnuts int    `json:"opponentHazelnuts"`
}

type gameStartMessage struct {
	MessageType string `json:"messageType"`
	Problem     string `json:"problem"`
}
