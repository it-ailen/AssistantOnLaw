package resource

import (
    "net/http"
    "foolhttp"
    "handlers/accounts"
    "content"
    "log"
)

type FaLvWenDaClassHandler struct{}

func (self *FaLvWenDaClassHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    foolhttp.DoServeHTTP(self, w, r)
}

func (self *FaLvWenDaClassHandler) POST(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
    accounts.CheckAccessibility(r, &accounts.AccessControl{
        SuperOnly: true,
    })

    type argsDefine struct {
        Name string `json:"name"`
    }
    schemaDefine := `{
		"type": "object",
		"properties": {
			"name": {"type": "string"}
		},
		"required": ["name"]
	}`
    args := argsDefine{}
    foolhttp.JsonSchemaCheck(r, schemaDefine, &args)

    mgr := content.GetManager()
    class, err := mgr.CreateFaLvWenDaClass(args.Name)
    if err != nil {
        panic(err)
    }
    foolhttp.WriteJson(w, class)
    return nil
}

func (self *FaLvWenDaClassHandler) PUT(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
    accounts.CheckAccessibility(r, &accounts.AccessControl{
        SuperOnly: true,
    })

    type argsDefine struct {
        Name string `json:"name"`
    }
    schemaDefine := `{
		"type": "object",
		"properties": {
			"name": {"type": "string"}
		},
		"required": ["name"]
	}`
    args := argsDefine{}
    foolhttp.JsonSchemaCheck(r, schemaDefine, &args)

    id := foolhttp.RouteArgument(r, "id")
    mgr := content.GetManager()

    filter := content.FaLvWenDaClassesFilter{
        ID: []string{id},
    }
    classes, err := mgr.LoadFaLvWenDaClasses(&filter)
    if err != nil {
        panic(err)
    }
    if len(classes) == 0 {
        panic(foolhttp.UnknownHTTPError("Unknown id"))
    }
    class := classes[0]
    err = mgr.UpdateFaLvWenDaClass(class, args.Name)
    if err != nil {
        panic(err)
    }
    foolhttp.WriteJson(w, class)
    return nil
}

func (self *FaLvWenDaClassHandler) GET(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
    mgr := content.GetManager()
    classes, err := mgr.LoadFaLvWenDaClasses(nil)
    if err != nil {
        panic(err);
    }
    foolhttp.WriteJson(w, classes)
    return nil
}

func (self *FaLvWenDaClassHandler) DELETE(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
    id := foolhttp.RouteArgument(r, "id")
    mgr := content.GetManager()

    filter := content.FaLvWenDaClassesFilter{
        ID: []string{id},
    }
    classes, err := mgr.LoadFaLvWenDaClasses(&filter)
    if err != nil {
        panic(err)
    }
    if len(classes) == 0 {
        panic(foolhttp.UnknownHTTPError("Unknown id"))
    }
    mgr.DeleteFaLvWenDaClasses(id)
    return nil
}

func NewFaLvWenDaClassHandler() *FaLvWenDaClassHandler {
    handler := new(FaLvWenDaClassHandler)
    return handler
}

type FaLvWenDaArticleHandler struct{}

func (self *FaLvWenDaArticleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    foolhttp.DoServeHTTP(self, w, r)
}

func (self *FaLvWenDaArticleHandler) POST(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
    accounts.CheckAccessibility(r, &accounts.AccessControl{
        SuperOnly: true,
    })

    type argsDefine struct {
        Name    string `json:"name"`
        ClassId string `json:"class_id"`
        Content string `json:"content"`
    }
    schemaDefine := `{
		"type": "object",
		"properties": {
			"name": {"type": "string"},
			"class_id": {"type": "string"},
			"content": {"type": "string"}
		},
		"required": ["name", "class_id", "content"]
	}`
    args := argsDefine{}
    foolhttp.JsonSchemaCheck(r, schemaDefine, &args)

    mgr := content.GetManager()
    article, err := mgr.CreateFaLvWenDaArticle(args.Name, args.ClassId, args.Content)
    if err != nil {
        panic(err)
    }
    foolhttp.WriteJson(w, article)
    return nil
}

func (self *FaLvWenDaArticleHandler) PUT(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
    accounts.CheckAccessibility(r, &accounts.AccessControl{
        SuperOnly: true,
    })

    schemaDefine := `{
		"type": "object",
		"properties": {
			"name": {"type": "string"},
			"class_id": {"type": "string"},
			"content": {"type": "string"}
		}
	}`
    args := make(map[string]interface{})
    foolhttp.JsonSchemaCheck(r, schemaDefine, &args)

    id := foolhttp.RouteArgument(r, "id")
    mgr := content.GetManager()
    filter := &content.FaLvWenDaArticlesFilter{
        ID: []string{id},
    }
    articles, err := mgr.LoadFaLvWenDaArticles(filter)
    if err != nil {
        panic(err)
    }
    if len(articles) == 0 {
        panic(foolhttp.NotFoundHTTPError("Unknown id"))
    }
    article := articles[0]
    log.Printf("error: %#v", err)
    err = mgr.UpdateFaLvWenDaArticle(article, args)
    if err != nil {
        log.Printf("err: %#v", err)
        panic(err)
    }
    foolhttp.WriteJson(w, article)
    return nil
}


func (self *FaLvWenDaArticleHandler) DELETE(w http.ResponseWriter, r *http.Request) *foolhttp.HTTPError {
    id := foolhttp.RouteArgument(r, "id")
    mgr := content.GetManager()

    filter := content.FaLvWenDaArticlesFilter{
        ID: []string{id},
    }
    articles, err := mgr.LoadFaLvWenDaArticles(&filter)
    if err != nil {
        panic(err)
    }
    if len(articles) == 0 {
        panic(foolhttp.UnknownHTTPError("Unknown id"))
    }
    mgr.DeleteFaLvWenDaArticle(id)
    return nil
}

func NewFaLvWenDaArticleHandler() *FaLvWenDaArticleHandler {
    handler := new(FaLvWenDaArticleHandler)
    return handler
}
