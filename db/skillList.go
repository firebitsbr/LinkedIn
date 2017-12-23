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
	// for app and result
	skills := types.Skills{}
	// for db model query
	skillRows := types.SkillRows{}

	// fetch rows from db into skillRows
	qb, _ := orm.NewQueryBuilder("mysql")
	sql := qb.Select("skill_list.id", "skill_list.name", "skill_list.count", "endorse.sender").
		From("skill_list").InnerJoin("endorse").
			On("skill_list.owner="+userId).And("endorse.owner=" + userId).
				GroupBy("skill_list.id", "endorse.sender").String()
	_, err := db.Raw(sql).QueryRows(&skillRows)
	checkErr(err)

	// convert SkillRow data type to Skill data type that store senders as array
	if len(skillRows) < 1 {
		log.Print("No result of skill rows")
		return skills
	}
	skillRow := skillRows[0]
	prevId := skillRow.Id
	skill := types.Skill{ Id: skillRow.Id, Name: skillRow.Name, Count: skillRow.Count, Sender: []int{} }

	for _, skillRow = range skillRows {
		if skillRow.Id != prevId {
			skills = append(skills, skill)
			skill = types.Skill{ Id: skillRow.Id, Name: skillRow.Name, Count: skillRow.Count, Sender: []int{} }
			prevId = skillRow.Id
		}
		skill.Sender = append(skill.Sender, skillRow.Sender)
	}
	skills = append(skills, skill)
	return skills
}
//SELECT skill_list.id, skill_list.name, skill_list.count, endorse.sender
//FROM skill_list INNER JOIN endorse
//ON skill_list.owner=3 AND endorse.owner=3
//GROUP BY skill_list.id, endorse.sender
