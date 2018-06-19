/*
	__tablename__ = 'budaya'
	id = Column(Integer, primary_key = True)
	name = Column(String(80), nullable=False)
	description = Column(String(250))
	image_url = Column(String(250))
	google_search_term = Column(String(250))


	province_id = Column(Integer, ForeignKey('province.id'))
	province = relationship(Province)

	category = Column(Integer, ForeignKey('categories.id'))
	categories = relationship(Categories)

*/

package models

type Budaya struct {
	Id               int    `db:"id" json:"id"`
	Name             string `db:"name" json:"name"`
	Description      string `db:"description" json:"description"`
	ImageUrl         string `db:"image_url" json:"image_url"`
	GoogleSearchTerm string `db:"google_search_term" json:"google_search_term"`

	Province   Province   `has_one:"province" fk_id:"id"`
	Categories Categories `has_one:"categories" fk_id:"id"`
}

func (b Budaya) TableName() string {
	return "budaya"
}
