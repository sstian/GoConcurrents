// NOT RUN!

package main

import (
	"fmt"
	"time"

	"xorm.io/xorm"
)

func main() {
	/* 1. 同步结构体到数据库 */
	var (
		userName  string = "root"
		password  string = "root"
		ipAddress string = "127.0.0.1"
		port      int    = 3306
		dbName    string = "go_test"
		charset   string = "utf8mb4"
	)
	// dataSourceName = root:root@tcp(127.0.0.1:3306)/go_test?charset=utf8mb4
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddress, port, dbName, charset)

	//engine xorm核心引擎 连接数据库
	engine, err := xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		fmt.Println("数据库连接失败！", err)
	}

	//建立结构体 把结构同步到数据库里面
	type User struct {
		Id      int64
		Name    string
		Age     int
		Passwd  string    `xorm:"varchar(200)"`
		Created time.Time `xorm:"created"` //创建的时候自动同步创建
		Updated time.Time `xorm:"updated"` //更新的时候自动同步更新
	}

	// engine.Sync
	err = engine.Sync(new(User))
	if err != nil {
		fmt.Println("表结构同步失败！")
	}

	/* 2. 插入数据 */
	//engine.Insert() 插入 对象 返回值：受影响的行数
	user := User{
		Id:      10003,
		Name:    "lf",
		Age:     22,
		Passwd:  "1234",
		Created: time.Time{}, //可写可不写
		Updated: time.Time{},
	}
	user1 := User{
		Id:      10002,
		Name:    "lf",
		Age:     22,
		Passwd:  "1234",
		Created: time.Time{}, //可写可不写
		Updated: time.Time{},
	}
	user2 := User{
		Id:      10001,
		Name:    "lf",
		Age:     22,
		Passwd:  "1234",
		Created: time.Time{}, //可写可不写
		Updated: time.Time{},
	}
	n, err := engine.Insert(&user, &user1, &user2)
	fmt.Println(n, err)

	if n >= 1 {
		fmt.Println("数据插入成功")
	}

	//切片数据 这里的users append了两个Users
	var users []User
	users = append(users, User{
		Id:      12,
		Name:    "1",
		Age:     0,
		Passwd:  "2",
		Created: time.Time{},
		Updated: time.Time{},
	})
	users = append(users, User{
		Id:      13,
		Name:    "1",
		Age:     0,
		Passwd:  "2",
		Created: time.Time{},
		Updated: time.Time{},
	})
	n, _ = engine.Insert(&user)

	/* 3. 更新和删除数据 */
	//update id 为 1 的 Name 为 as .Update(&user)
	user33 := User{Name: "as", Age: 80}
	n33, _ := engine.ID(1).Update(&user33)
	fmt.Println(n33)

	user = User{Name: "lf111111"}
	//把Id10002 name 改为 lf111111
	affected, err := engine.Update(&user, &User{Id: 10002})
	fmt.Println(affected)

	_, _ = engine.Exec("update user set age = ? where id = ?", 10, 10002)

	// delete id = 10003 && name = lf .Delete(&user)
	user = User{Name: "lf"}
	_, _ = engine.ID(10003).Delete(&user)

	user = User{Name: "lf"}
	//删除 name = lf Id = 3 的数据
	affected, err = engine.Where(`Id = 3`).Delete(&user)
	fmt.Println(affected)

	/* 4. 查询数据 */
	// Query查询 会返回结果集
	//返回值是byte
	results, err := engine.Query("select * from user")
	fmt.Println(results)
	//返回值是string
	results2, err := engine.QueryString("select * from user")
	fmt.Println(results2)
	//返回值是interface
	results3, err := engine.QueryInterface("select * from user")
	fmt.Println(results3)

	//Get 获取某一条数据 数据直接保存到user的结构体中
	user41 := User{}
	_, _ = engine.Get(&user41)
	fmt.Println(user41)

	//指定条件来查询用户 Name: "lf" 是约束条件
	user42 := User{Name: "lf"}
	_, _ = engine.Where("passwd=?", 1234).Desc("id").Get(&user42)
	fmt.Println(user42)

	//获取指定字段值 此时只打印name = 3
	var name string
	_, _ = engine.Table(&user).Where("id = ?", 4).Cols("name").Get(&name)
	fmt.Println(name)

	var users4 []User
	//limit start 可以设置分页 查询多个使用find
	_ = engine.Where("passwd=1234").And("age=22").Limit(10, 0).Find(&users4)
	fmt.Println(users4)

	//Count 获取记录条数
	user = User{Passwd: "1234"}
	count, err := engine.Count(&user)
	fmt.Println(count)

	//Iterate 和 Rows 根据条件遍历数据
	_ = engine.Iterate(&User{Passwd: "1234"}, func(idx int, bean interface{}) error {
		users := bean.(*User)
		fmt.Println(users)
		return nil
	})

	fmt.Println("==================")
	rows, err := engine.Rows(&User{Passwd: "1234"})
	defer rows.Close()
	userBean := new(User) //scan需要一个指针 通过new一个结构体来实现
	for rows.Next() {     //bool类型
		_ = rows.Scan(userBean) //保存到user里面
		fmt.Println(userBean)
	}

	/* 5. 事务处理
	一堆语句执行时只要有一个报错就整体不提交
	如果事务中的某个点发生故障，则所有更新都可以回滚到事务开始之前的状态。如果没有发生故障，则通过以完成状态提交事务来完成更新。
	panic recover
	*/
	//事务
	session := engine.NewSession()
	defer session.Close()

	//开启事务
	_ = session.Begin()

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Rollback")
			fmt.Println(err)
			//回滚
			_ = session.Rollback()
		} else {
			_ = session.Commit()
		}
	}()

	user51 := User{Id: 10005, Name: "Alice", Age: 18, Passwd: "1234"}
	if _, err := session.Insert(&user51); err != nil {
		panic(err)
	}

	user52 := User{Name: "3", Age: 22, Passwd: "1234"}
	if _, err := session.Where("id=100000000").Update(&user52); err != nil {
		panic(err)
	}

	if _, err := session.Exec("delete from user where name = '7'"); err != nil {
		panic(err)
	}

}
