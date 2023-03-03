package gormutil

import "gorm.io/gorm"

// Paginate gorm paging tool
func Paginate(pageCurrent, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageCurrent == -1 && pageSize == -1 {
			return db.Offset(-1).Limit(-1)
		}

		if pageCurrent <= 0 {
			pageCurrent = 1
		}
		if pageSize <= 0 {
			pageSize = 10
		}

		offset := (pageCurrent - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
