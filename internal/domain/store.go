package domain

import dt "gorm.io/datatypes"

// gorm: db的column name，不需要在db做操作的話就用-
// json: 最後吐出去的json key name，不需要在json顯示的話就用-
// autoIncrement;
// not null;
// primaryKey;
type Store struct {
	Id                string             `gorm:"column:id;not null;primaryKey;" json:"id"`
	UserId            string             `gorm:"column:user_id;" json:"userId,omitempty"`
	Name              string             `gorm:"column:name;" json:"name"`
	Description       string             `gorm:"column:description;" json:"description"`
	Email             string             `gorm:"-" json:"-"` // 從User那來
	Phone             string             `gorm:"column:phone;" json:"phone"`
	Address           string             `gorm:"column:address;" json:"address"`
	Timezone          string             `gorm:"column:timezone;" json:"timezone"`
	IsBreak           bool               `gorm:"column:is_break;" json:"isBreak"`
	StoreOpeningHours []StoreOpeningHour `gorm:"foreignKey:StoreId;references:id;" json:"storeOpeningHours"`
}

// NOTE:
// StoreOpeningHour裡面的StoreId為fKey，目前是用store，預先加載裡面的StoreOpeningHour。
// if err := r.db.Preload("StoreOpeningHours").Find(&stores).Error; err != nil {
// 	return nil, err
// }
type StoreOpeningHour struct {
	Id        int     `gorm:"column:id;not null;primaryKey;" json:"-"`
	StoreId   string  `gorm:"column:store_id;" json:"-"`
	DayOfWeek int     `gorm:"column:day_of_week;" json:"dayOfWeek"`
	OpenTime  dt.Time `gorm:"column:open_time" json:"openTime"`
	CloseTime dt.Time `gorm:"column:close_time" json:"closeTime"`
}
