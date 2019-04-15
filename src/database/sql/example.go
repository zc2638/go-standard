package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

// 提供了保证SQL或类SQL数据库的泛用接口。使用sql包时必须注入（至少）一个数据库驱动
func main() {

	// 注册并命名一个数据库，可以在Open函数中使用该命名启用该驱动
	// 如果注册同一名称两次，或者driver参数为nil，会导致panic
	// 注册mysql驱动，可以直接import引入使用，mysql包的init方法已经做了Register
	sql.Register("myMysql", &mysql.MySQLDriver{})

	// 查看所有已经注册的驱动
	// 返回所有驱动注册名称
	fmt.Println(sql.Drivers())

	// 提供了一种更简洁的方式来创建 sql.NamedArg
	// 可以直接使用sql.NamedArg{}
	// NamedArg是一个命名参数。NamedArg值可以用作Query或Exec的参数，并绑定到SQL语句中的相应命名参数
	sql.Named("nickname", "")

	// 连接数据库, driverName数据库驱动名称，dataSourceName数据库连接信息
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	// 关闭数据库
	defer db.Close()

	// 返回数据库下层驱动
	fmt.Println(db.Driver())

	// 检查与数据库的连接是否仍有效，如果需要会创建连接
	fmt.Println(db.Ping())

	// 设置与数据库建立连接的最大数目
	// 如果n大于0且小于最大闲置连接数，会将最大闲置连接数减小到匹配最大开启连接数的限制
	// 如果n <= 0，不会限制最大开启连接数，默认为0（无限制）
	db.SetMaxOpenConns(0)

	// 设置连接池中的最大闲置连接数
	// 如果n大于最大开启连接数，则新的最大闲置连接数会减小到匹配最大开启连接数的限制
	// 如果n <= 0，不会保留闲置连接
	db.SetMaxIdleConns(0)

	// 设置连接最大生存时间
	// 过期的连接可能会在重用之前被缓慢地关闭
	// 如果d <= 0，连接将永久重用
	db.SetConnMaxLifetime(0)

	// 返回数据库统计信息
	stats(db)

	// 查询
	query(db)

	// 查询一行
	queryRows(db)

	// 执行一次命令，不返回任何结果
	exec(db)

	// 提前准备一条语句，可用于多次执行数据库操作
	prepare(db)

	// 事务
	transaction(db)
}

func stats(db *sql.DB) {

	// 返回数据库统计信息
	s := db.Stats()

	// 空闲连接数
	fmt.Println(s.Idle)

	// 当前使用的连接数
	fmt.Println(s.InUse)

	// 由于setMaxIdleConns而关闭的连接总数
	fmt.Println(s.MaxIdleClosed)

	// 由于setConnmaxLifetime而关闭的连接总数
	fmt.Println(s.MaxLifetimeClosed)

	// 与数据库的最大打开连接数
	fmt.Println(s.MaxOpenConnections)

	// 已建立的在用和空闲的连接数
	fmt.Println(s.OpenConnections)

	// 等待的总连接数
	fmt.Println(s.WaitCount)

	// 等待新连接所阻塞的总时间
	fmt.Println(s.WaitDuration)
}

func query(db *sql.DB) {

	// 查询
	// 等同于 db.QueryContext(context.Background(), "select nickname from users where id=?", 1)
	rows, err := db.Query("SELECT nickname FROM users WHERE id=?", 1)
	if err != nil {
		log.Fatal(err)
	}

	res := make([]string, 0)
	// 循环取值
	for rows.Next() {
		var nickname string
		// 赋值，rows中的列的数量必须与值的数量相同
		if err := rows.Scan(&nickname); err != nil {
			log.Fatal(err)
		}
		res = append(res, nickname)
	}

	// 关闭查询
	if err := rows.Close(); err != nil {
		log.Fatal(err)
	}

	// 返回扫描赋值时的最后个错误
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}

func queryRows(db *sql.DB) {

	// 查询最多一行信息
	// 等同于 db.QueryRowContext(context.Background(), "select nickname from users where id=?", 1)
	row := db.QueryRow("SELECT nickname FROM users WHERE id=?", 1)
	var name string
	if err := row.Scan(&name); err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
}

func exec(db *sql.DB) {

	// 执行一次命令（包括查询、删除、更新、插入等），不返回任何执行结果。参数args表示query中的占位参数
	// 等同于 db.ExecContext(context.Background(), `select nickname from users where id = 1`)
	res, err := db.Exec(`SELECT nickname FROM users WHERE id = 1`)
	if err != nil {
		log.Fatal(err)
	}

	// 返回一个数据库生成的回应命令的整数
	// 当插入新行时，一般来自一个"自增ID"
	// 不是所有的数据库都支持该功能，该状态的语法也各有不同
	fmt.Println(res.LastInsertId())

	// 返回被update、insert或delete命令影响的行数
	// 不是所有的数据库都支持该功能
	fmt.Println(res.RowsAffected())
}

func prepare(db *sql.DB) {

	// 创建一个准备好的状态用于之后的查询和命令
	// 返回值可以同时执行多个查询和命令
	stmt, err := db.Prepare("SELECT nickname FROM users WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}

	// 使用提供的参数执行准备好的查询状态，返回Rows类型查询结果
	// 其它Query, Exec用法参见db相关方法
	row := stmt.QueryRow(1)
	var name string
	if err := row.Scan(&name); err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
}

func transaction(db *sql.DB) {

	// 开启事务
	// 等同于 db.BeginTx(context.Background(), nil)
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	// 使用提供的参数执行准备好的查询状态，返回Rows类型查询结果
	// 其它Query, Exec等用法参见db相关方法
	row := tx.QueryRow("SELECT nickname FROM users WHERE id = ?", 1)
	var name string
	if err := row.Scan(&name); err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)

	// 提交事务并关闭
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	// 回滚事务并关闭
	//tx.Rollback()
}