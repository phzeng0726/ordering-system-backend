package domain

type Category struct {
	Id         int    `gorm:"column:id;not null;primaryKey;autoIncrement;" json:"id"`
	Identifier string `gorm:"column:identifier;" json:"identifier"`
}

type CategoryUserMapping struct {
	Id         int      `gorm:"column:id;not null;primaryKey;autoIncrement;" json:"id"`
	CategoryId int      `gorm:"column:category_id;" json:"-"`
	UserId     string   `gorm:"column:user_id;" json:"userId"`
	OrderBy    int      `gorm:"column:order_by;" json:"orderBy"`
	Category   Category `gorm:"foreignKey:CategoryId;references:id;" json:"-"`
	User       User     `gorm:"foreignKey:UserId;references:id;" json:"-"`
}

type CategoryLanguage struct {
	Id         int      `gorm:"column:id;not null;primaryKey;autoIncrement;" json:"id"`
	CategoryId int      `gorm:"column:category_id;" json:"categoryId"`
	LanguageId int      `gorm:"column:language_id;" json:"languageId"`
	Title      string   `gorm:"column:title;" json:"title"`
	Category   Category `gorm:"foreignKey:CategoryId;references:id;" json:"-"`
}

func (CategoryUserMapping) TableName() string {
	return "category_user_mapping"
}
