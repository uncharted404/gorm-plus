package wrapper

import (
	"gorm.io/gorm"
	"reflect"
)

type QueryWrapper struct {
	conditions []condition
	isCheck    bool
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

func (qw *QueryWrapper) CheckZero(args ...interface{}) *QueryWrapper {
	for _, arg := range args {
		if arg == nil || reflect.ValueOf(args).IsZero() {
			qw.isCheck = true
			return qw
		}
	}
	qw.isCheck = false
	return qw
}

func (qw *QueryWrapper) Check(checkFunc func() bool) *QueryWrapper {
	qw.isCheck = checkFunc()
	return qw
}

func (qw *QueryWrapper) Where(query interface{}, args ...interface{}) *QueryWrapper {
	if qw.isCheck {
		qw.isCheck = false
		return qw
	}
	qw.addCondition(where, query, args...)
	return qw
}

func (qw *QueryWrapper) Eq(field string, arg interface{}) *QueryWrapper {
	if qw.isCheck {
		qw.isCheck = false
		return qw
	}
	qw.addCondition(eq, field, arg)
	return qw
}

func (qw *QueryWrapper) NEq(field string, arg interface{}) *QueryWrapper {
	if qw.isCheck {
		qw.isCheck = false
		return qw
	}
	qw.addCondition(nEq, field, arg)
	return qw
}

func (qw *QueryWrapper) Gt(field string, arg interface{}) *QueryWrapper {
	if qw.isCheck {
		qw.isCheck = false
		return qw
	}
	qw.addCondition(gt, field, arg)
	return qw
}

func (qw *QueryWrapper) Ge(field string, arg interface{}) *QueryWrapper {
	if qw.isCheck {
		qw.isCheck = false
		return qw
	}
	qw.addCondition(ge, field, arg)
	return qw
}

func (qw *QueryWrapper) Lt(field string, arg interface{}) *QueryWrapper {
	if qw.isCheck {
		qw.isCheck = false
		return qw
	}
	qw.addCondition(lt, field, arg)
	return qw
}

func (qw *QueryWrapper) Le(field string, arg interface{}) *QueryWrapper {
	if qw.isCheck {
		qw.isCheck = false
		return qw
	}
	qw.addCondition(le, field, arg)
	return qw
}

func (qw *QueryWrapper) Between(field string, left, right interface{}) *QueryWrapper {
	if qw.isCheck {
		qw.isCheck = false
		return qw
	}
	qw.addCondition(between, field, left, right)
	return qw
}

func (qw *QueryWrapper) NBetween(field string, left, right interface{}) *QueryWrapper {
	if qw.isCheck {
		qw.isCheck = false
		return qw
	}
	qw.addCondition(nBetween, field, left, right)
	return qw
}

func (qw *QueryWrapper) Like(field string, arg interface{}) *QueryWrapper {
	if qw.isCheck {
		qw.isCheck = false
		return qw
	}
	qw.addCondition(like, field, arg)
	return qw
}

func (qw *QueryWrapper) NLike(field string, arg interface{}) *QueryWrapper {
	if qw.isCheck {
		qw.isCheck = false
		return qw
	}
	qw.addCondition(nLike, field, arg)
	return qw
}

func (qw *QueryWrapper) IsNull(field string) *QueryWrapper {
	if qw.isCheck {
		qw.isCheck = false
		return qw
	}
	qw.addCondition(isNull, field)
	return qw
}

func (qw *QueryWrapper) IsNNull(field string) *QueryWrapper {
	if qw.isCheck {
		qw.isCheck = false
		return qw
	}
	qw.addCondition(isNNull, field)
	return qw
}

func (qw *QueryWrapper) In(field string, args interface{}) *QueryWrapper {
	if qw.isCheck {
		qw.isCheck = false
		return qw
	}
	qw.addCondition(in, field, args)
	return qw
}

func (qw *QueryWrapper) NIn(field string, args interface{}) *QueryWrapper {
	if qw.isCheck {
		qw.isCheck = false
		return qw
	}
	qw.addCondition(nIn, field, args)
	return qw
}

func (qw *QueryWrapper) GroupBy(field string) *QueryWrapper {
	if qw.isCheck {
		qw.isCheck = false
		return qw
	}
	qw.addCondition(groupBy, field)
	return qw
}

func (qw *QueryWrapper) OrderByAsc(field string) *QueryWrapper {
	if qw.isCheck {
		qw.isCheck = false
		return qw
	}
	qw.addCondition(orderByAsc, field)
	return qw
}

func (qw *QueryWrapper) OrderByDesc(field string) *QueryWrapper {
	if qw.isCheck {
		qw.isCheck = false
		return qw
	}
	qw.addCondition(orderByDesc, field)
	return qw
}

func (qw *QueryWrapper) Having(query string, arg ...interface{}) *QueryWrapper {
	if qw.isCheck {
		qw.isCheck = false
		return qw
	}
	qw.addCondition(having, query, arg)
	return qw
}

func (qw *QueryWrapper) Or(query string, arg ...interface{}) *QueryWrapper {
	if qw.isCheck {
		qw.isCheck = false
		return qw
	}
	qw.addCondition(having, query, arg)
	return qw
}

func (qw *QueryWrapper) Distinct(arg ...interface{}) *QueryWrapper {
	if qw.isCheck {
		qw.isCheck = false
		return qw
	}
	qw.addCondition(having, nil, arg...)
	return qw
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
