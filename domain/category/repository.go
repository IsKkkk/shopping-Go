package category

import (
	"log"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

// 创建商品分类
func NewCategoryRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// 生成商品分类表
func (r *Repository) Migration() {
	err := r.db.AutoMigrate(&Category{})
	if err != nil {
		log.Print(err)
	}
}

// 生成商品分类测试数据
func (r *Repository) InsertSampleData() {
	categories := []Category{
		{Name: "CAT1", Desc: "Category 1"},
		{Name: "CAT2", Desc: "Category 2"},
	}

	for _, c := range categories {
		r.db.Where(Category{Name: c.Name}).Attrs(Category{Name: c.Name}).FirstOrCreate(&c)
	}
}

// 创建商品分类
func (r *Repository) Create(c *Category) error {
	result := r.db.Create(c)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// 通过名称查询商品分类
func (r *Repository) GetByName(name string) []Category {
	var categories []Category
	r.db.Where("Name = ?", name).Find(&categories)

	return categories
}

// 批量创建商品分类
func (r *Repository) BulkCreate(categories []*Category) (int, error) {
	var count int64
	err := r.db.Create(&categories).Count(&count).Error
	return int(count), err
}

// 获得分页商品分类
func (r *Repository) GetAll(pageIndex, pageSize int) ([]Category, int) {
	var categories []Category
	var count int64

	r.db.Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(&categories).Count(&count)

	return categories, int(count)
}
