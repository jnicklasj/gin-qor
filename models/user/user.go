package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email     string     `json: "email" form:"email"`
	Password  string     `json:"password"`
	Name      string     `json: "name form:"name" gorm:"type:varchar(32);unique_index"`
	Gender    string     `json:"gender"`
	Role      string     `json:"role,omitempty"`
	Birthday  *time.Time `json:"birthday,omitempty"`
	Balance   float32    `json:"balance"`
	Confirmed int        `-`
}

func (user User) DisplayName() string {
	return user.Email
}

/*
type User struct {
  gorm.Model
  Name         string
  Age          sql.NullInt64
  Birthday     *time.Time
  Email        string  `gorm:"type:varchar(100);unique_index"`
  Role         string  `gorm:"size:255"` // 设置字段大小为255
  MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
  Num          int     `gorm:"AUTO_INCREMENT"` // 设置 num 为自增类型
  Address      string  `gorm:"index:addr"` // 给address字段创建名为addr的索引
  IgnoreMe     int     `gorm:"-"` // 忽略本字段
}

GORM框架中的默认行为

GORM 默认会使用名为ID的字段作为表的主键
表名默认就是结构体名称的复数例如结构体名称为user那么表名对应的是users
列名由字段名称进行下划线分割来生成例如:CreatedAt对应生成created_at
如果模型有CreatedAt字段，该字段的值将会是初次创建记录的时间
如果模型有UpdatedAt字段，该字段的值将会是每次更新记录的时间
如果模型有DeletedAt字段，调用Delete删除该记录时，将会设置DeletedAt字段为当前时间，而不是直接将记录从数据库中删除。

GORM设置表属性
修改表字段
结论：
修改表属性，只能在第一次建表的时候，有效！
或者给表增加新字段的时候，有效！
其他场景，修改表属性 ，在 gorm 操作中，无效！
*/
