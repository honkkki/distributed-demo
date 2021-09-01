package user

import (
	"bytes"
	"net/http"
)

func RegisterHandlers() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		addr := request.RemoteAddr
		buf := new(bytes.Buffer)
		_, err := buf.WriteString("your ip: " + addr)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = writer.Write(buf.Bytes())
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
	})

}
