package gomysql

import (
	"fmt"
	"net/http"
	"net/url"
)

func (db *MySQLConnection) queryLog(err string) {
	if db.Config.ErrorLog.ErrorApiurl == "" || db.Config.ErrorLog.ErrorFromFeild == "" {
		fmt.Println(err)
		return
	}
	if db.Config.ErrorLog.IsPostRequest {
		post(db.Config.ErrorLog.ErrorApiurl, err, db.Config.ErrorLog.ErrorFromFeild)
	} else {
		get(db.Config.ErrorLog.ErrorApiurl, err, db.Config.ErrorLog.ErrorFromFeild)
	}

}

func get(eurl string, err string, param string) {
	req, _ := http.NewRequest("GET", eurl, nil)
	q := req.URL.Query()
	q.Add(param, err)
	req.URL.RawQuery = q.Encode()
	http.Get(req.URL.String())
}

func post(eurl string, err string, param string) {
	data := url.Values{
		param: {err},
	}
	http.PostForm(eurl, data)
}
