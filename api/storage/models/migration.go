package models

import (
	"fmt"
	"time"
)

type Migration struct {
	ID        int    `json:"ID"`
	SqlDump   string `json:"SqlDump"`
	TableName string `json:"TableName"`
	*Timestamp
}

func CreateMigrationModel(tablename, sql string) *Migration {
	mm := &Migration{
		TableName: tablename,
		SqlDump:   sql,
		Timestamp: &Timestamp{},
	}

	mm.SetTimestamp()
	return mm
}

func (m *Migration) CreateRecord() string {
	return fmt.Sprintf(`INSERT INTO migrations ("sqlDump", "tableName", "CreatedAt", "UpdatedAt") VALUES ('%s','%s','%s','%s')`, m.SqlDump, m.TableName, m.CreatedAt, m.UpdatedAt)
}

func (m *Migration) SetTimestamp() {
	m.UpdatedAt = time.Now().Format(time.RFC3339)
	if len(m.CreatedAt) == 0 {
		m.CreatedAt = time.Now().Format(time.RFC3339)
	}
}

func (m *Migration) TableCreate() *TableCreation {
	return &TableCreation{
		Sql: `
			create table migrations
			(
				ID          serial
					constraint migrations_pk
						primary key,
				"sqlDump"  text      not null,
				"tableName"  text      not null,
				"CreatedAt" timestamp not null,
				"UpdatedAt" timestamp not null
			);
		`,
		Dependencies: []string{},
	}
}
