// Package httpclient provides convenient functions to build request and
// other things about http client. This does NOT implement all features
// but pack some complicated procedures in http.
package httpclient

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
	"os"
	"strings"
)

func BuildFromFiles(uri string, queryParams, postParams map[string]string, paths map[string][]string) (*http.Request, error) {
	return BuildPost(uri, queryParams, postParams, paths, nil)
}

func BuildFromIOReaders(uri string, queryParams, postParams map[string]string, readers map[string]map[string]io.Reader) (*http.Request, error) {
	return BuildPost(uri, queryParams, postParams, nil, readers)
}

// BuildRequest builds request based on uri, query and post params, files, io readers, Please choose the functions above instead if you have no idea how to invoke this.
// paths "field name": ["filepath 1", "filepath 2"]
// ioReaders "field name": ["reader 1", "reader 2"]
func BuildPost(uri string, queryParams, postParams map[string]string, paths map[string][]string, ioReaders map[string]map[string]io.Reader) (req *http.Request, err error) {
	switch {
	case paths != nil || ioReaders != nil:
		body, contentType, err := buildMultiplerPart(postParams, paths, ioReaders)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest(http.MethodPost, uri, body)
		if err != nil {
			return nil, err
		}
		req.Header.Add("Content-Type", contentType)
	case postParams != nil:
		form := url.Values{}
		for k, v := range postParams {
			form.Add(k, v)
		}
		req, err = http.NewRequest(http.MethodPost, uri, strings.NewReader(form.Encode()))
		if err != nil {
			return nil, err
		}
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	default:
		req, err = http.NewRequest(http.MethodPost, uri, nil)
		if err != nil {
			return nil, err
		}
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}
	q := req.URL.Query()
	for k, v := range queryParams {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	return req, nil
}

func BuildJSONBody(uri string, queryParams map[string]string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(http.MethodPost, uri, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	q := req.URL.Query()
	for k, v := range queryParams {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	return req, nil
}

func buildMultiplerPart(params map[string]string, files map[string][]string, ioReaders map[string]map[string]io.Reader) (*bytes.Buffer, string, error) {
	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)

	for name, ps := range files {
		for _, p := range ps {
			if err := writeFileToWriter(writer, name, p); err != nil {
				return nil, "", err
			}
		}
	}

	for fieldName, readers := range ioReaders {
		for filename, r := range readers {
			if err := writeReaderToWriter(writer, fieldName, filename, r); err != nil {
				return nil, "", err
			}
		}
	}
	for key, val := range params {
		if err := writer.WriteField(key, val); err != nil {
			return nil, "", err
		}
	}
	if err := writer.Close(); err != nil {
		return nil, "", err
	}
	return body, writer.FormDataContentType(), nil
}

func writeFileToWriter(writer *multipart.Writer, name, fpath string) error {
	f, err := os.Open(fpath)
	if err != nil {
		return err
	}
	defer f.Close()
	fs, err := f.Stat()
	if err != nil {
		return err
	}

	header := textproto.MIMEHeader{}
	buffer := make([]byte, 512)
	_, err = f.Read(buffer)
	if err != nil {
		return err
	}
	f.Seek(0, 0)
	header.Add("Content-Type", http.DetectContentType(buffer))
	header.Add("Content-Disposition", fmt.Sprintf(`form-data; filename="%s"; name="%s"`, fs.Name(), name))

	part, err := writer.CreatePart(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(part, f)
	if err != nil {
		return err
	}
	return nil
}

func writeReaderToWriter(writer *multipart.Writer, filedName, filename string, reader io.Reader) error {
	header := textproto.MIMEHeader{}
	copy, ct, err := retriveContentType(reader)
	if err != nil {
		return err
	}
	header.Add("Content-Type", ct)
	header.Add("Content-Disposition", fmt.Sprintf(`form-data; filename="%s"; name="%s"`, filename, filedName))
	part, err := writer.CreatePart(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(part, copy)
	if err != nil {
		return err
	}
	return nil
}

func retriveContentType(r io.Reader) (io.Reader, string, error) {
	first512 := make([]byte, 512)
	_, err := r.Read(first512)
	if err != nil {
		return nil, "", err
	}
	buffer := bytes.NewBuffer(first512)
	ct := http.DetectContentType(first512)
	var copy bytes.Buffer
	_, err = io.Copy(&copy, buffer)
	if err != nil {
		return nil, "", err
	}
	_, err = io.Copy(&copy, r)
	if err != nil {
		return nil, "", err
	}
	return &copy, ct, nil
}
