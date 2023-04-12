package data

import (
	"kstyleAPI/features/member"
	"time"

	"gorm.io/gorm"
)

type Member struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
	Username  string
	Gender    string
	Skintype  string
	Skincolor string
}

func ToCore(data Member) member.Core {
	return member.Core{
		ID:        data.ID,
		Username:  data.Username,
		Gender:    data.Gender,
		Skintype:  data.Skintype,
		Skincolor: data.Skincolor,
	}
}

func CoreToData(data member.Core) Member {
	return Member{
		ID:        data.ID,
		Username:  data.Username,
		Gender:    data.Gender,
		Skintype:  data.Skintype,
		Skincolor: data.Skincolor,
	}
}
