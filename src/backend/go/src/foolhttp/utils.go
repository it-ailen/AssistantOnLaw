package foolhttp

import (
    "github.com/xeipuuv/gojsonschema"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "errors"
)

func JsonSchemaCheck(r *http.Request, schema string, dst interface{}) {
	data, err := ioutil.ReadAll(r.Body)
    if err != nil {
        panic(err)
    }
    schemaLoader := gojsonschema.NewStringLoader(schema)
    sourceLoader := gojsonschema.NewBytesLoader(data)
    result, err := gojsonschema.Validate(schemaLoader, sourceLoader)
    if err != nil {
        panic(err)
    }
    if !result.Valid() {
        for _, e := range result.Errors() {
            panic(BadArgHTTPError(e.String()))
        }
    }
    err = json.Unmarshal(data, dst)
    if err != nil {
        panic(err)
    }
}

func JsonStringCheck(src, schema string, dst interface{}) error {
    schemaLoader := gojsonschema.NewStringLoader(schema)
    sourceLoader := gojsonschema.NewStringLoader(src)
    result, err := gojsonschema.Validate(schemaLoader, sourceLoader)
    if err != nil {
        return err
    }
    if !result.Valid() {
        for _, e := range result.Errors() {
            return errors.New(e.String())
        }
    }
    err = json.Unmarshal([]byte(src), dst)
    return err
}
