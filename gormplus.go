package gormplus

import (
	"context"
	"errors"
	"gorm-plus/wrapper"
	"gorm.io/gorm"
)

type IRepository[T any] interface {
	// GetOne 指定查询条件，获取某一列数据
	GetOne(ctx context.Context, qw *wrapper.QueryWrapper) (*T, error)
	// GetOneById 根据主键获取某一些数据
	GetOneById(ctx context.Context, id interface{}) (*T, error)
	// GetList 指定查询条件，获取多列数据
	GetList(ctx context.Context, qw *wrapper.QueryWrapper) ([]*T, error)
	// Page 指定分页参数以及查询条件，分页查询
	Page(ctx context.Context, page *wrapper.Page, qw *wrapper.QueryWrapper) ([]*T, int64, error)
	// Count 指定统计条件，统计
	Count(ctx context.Context, qw *wrapper.QueryWrapper) (int64, error)
	// Save 添加一列数据
	Save(ctx context.Context, entity *T) error
	// SaveBatch 批量添加多列数据
	SaveBatch(ctx context.Context, entityList []*T) error
	// Delete 指定删除条件，删除数据
	Delete(ctx context.Context, qw *wrapper.QueryWrapper) error
	// DeleteById 根据主键删除某一些数据
	DeleteById(ctx context.Context, id interface{}) error
	// DeleteByIds 根据主键批量删除数据
	DeleteByIds(ctx context.Context, ids []interface{}) error
	// Update 指定更新条件以及参数，更新数据
	Update(ctx context.Context, uw *wrapper.UpdateWrapper) error
	// UpdateById 根据主键更新所有的字段（即使字段是零值），如果结构体不包含主键，它将执行 Create 操作
	UpdateById(ctx context.Context, entity *T) error
}

type Repository[T any] struct {
	db *gorm.DB
}

func NewRepository[T any](db *gorm.DB) Repository[T] {
	return Repository[T]{db: db}
}

func (r *Repository[T]) GetOne(ctx context.Context, qw *wrapper.QueryWrapper) (*T, error) {
	var entity T
	db := r.db.WithContext(ctx)
	qw.FillCondition(db)
	err := db.First(&entity).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
	}
	return &entity, err
}

func (r *Repository[T]) GetOneById(ctx context.Context, id interface{}) (*T, error) {
	var entity T
	err := r.db.WithContext(ctx).Where("id = ?", id).First(&entity).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &entity, err
}

func (r *Repository[T]) GetList(ctx context.Context, qw *wrapper.QueryWrapper) ([]*T, error) {
	var entityList []*T
	db := r.db.WithContext(ctx)
	qw.FillCondition(db)
	err := db.Find(&entityList).Error
	return entityList, err
}

func (r *Repository[T]) Page(ctx context.Context, page *wrapper.Page, qw *wrapper.QueryWrapper) ([]*T, int64, error) {
	var entityList []*T
	var total int64
	db := r.db.WithContext(ctx)
	qw.FillCondition(db)
	err := db.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	err = db.Limit(page.PageSize).Offset((page.Page - 1) * page.PageSize).
		Find(&entityList).Error
	return entityList, total, err
}

func (r *Repository[T]) Count(ctx context.Context, qw *wrapper.QueryWrapper) (int64, error) {
	var total int64
	db := r.db.WithContext(ctx)
	qw.FillCondition(db)
	err := db.Count(&total).Error
	return total, err
}

func (r *Repository[T]) Save(ctx context.Context, entity *T) error {
	err := r.db.WithContext(ctx).Create(entity).Error
	return err
}

func (r *Repository[T]) SaveBatch(ctx context.Context, entityList []*T) error {
	err := r.db.WithContext(ctx).Create(entityList).Error
	return err
}

func (r *Repository[T]) Delete(ctx context.Context, qw *wrapper.QueryWrapper) error {
	var entity T
	db := r.db.WithContext(ctx)
	qw.FillCondition(db)
	err := db.Delete(&entity).Error
	return err
}

func (r *Repository[T]) DeleteById(ctx context.Context, id interface{}) error {
	var entity T
	err := r.db.WithContext(ctx).Delete(&entity, id).Error
	return err
}

func (r *Repository[T]) DeleteByIds(ctx context.Context, ids []interface{}) error {
	var entityList []T
	err := r.db.WithContext(ctx).Delete(&entityList, ids).Error
	return err
}

func (r *Repository[T]) Update(ctx context.Context, uw *wrapper.UpdateWrapper) error {
	var entity T
	db := r.db.WithContext(ctx).Model(&entity)
	uw.FillCondition(db)
	err := uw.FillSet(db).Error
	return err
}

func (r *Repository[T]) UpdateById(ctx context.Context, entity *T) error {
	err := r.db.WithContext(ctx).Save(&entity).Error
	return err
}
