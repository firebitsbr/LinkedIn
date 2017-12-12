package db

import (
	"github.com/izayacity/LinkedIn/types"
	"github.com/astaxie/beego/orm"
	"log"
)

func init() {
	// register model
	orm.RegisterModel(new(types.Endorse))
}

func Endorse(sid, sender, owner int) error {
	endorse := types.Endorse{Sid: sid, Owner: owner, Sender: sender}
	_, err := db.Insert(&endorse)
	return err
}

// return true if it is endorsed; return false if it hasn't been endorsed
func IsEndorsed(sid, sender, owner int ) bool {
	endorse := types.Endorse{Id: sid, Owner: owner, Sender: sender}
	err := db.QueryTable("endorse").
		Filter("sid", sid).Filter("sender", sender).Filter("owner", owner).One(&endorse)

	if err == orm.ErrNoRows {
		log.Print("Skill ", sid, " has not been endorsed.")
		return false
	}
	log.Print("Skill ", sid, " has been endorsed.")
	return true
}

func UpdateCount(sid int) error {
	skill := getSkill(sid)
	var err error
	if &skill != nil {
		skill.Count++
		_, err = db.Update(&skill)
		checkErr(err)
	}
	return err
}

func getSkill(sid int) types.SkillList {
	skill := types.SkillList{}
	err := db.QueryTable("skill_list").Filter("id", sid).One(&skill)
	checkErr(err)
	return skill
}
