package user

import "time"

type UserDormantLog struct {
	ID        uint      `gorm:"primaryKey;autoIncrement:true;comment:'기본키'" json:"id"`
	UserID    uint      `json:"user_id,omitempty" gorm:"comment:'유저 id'"`
	Type      string    `json:"type,omitempty" gorm:"comment:'전환타입 - 휴면 전환, 휴면 해제, 개인정보제거';type:enum('DORMANT', 'EXIT_DORMANT', 'DELETED');not null"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;comment:'생성일자';not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;comment:'수정일자';not null;default:CURRENT_TIMESTAMP"`
	DeletedAt time.Time `json:"deleted_at" gorm:"type:timestamp;comment:'삭제일자'"`
}

func (d *UserDormantLog) Alias() string {
	return "user_dormant_log udl"
}

func (d *UserDormantLog) TableName() string {
	return "user_dormant_log"
}
