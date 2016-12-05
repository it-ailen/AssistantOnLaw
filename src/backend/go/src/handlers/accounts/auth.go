package accounts

import (
    "foolhttp"
    "net/http"
    "content"
    "content/definition"
)

func Authorize(r *http.Request) *definition.Account {
    mgr := content.GetManager()
    session, err := r.Cookie("session")
    if err != nil || session == nil {
        panic(foolhttp.NewHTTPError(401, "Miss session", err.Error()))
    }
    account, e := mgr.AccountsAuthSession(session.Value)
    if e != nil {
        panic(foolhttp.NewHTTPError(500, e.Err, e.Detail))
    }
    if account == nil {
        panic(foolhttp.NewHTTPError(401, "Wrong session", "Auth failed"))
    }
    account.Session = session.Value
    return account
}

type AccessControl struct {
    SuperOnly bool `json:"super_only"`
}
func CheckAccessibility(r *http.Request, access *AccessControl) *definition.Account {
    me := Authorize(r)
    if access.SuperOnly {
        if me.Type != definition.C_ACC_TYPE_SUPER {
            panic(foolhttp.ForbiddenHTTPError("Super only"))
        }
    }
    return me
}
