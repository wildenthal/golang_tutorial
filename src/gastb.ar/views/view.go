package views

import (
	"html/template"
	"path/filepath"
	"net/http"
)

//Function to read all .gohtml files in layouts directory
var (
	LayoutDir   string = "views/layouts/"
	TemplateDir string = "views/"
	TemplateExt string = ".gohtml"
)

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}

//Interface "View" for rendering pages
type View struct {
	Template *template.Template
	Layout   string
}

func NewView(layout string, files ...string) *View {
	addTemplatePath(files)
	addTemplateExt(files)
	files = append(files,layoutFiles()...)
	t,err := template.ParseFiles(files...)
	if err != nil{
		panic(err)
	}

	return &View{
		Template: t,
		Layout: layout,
	}
}

func (v *View) Render(w http.ResponseWriter,data interface{}) error {
	w.Header().Set("Content-Type", "text/html")
	return v.Template.ExecuteTemplate(w,v.Layout,data)
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := v.Render(w, nil); err != nil {
		panic(err)
	}
}


// addTemplatePath takes in a slice of strings
// representing file paths for templates, and it prepends
// the TemplateDir directory to each string in the slice
//
// e.g. the input {"home"} would result in the output
// {"views/home"} if TemplateDir == "views/"
func addTemplatePath(files []string) {
	for i,f := range files {
		files[i] = TemplateDir + f
	}
}

// addTemplateExt takes in a slice of strings 
// representing file paths for templates and it appends
// the TemplateExt extension to each string in the slice
//
// e.g. the input {"views/home"} would result in the output
// {"views/home.gohtml"} if TemplateExt == ".gohtml"
func addTemplateExt(files []string) {
	for i,f := range files {
		files[i] = f + TemplateExt
	}
}
