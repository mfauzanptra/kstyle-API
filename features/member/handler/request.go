package handler

import "kstyleAPI/features/member"

type InsertMemberReq struct {
	Username  string `json:"username" form:"username"`
	Gender    string `json:"gender" form:"gender"`
	Skintype  string `json:"skintype" form:"skintype"`
	Skincolor string `json:"skincolor" form:"skincolor"`
}

func ToCore(data InsertMemberReq) *member.Core {
	return &member.Core{
		Username:  data.Username,
		Gender:    data.Gender,
		Skintype:  data.Skintype,
		Skincolor: data.Skincolor,
	}
}
