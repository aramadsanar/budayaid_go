/*__tablename__ = 'categories'
id = Column(Integer, primary_key = True)
name = Column(String(250), nullable = False)
friendly_name = Column(String(250), nullable = False)*/

package models

type Categories struct {
	Id           int    `db:"id" json:""id`
	Name         string `db:"name" json:"name"`
	FriendlyName string `db:"friendly_name" json:"friendly_name"`
}

func (c Categories) TableName() string {
	return "categories"
}
