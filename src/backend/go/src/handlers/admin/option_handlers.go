package admin

import (
	"foolhttp"
	"net/http"
	"content"
	"encoding/json"
	"log"
	"content/definition"
)

type OptionHandler struct {}

func (self *OptionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	foolhttp.DoServeHTTP(self, w, r)
}

func (self *OptionHandler) GET(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
	mgr := content.GetManager()
	parentId := foolhttp.QueryArgument(r, "parent_id")
	filter := content.OptionsFilter{
		Parents: []string{parentId},
	}
	options, err := mgr.OptionsGet(&filter)
	if err != nil {
		return foolhttp.UnknownHTTPError(err.Error())
	}
	res, err := json.Marshal(options)
	if err != nil {
		return foolhttp.UnknownHTTPError(err.Error())
	}
	w.Write(res)
	return nil
}

func (self *OptionHandler) POST(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
	mgr := content.GetManager()
	option := definition.Option{
		ID: mgr.AllocateId(true),
	}
	err := foolhttp.ParseJsonArgs(r, &option)
	if err != nil {
		return err
	}
	if option.Type == definition.C_ST_report {
		option.Report.ID = mgr.AllocateId(true)
		for _, c := range option.Report.Cases {
			c.ID = mgr.AllocateId(true)
		}
		for _, decree := range option.Report.Decrees {
			decree.ID = mgr.AllocateId(true)
		}
	}
	log.Printf("%#v", &option)
	e := option.Validate()
	if e != nil {
		return foolhttp.NewHTTPError(400, foolhttp.ErrorBadArgument, e.Error())
	}
	e = mgr.OptionCreate(&option)
	if e != nil {
		return foolhttp.UnknownHTTPError(e.Error())
	}
	return nil
}


func (self *OptionHandler) PUT(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
	mgr := content.GetManager()
	option := definition.Option{
		ID: mgr.AllocateId(true),
	}
	err := foolhttp.ParseJsonArgs(r, &option)
	if err != nil {
		return err
	}
	if option.Type == definition.C_ST_report {
		if len(option.Report.ID) == 0 {
			option.Report.ID = mgr.AllocateId(true)
		}
		for _, c := range option.Report.Cases {
			if len(c.ID) == 0 {
				c.ID = mgr.AllocateId(true)
			}
		}
		for _, decree := range option.Report.Decrees {
			if len(decree.ID) == 0 {
				decree.ID = mgr.AllocateId(true)
			}
		}
	}
	log.Printf("%#v", &option)
	e := option.Validate()
	if e != nil {
		return foolhttp.NewHTTPError(400, foolhttp.ErrorBadArgument, e.Error())
	}
	return nil
}

func (self *OptionHandler) DELETE(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
	id := foolhttp.RouteArgument(r, "id")
	mgr := content.GetManager()
	option, err := mgr.OptionGet(id)
	if err != nil {
		return foolhttp.UnknownHTTPError(err.Error())
	}
	if option == nil {
		return foolhttp.NotFoundHTTPError("Unknown option id")
	}
	err = mgr.OptionDelete(option)
	if err != nil {
		return foolhttp.UnknownHTTPError(err.Error())
	}
	return nil
}

func NewOptionHandler() *OptionHandler {
	handler := new(OptionHandler)
	return handler
}
