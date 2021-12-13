package beegoutil

import (
	"net/http"

	"github.com/grokify/mogo/net/httputilmore"
)

func WriteHtml(rw http.ResponseWriter, html []byte) (int, error) {
	rw.Header().Add(httputilmore.HeaderContentType, httputilmore.ContentTypeTextHtmlUtf8)
	return rw.Write([]byte(html))
}
