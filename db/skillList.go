package db

import (
	"github.com/izayacity/LinkedIn/types"
	"github.com/astaxie/beego/orm"
	"log"
)

func init() {
	// register model
	orm.RegisterModel(new(types.SkillList))
}

func AddSkill(userId int, skillName string) error {
	skill := types.SkillList{Name: skillName, Owner: userId}
	_, err := db.Insert(&skill)
	return err
}

func RemoveSkill(userId, skillId int) error {
	skill := types.SkillList{Id: skillId, Owner: userId}
	var err error

	if _, err = db.Delete(&skill); err == nil {
		log.Print("Skill ", skillId, " of user ", userId, " is deleted.")
	}
	return err
}
