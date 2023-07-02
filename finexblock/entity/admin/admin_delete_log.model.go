package admin

import (
	"time"
)

type AdminDeleteLog struct {
	ID         uint      `gorm:"primaryKey;autoIncrement:true;comment:'기본키'" json:"id"`
	ExecutorID uint      `json:"executor_id,omitempty" gorm:"comment:'삭제한 운영진 id'"`
	TargetID   uint      `json:"target_id,omitempty" gorm:"comment:'삭제된 운영진 id'"`
	CreatedAt  time.Time `json:"created_at" gorm:"comment:'생성일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"comment:'수정일자';not null;type:timestamp;default:CURRENT_TIMESTAMP;type:timestamp"`

	//Executor Admin `gorm:"foreignKey:ExecutorID;references:ID;constraint:OnUpdate:CASCADE"`
	//Target   Admin `gorm:"foreignKey:TargetID;references:ID;constraint:OnUpdate:CASCADE"`
}

func (d *AdminDeleteLog) Alias() string {
	return "admin_delete_log adl"
}

func (d *AdminDeleteLog) TableName() string {
	return "admin_delete_log"
}
