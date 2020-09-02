package lib

import (
	"github.com/imroc/req"
)

func (d *Device) GetSN() string {
	if resp, err := checkResp(req.Get(d.getURL("getDeviceKey"))); err != nil {
		return ""
	} else {
		return string(resp.Data)
	}
}

func (d *Device) SetPhotoCallback(url string) error {
	_, err := checkResp(req.Post(d.getURL("setImgRegCallBack"), req.Param{
		"pass": d.Pass,
		"url":  url,
	}))
	return err
}

func (d *Device) SetIdentifyCallBack(url string) error {
	_, err := checkResp(req.Post(d.getURL("setIdentifyCallBack"), req.Param{
		"pass":        d.Pass,
		"callbackUrl": url,
	}))
	return err
}

func (d *Device) SetHeartBeatURL(url string) error {
	_, err := checkResp(req.Post(d.getURL("setDeviceHeartBeat"), req.Param{
		"pass": d.Pass,
		"url":  url,
	}))
	return err
}
