package oss

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

type common struct {
	TotalCount int    `json:"totalCount"`
	Success    bool   `json:"success"`
	Error      string `json:"error"`
	SingleInfo string `json:"singleInfo"`
	Data       []File `json:"data"`
}

type File struct {
	Error      string `json:"error"`
	Data       string `json:"data"`
	FileName   string `json:"fileName"`
	FileURL    string `json:"fileUrl"`
	FullURL    string `json:"fullUrl"`
	FileSize   string `json:"fileSize"`
	Width      string `json:"width"`
	Height     string `json:"height"`
	Lat        string `json:"lat"`
	Lon        string `json:"lon"`
	ThumbURL   string `json:"thumbUrl"`
	MetalURL   string `json:"metalUrl"`
	WoodURL    string `json:"woodUrl"`
	PoiID      string `json:"poiId"`
	PoiType    string `json:"poiType"`
	DistrictID string `json:"districtId"`
	Ttt        string `json:"ttt"`
}

const (
	ctripFileUploadPath = "https://youphotogalleryuploader.ctrip.com/photo/uploadImageAdaptor"
)

var InterfaceError = errors.New("the upload API returns an error")

func Upload(filepath string) (error, *File) {
	file, err := os.Open(filepath)
	if err != nil {
		return err, nil
	}
	return UploadFromReader(file)
}

func UploadFromReader(reader io.Reader) (error, *File) {
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("type", "journals")
	_ = writer.WriteField("typeID", "13")

	fileField, _ := writer.CreateFormFile("userfile", "O1CN01ev1lwQ1sc2FDntCSy_!!578305786.jpg")

	_, _ = io.Copy(fileField, reader)

	_ = writer.Close()

	client := &http.Client{}
	req, err := http.NewRequest("POST", ctripFileUploadPath, payload)

	if err != nil {
		return err, nil
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		return err, nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err, nil
	}

	err, com := formatToStruct(string(body))
	if err != nil {
		return err, nil
	}
	if (!com.Success) || len(com.Data) < 1 {
		return InterfaceError, nil
	}
	return nil, &com.Data[0]
}

func formatToStruct(text string) (error, *common) {
	result := strings.SplitN(text, "</script>", 2)
	if len(result) != 2 {
		return InterfaceError, nil
	}

	c := &common{}
	err := json.Unmarshal([]byte(result[1]), c)
	if err != nil {
		return InterfaceError, nil
	}
	return nil, c
}
