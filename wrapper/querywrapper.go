package wrapper

import (
	"gorm.io/gorm"
)

type QueryWrapper struct {
	conditions []condition
}

func Query(capacity ...int) *QueryWrapper {
	var conditions []condition
	if len(capacity) > 0 {
		conditions = make([]condition, 0, capacity[0])
	} else {
		conditions = make([]condition, 0)
	}
	return &QueryWrapper{conditions: conditions}
}

func (qw *QueryWrapper) Where(query interface{}, args ...interface{}) *QueryWrapper {
	qw.addCondition(where, query, args...)
	return qw
}

func (qw *QueryWrapper) Eq(field string, arg interface{}) *QueryWrapper {
	qw.addCondition(eq, field, arg)
	return qw
}

func (qw *QueryWrapper) NEq(field string, arg interface{}) *QueryWrapper {
	qw.addCondition(nEq, field, arg)
	return qw
}

func (qw *QueryWrapper) Gt(field string, arg interface{}) *QueryWrapper {
	qw.addCondition(gt, field, arg)
	return qw
}

func (qw *QueryWrapper) Ge(field string, arg interface{}) *QueryWrapper {
	qw.addCondition(ge, field, arg)
	return qw
}

func (qw *QueryWrapper) Lt(field string, arg interface{}) {
	qw.addCondition(lt, field, arg)
}

func (qw *QueryWrapper) Le(field string, arg interface{}) {
	qw.addCondition(le, field, arg)
}

func (qw *QueryWrapper) Between(field string, arg ...interface{}) {
	qw.addCondition(between, field, arg...)
}

func (qw *QueryWrapper) NBetween(field string, arg ...interface{}) {
	qw.addCondition(nBetween, field, arg...)

}

func (qw *QueryWrapper) Like(field string, arg interface{}) {
	qw.addCondition(like, field, arg)
}

func (qw *QueryWrapper) NLike(field string, arg interface{}) {
	qw.addCondition(nLike, field, arg)
}

func (qw *QueryWrapper) IsNull(field string) {
	qw.addCondition(isNull, field)
}

func (qw *QueryWrapper) IsNNull(field string) {
	qw.addCondition(isNNull, field)
}

func (qw *QueryWrapper) In(field string, args interface{}) {
	qw.addCondition(in, field, args)
}

func (qw *QueryWrapper) NIn(field string, args interface{}) {
	qw.addCondition(nIn, field, args)
}

func (qw *QueryWrapper) GroupBy(field string) {
	qw.addCondition(groupBy, field)
}

func (qw *QueryWrapper) OrderByAsc(field string) {
	qw.addCondition(orderByAsc, field)
}

func (qw *QueryWrapper) OrderByDesc(field string) {
	qw.addCondition(orderByDesc, field)
}

func (qw *QueryWrapper) Having(query string, arg ...interface{}) {
	qw.addCondition(having, query, arg)
}

func (qw *QueryWrapper) Or(query string, arg ...interface{}) {
	qw.addCondition(having, query, arg)
}

func (qw *QueryWrapper) Distinct(arg ...interface{}) {
	qw.addCondition(having, nil, arg...)
}

func (qw *QueryWrapper) FillCondition(db *gorm.DB) {
	for _, c := range qw.conditions {
		c.fill(db)
	}
}

func (qw *QueryWrapper) addCondition(keyType int, query interface{}, args ...interface{}) {
	qw.conditions = append(qw.conditions, condition{
		keyType: keyType,
		query:   query,
		args:    args,
	})
}
