package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"errors"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"
	"strings"
)

const (
	invalidImage      = "invalid image"
	unsupportedFormat = "does not supported format"
	defaultFileMode   = 0644
)

// B64ToFile save a base64 picture as the file
func B64ToFile(fileName, data string) error {
	idx := strings.Index(data, ";base64,")
	if idx < 0 {
		return errors.New(invalidImage)
	}

	reader := base64.NewDecoder(
		base64.StdEncoding, strings.NewReader(data[idx+8:]))
	buff := bytes.Buffer{}
	_, err := buff.ReadFrom(reader)
	if err != nil {
		return err
	}
	_, format, err := image.DecodeConfig(bytes.NewReader(buff.Bytes()))
	if err != nil {
		return err
	}

	if (strings.Compare(format, "jpeg") < 0) &&
		(strings.Compare(format, "png") < 0) &&
		(strings.Compare(format, "gif") < 0) {
		return errors.New(unsupportedFormat)
	}

	fileName = fileName + "." + format
	err = ioutil.WriteFile(
		fileName, buff.Bytes(), os.FileMode(defaultFileMode))

	return err
}

// FileToB64 get the base64 string from an image file
func FileToB64(file string) (imgB64 string, err error) {
	img, err := os.Open(file)
	if err != nil {
		return imgB64, err
	}
	defer img.Close()

	info, _ := img.Stat()
	size := info.Size()
	buf := make([]byte, size)

	reader := bufio.NewReader(img)
	reader.Read(buf)
	return base64.StdEncoding.EncodeToString(buf), err
}
