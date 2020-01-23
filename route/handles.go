package route

import (
	"io/ioutil"
	"log"
	"net/http"
	"webStressTesting/config"
	"webStressTesting/template"
)

func Index(w http.ResponseWriter, r *http.Request) {
	//t, _ := template.ParseFiles("./template/index.html")
	//_ = t.Execute(w, nil)
	//直接将html字符串输出，不用导入本地html文件
	_, _ = w.Write([]byte(template.GetIndexHtml()))
}
func Work(w http.ResponseWriter, r *http.Request) {
	//post only
	log.Println("work")
	if r.Method == "POST" {
		log.Println("post")
		commandData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			_, _ = w.Write(config.Work(commandData))
		}

	} else {
		w.WriteHeader(http.StatusForbidden)
	}

}
