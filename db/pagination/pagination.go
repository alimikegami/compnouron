package pagination

import "gorm.io/gorm"

func Paginate(limit int, offset int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(limit)
	}
}
