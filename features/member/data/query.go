package data

import (
	"errors"
	"kstyleAPI/features/member"
	"log"

	"gorm.io/gorm"
)

type memberQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) member.MemberData {
	return &memberQuery{
		db: db,
	}
}

func (mq *memberQuery) checkDuplicate(username string) bool {
	m := Member{}
	mq.db.Where("username = ?", username).First(&m)

	return m.ID != 0
}

func (mq *memberQuery) checkMember(id uint) bool {
	m := Member{}
	mq.db.Where("id = ?", id).First(&m)

	return m.ID != 0
}

func (mq *memberQuery) Insert(newMember member.Core) (member.Core, error) {
	if mq.checkDuplicate(newMember.Username) {
		log.Println("duplicated username")
		return member.Core{}, errors.New("duplicated")
	}

	cnv := CoreToData(newMember)
	if err := mq.db.Create(&cnv).Error; err != nil {
		log.Println("error insert new member: ", err)
		return member.Core{}, err
	}

	newMember.ID = cnv.ID

	return newMember, nil
}

func (mq *memberQuery) Update(updMember member.Core) (member.Core, error) {
	if !mq.checkMember(updMember.ID) {
		log.Println("member not found")
		return member.Core{}, errors.New("member not found")
	}

	cnv := CoreToData(updMember)
	if err := mq.db.Updates(&cnv).Error; err != nil {
		log.Println("error update member: ", err)
		return member.Core{}, err
	}

	updMember.ID = cnv.ID

	return updMember, nil
}

func (mq *memberQuery) Delete(id uint) error {
	if !mq.checkMember(id) {
		log.Println("member not found")
		return errors.New("member not found")
	}
	mem := Member{}
	err := mq.db.Model(&mem).Delete("id_member = ?", id).Error
	if err != nil {
		log.Println("error delete member: ", err)
		return err
	}
	return nil
}

func (mq *memberQuery) GetMembers() ([]member.Core, error) {
	members := []member.Core{}
	err := mq.db.Raw("SELECT * FROM members WHERE deleted_at IS NULL").Scan(&members).Error
	if err != nil {
		log.Println("error get members: ", err)
	}

	return members, nil
}
