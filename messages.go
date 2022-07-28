package js8call

import "encoding/json"

/*
Message is a generic type to allow us to get the value of Type */
type Message struct {
	Params json.RawMessage `json:"params"`
	Type   string          `json:"type"`
	Value  string          `json:"value"`
}

/*
CLOSE
INBOX.MESSAGES
INBOX.MESSAGE
MODE.SPEED
PING
*/

/*
RigFreq contains the specific RIG.FREQ message structure */
type RigFreq struct {
	Type   string `json:"type"`
	Value  string `json:"value"`
	Params struct {
    BAND string `json:"BAND"`
    DIAL int64 `json:"DIAL"`
    FREQ int64 `json:"FREQ"`
    OFFSET int `json:"OFFSET"`
    ID int `json:"_ID"`
  }
}

/*
RigPTT contains the specific RIG.PTT message structure */
type RigPTT struct {
	Type   string `json:"type"`
	Value  string `json:"value"`
	Params struct {
		PTT bool  `json:"PTT"`
		UTC int64 `json:"UTC"`
		ID  int   `json:"_ID"`
	}
}

/*
RXActivity contains the specific RX.ACTIVITY message structure */
type RXActivity struct {
	Type   string `json:"type"`
	Value  string `json:"value"`
	Params struct {
		DIAL   int     `json:"DIAL"`
		FREQ   int     `json:"FREQ"`
		OFFSET int     `json:"OFFSET"`
		SNR    int     `json:"SNR"`
		SPEED  int     `json:"SPEED"`
		TDRIFT float64 `json:"TDRIFT"`
		UTC    int64   `json:"UTC"`
		ID     int     `json:"_ID"`
	}
}

/*
RX.BAND_ACTIVITY
RX.CALL_ACTIVITY
RX.CALL_SELECTED
*/
/*
RXDirected contains the specific RX.DIRECTED message structure */
type RXDirected struct {
	Type   string `json:"type"`
	Value  string `json:"value"`
	Params struct {
		CMD    string  `json:"CMD"`
		DIAL   int     `json:"DIAL"`
		EXTRA  string  `json:"EXTRA"`
		FREQ   int     `json:"FREQ"`
		FROM   string  `json:"FROM"`
		GRID   string  `json:"GRID"`
		OFFSET int     `json:"OFFSET"`
		SNR    int     `json:"SNR"`
		SPEED  int     `json:"SPEED"`
		TDRIFT float64 `json:"TDRIFT"`
		TEXT   string  `json:"TEXT"`
		TO     string  `json:"TO"`
		UTC    int64   `json:"UTC"`
		ID     int     `json:"_ID"`
	}
}

/*
RXSpot contains the specific RX.SPOT message structure */
type RXSpot struct {
	Type   string `json:"type"`
	Value  string `json:"value"`
	Params struct {
		CALL   string `json:"CALL"`
		DIAL   int    `json:"DIAL"`
		FREQ   int    `json:"FREQ"`
		GRID   string `json:"GRID"`
		OFFSET int    `json:"OFFSET"`
		SNR    int    `json:"SNR"`
		ID     int    `json:"_ID"`
	}
}

/*
RX.TEXT
STATION.CALLSIGN
STATION.GRID
STATION.INFO
STATION.STATUS
*/

/*
TXFrame contains the specific TX.FRAME message structure */
type TXFrame struct {
	Type   string `json:"type"`
	Value  string `json:"value"`
	Params struct {
		TONES []int `json:"TONES"`
	}
}

/*
TX.TEXT
*/

/*
INBOX.GET_MESSAGES – Get a list of inbox messages
INBOX.STORE_MESSAGE – Store a message in the inbox
MODE.GET_SPEED – Get the TX speed
MODE.SET_SPEED – Set the TX speed
RIG.GET_FREQ – Get the current dial freq and offset
RIG.SET_FREQ – Set the current dial freq and offset
RX.GET_BAND_ACTIVITY – Get the contents of the band activity window
RX.GET_CALL_ACTIVITY – Get the contents of the call activity window
RX.GET_CALL_SELECTED – Get the currently selected callsign
RX.GET_TEXT – Get the contents of the QSO qindow
STATION.GET_CALLSIGN – Get the station callsign
STATION.GET_GRID – Get the station grid
STATION.SET_GRID – Set the station grid
STATION.GET_INFO – Get the station QTH/info
STATION.SET_INFO – Set the station QTH/info
STATION.GET_STATUS – Get the station status
STATION.SET_STATUS – Set the station status
TX.SEND_MESSAGE – Send a message via JS8Call
TX.SET_TEXT – Sets the text in the outgoing message box, but does not send it.
WINDOW.RAISE – Focus the JS8Call window
*/
