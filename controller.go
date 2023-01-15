package beegoutil

import (
	"net/http"

	"github.com/grokify/mogo/net/http/httputilmore"
)

func WriteHTML(rw http.ResponseWriter, html []byte) (int, error) {
	rw.Header().Add(httputilmore.HeaderContentType, httputilmore.ContentTypeTextHTMLUtf8)
	return rw.Write(html)
}
