package log

import (
	"io/ioutil"
	stlog "log"
	"net/http"
	"os"
)

var log *stlog.Logger

type fileLog string

func (fl fileLog) Write(data []byte) (int, error) {
	f, err := os.OpenFile(string(fl), os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0600)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return f.Write(data)
}

// Init init log
func Init(dest string)  {
	log = stlog.New(fileLog(dest), "distributed： ", stlog.LstdFlags)
}

func write(msg string)  {
	log.Printf("%s\n", msg)
}

func RegisterHandlers()  {
	http.HandleFunc("/log", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case http.MethodPost:
			msg, err := ioutil.ReadAll(request.Body)
			if err != nil {
				writer.WriteHeader(http.StatusBadRequest)
				return
			}
			write(string(msg))
		default:
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
	})

}

