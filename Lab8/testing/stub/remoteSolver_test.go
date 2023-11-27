package stub

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRemoteSolver_Resolve(t *testing.T) {

	type info struct {
		expression string
		statusCode int
		body       string
	}

	var input info
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		expression := r.URL.Query().Get("expression")
		if expression != input.expression {
			t.Fatalf("Expected expression to be %s, but got %s", input.expression, expression)
		}
		rw.WriteHeader(input.statusCode)
		rw.Write([]byte(input.body))
	}))
	defer server.Close()

	rs := RemoteSolver{
		url:    server.URL,
		client: server.Client(),
	}

	data := []struct {
		name   string
		input  info
		result int
		err    string
	}{
		{"ok", info{"1+2", http.StatusOK, "3"}, 3, ""},
		{"invalid_status", info{"1+2", http.StatusBadRequest, "3"}, 0, "invalid status code: 400"},
		{"invalid_body", info{"1+2", http.StatusOK, "x"}, 0, "strconv.ParseInt: parsing \"x\": invalid syntax"},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			input = d.input
			result, err := rs.Resolve(d.input.expression)
			if err != nil {
				if err.Error() != d.err {
					t.Fatalf("Expected error to be %s, but got %s", d.err, err.Error())
				}
			} else {
				if result != d.result {
					t.Fatalf("Expected result to be %d, but got %d", d.result, result)
				}
			}
		})
	}

}
