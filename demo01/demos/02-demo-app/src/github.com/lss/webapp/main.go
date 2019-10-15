package main

import (
	"html/template"
	"io/ioutil"

	"net/http"
	"os"

	"pluralSightWebAppWithGo/demo01/demos/02-demo-app/src/github.com/lss/webapp/viewmodel"
)

func main() {
	templates := populateTemplates()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestedFile := r.URL.Path[1:]
		//template := templates.Lookup(requestedFile + ".html")
		template := templates[requestedFile+".html"]
		context := viewmodel.NewBase()
		if template != nil {
			//template.Execute(w, nil)
			//add context to work as data cource
			template.Execute(w, context)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	})

	http.Handle("/img/", http.FileServer(http.Dir("/Users/foureverhh/go/src/pluralSightWebAppWithGo/demo01/demos/02-demo-app/public")))
	http.Handle("/css/", http.FileServer(http.Dir("/Users/foureverhh/go/src/pluralSightWebAppWithGo/demo01/demos/02-demo-app/public")))
	http.ListenAndServe(":8000", nil)
}
func populateTemplates() map[string]*template.Template {
	result := make(map[string]*template.Template)
	const basePath = "/Users/foureverhh/go/src/pluralSightWebAppWithGo/demo01/demos/02-demo-app/templates"
	layout := template.Must(template.ParseFiles(basePath + "/_layout.html"))
	template.Must(layout.ParseFiles(basePath+"/_header.html", basePath+"/_footer.html"))
	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}
	files, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory")
	}
	for _, file := range files {
		f, err := os.Open(basePath + "/content/" + file.Name())
		if err != nil {
			panic("Failed to open template '" + file.Name() + "'")
		}
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to read content from file '" + file.Name() + "'")
		}
		f.Close()
		teml := template.Must(layout.Clone())
		_, err = teml.Parse(string(content))
		if err != nil {
			panic("Failed to parse contents of '" + file.Name() + "'")
		}
		result[file.Name()] = teml
	}
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
