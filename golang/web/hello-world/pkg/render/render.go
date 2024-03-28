package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/larry-lw-chan/helloworld/pkg/config"
)

const TMPL_PATH string = "./templates/"
const TMPL_PATTERN string = TMPL_PATH + "*.page.tmpl"
const LAYOUT_PATTERN string = TMPL_PATH + "*.layout.tmpl"

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

// Render is the function to render the template
func Html(w http.ResponseWriter, tmpl string) {
	// create template cache
	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// get requested template
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	err = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	// render template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob(TMPL_PATTERN)

	if err != nil {
		return myCache, err
	}

	// loop through all files ending *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)

		// First we parse the page template
		ts, err := template.ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// Next we make sure layout files exist
		matches, err := filepath.Glob(LAYOUT_PATTERN)
		if err != nil {
			log.Println(err)
			// return myCache, err
		}

		// If layout files found, we now parse them too
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		// add template to cache
		myCache[name] = ts
	}

	// return the cache
	return myCache, nil
}

// var tc = make(map[string]*template.Template)

// func Html(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error

// 	// Check to see if the template is in the map
// 	_, inMap := tc[t]

// 	if !inMap {
// 		//need to create template
// 		log.Println("creating template and adding to cache")
// 		err = createTemplateCache(t)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	} else {
// 		// we have template in cache
// 		log.Println("using cached template")
// 	}

// 	tmpl = tc[t]
// 	err = tmpl.Execute(w, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func createTemplateCache(t string) error {
// 	tmplfile := TMPL_PATH + t
// 	layout := TMPL_PATH + DEFAULT_LAYOUT
// 	templates := []string{tmplfile, layout}

// 	// parse the template
// 	tmpl, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}

// 	// add template to cache (map)
// 	tc[t] = tmpl
// 	return nil
// }
