package main

import (
	"encoding/hex"
	"errors"
	"github.com/zc2638/go-standard/cli/gosl/htmlTemp"
	"gopkg.in/russross/blackfriday.v2"
	"log"
	"net/http"
	"os"
	"path"
	"reflect"
	"strconv"
	"strings"
)

func main() {

	args := os.Args[1:]
	if len(args) > 2 {
		os.Stdout.WriteString("gosl error: '" + strings.Join(args, " ") + "' is not allowed\n")
		os.Stdout.Close()
		return
	}

	var port = "8080"
	if len(args) != 0 {
		switch args[0] {
		case "-v", "-version":
			os.Stdout.WriteString("v0.0.1\n")
			os.Stdout.Close()
			return
		case "-h", "-help":
			help := `
NAME:
    gosl - Go Standard Library Sample Tool
USAGE:
    gosl [global options] [command options]
VERSION:
    v0.0.1
COMMANDS:
    -port           http server port
GLOBAL OPTIONS:
    -help, -h       show help
    -version, -v    print the version
`
			os.Stdout.WriteString(strings.TrimPrefix(help, "\n"))
			os.Stdout.Close()
			return
		case "-port":
			_, err := strconv.Atoi(args[1])
			if err != nil {
				os.Stdout.WriteString("gosl error: '" + strings.Join(args, " ") + "' is not allowed\n")
				os.Stdout.Close()
				return
			}
			port = args[1]
		}
	}

	listen(port)
}

func listen(port string) {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			w.Write([]byte("404 not found"))
			return
		}

		var host string
		if strings.Contains(r.Host, "http://") || strings.Contains(r.Host, "https://") {
			host = r.Host
		} else {
			host = "http://" + r.Host
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		if r.RequestURI == "/" {
			menu := htmlTemp.MenuList(host)
			w.Write([]byte(menu))
			return
		}

		var uri, extraPath string
		extra := false
		backUri := host + path.Dir(r.RequestURI)
		for _, v := range htmlTemp.PackList {
			if "/"+v == r.RequestURI || "/"+v+"/" == r.RequestURI {
				backUri = host
				uri = v
				break
			}
		}
		if uri == "" {
			for _, v := range htmlTemp.PackList {
				if strings.Contains(r.RequestURI, v) {
					extra = true
					uri = v
					extraPath = strings.TrimPrefix(r.RequestURI, "/"+v)
					break
				}
			}
		}

		data, ok := htmlTemp.PackDataMap[uri]
		if !ok && !extra {
			w.Write([]byte("目录不存在"))
			return
		}

		if extra {
			extraArr := strings.Split(extraPath, "/")
			for _, v := range extraArr {
				if v != "" {
					data, ok = data.(map[string]interface{})[v]
					if !ok {
						w.Write([]byte("目录不存在"))
						return
					}
				}
			}
		}

		b, err := parseMap(data, host+r.RequestURI)
		if err != nil {
			w.Write([]byte("目录不存在"))
			return
		}

		w.Write([]byte(htmlTemp.HeaderHtml(host, backUri)))
		w.Write(blackfriday.Run(b))
		w.Write([]byte(htmlTemp.Footer))
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func parseMap(m interface{}, p string) ([]byte, error) {

	var contentBuf []byte
	var buf []byte
	for k, v := range m.(map[string]interface{}) {
		t := reflect.TypeOf(v)
		switch t.Kind() {
		case reflect.String:
			b, err := hex.DecodeString(v.(string))
			if err != nil {
				return nil, err
			}
			contentBuf = append(contentBuf, []byte(k+"\n")...)
			contentBuf = append(contentBuf, b...)
			break
		case reflect.Map:
			buf = append(buf, []byte(`<a href="`+p+`/`+k+`">`+k+`</a><br/><br/>`)...)
			break
		default:
			return nil, errors.New("解析异常")
		}
	}
	buf = append(buf, contentBuf...)

	return buf, nil
}
