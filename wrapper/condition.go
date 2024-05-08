package wrapper

import (
	"fmt"
	"gorm.io/gorm"
)

const (
	where       = iota // 封装 db.Where()
	eq                 // =
	nEq                // !=
	gt                 // >
	ge                 // >=
	lt                 // <
	le                 // <=
	between            // between v1 and v2
	nBetween           // not between v1 and v2
	like               // like "%v%"
	nLike              // not like "%v%"
	isNull             // is null
	isNNull            // is not null
	in                 // in (v1, v2 ...)
	nIn                // not in (v1, v2 ...)
	groupBy            // group by field1, field2
	orderByAsc         // order by field1 asc
	orderByDesc        // order by field1 desc
	having             // 封装 db.Having()
	or                 // 封装 db.Or()
	distinct
)

type condition struct {
	keyType int
	query   interface{}
	args    []interface{}
}

func (c *condition) fill(db *gorm.DB) {
	switch c.keyType {
	case where:
		c.fillWhere(db)
	case eq:
		c.fillEq(db)
	case nEq:
		c.fillNEq(db)
	case gt:
		c.fillGt(db)
	case ge:
		c.fillGe(db)
	case lt:
		c.fillLt(db)
	case le:
		c.fillLe(db)
	case between:
		c.fillBetween(db)
	case nBetween:
		c.fillNBetween(db)
	case like:
		c.fillLike(db)
	case nLike:
		c.fillNLike(db)
	case isNull:
		c.fillIsNull(db)
	case isNNull:
		c.fillIsNNull(db)
	case in:
		c.fillIn(db)
	case nIn:
		c.fillNIn(db)
	case groupBy:
		c.fillGroupBy(db)
	case orderByAsc:
		c.fillOrderByAsc(db)
	case orderByDesc:
		c.fillOrderByDesc(db)
	case having:
		c.fillHaving(db)
	case or:
		c.fillOr(db)
	case distinct:
		c.fillDistinct(db)
	default:
	}

}

func (c *condition) fillWhere(db *gorm.DB) {
	db.Where(c.query, c.args...)
}

func (c *condition) fillEq(db *gorm.DB) {
	db.Where("? = ?", c.query, c.args[0])
}

func (c *condition) fillNEq(db *gorm.DB) {
	db.Where("? != ?", c.query, c.args[0])
}

func (c *condition) fillGt(db *gorm.DB) {
	db.Where("? > ?", c.query, c.args[0])
}

func (c *condition) fillGe(db *gorm.DB) {
	db.Where("? >= ?", c.query, c.args[0])
}

func (c *condition) fillLt(db *gorm.DB) {
	db.Where("? < ?", c.query, c.args[0])
}

func (c *condition) fillLe(db *gorm.DB) {
	db.Where("? <= ?", c.query, c.args[0])
}

func (c *condition) fillBetween(db *gorm.DB) {
	db.Where("? between ? and ?", c.query, c.args[0], c.args[1])
}

func (c *condition) fillNBetween(db *gorm.DB) {
	db.Where("? not between ? and ?", c.query, c.args[0], c.args[1])
}

func (c *condition) fillLike(db *gorm.DB) {
	db.Where("? like ?", c.query, fmt.Sprintf("%%%s%%", c.args[0]))
}

func (c *condition) fillNLike(db *gorm.DB) {
	db.Where("? not like ?", c.query, fmt.Sprintf("%%%s%%", c.args[0]))
}

func (c *condition) fillIsNull(db *gorm.DB) {
	db.Where("? is null", c.query)
}

func (c *condition) fillIsNNull(db *gorm.DB) {
	db.Where("? is not null", c.query)
}

func (c *condition) fillIn(db *gorm.DB) {
	db.Where("? in (?)", c.query, c.args)
}

func (c *condition) fillNIn(db *gorm.DB) {
	db.Where("? not in (?)", c.query, c.args)
}

func (c *condition) fillGroupBy(db *gorm.DB) {
	db.Group(c.query.(string))
}

func (c *condition) fillOrderByAsc(db *gorm.DB) {
	db.Order(fmt.Sprintf("%s asc", c.query))
}

func (c *condition) fillOrderByDesc(db *gorm.DB) {
	db.Order(fmt.Sprintf("%s desc", c.query))
}

func (c *condition) fillHaving(db *gorm.DB) {
	db.Having(c.query, c.args...)
}

func (c *condition) fillOr(db *gorm.DB) {
	db.Or(c.query, c.args...)
}

func (c *condition) fillDistinct(db *gorm.DB) {
	db.Distinct(c.args...)
}
