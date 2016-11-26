package mobile

import (
	"foolhttp"
	"net/http"
	"encoding/json"
	"content"
)

type HomeLayoutHandler struct {}

func (self *HomeLayoutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	foolhttp.DoServeHTTP(self, w, r)
}

func (self *HomeLayoutHandler) GET(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
	type slide struct {
		Src string `json:"src"`
	}
	type channelData struct {
		ID string `json:"id"`
		Icon string `json:"icon"`
		Text string `json:"text"`
	}
	res := struct {
		Slides []*slide `json:"slides"`
		Channels []*channelData `json:"channels"`
	}{
		Slides: []*slide{},
		Channels: []*channelData{},
	}

	mgr := content.GetManager()
	channels, err := mgr.ChannelsGet(nil)
	if err != nil {
		return foolhttp.UnknownHTTPError(err.Error())
	}
	for _, c := range channels {
        data := channelData{
            ID: c.ID,
	        Icon: c.Icon,
	        Text: c.Name,
        }
		res.Channels = append(res.Channels, &data)
	}

	body, _ := json.Marshal(&res)
	w.Write(body)
	return nil
}

func NewHomeLayoutHandler() *HomeLayoutHandler {
	handler := &HomeLayoutHandler{}
	return handler
}


type ChannelLayoutHandler struct {}

func (self *ChannelLayoutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	foolhttp.DoServeHTTP(self, w, r)
}

func (self *ChannelLayoutHandler) GET(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
	type entry struct {
		ID string `json:"id"`
		Text string `json:"text"`
		LayoutType string `json:"layout_type"`
		ChannelId string `json:"channel_id"`
	}
	type response struct {
		ID string `json:"id"`
		Name string `json:"name"`
		Icon string `json:"icon"`
		Entries []*entry `json:"entries"`
	}
	id := foolhttp.RouteArgument(r, "id")
	mgr := content.GetManager()
	channel, err := mgr.ChannelGet(id)
	if err != nil {
		return foolhttp.NewHTTPError(500, foolhttp.ErrorUnknown, err.Error())
	}
	if channel == nil {
		return foolhttp.NewHTTPError(404, foolhttp.ErrorUnknownResource, "")
	}
	entries, err := mgr.EntriesGet(id)
	res := response{
		ID: channel.ID,
		Name: channel.Name,
		Icon: channel.Icon,
		Entries: []*entry{},
	}
	for _, e := range entries {
		entryRes := entry{
			ID: e.ID,
			Text: e.Text,
			LayoutType: e.LayoutType,
			ChannelId: channel.ID,
		}
		res.Entries = append(res.Entries, &entryRes)
	}
	body, _ := json.Marshal(&res)
	w.Write(body)
	return nil
}

func NewChannelLayoutHandler() *ChannelLayoutHandler {
	handler := &ChannelLayoutHandler{}
	return handler
}

type EntryLayoutHandler struct {}

func (self *EntryLayoutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	foolhttp.DoServeHTTP(self, w, r)
}

func (self *EntryLayoutHandler) GET(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
	id := foolhttp.RouteArgument(r, "id")
	mgr := content.GetManager()
	tree, err := mgr.EntryTreeGet(id)
	if err != nil {
		return foolhttp.UnknownHTTPError(err.Error())
	}
	body, _ := json.Marshal(&tree)
	w.Write(body)
	return nil
}

func NewEntryLayoutHandler() *EntryLayoutHandler {
	handler := &EntryLayoutHandler{}
	return handler
}