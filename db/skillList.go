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

// write a new skill_list object into the database
func AddSkill(userId int, skillName string) error {
	skill := types.SkillList{Name: skillName, Owner: userId}
	_, err := db.Insert(&skill)
	return err
}

// delete the skill_list object from the database
func RemoveSkill(userId, skillId int) error {
	skill := types.SkillList{Id: skillId, Owner: userId}
	var err error

	if _, err = db.Delete(&skill); err == nil {
		log.Print("Skill ", skillId, " of user ", userId, " is deleted.")
	}
	return err
}
