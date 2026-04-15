package rustore

import (
	"fmt"

	"github.com/bytedance/sonic"
	"github.com/valyala/fasthttp"
)

const baseURL = "https://backapi.rustore.ru"

var httpClient = &fasthttp.Client{}

type ApiError struct {
	Code    string
	Message string
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("api error: %s - %s", e.Code, e.Message)
}

func doRequest[T any](req *fasthttp.Request) (T, error) {
	var zero T

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := httpClient.Do(req, resp); err != nil {
		return zero, err
	}

	if resp.StatusCode() != fasthttp.StatusOK {
		return zero, fmt.Errorf("http error: %d", resp.StatusCode())
	}

	var apiResp ApiResponse[T]
	if err := sonic.Unmarshal(resp.Body(), &apiResp); err != nil {
		return zero, err
	}

	if apiResp.Code != "OK" {
		msg := ""
		if apiResp.Message != nil {
			msg = *apiResp.Message
		}
		return zero, &ApiError{Code: apiResp.Code, Message: msg}
	}

	return apiResp.Body, nil
}

func apiGet[T any](endpoint string) (T, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(baseURL + endpoint)
	req.Header.SetMethod(fasthttp.MethodGet)

	return doRequest[T](req)
}

func apiPost[T any, B any](endpoint string, reqBody B) (T, error) {
	data, err := sonic.Marshal(reqBody)
	if err != nil {
		var zero T
		return zero, err
	}

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(baseURL + endpoint)
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetContentType("application/json")
	req.SetBody(data)

	return doRequest[T](req)
}

func FormatFileSize(bytes int64) string {
	const kb = 1024
	const mb = kb * 1024
	const gb = mb * 1024

	switch {
	case bytes < kb:
		return fmt.Sprintf("%d B", bytes)
	case bytes < mb:
		return fmt.Sprintf("%.1f KB", float64(bytes)/kb)
	case bytes < gb:
		return fmt.Sprintf("%.1f MB", float64(bytes)/mb)
	default:
		return fmt.Sprintf("%.2f GB", float64(bytes)/gb)
	}
}
