package alwaysDel

import (
	"fmt"
	"testing"
	"time"
)

type Car struct {
	Name string
	Age  int
}

func (c *Car) Set(name string, age int) {
	c.Name = name
	c.Age = age
}

type Car2 struct {
	Name string
}

//Go有匿名字段特性
type Train struct {
	Car
	Car2
	createTime time.Time
	//count int   正常写法，Go的特性可以写成
	int
}

//给Train加方法，t指定接受变量的名字，变量可以叫this，t，p
func (t Train) Set(age int) {
	t.int = age
}

func TestIt(t *testing.T) {
	//var train Train
	//train.int = 300 //这里用的匿名字段写法，给Age赋值
	////(&train).Set(1000)
	//train.Car.Set("huas", 100)
	//train.Car.Name = "test" //这里Name必须得指定结构体
	//fmt.Println(train)

	var train Train
	fmt.Println(train.int)

}

type data struct {
	name string
}

type printer interface {
	print()
}

func (p *data) print() {
	fmt.Println("name: ", p.name)
	p.name = "new data..."
}

func TestIts(t *testing.T) {
	d1 := data{"one"}
	d1.print() // d1 变量可寻址，可直接调用指针 receiver 的方法

	var in printer = &data{"two"}
	in.print() // 类型不匹配

	m := map[string]data{
		"x": data{"three"},
	}
	d := m["x"]
	d.print()
	fmt.Println(d)
	m["x"].print() // m["x"] 是不可寻址的	// 变动频繁

}
