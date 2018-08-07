package postgres

import (
	"fmt"
	"testing"
)

var (
	sql1 = "select id, user_id, other_user_id, sticker_id, question_id, reference_id, moment_id, location, value, coalesce(recalled,false) as recalled, created_time, coalesce(sent_from, '') as sent_from, status  FROM rel_8192_1.select_messages(?,?,?,?,?,?,?,?);"
	sql2 = "select id, xe. fd a fff from pg.postgres(?)"
	sql3 = "select id from select_user_by_id(?)"
	sql4 = "select update_user_by_id(?)"
	sql5 = "select select_user_by_id(?)"
)

func TestGetFunctionName(t *testing.T) {

	name := getFuncName(sql1)
	fmt.Printf("--> \"%s\"\n", name)
	if name != "sql.select_messages" {
		t.Fatal(name)
	}

	name = getFuncName(sql2)
	fmt.Printf("--> \"%s\"\n", name)
	if name != "sql.postgres" {
		t.Fatal(name)
	}

	name = getFuncName(sql3)
	fmt.Printf("--> \"%s\"\n", name)
	if name != "sql.select_user_by_id" {
		t.Fatal(name)
	}


	name = getFuncName(sql4)
	fmt.Printf("--> \"%s\"\n", name)
	if name != "sql.update_user_by_id" {
		t.Fatal(name)
	}

	name = getFuncName(sql5)
	fmt.Printf("--> \"%s\"\n", name)
	if name != "sql.select_user_by_id" {
		t.Fatal(name)
	}
}

func TestGetFunctionNameMaxMapSize(t *testing.T) {
	sqltmp := "select id, user_id, other_user_id, sticker_id, question_id, reference_id, moment_id, location, value, coalesce(recalled,false) as recalled, created_time, coalesce(sent_from, '') as sent_from, status  FROM rel_8192_1.select_messages%v(?,?,?,?,?,?,?,?);"
	for i := 0; i < 101; i++ {
		sql := fmt.Sprintf(sqltmp, i)
		name := getFuncName(sql)
		fmt.Printf("%d : \"%s\"\n", i, name)
	}
}
