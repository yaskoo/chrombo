package chrombo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Browser struct {
	Addr  string
	Pages []Page
}

func (b *Browser) NewPage(url string) (*Page, error) {
	res, err := http.Get(b.Addr + "/json/new?url=" + url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var data []byte
	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var page Page
	if err = json.Unmarshal(data, &page); err != nil {
		return nil, err
	}

	pages := append(b.Pages, page)
	b.Pages = pages
	return &page, nil
}

func NewBrowser(addr string) (*Browser, error) {
	res, err := http.Get(addr + "/json")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var data []byte
	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var pages []Page
	json.Unmarshal(data, &pages)

	b := &Browser{
		Addr:  addr,
		Pages: pages,
	}

	return b, nil
}
