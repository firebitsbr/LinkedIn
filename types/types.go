package types

type User struct {
	Id           		int           	`json:"id"`
	Username	string		`json:"username"`
	Password		string		`json:"password"`
	Email			string		`json:"email"`
	Status			string		`json:"status"`
}

type Users []User

type Endorse struct {
	Id           					int           	`json:"id"`
	Sid							int			`json:"sid"`
	Owner						int			`json:"owner"`
	Sender						int			`json:"sender"`
	LastModified			string		`json:"last_modified"`
}

type Endorses []Endorse

type SkillList struct {
	Id           				int           	`json:"id"`
	Name					string		`json:"name"`
	Count					int			`json:"count"`
	Owner					int			`json:"owner"`
	CreatedTime		string		`json:"created_time"`
}

type SkillLists []SkillList

type SkillTag struct {
	Id           	int           	`json:"id"`
	Name		string		`json:"name"`
	Total			int			`json:"total"`
	Expert		int			`json:"expert"`
}

type Skill struct {
	Id           		int
	Name			string
	Count			int
	Sender			int
}

type Skills []Skill
