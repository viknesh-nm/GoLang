package models

type Article struct {
	Id     int    `form:"-"`
	Name   string `form:"name,text,name:" valid:"MinSize(5);MaxSize(20)"`
	Contact string `form:"contact,text,contact:"`
	Email    string `form:"email,text,email:"`
}
// TableName() - define the tablename which was used in the database
func (a *Article) TableName() string {
	return "formss"
}
