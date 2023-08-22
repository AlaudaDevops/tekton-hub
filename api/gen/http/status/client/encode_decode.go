// Code generated by goa v3.12.4, DO NOT EDIT.
//
// status HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/tektoncd/hub/api/design

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	status "github.com/tektoncd/hub/api/gen/status"
	goahttp "goa.design/goa/v3/http"
)

// BuildStatusRequest instantiates a HTTP request object with method and path
// set to call the "status" service "Status" endpoint
func (c *Client) BuildStatusRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: StatusStatusPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("status", "Status", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeStatusResponse returns a decoder for responses returned by the status
// Status endpoint. restoreBody controls whether the response body should be
// restored after having been read.
func DecodeStatusResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body StatusResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("status", "Status", err)
			}
			err = ValidateStatusResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("status", "Status", err)
			}
			res := NewStatusResultOK(&body)
			return res, nil
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("status", "Status", resp.StatusCode, string(body))
		}
	}
}

// unmarshalHubServiceResponseBodyToStatusHubService builds a value of type
// *status.HubService from a value of type *HubServiceResponseBody.
func unmarshalHubServiceResponseBodyToStatusHubService(v *HubServiceResponseBody) *status.HubService {
	if v == nil {
		return nil
	}
	res := &status.HubService{
		Name:   *v.Name,
		Status: *v.Status,
		Error:  v.Error,
	}

	return res
}
