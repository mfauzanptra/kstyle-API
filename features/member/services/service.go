package services

import (
	"errors"
	"kstyleAPI/features/member"
	"strings"
)

type memberUseCase struct {
	qry member.MemberData
}

func New(md member.MemberData) member.MemberService {
	return &memberUseCase{
		qry: md,
	}
}

func (muc *memberUseCase) Insert(newMember member.Core) (member.Core, error) {
	res, err := muc.qry.Insert(newMember)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "duplicated") {
			msg = "duplicated username"
		} else {
			msg = "server error"
		}
		return member.Core{}, errors.New(msg)
	}

	return res, nil
}
func (muc *memberUseCase) Update(IdMember uint, updMember member.Core) (member.Core, error) {
	updMember.ID = IdMember
	res, err := muc.qry.Update(updMember)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "member not found"
		} else {
			msg = "server error"
		}
		return member.Core{}, errors.New(msg)
	}

	return res, nil
}
func (muc *memberUseCase) Delete(id uint) error {
	if err := muc.qry.Delete(id); err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "member not found"
		} else {
			msg = "server error"
		}
		return errors.New(msg)
	}

	return nil
}
func (muc *memberUseCase) GetMembers() ([]member.Core, error) {
	res, err := muc.qry.GetMembers()
	if err != nil {
		return []member.Core{}, err
	}

	return res, nil
}
