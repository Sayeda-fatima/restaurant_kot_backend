package common

import "gorm.io/gorm"

func Paginate(page int) func(db *gorm.DB) *gorm.DB {

	return func(db *gorm.DB) *gorm.DB{
		if page <= 0 {
			page = 1
		}
		
		// defining a fixed pagesize of 20
		pageSize := 20
	
		offset := (page - 1) * pageSize
	
		return db.Offset(offset).Limit(pageSize)
	}
}

// for raw queries
func ApplyPagination(page int) (limit, offset int) {
	if page <= 0 {
		page = 1
	}
	pageSize := 20
	limit = pageSize
	offset = (page - 1) * pageSize
	return limit, offset
}
