package lib

//各人脸设备通用的DTO
type Face struct {
	FaceId   string `json:"faceId"`   //照片Id
	PersonId string `json:"personId"` //人的Id，一般是身份证
	FaceData string `json:"image"`    //人脸信息
}

type Person struct {
	*Face
	AttendCard string `json:"attendCard"` //卡号
	Name       string `json:"name"`       //姓名
}
