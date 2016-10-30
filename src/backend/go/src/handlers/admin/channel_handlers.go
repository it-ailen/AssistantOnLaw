package admin

import (
	"foolhttp"
	"net/http"
	"log"
	"content"
	"encoding/json"
)

type ChannelHandler struct {}

func (self *ChannelHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	foolhttp.DoServeHTTP(self, w, r)
}

func (self *ChannelHandler) GET(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
	mgr := content.GetManager()
	channels, err := mgr.ChannelsGet(nil)
	if err != nil {
		return foolhttp.UnknownHTTPError(err.Error())
	}
	res, err := json.Marshal(channels)
	if err != nil {
		return foolhttp.UnknownHTTPError(err.Error())
	}
	headers := w.Header()
	headers["Content-Type"] = []string{"application/json;charset=utf-8"}
	w.Write(res)
	return nil
}

func (self *ChannelHandler) POST(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
	type argsChannelPost struct {
		Name string `json:"name"`
		Icon string `json:"icon"`
	}
	args := argsChannelPost{}
	err := foolhttp.ParseJsonArgs(r, &args)
	if err != nil {
		return err
	}
	if len(args.Name) == 0 {
		return foolhttp.NewHTTPError(400, foolhttp.ErrorBadArgument, "Missing `name`")
	}
	if len(args.Icon) == 0 {
		return foolhttp.NewHTTPError(400, foolhttp.ErrorBadArgument, "Missing `icon`")
	}
	log.Printf("name: %#v", args)
	mgr := content.GetManager()
	channel := content.Channel{
		ID: mgr.AllocateId(true),
		Name: args.Name,
		Icon: args.Icon,
	}
	e := channel.Validate()
	if e != nil {
		return foolhttp.NewHTTPError(400, foolhttp.ErrorBadArgument, e.Error())
	}
	e = mgr.ChannelCreate(&channel)
	if e != nil {
		return foolhttp.NewHTTPError(500, foolhttp.ErrorUnknown, e.Error())
	}
	return nil
}


func (self *ChannelHandler) PUT(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
	type argsChannelPut struct {
		Name string `json:"name"`
		Icon string `json:"icon"`
	}
	args := argsChannelPut{}
	e := foolhttp.ParseJsonArgs(r, &args)
	if e != nil {
		return e
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
	channel.Name = args.Name
	channel.Icon = args.Icon
	err = channel.Validate()
	if err != nil {
		return foolhttp.NewHTTPError(400, foolhttp.ErrorBadArgument, err.Error())
	}
	err = mgr.ChannelUpdate(channel)
	if err != nil {
		return foolhttp.NewHTTPError(500, foolhttp.ErrorUnknown, err.Error())
	}
	return nil
}

func (self *ChannelHandler) DELETE(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
	id := foolhttp.RouteArgument(r, "id")
	mgr := content.GetManager()
	channel, err := mgr.ChannelGet(id)
	if err != nil {
		return foolhttp.NewHTTPError(500, foolhttp.ErrorUnknown, err.Error())
	}
	if channel == nil {
		return foolhttp.NewHTTPError(404, foolhttp.ErrorUnknownResource, "")
	}
	err = mgr.ChannelDelete(channel)
	if err != nil {
		return foolhttp.NewHTTPError(500, foolhttp.ErrorUnknown, err.Error())
	}
	return nil
}

func NewChannelHandler() *ChannelHandler {
	handler := new(ChannelHandler)
	return handler
}
