package endpoint

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/pkg/errors"
)

// Service represents endpoint service
type Service struct {
	signer Signer
}

// NewService creates new endpoint service
func NewService(signer Signer) Service {
	return Service{signer: signer}
}

func (s Service) NewDenyEndpoint(rootURL, taskToken, candidateEmail string) (string, error) {
	v := url.Values{}
	v.Add("taskToken", taskToken)
	v.Add("candidateEmail", candidateEmail)
	v.Add("decision", "deny")

	base := fmt.Sprintf("%v?%v", rootURL, v.Encode())
	return s.preSign(base)
}

func (s Service) NewApproveEndpoint(rootURL, taskToken, candidateEmail string) (string, error) {
	v := url.Values{}
	v.Add("taskToken", taskToken)
	v.Add("candidateEmail", candidateEmail)
	v.Add("decision", "approve")

	base := fmt.Sprintf("%v?%v", rootURL, v.Encode())
	return s.preSign(base)
}

func (s Service) preSign(url string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", errors.Wrap(err, "while creating new request during preSign")
	}

	expDuration := time.Until(time.Now().Add(5 * time.Minute))
	_, err = s.signer.Presign(req, nil, "execute-api", "eu-central-1", expDuration, time.Now())
	if err != nil {
		return "", errors.Wrap(err, "while presigning")
	}

	return req.URL.String(), nil
}
