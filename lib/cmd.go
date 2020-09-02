package lib

import (
	"encoding/json"
	"strings"


	"github.com/YiuTerran/leaf/util/image"
	"github.com/imroc/req"
)

func (d *Device) UpsertPerson(dto *Person, isInsert bool) (string, error) {
	var url string
	person := req.Param{
		"id":        dto.PersonId,
		"name":      dto.Name,
		"idcardNum": dto.AttendCard,
	}
	personStr, _ := json.Marshal(person)
	param := req.Param{
		"person": string(personStr),
		"pass":   d.Pass,
	}
	if dto.PersonId == "" || isInsert {
		param["id"] = dto.PersonId
		url = d.getURL("person/create")
	} else {
		url = d.getURL("person/update")
	}
	resp, err := checkResp(req.Post(url, param))
	if err != nil {
		return "", err
	}
	if dto.PersonId != "" {
		return "", nil
	}
	var info personResp
	if err = json.Unmarshal(resp.Data, &info); err != nil {
		return "", err
	}
	return info.ID, nil
}

func (d *Device) DeletePerson(ids []string) ([]string, error) {
	var target string
	if ids == nil {
		//删除所有人
		target = "-1"
	} else {
		target = strings.Join(ids, ",")
	}
	if target == "" {
		return nil, nil
	}
	resp, err := checkResp(req.Post(d.getURL("person/delete"), req.Param{
		"id":   target,
		"pass": d.Pass,
	}))
	if err != nil {
		return nil, err
	}
	info := make(map[string]string, 2)
	if err = json.Unmarshal(resp.Data, &info); err != nil {
		return nil, err
	}
	return strings.Split(info["effective"], ","), nil
}

func (d *Device) Clear() error {
	_, err := d.DeletePerson(nil)
	return err
}

func (d *Device) ListPerson() ([]*Person, error) {
	resp, err := checkResp(req.Get(d.getURL("person/find"), req.Param{
		"id":   "-1",
		"pass": d.Pass,
	}))
	if err != nil {
		return nil, err
	}
	var persons []*personResp
	if err = json.Unmarshal(resp.Data, &persons); err != nil {
		return nil, err
	}
	result := make([]*Person, 0, len(persons))
	for _, p := range persons {
		result = append(result, &Person{
			Face: &Face{
				PersonId: p.ID,
			},
			Name:       p.Name,
			AttendCard: p.IdCardNum,
		})
	}
	return result, nil
}

//img为图片的base64编码
func (d *Device) UpsertPhoto(dto *Face) (string, error) {
	var url string
	param := req.Param{
		"personId":  dto.PersonId,
		"imgBase64": image.RemoveImgBase64Header(dto.FaceData),
		"pass":      d.Pass,
		"faceId":    dto.FaceId,
	}
	if dto.FaceId == "" {
		url = d.getURL("face/create")
	} else {
		url = d.getURL("face/update")
	}
	resp, err := checkResp(req.Post(url, param))
	if err != nil {
		return "", err
	}
	if dto.FaceId != "" {
		return "", nil
	}
	return string(resp.Data), nil
}

func (d *Device) DeletePhoto(peopleId string, faceId string) error {
	var err error
	if peopleId != "" {
		_, err = checkResp(req.Post(d.getURL("face/deletePerson"), req.Param{
			"personId": peopleId,
			"pass":     d.Pass,
		}))
	} else if faceId != "" {
		_, err = checkResp(req.Post(d.getURL("face/delete"), req.Param{
			"faceId": faceId,
			"pass":   d.Pass,
		}))
	}
	return err
}

func (d *Device) ListPhoto(personId string) ([]*Face, error) {
	resp, err := checkResp(req.Post(d.getURL("face/find"), req.Param{
		"personId": personId,
		"pass":     d.Pass,
	}))
	if err != nil {
		return nil, err
	}
	var faces []*faceResp
	if err = json.Unmarshal(resp.Data, &faces); err != nil {
		return nil, err
	}
	result := make([]*Face, 0, len(faces))
	for _, f := range faces {
		result = append(result, &Face{
			FaceId:   f.FaceID,
			PersonId: f.PersonId,
			FaceData: f.Feature,
		})
	}
	return result, nil
}

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
