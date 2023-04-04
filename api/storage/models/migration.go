package models

import "time"

type MigrationModel struct {
	ID        int    `json:"ID"`
	SqlDump   string `json:"SqlDump"`
	TableName string `json:"TableName"`
	*Timestamp
}

func (m *MigrationModel) SetTimestamp() {
	m.UpdatedAt = time.Now().Format(time.RFC3339)
	if len(m.CreatedAt) != 0 {
		m.CreatedAt = time.Now().Format(time.RFC3339)
	}
}

func (m *MigrationModel) TableCreate() string {
	return `
		create table if not exists migrations
		(
			id          serial
				constraint migrations_pk
					primary key,
			"sqlDump"  text      not null,
			"tableName"  text      not null,
			"CreatedAt" timestamp not null,
			"UpdatedAt" timestamp not null
		);
	`
}
