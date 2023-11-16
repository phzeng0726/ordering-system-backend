package domain

type Category struct {
	Id               int              `gorm:"column:id;not null;primaryKey;autoIncrement;" json:"id"`
	Identifier       string           `gorm:"column:identifier;" json:"identifier"`
	IsDefault        *bool            `gorm:"column:is_default;" json:"isDefault"`
	Title            string           `gorm:"-" json:"title"`
	CategoryLanguage CategoryLanguage `gorm:"foreignKey:Id;references:category_id;" json:"-"`
}

// 下面兩個本來用Preload，但改成raw scan，暫時沒用到
type CategoryUserMapping struct {
	Id         int      `gorm:"column:id;not null;primaryKey;autoIncrement;" json:"id"`
	CategoryId int      `gorm:"column:category_id;" json:"-"`
	UserId     string   `gorm:"column:user_id;" json:"userId"`
	OrderBy    int      `gorm:"column:order_by;" json:"orderBy"`
	Category   Category `gorm:"foreignKey:CategoryId;references:id;" json:"-"`
}

type CategoryLanguage struct {
	Id         int    `gorm:"column:id;not null;primaryKey;autoIncrement;" json:"-"`
	CategoryId int    `gorm:"column:category_id;" json:"-"`
	LanguageId *int   `gorm:"column:language_id;" json:"languageId"`
	Title      string `gorm:"column:title;" json:"title"`
}

func (CategoryUserMapping) TableName() string {
	return "category_user_mapping"
}

func (CategoryLanguage) TableName() string {
	return "category_language"
}
