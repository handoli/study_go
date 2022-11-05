package main

import (
	"encoding/json"
	"fmt"
)

//嵌套结构体
type mh struct{
	server string
	m mhsy
}

type mhsy struct{
	name string
	grade int
	sect string
	score int
}

type Address struct {
    Province string
    City     string
}

//User 用户结构体
type User struct {
    Name    string
    Gender  string
    Address //匿名结构体
}

//结构体匿名字段
type person struct{
	string
	int
}

func main() {
	//值类型
	var yg mhsy
	yg.name = "时空之隙"
	yg.grade = 69
	yg.sect = "月宫"
	yg.score = 45678
	fmt.Printf("%#v\n", yg)
	fmt.Printf("%T\n", yg)
	fmt.Println("------------------------------")

	// 指针类型
	var mw = new(mhsy)
	mw.name = "三味真火"
	mw.grade = 115
	mw.sect = "魔王"
	mw.score = 67890
	fmt.Printf("%#v\n", mw)
	fmt.Printf("%T\n",mw)
	fmt.Println(mw.name)
	fmt.Println(&mw.name)
	fmt.Println(&(mw.name))
	fmt.Println(&(mw).name)
	fmt.Println(&mw.grade)
	fmt.Println(&mw.sect)
	fmt.Println(&mw.score)
	fmt.Println("------------------------------")

	//使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作
	var hgs = &mhsy{}
	hgs.name = "火眼金睛"
	hgs.grade = 115
	hgs.sect = "花果山"
	hgs.score = 67890
	fmt.Printf("%#v\n", hgs)
	fmt.Printf("%T\n",hgs)
	fmt.Println("------------------------------")

	//匿名结构体
	var cat struct{
		name string
		age int
	}
	cat.name = "mao"
	cat.age = 1
	fmt.Printf("%#v\n",cat) 
	fmt.Println("------------------------------")

	//键值对初始化 ，值类型
	ly := mhsy{
		name: "一念成佛",
		grade: 89,
		sect: "雷音寺",
	}
	fmt.Println(ly)
	fmt.Println("------------------------------")

	//键值对初始化 ，指针类型
	lg := &mhsy{
		name: "双龙戏珠",
		sect: "龙宫",
		score: 55678,
	}
	fmt.Println(lg)
	fmt.Println("------------------------------")

	//结构体参数列表初始化，也可值类型， 也可指针类型
	fc := &mhsy{
		"五雷咒",
		69,
		"方寸山",
		43526,
	}
	fmt.Println(fc)
	fmt.Println("------------------------------")

	p := &person{
		"hdl",
		28,
	}
	fmt.Println(p)
	fmt.Println("------------------------------")

	//嵌套结构体,初始化一
	m := &mh{
		server : "2018",
		m : mhsy{
			name: "唧唧歪歪",
			grade: 99,
			sect: "化生寺",
			score: 67777,
		},
	}
	fmt.Println(m)
	fmt.Println("------------------------------")
	//嵌套结构体,初始化二（匿名）
	//匿名字段默认采用类型名作为字段名，结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个
	user := new(User)
	user.Name = "pprof"
    user.Gender = "女"
    user.Address.Province = "黑龙江"    //通过匿名结构体.字段名访问
    user.City = "哈尔滨"                //直接访问匿名结构体的字段名
    fmt.Printf("user2=%#v\n", user)
	fmt.Println(user)
	fmt.Println("------------------------------")
	//嵌套匿名结构体，
	var x struct{
		a string
		b string
		c struct{
			d string
		}
	}
	x.a = "1"
	x.b = "2"
	x.c.d = "4"
	fmt.Println(x)
	fmt.Println("------------------------------")
	//结构体继承属性和方法（嵌套）
	dog := &Dog{
		5,
		&Animal{
			"小花",
		},
	}
	dog.move()
	dog.wang()
	fmt.Println(dog)
	fmt.Println("------------------------------")

	//结构体与json,不能序列化私有属性（属性首字母小写的）
	c := &Class{
        Title:    "101",
        Students: make([]*Student, 0, 200),
    }
    for i := 0; i < 10; i++ {
        stu := &Student{
            Name:   fmt.Sprintf("stu%02d", i),
            Gender: "男",
            ID:     i,
        }
        c.Students = append(c.Students, stu)
    }
	//结构体序列化成json
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("结构体序列化成json出错！！！")
		return
	}
	fmt.Printf("json:%s\n", data)

	//json反序列化成结构体
	c1 := &Class{}
	err = json.Unmarshal([]byte(data), c1)
	if err != nil {
		fmt.Println("结构体序列化成json出错！！！")
		return
	}
	fmt.Printf("%#v\n", c1)
	fmt.Println("------------------------------")

	//结构体标签（Tag）
	s1 := StudentTag{
        ID:     1,
        Gender: "女",
        name:   "pprof",
    }
    data, err = json.Marshal(s1)
    if err != nil {
        fmt.Println("json marshal failed!")
        return
    }
    fmt.Printf("json str:%s\n", data) //json str:{"id":1,"Gender":"女"}

 }

//Animal 动物
type Animal struct {
    name string
}

func (a *Animal) move() {
    fmt.Printf("%s会动！\n", a.name)
}

//Dog 狗
type Dog struct {
    Feet    int8
    *Animal //通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
    fmt.Printf("%s会汪汪汪~\n", d.name)
}

//Student 学生
type Student struct {
    ID     int
    Gender string
    Name   string
}

//Class 班级
type Class struct {
    Title    string
    Students []*Student
}

type StudentTag struct {
    ID     int    `json:"id"` //通过指定tag实现json序列化该字段时的key
    Gender string //json序列化是默认使用字段名作为key
    name   string //私有不能被json包访问
}