package po

import "time"

type ExecutionEvent struct {
	ID             int64     `json:"id" gorm:"primaryKey;autoIncrement;type:INT(11)"`
	ExecutionID    int       `json:"execution_id" gorm:"index;not null;type:INT(11)"` //所在的Executeion
	StepID         int       `json:"step_id"  gorm:"type:INT(11);index;not null"`
	StepName       string    `json:"step_name" gorm:"size:100"`
	EventType      string    `json:"event_type"  gorm:"not null;size:255"`
	Data           string    `json:"data"  gorm:"type:JSON"`
	NanoSeconds    int64     `json:"nano_seconds"  gorm:"not null;index;type:BIGINT UNSIGNED"`
	StartTime      time.Time `json:"start_time" gorm:"type:TIMESTAMP NULL;default:NULL" `  //开始时间
	FinishTime     time.Time `json:"finish_time"  gorm:"type:TIMESTAMP NULL;default:NULL"` // 结束时间
	DispatcherName string    `json:"dispatcher_name" gorm:"size:128"`                      //dispatcher name
	GmtCreated     time.Time `json:"gmt_created" gorm:"<-:create;autoCreateTime;type:TIMESTAMP" `
}