package chrombo

import (
	"encoding/json"

	"github.com/gorilla/websocket"
)

type WsClient struct {
	rid  int
	conn *websocket.Conn
}

type Request struct {
	Id     int                    `json:"id"`
	Method string                 `json:"method"`
	Params map[string]interface{} `json:"params"`
}

func (w *WsClient) Send(req *Request) error {
	w.rid++
	req.Id = w.rid
	data, err := json.Marshal(req)
	if err != nil {
		return err
	}

	if err := w.conn.WriteMessage(websocket.TextMessage, data); err != nil {
		return err
	}

	_, data, err = w.conn.ReadMessage()
	println(string(data))
	return err
}

func NewWsClient(url string) (*WsClient, error) {
	MaxReadBufferSize := 0
	MaxWriteBufferSize := 100 * 1024
	d := &websocket.Dialer{
		ReadBufferSize:  MaxReadBufferSize,
		WriteBufferSize: MaxWriteBufferSize,
	}

	var c *websocket.Conn
	c, _, err := d.Dial(url, nil)
	if err != nil {
		return nil, err
	}

	return &WsClient{conn: c}, nil
}
