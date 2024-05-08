package gormplus

import (
	"context"
	"gorm-plus/wrapper"
	"gorm.io/gorm"
	"log"
	"testing"
)

type Person struct {
	Id   int64  `gorm:"column:id;primaryKey;autoIncrement;not null" json:"id"`
	Name string `gorm:"column:name;default:null" json:"name"`
	Age  int64  `gorm:"column:age;default:null" json:"age"`
}

func (Person) TableName() string {
	return "person"
}

type IPersonRepo interface {
	IRepository[Person]
	GetOneByName(ctx context.Context, name string) (*Person, error)
}

type personRepo struct {
	Repository[Person]
	db *gorm.DB
}

func NewPersonRepo(db *gorm.DB) IPersonRepo {
	return &personRepo{
		Repository: NewRepository[Person](db),
		db:         db,
	}
}

func (r *personRepo) GetOneByName(ctx context.Context, name string) (*Person, error) {
	qw := wrapper.Query().Eq("name", name)
	person, err := r.GetOne(ctx, qw)
	return person, err
}

func TestRepo(t *testing.T) {
	ctx := context.Background()
	repo := NewPersonRepo(&gorm.DB{})
	person, err := repo.GetOneByName(ctx, "大黄")
	if err != nil {
		log.Printf("err:%v", err)
		return
	}
	log.Printf("person:%+v", person)
}
