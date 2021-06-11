package gomysql


// func queryLog(err string)  {
	
// }

// func (db *MySQLConnection) queryLog(method string,query string,err string)  {
// 	db.insert("INSERT INTO log.auth_log(error_at,error) VALUES(?,?)", []string{f, e})
// 	a := "#AuthAcad  #" + inc.Getenv("INSTUTION_NAME") + "   " + inc.Getenv("ACAD_DOMAIN") + "  Mysql (" + f + ")"
// 	Telgram(a, e)
// 	fmt.Println("Error For => " + a)
// 	fmt.Println("Error At => " + time.Now().Local().String() + "e")
// }

// func Dlog(e string) {
// 	a := "#AuthAcad  #" + inc.Getenv("INSTUTION_NAME") + "   " + inc.Getenv("ACAD_DOMAIN") + "  Mysql (DB Connection)"
// 	Telgram(a, e)
// }

// func Telgram(a string, e string) {

// 	url := "https://api.telegram.org/bot" + inc.Getenv("BOT_KEY") + "/sendMessage?chat_id=" + inc.Getenv("CHAT_ID") + "&parse_mode=Markdown&text==========__Error__==========%0a%0a" + url.QueryEscape(a) + "%0a%0a%0a" + url.QueryEscape(e) + "%0a%0a%0a Time => " + url.QueryEscape(time.Now().Local().String()) + "%0a%0a%0a==========__Error__=========="
// 	http.Get(url)

// }
