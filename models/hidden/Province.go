package models

type Province struct {
	/*
		__tablename__ = 'province'
		id = Column(Integer, primary_key = True)
		name = Column(String(250), nullable = False)
		friendly_name = Column(String(250), nullable = False)
		island = Column(String(50))*/
	Id           int    `db:"id" json:"id"`
	Name         string `db:"name" json:"name"`
	FriendlyName string `db:"friendly_name" json:"firendly_name"`
	Island       string `db:"island" json:"island"`
}

func (p Province) TableName() string {
	return "province"
}
