package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type Priority string

func (p Priority) String() string {
	return string(p)
}

func (p Priority) Validate() error {
	switch p {
	case HIGH, MEDIUM, LOW:
		return nil
	}
	return errors.New("invalid priority")
}

type Metadata map[string]interface{}

const (
	HIGH   Priority = "HIGH"
	MEDIUM Priority = "MEDIUM"
	LOW    Priority = "LOW"
)

type FinexblockErrorLog struct {
	ID          uint     `gorm:"not null;primaryKey;autoIncrement;comment:'기본키'" json:"id"`
	ServerID    uint     `gorm:"not null;comment:서버 id" json:"serverId"`
	Process     string   `gorm:"type:LONGTEXT;not null;comment:프로세스" json:"process"`
	Priority    Priority `gorm:"type:enum('HIGH', 'MEDIUM', 'LOW');not null;comment:중요도" json:"priority"`
	Description string   `gorm:"type:LONGTEXT;not null;comment:부가 설명" json:"description"`
	Err         string   `gorm:"type:LONGTEXT;not null;comment:에러 메세지" json:"err"`
	Metadata    Metadata `gorm:"type:json;comment:첨부 metadata" json:"metadata"`
}

func (m Metadata) Value() (driver.Value, error) {
	if m == nil {
		return nil, nil
	}
	b, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return string(b), nil
}

func (m *Metadata) Scan(value interface{}) error {
	if value == nil {
		*m = nil
		return nil
	}
	temp, ok := value.(string)
	if !ok {
		return errors.New("Failed to unmarshal JSON value")
	}
	return json.Unmarshal([]byte(temp), m)
}

func (e *FinexblockErrorLog) TableName() string {
	return "finexblock_error_log"
}

func (e *FinexblockErrorLog) Alias() string {
	return "finexblock_error_log fel"
}