package qr

import (
	"context"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const urlService = "https://api.qrserver.com/v1/create-qr-code/"

type Service struct {
	Client *http.Client
}

func NewService() *Service {
	return &Service{&http.Client{}}
}

func (s Service) Encode(ctx context.Context, data string) (io.Reader, string, error)  {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, urlService, nil)
	if err != nil {
		return nil, "", err
	}

	q:= url.Values{}
	q.Add("data", data)
	req.URL.RawQuery = q.Encode()

	response, err := s.Client.Do(req)
	if err != nil {
		return nil, "", err
	}

	fileType := response.Header.Get("Content-Type")

	return response.Body, fileType, nil
}

func SaveToFile(filetype string, reader io.Reader, filename string) error {
	if filetype == "image/png" {
		file, err := ioutil.ReadAll(reader)
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(filename+".png", file, os.ModePerm)
		if err != nil {
			return err
		}
		return nil
	}

	return errors.New("incorrect file type'")
}