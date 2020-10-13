package multipart

import (
	"mime/multipart"
	"path/filepath"
	"net/textproto"
	"io"
	"fmt"
)

type Param struct {
	Key    string
	Value  interface{} // if Reader != nil, Value must be the filepath/filename
	Reader io.Reader
}

// create multipart body with/without boundary
//  - body      the writer to store the multipart result
//  - boundary  if empty, a random boundary will be generated
//  - params    fields to be included in the multipart result
func Create(body io.Writer, boundary string, params []Param) (contentType string, err error) {
	if len(params) == 0 {
		return Generate(body, boundary, nil)
	}

	c := make(chan *Param)
	go func() {
		for i, _ := range params {
			c <- &params[i]
		}
		close(c)
	}()
	return Generate(body, boundary, c)
}

// create multipart body with/without boundary
//  - body      the writer to store the multipart result
//  - boundary  if empty, a random boundary will be generated
//  - params    fields to be included in the multipart result
func Generate(body io.Writer, boundary string, params <-chan *Param) (contentType string, err error) {
	w := multipart.NewWriter(body)

	if boundary != "" {
		if err = w.SetBoundary(boundary); err != nil {
			return
		}
	}

	if params == nil {
		goto EXIT
	}

	for param := range params {
		if param == nil {
			continue
		}
		if err = param.createPart(w); err != nil {
			return
		}
	}

EXIT:
	if err = w.Close(); err != nil {
		return
	}
	contentType = w.FormDataContentType()
	return
}

func (param *Param) createPart(w *multipart.Writer) error {
	if param.Reader == nil {
		w.WriteField(param.Key, fmt.Sprintf("%v", param.Value))
		return nil
	}
	return param.createFormFile(w)
}

func (param *Param) createFormFile(w *multipart.Writer) error {
	fileName := filepath.Base(fmt.Sprintf("%v", param.Value))
	part, err := w.CreatePart(textproto.MIMEHeader{
		"Content-Disposition": []string{fmt.Sprintf(`form-data; name="%s"; filename="%s"`, param.Key, fileName)},
		"Content-Type": []string{FileContentType(fileName)},
	})
	if err != nil {
		return err
	}

	_, err = io.Copy(part, param.Reader)
	return err
}
