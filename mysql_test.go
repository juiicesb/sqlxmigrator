// +build mysql

package sqlxmigrator

import (
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	databases = append(databases, database{
		name:    "mysql",
		connEnv: "MYSQL_CONN_STRING",
	})
} 
