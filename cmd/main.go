package main

import (
	"bgo-homeworks-10-2/pkg/qr"
	"context"
	"errors"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {

	timeout, data, filename, err := getEnvVariable()

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	ctx, _ := context.WithTimeout(context.Background(), timeout)
	svc := qr.NewService()
	reader, filetype, err := svc.Encode(ctx, data)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	err = qr.SaveToFile(filetype, reader, filename)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}

var (
	ErrEmptyEnvParamTimeout = errors.New("empty env param 'TIMEOUT'")
	ErrEmptyEnvParamData    = errors.New("empty env param 'DATA'")
	DefaultFileName = "qrcode"
)

func getEnvVariable() (timeout time.Duration, data string, filename string, err error) {
	timeout, err = getEnvTimeoutInMillisecond()
	if err != nil {
		return 0, "", "", err
	}

	data, err = getEnvDataStringToConvert()
	if err != nil {
		return 0, "", "", err
	}

	filename = getEnvSaveFileName()

	return timeout, data, filename, nil
}

func getEnvTimeoutInMillisecond() (time.Duration, error) {
	env, ok := os.LookupEnv("TIMEOUT")
	if ok {
		timeout, err := strconv.Atoi(env)
		if err != nil {
			return 0, err
		}
		return time.Duration(timeout) * time.Millisecond, nil
	}
	return 0, ErrEmptyEnvParamTimeout
}

func getEnvDataStringToConvert() (string, error) {
	env, ok := os.LookupEnv("DATA")
	if ok {
		return env, nil
	}
	return "", ErrEmptyEnvParamData
}

func getEnvSaveFileName() string {
	env, ok := os.LookupEnv("FILE_NAME")
	if ok {
		return env
	}
	return DefaultFileName
}


