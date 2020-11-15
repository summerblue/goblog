package pagination

import (
	"github.com/vcraescu/go-paginator"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// GORMAdapter gorm adapter 会被传送给 paginator 构造器
type GORMAdapter struct {
	db *gorm.DB
}

// NewGORMAdapter db 参数是 db 请求语句
func NewGORMAdapter(db *gorm.DB) paginator.Adapter {
	return &GORMAdapter{db: db}
}

// Nums 返回数据条数
func (a *GORMAdapter) Nums() (int64, error) {
	var count int64
	if err := a.db.Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// Slice stores into data argument a slice of the results.
// data must be a pointer to a slice of models.
func (a *GORMAdapter) Slice(offset, length int, data interface{}) error {
	// Work on a dedicated session to not offset the total count nums
	return a.db.Session(&gorm.Session{}).Preload(clause.Associations).Limit(length).Offset(offset).Find(data).Error
}
