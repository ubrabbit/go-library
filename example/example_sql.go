package main

import (
	"context"
	"fmt"
	"time"
)

import (
	. "go-library/database/sql"
	log "go-library/log"
)

const (
	USERNAME string = "umail"
	PASSWORD string = "123456"
	HOST     string = "192.168.1.78"
	PORT     int    = 6033
	DBNAME   string = "umail"
)

func test_db_1() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", USERNAME, PASSWORD, HOST, PORT, DBNAME)
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
	defer db.Close()

	if err := db.Ping(context.TODO()); err != nil {
		log.Error("MySQL: ping error(%v)", err)
	} else {
		fmt.Println("MySQL: ping ok")
	}

	table := "CREATE TABLE IF NOT EXISTS `test` (`id` int(11) NOT NULL AUTO_INCREMENT COMMENT '自增ID', `name` varchar(16) NOT NULL DEFAULT '' COMMENT '名称', PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8"
	if _, err := db.Exec(context.TODO(), table); err != nil {
		log.Error("MySQL: create table error(%v)", err)
	} else {
		fmt.Println("MySQL: create table ok")
	}

	sql := "INSERT INTO test(name) VALUES(?)"
	if _, err := db.Exec(context.TODO(), sql, "test"); err != nil {
		log.Error("MySQL: insert error(%v)", err)
	} else {
		fmt.Println("MySQL: insert ok")
	}
}

func main() {
	fmt.Println("start")

	test_db_1()
}
