package endpoint

import (
	"io"
	"net/http"
	"time"
)

type Signer interface {
	Presign(r *http.Request, body io.ReadSeeker, service, region string, exp time.Duration, signTime time.Time) (http.Header, error)
}
