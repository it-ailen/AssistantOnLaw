package accounts

import (
    "foolhttp"
    "net/http"
    "content"
)

func Authorize(r *http.Request) (*content.Account, *foolhttp.HTTPError) {
    mgr := content.GetManager()
    session, err := r.Cookie("session")
    if err != nil || session == nil {
        return nil, foolhttp.NewHTTPError(401, "Miss session", err.Error())
    }
    account, e := mgr.AccountsAuthSession(session.Value)
    if e != nil {
        return nil, foolhttp.NewHTTPError(500, e.Err, e.Detail)
    }
    if account == nil {
        return nil, foolhttp.NewHTTPError(401, "Wrong session", err.Error())
    }
    account.Session = session.Value
    return account, nil
}
