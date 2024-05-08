## 快速入门

```go
// Person 定义model
type Person struct {
	Id   int64  `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`
	Name string `gorm:"column:name;default:null" json:"name"`
	Age  int64  `gorm:"column:age;default:null" json:"age"`
}

func (Person) TableName() string {
	return "person"
}

// Repo接口
type IPersonRepo interface {
	// 继承此接口，并指定泛型类型
	gormplus.IRepository[Person]
	// 自定义接口
	GetOneByName(ctx context.Context, name string) (*Person, error)
}

// Repo实现
type personRepo struct {
	// 继承此实现，并指定泛型类型
	gormplus.Repository[Person]
	db *gorm.DB
}

// NewPersonRepo 初始化函数
func NewPersonRepo(db *gorm.DB) IPersonRepo {
	return &personRepo{
		Repository: gormplus.NewRepository[Person](db),
		db:         db,
	}
}

// GetOneByName 可以自定义实现，也可以使用已实现的方法
func (r *personRepo) GetOneByName(ctx context.Context, name string) (*Person, error) {
	// 查询包装器
	qw := wrapper.Query().Eq("name", name)
	person, err := r.GetOne(ctx, qw)
	return person, err
}

// TestRepo 测试
func TestRepo(t *testing.T) {
	var db *gorm.DB
	ctx := context.Background()
	repo := NewPersonRepo(db)
	person, err := repo.GetOneByName(ctx, "大黄")
	if err != nil {
		log.Printf("err:%v", err)
		return
	}
	log.Printf("person:%+v", person)
}
```

## 方法介绍
```go
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
```

