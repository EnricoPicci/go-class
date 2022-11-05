package remote

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
)

const serverUrl = "http://localhost:8081/"

var NoServerErr error = errors.New("No server available at " + serverUrl)
var TimeoutErr error = errors.New("A timeout occurred for server at " + serverUrl)

type File struct {
	Name string
}

func NewFile(fName string) (file *File, respBody string, err error) {
	encoded, err := encodeFileName(fName)
	if err != nil {
		return
	}
	resp, _err := http.Post(serverUrl+"create", "application/json", bytes.NewBuffer(encoded))
	if _err != nil {
		err = buildError(_err)
		return
	}
	defer resp.Body.Close()
	body, _err := io.ReadAll(resp.Body)
	if _err != nil {
		err = buildError(_err)
		return
	}

	respBody = string(body)
	file = &File{fName}
	return
}

func (f *File) Write(b []byte) (int, error) {
	resp, _err := http.Post(serverUrl+"write", "application/json", bytes.NewBuffer(b))
	if _err != nil {
		err := buildError(_err)
		return 0, err
	}
	defer resp.Body.Close()

	type writeResp struct {
		N   int
		Err string
	}
	wr := &writeResp{}
	_err = json.NewDecoder(resp.Body).Decode(wr)
	if _err != nil {
		return 0, _err
	}

	var err error
	if wr.Err != "" {
		err = errors.New(wr.Err)
	}
	return wr.N, err
}

func (f *File) Close() error {
	fmt.Printf("Closing file %v\n", f.Name)

	encoded, _err := encodeFileName(f.Name)
	if _err != nil {
		return _err
	}
	resp, _err := http.Post(serverUrl+"close", "application/json", bytes.NewBuffer(encoded))
	if _err != nil {
		err := buildError(_err)
		return err
	}
	defer resp.Body.Close()

	type closeeResp struct {
		Err string
	}
	cr := &closeeResp{}
	_err = json.NewDecoder(resp.Body).Decode(cr)
	if _err != nil {
		return _err
	}

	var err error
	if cr.Err != "" {
		err = errors.New(cr.Err)
	}

	return err
}

func encodeFileName(fName string) ([]byte, error) {
	fileData := map[string]any{"FileName": fName}
	jsonValue, err := json.Marshal(fileData)
	return jsonValue, err
}

func buildError(e error) error {
	var netErr net.Error
	if ok := errors.As(e, &netErr); ok && netErr.Timeout() {
		return TimeoutErr
	}

	var netOpErr *net.OpError
	if ok := errors.As(e, &netOpErr); ok {
		return NoServerErr
	}

	return e
}
