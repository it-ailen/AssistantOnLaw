package admin

import (
	"foolhttp"
	"net/http"
	"content"
	"encoding/json"
)

type EntryHandler struct {}

func (self *EntryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	foolhttp.DoServeHTTP(self, w, r)
}

func (self *EntryHandler) GET(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
	mgr := content.GetManager()
	id := foolhttp.RouteArgumentWithDefault(r, "id", "")
	if id == "" {
		channelId := foolhttp.QueryArgument(r, "channel_id")
		entries, err := mgr.EntriesGet(channelId)
		if err != nil {
			return foolhttp.UnknownHTTPError(err.Error())
		}
		res, err := json.Marshal(entries)
		if err != nil {
			return foolhttp.UnknownHTTPError(err.Error())
		}
		w.Write(res)
	} else {
		entry, err := mgr.EntryGet(id)
		if err != nil {
			return foolhttp.UnknownHTTPError(err.Error())
		}
		if entry == nil {
			return foolhttp.NewHTTPError(404, foolhttp.ErrorNotFound, "Unknown id")
		}
		res, err := json.Marshal(entry)
		if err != nil {
			return foolhttp.UnknownHTTPError(err.Error())
		}
		w.Write(res)
	}
	return nil
}

func (self *EntryHandler) POST(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
	mgr := content.GetManager()
	entry := content.Entry{
		ID: mgr.AllocateId(true),
	}
	err := foolhttp.ParseJsonArgs(r, &entry)
	if err != nil {
		return err
	}
	e := entry.Validate()
	if e != nil {
		return foolhttp.BadArgHTTPError(e.Error())
	}
	e = mgr.EntryCreate(&entry)
	if e != nil {
		return foolhttp.UnknownHTTPError(e.Error())
	}
	return nil
}


func (self *EntryHandler) PUT(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
	mgr := content.GetManager()
	entry := content.Entry{}
	err := foolhttp.ParseJsonArgs(r, &entry)
	if err != nil {
		return err
	}
	e := entry.Validate()
	if e != nil {
		return foolhttp.BadArgHTTPError(e.Error())
	}
	e = mgr.EntryUpdate(&entry)
	if e != nil {
		return foolhttp.UnknownHTTPError(e.Error())
	}
	return nil
}

func (self *EntryHandler) DELETE(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
	//id := foolhttp.RouteArgument(r, "id")
	//mgr := content.GetManager()
	//channel, err := mgr.ChannelGet(id)
	//if err != nil {
	//	return foolhttp.NewHTTPError(500, foolhttp.ErrorUnknown, err.Error())
	//}
	//if channel == nil {
	//	return foolhttp.NewHTTPError(404, foolhttp.ErrorUnknownResource, "")
	//}
	//err = mgr.ChannelDelete(channel)
	//if err != nil {
	//	return foolhttp.NewHTTPError(500, foolhttp.ErrorUnknown, err.Error())
	//}
	return nil
}

func NewEntryHandler() *EntryHandler {
	handler := new(EntryHandler)
	return handler
}
