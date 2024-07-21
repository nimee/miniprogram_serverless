package album

import (
	"time"
)

type Album struct {
	ID            uint      `gorm:"AUTO_INCREMENT;comment:主键" json:"id"`
	Theme         string    `gorm:"column:theme;type:varchar(255)" json:"theme"` //index;
	Count         int16     `gorm:"column:count;type:int(11)" json:"count"`
	Photographer  string    `gorm:"column:photographer;type:varchar(255)" json:"photographer"`
	Mua           string    `gorm:"column:mua;type:varchar(255)" json:"mua"`
	T             time.Time `gorm:"type:timestamp;column:t" json:"t"`
	Location      string    `gorm:"column:location;type:varchar(255)" json:"location"`
	GroupName     string    `gorm:"column:groupName;type:varchar(255)" json:"groupName"`
	CoverId       string    `gorm:"column:coverId;type:varchar(255)" json:"coverId"`
	StorageType   string    `gorm:"column:storageType;type:varchar(255)" json:"storageType"`
	StorageParams []byte    `gorm:"column:storageParams;type:varchar(1024)" json:"storageParams"`
	Enable        bool      `gorm:"column:enable;type:TINYINT(1);default:1" json:"enable"`
	Note          string    `gorm:"column:note;type:varchar(1024)" json:"note"`
}

type Photo struct {
	AlbumId uint `gorm:"column:albumId;type:int(11)" json:"albumId"`
	//Theme   string `gorm:"column:theme;type:varchar(255)" json:"theme"`
	PhotoId string `gorm:"primaryKey;index;column:photoId;type:varchar(255)" json:"photoId"`
	Enable  bool   `gorm:"column:enable;type:TINYINT(1);default:1" json:"enable"`
	Params  []byte `gorm:"column:Params;type:varchar(1024)" json:"Params"`
}
