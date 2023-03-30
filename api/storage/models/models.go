package models

type Validate interface {
	Validate() bool
}

// Migration interface will be useful for database storage
type Migration interface {
	TableCreate() string
}

// SoftDeletion will useful for checking, that model is 'softDeleted'
type SoftDeletion interface {
	Soft()
	Undo()
}

// TimestampSetter will useful for setting timestamp
type TimestampSetter interface {
	SetTimestamp()
}

type Timestamp struct {
	CreatedAt string `json:"startedAt"`
	UpdatedAt string `json:"updatedAt"`
}

type SoftDelete struct {
	DeletedAt string `json:"deletedAt"`
}
