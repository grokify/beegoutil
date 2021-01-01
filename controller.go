package beegoutil

import (
	"net/http"

	hum "github.com/grokify/simplego/net/httputilmore"
)

func WriteHtml(rw http.ResponseWriter, html []byte) (int, error) {
	rw.Header().Add(hum.HeaderContentType, hum.ContentTypeTextHtmlUtf8)
	return rw.Write([]byte(html))
}
