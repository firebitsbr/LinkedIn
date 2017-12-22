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

// return user name, skill name, count, and endorsers
func GetSkills(userId string) types.Skills {
	skills := types.Skills{}
	qb, _ := orm.NewQueryBuilder("mysql")

	sql := qb.Select("skill_list.id", "skill_list.name", "skill_list.count", "endorse.sender").
		From("skill_list").InnerJoin("endorse").
			On("skill_list.owner="+userId+" AND endorse.owner=" + userId).
				GroupBy("skill_list.id", "endorse.sender").String()
	_, err := db.Raw(sql).QueryRows(&skills)
	checkErr(err)

	for _, skill := range skills {
		log.Print(skill)
	}
	return skills
}
//SELECT skill_list.id, skill_list.name, skill_list.count, endorse.sender
//FROM skill_list INNER JOIN endorse
//ON skill_list.owner=3 AND endorse.owner=3
//GROUP BY skill_list.id, endorse.sender
