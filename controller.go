package beegoutil

import (
	"net/http"

	"github.com/grokify/mogo/net/http/httputilmore"
)

func WriteHtml(rw http.ResponseWriter, html []byte) (int, error) {
	rw.Header().Add(httputilmore.HeaderContentType, httputilmore.ContentTypeTextHTMLUtf8)
	return rw.Write([]byte(html))
}
