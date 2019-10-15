package main

import (
	"html/template"
	"log"
	"net/http"
)


func main() {
	templates := populateTemplates()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestedFile := r.URL.Path[1:]
		template := templates.Lookup(requestedFile+".html")
		if template != nil {
		   template.Execute(w, nil)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})

	http.Handle("/img/",http.FileServer(http.Dir("/Users/foureverhh/go/src/1-creating-web-applications-go-update-m1-exercise-files/demos/02-demo-app/public")))
	http.Handle("/css/",http.FileServer(http.Dir("/Users/foureverhh/go/src/1-creating-web-applications-go-update-m1-exercise-files/demos/02-demo-app/public")))
	http.ListenAndServe(":8000",nil)
}

func populateTemplates() *template.Template {
	result:= template.New("template")
	const basePath = "templates"
	template.Must(result.ParseGlob(basePath + "/*.html"))
	return result
}
/* 
//This does not work with template - content
func populateTemplates() *template.Template {
	result := template.New("templates")
	const basePath = "/Users/foureverhh/go/src/1-creating-web-applications-go-update-m1-exercise-files/demos/02-demo-app/templates"
	template.Must(result.ParseGlob(basePath + "/*.html"))
	return result
}
 */
/*
func main() {
	http.ListenAndServe(":8000", http.FileServer(http.Dir("/Users/foureverhh/go/src/1-creating-web-applications-go-update-m1-exercise-files/demos/02-demo-app/public")))

} */

/*
import (
	"net/http"
)

func main()  {
	http.HandleFunc("/",func (w http.ResponseWriter, r *http.Requset)  {
		http.ServeFile(w,r,"/Users/foureverhh/go/src/1-creating-web-applications-go-update-m1-exercise-files/demos/02-demo-app/public"+r.URL.Path)
	})
	http.ListenAndServe(":8000",nil)
}
*/

/*
import (
		"io"
		"os"
		"log"
		"net/http"
		"strings"
	)
func main()  {
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		//http.ServeFile(w,r,"/Users/foureverhh/go/src/1-creating-web-applications-go-update-m1-exercise-files/demos/02-demo-app/public"+r.URL.Path)

	//To give responce to request
	f, err := os.Open("/Users/foureverhh/go/src/1-creating-web-applications-go-update-m1-exercise-files/demos/02-demo-app/public"+r.URL.Path)
		if err!=nil{
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}
		defer f.Close()
		var contentType string
		switch  {
		case strings.HasSuffix(r.URL.Path,"css"):
			contentType="text/css"
		case strings.HasSuffix(r.URL.Path,"html"):
			contentType="text/html"
		case strings.HasSuffix(r.URL.Path,"png"):
			contentType="image/png"
		default:
			contentType="text/plain"
		}
		w.Header().Add("Content-Type",contentType)
		io.Copy(w,f)

	})
	http.ListenAndServe(":8000",nil)
} */
