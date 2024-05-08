package wrapper

import "gorm.io/gorm"

type UpdateWrapper struct {
	QueryWrapper
	sets map[string]interface{}
}

func Update(capacity ...int) *UpdateWrapper {
	var conditions []condition
	if len(capacity) > 0 {
		conditions = make([]condition, 0, capacity[0])
	} else {
		conditions = make([]condition, 0)
	}
	return &UpdateWrapper{
		QueryWrapper: QueryWrapper{conditions: conditions},
		sets:         make(map[string]interface{}),
	}
}

func (uw *UpdateWrapper) SetMap(sets map[string]interface{}) {
	uw.sets = sets
}

func (uw *UpdateWrapper) Set(field string, arg interface{}) {
	uw.sets[field] = arg
}

func (uw *UpdateWrapper) FillSet(db *gorm.DB) *gorm.DB {
	return db.Updates(uw.sets)
}
