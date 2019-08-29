package common

import (
	"context"
	"fmt"
	"time"
)

import (
	. "go-library/database/sql"
	log "go-library/log"
)

func ExampleConnect(user string, pwd string, host string, port int, dbname string) *DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", user, pwd, host, port, dbname)
	dsn = dsn + "?timeout=5s&readTimeout=5s&writeTimeout=5s&parseTime=true&loc=Local&charset=utf8"
	c := &Config{
		Addr:         "test",
		DSN:          dsn,
		Active:       10,
		Idle:         5,
		IdleTimeout:  time.Duration(time.Minute),
		QueryTimeout: time.Duration(time.Minute),
		ExecTimeout:  time.Duration(time.Minute),
		TranTimeout:  time.Duration(time.Minute),
	}
	db := NewMySQL(c)
	return db
}

func ExamplePing(db *DB) {
	if err := db.Ping(context.TODO()); err != nil {
		log.Fatal("MySQL: ping error(%v)", err)
	} else {
		log.Info("MySQL: ping ok")
	}
}

func ExampleTable(db *DB) {
	table := "CREATE TABLE IF NOT EXISTS `test` (`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增ID', `name` varchar(16) NOT NULL DEFAULT '' COMMENT '名称', PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8"
	if _, err := db.Exec(context.TODO(), table); err != nil {
		log.Fatal("MySQL: create table error(%v)", err)
	} else {
		log.Info("MySQL: create table ok")
	}
}

func ExampleInsert(db *DB) {
	sql := "INSERT INTO test(name) VALUES(?)"
	if _, err := db.Exec(context.TODO(), sql, "test"); err != nil {
		log.Fatal("MySQL: insert error(%v)", err)
	} else {
		log.Info("MySQL: insert ok")
	}
}

func ExampleQuery(db *DB) {
	sql := "SELECT name FROM test WHERE name=?"
	rows, err := db.Query(context.TODO(), sql, "test")
	if err != nil {
		log.Fatal("MySQL: query error(%v)", err)
	}
	defer rows.Close()
	for rows.Next() {
		name := ""
		if err := rows.Scan(&name); err != nil {
			log.Fatal("MySQL: query scan error(%v)", err)
		} else if name != "test" {
			log.Fatal("MySQL: query name error: %s", name)
		}
	}
	log.Info("MySQL: query ok")
}
