package api
import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"strconv"
	"time"
)
type Article struct {
	Id          int    "json:`id`"
	Title       string "json:`title`"
	Author      string "json:`author`"
	Content     string "json:`content`"
	State       int    "json:`state`"
	Created      int    "json:`create`"
	Updated      int    "json:`update`"
}
func Getpersons(c *gin.Context)  {
	    //数据库连接
		db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8")
		defer db.Close()
		if err != nil{
			log.Fatalln(err)
		}
		db.SetMaxIdleConns(20)
		db.SetMaxOpenConns(20)
		if err := db.Ping(); err != nil{
			log.Fatalln(err)
		}
	    rows, err := db.Query("SELECT title,author,content,created,updated,state FROM article")
		defer rows.Close()

		if err != nil {
			log.Fatalln(err)
		}
		data := make([]Article, 0)

		for rows.Next() {
			var article Article
			rows.Scan(&article.Title, &article.Author, &article.Content, &article.Created, &article.Updated, &article.State)
			data = append(data, article)
		}
		if err = rows.Err(); err != nil {
			log.Fatalln(err)
			c.JSON(http.StatusOK, gin.H{
				"status": 0,
				"data":  nil,
				"msg":    "查询失败",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"status": 1,
			"data":  data,
			"msg":    "查询成功",
		})

}
func Addpersons(c *gin.Context)  {
	//数据库连接
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8")
	defer db.Close()
	if err != nil{
		log.Fatalln(err)
	}


	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)

	if err := db.Ping(); err != nil{
		log.Fatalln(err)
	}
	title :=c.Request.FormValue("title")
	author := c.Request.FormValue("author")
	content := c.Request.FormValue("content")
	t := time.Now()
	updated := t.Format("2006-01-02 15:04:05")
	created := updated
	state := c.Request.FormValue("state")

	rs, err := db.Exec("insert into article(title,author,content,created,updated,state) values(?,?,?,?,?,?)", title,author,content,created,updated,state)
	if err != nil {
		log.Fatalln(err)
	}
	id, err := rs.LastInsertId()
	if err != nil {
		log.Fatalln(err)
		msg := fmt.Sprintf("文章保存失败！")
		c.JSON(http.StatusOK, gin.H{
			"msg":    msg,
			"data":   nil,
			"status": 0,
		})
	}
	fmt.Println("insert artilce Id {}", id)
	msg := fmt.Sprintf("insert successful %d", id)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
		"data":nil,
		"status":1,
	})
}
func UpdateArticle(c *gin.Context){
	//数据库连接
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8")
	defer db.Close()
	if err != nil{
		log.Fatalln(err)
	}


	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)

	if err := db.Ping(); err != nil{
		log.Fatalln(err)
	}
	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	title := c.Request.FormValue("title")
	author := c.Request.FormValue("author")
	content := c.Request.FormValue("content")
	t := time.Now()
	updated := t.Format("2006-01-02 15:04:05")

	state := c.Request.FormValue("state")
	article := Article{Id: id}
	err = c.Bind(&article)
	if err != nil {
		log.Fatalln(err)
	}

	stmt, err := db.Prepare ("update article set title=?,author=?,content=?,updated=?,state=? where id=?")
	defer stmt.Close()
	if err != nil {
		log.Fatalln(err)
	}
	rs, err := stmt.Exec(title,author,content,updated,state,article.Id)
	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": "数据更新失败",
			"data":nil,
			"status":0,
		})
	}
	ra, err := rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("更新文章id为%d成功%d", article.Id,ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
		"data":nil,
		"status":1,
	})
}
func DelArticle(c *gin.Context)  {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/test?charset=utf8")
	defer db.Close()
	if err != nil{
		log.Fatalln(err)
	}


	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)

	if err := db.Ping(); err != nil{
		log.Fatalln(err)
	}

	cid := c.Param("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		log.Fatalln(err)
	}
	rs, err := db.Exec("DELETE FROM article WHERE id=?", id)
	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusOK, gin.H{
			"msg": "数据库查询不到此题数据",
			"status":0,
		})
	}
	ra, err := rs.RowsAffected()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg": "删除数据失败",
			"status":0,
		})
	}
	msg := fmt.Sprintf("删除文章id为 %d 成功 %d", id, ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
		"status":1,
	})
}
