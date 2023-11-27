package stub

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type RemoteSolver struct {
	url    string
	client *http.Client
}

func (rs RemoteSolver) Resolve(expression string) (int, error) {
	url := fmt.Sprintf("%s?expression=%s", rs.url, url.QueryEscape(expression))
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, url, nil)
	if err != nil {
		return 0, err
	}
	resp, err := rs.client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("invalid status code: %d", resp.StatusCode)
	}

	res, err := strconv.ParseInt(string(result), 10, 32)
	if err != nil {
		return 0, err
	}
	return int(res), nil
}
