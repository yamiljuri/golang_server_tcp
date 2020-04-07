package frame

import "time"

const (
	Input  = "input"
	Output = "output"
)

type Frame struct {
	id       int       `bson:"_id"`
	DeviceId int       `bson:"deviceid"`
	Type     string    `bson:"type"`
	Frame    []byte    `bson:"frame"`
	Date     time.Time `bson:"date"`
}
