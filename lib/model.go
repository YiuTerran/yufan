package lib

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/imroc/req"
)

const (
	devicePort  = 8090
	DefaultPass = "755626"
)

type Device struct {
	Pass string
	IP   string
}

type Resp struct {
	Result  int             `json:"result"`
	Success bool            `json:"success"`
	Msg     string          `json:"msg"`
	Data    json.RawMessage `json:"data"`
}

type personResp struct {
	ID        string `json:"id,omitempty"`
	IdCardNum string `json:"idCardNum,omitempty"`
	Name      string `json:"name,omitempty"`
}

type faceResp struct {
	FaceID   string `json:"faceId"`
	Feature  string `json:"feature"`
	Path     string `json:"path"`
	PersonId string `json:"personId"`
}

func checkResp(r *req.Resp, err error) (*Resp, error) {
	if err != nil {
		return nil, err
	}
	var resp Resp
	if err := r.ToJSON(&resp); err != nil {
		return nil, err
	}
	if !resp.Success {
		return nil, fmt.Errorf(resp.Msg)
	}
	return &resp, nil
}

func CombineHttpUrl(host, path string) string {
	if path != "" && path[0] == '/' {
		path = path[1:]
	}
	if strings.HasPrefix(host, "http") {
		return fmt.Sprintf("%s/%s", host, path)
	} else {
		return fmt.Sprintf("http://%s/%s", host, path)
	}
}

func (d *Device) getURL(path string) string {
	return CombineHttpUrl(fmt.Sprintf("%s:%d", d.IP, devicePort), path)
}
