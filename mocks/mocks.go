// Provices mock objects until Square has their own api sandbox
package mocks

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"io/ioutil"
	"path/filepath"
	//"fmt"
	"log"
)

const (
	G = "GET"
	P = "POST"
	U = "PUT"
	D = "DELETE"
)

type res struct {
	code int
	list bool
	filename string
}

type Mock struct {
	Server *httptest.Server
	Client *http.Client
	things map[string]*map[string]res
}

var notfound = []byte(`
{
	"type":"",
	"message":"Not Found"
}
`)

func New() *Mock{
	s := Mock{}
	
	s.things = make(map[string]*map[string]res)
	
	s.Server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    s.handle(w, r)
  }))
  
  transport := &http.Transport{
    Proxy: func(req *http.Request) (*url.URL, error) {
      return url.Parse(s.Server.URL)
    },
  }

  s.Client = &http.Client{Transport: transport}	
  s.Init()
  return &s
}

func (this *Mock) handle(w http.ResponseWriter, r *http.Request) {
	code, data, list, err := this.read(r.URL.Path, r.Method)
	if err != nil { log.Panic("cannot read url: ", r.URL.Path) }

  w.Header().Set("Content-Type", "application/json")
  if (list) {
	  w.Header().Set("Link", "<http://connect.square.com" + r.URL.Path + "/next>;rel='next'")
	}

  w.WriteHeader(code)
	w.Write(*data)
}

func (this *Mock) add(endpoint string, method string, code int, filename string, list bool) {
	
	path := "/v1/me/" + endpoint
	
	r := res{
		code: code,
		list: list,
		filename: filename,
	}
	
	t := this.things[path]
	if t != nil {
		(*t)[method] = r
		return
	}
	
	m := make(map[string]res)
	m[method] = r
	this.things[path] = &m	
}

func (this *Mock) addFolder(folder string) {
	dir, err := filepath.Abs("../mocks/" + folder)
	if err != nil { log.Println("Dir not found", folder); return }
	
	files, err := ioutil.ReadDir(dir)
	if err != nil { log.Println("Dir not found", folder); return }

	for _, fileInfo := range files {
		if !fileInfo.IsDir() {
			filename := fileInfo.Name()
			file := filepath.Join(dir, filename)

			switch filename {
				case "list.json":
					this.add(folder, G, http.StatusOK, file, true)
				case "retrieve.json":
					this.add(folder + "/yes", G, http.StatusOK, file, false)
				case "create.json":
					this.add(folder, P, http.StatusOK, file, false)
				case "update.json":
					this.add(folder + "/yes", U, http.StatusOK, file, false)
				case "next.json":
					this.add(folder + "/next", G, http.StatusOK, file, false)
			}

		}
	}
}

func (this *Mock) addAll() {
	dir, _ := filepath.Abs("../mocks")
	
	files, _ := ioutil.ReadDir(dir)
	
	for _, fileInfo := range files {
		if fileInfo.IsDir() {
			this.addFolder(fileInfo.Name())
		}
	}
}

func (this *Mock) read(endpoint string, method string) (int, *[]byte, bool, error) {
	thing := this.things[endpoint]
	if thing == nil {
		return http.StatusNotFound, &notfound, false, nil
	}
	
	r := (*thing)[method]
	
	data, err := ioutil.ReadFile(r.filename)
	if err != nil { return 0, nil, false, err }
	
	return r.code, &data, r.list, nil
}

func (this *Mock) Init() {
	this.addAll()
}

