package main

import (
	"bufio"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strings"
	"time"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@/test?charset=utf8")
	checkErr(err)
	defer db.Close()
	sql1 := "SELECT title,text,created FROM typecho_contents WHERE type = ?"
	stmt, err := db.Prepare(sql1)
	checkErr(err)
	res, err := stmt.Query("post")
	checkErr(err)

	for res.Next() {
		var created int64
		var title,text string
		err = res.Scan(&title, &text, &created)
		checkErr(err)

		title = strings.Replace(title, "/", "&", -1)
		text = strings.Replace(text, "<!--markdown-->", "", -1)
		text = strings.Replace(text, "https://qn.samlog.club", "https://gitee.com/chengpro_admin/image/raw/master/test", -1)
		var path string = "/Users/yanghucheng/blog/source/_posts/"
		var filename string
		filename = title + ".md"
		file, err := os.OpenFile(path + filename, os.O_RDWR | os.O_APPEND | os.O_CREATE, 0666)
		checkErr(err)
		defer file.Close()
		var content string = "---\ntitle: " + title + "\ndate: " + time.Unix(created, 0).Format("2006-01-02 15:04:05") +"\ntags:\n---\n"
		content = content + text
		writer := bufio.NewWriter(file) // 写入缓冲
		writer.WriteString(content)
		writer.Flush()
	}

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
