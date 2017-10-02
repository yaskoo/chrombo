package chrombo

type Page struct {
	Id       string `json:id`
	Url      string `json:"url"`
	DebugUrl string `json:"webSocketDebuggerUrl"`

	ws *WsClient
}

func (p *Page) Navigate(url string) error {
	p.SendRequest(&Request{
		Method: "Page.navigate",
		Params: map[string]interface{}{"url": url},
	})
	return nil
}

func (p *Page) Evaluate(script string) error {
	p.SendRequest(&Request{
		Method: "Runtime.evaluate",
		Params: map[string]interface{}{"expression": script, "returnByValue": true},
	})
	return nil
}

func (p *Page) SendRequest(req *Request) error {
	if p.ws == nil {
		ws, err := NewWsClient(p.DebugUrl)
		if err != nil {
			return err
		}

		p.ws = ws
	}

	p.ws.Send(req)
	return nil
}
