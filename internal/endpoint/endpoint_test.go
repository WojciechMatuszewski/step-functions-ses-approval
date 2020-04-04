package endpoint_test

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"step-functions/internal/endpoint"
	"step-functions/internal/endpoint/mock"
)

const (
	ROOT_URL        = "/ROOT_URL"
	TASK_TOKEN      = "TASK_TOKEN"
	CANDIDATE_EMAIL = "MOCK_EMAIL"
)

func TestService_NewApproveEndpoint(t *testing.T) {
	t.Parallel()

	t.Run("signer failure", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		signer := mock.NewMockSigner(ctrl)
		service := endpoint.NewService(signer)
		encodedPath := encodedTestPath("approve")

		req, _ := http.NewRequest(http.MethodGet, encodedPath, nil)
		signer.EXPECT().Presign(gomock.Any(), nil, "execute-api", "eu-central-1", gomock.Any(), gomock.Any()).Return(req.Header, errors.New("boom"))

		approveURL, err := service.NewApproveEndpoint(ROOT_URL, TASK_TOKEN, CANDIDATE_EMAIL)

		assert.Error(t, err)
		assert.Empty(t, approveURL)
	})

	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		signer := mock.NewMockSigner(ctrl)
		service := endpoint.NewService(signer)
		encodedPath := encodedTestPath("approve")

		req, _ := http.NewRequest(http.MethodGet, encodedPath, nil)
		signer.EXPECT().Presign(req, nil, "execute-api", "eu-central-1", gomock.Any(), gomock.Any()).Return(req.Header, nil)

		approveURL, err := service.NewApproveEndpoint(ROOT_URL, TASK_TOKEN, CANDIDATE_EMAIL)
		assert.NoError(t, err)
		assert.Equal(t, approveURL, encodedPath)
	})
}

func TestService_NewDenyEndpoint(t *testing.T) {
	t.Parallel()

	t.Run("signer failure", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		signer := mock.NewMockSigner(ctrl)
		service := endpoint.NewService(signer)
		encodedPath := encodedTestPath("deny")

		req, _ := http.NewRequest(http.MethodGet, encodedPath, nil)
		signer.EXPECT().Presign(req, nil, "execute-api", "eu-central-1", gomock.Any(), gomock.Any()).Return(req.Header, errors.New("boom"))

		denyURL, err := service.NewDenyEndpoint(ROOT_URL, TASK_TOKEN, CANDIDATE_EMAIL)
		assert.Error(t, err)
		assert.Empty(t, denyURL)
	})

	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		signer := mock.NewMockSigner(ctrl)
		service := endpoint.NewService(signer)
		encodedPath := encodedTestPath("deny")

		req, _ := http.NewRequest(http.MethodGet, encodedPath, nil)
		signer.EXPECT().Presign(req, nil, "execute-api", "eu-central-1", gomock.Any(), gomock.Any()).Return(req.Header, nil)

		denyURL, err := service.NewDenyEndpoint(ROOT_URL, TASK_TOKEN, CANDIDATE_EMAIL)
		assert.NoError(t, err)
		assert.Equal(t, denyURL, encodedPath)
	})
}

func encodedTestPath(decision string) string {
	v := url.Values{}
	v.Add("taskToken", TASK_TOKEN)
	v.Add("candidateEmail", CANDIDATE_EMAIL)
	v.Add("decision", decision)
	return fmt.Sprintf("%v?%v", ROOT_URL, v.Encode())
}
