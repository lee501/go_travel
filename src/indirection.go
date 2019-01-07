// 指针为接收者的方法, 接受者乐意为值也可以为指针， Go 会将语句 v.Scale(5) 解释为 (&v).Scale(5)。
package main 

import "fmt"

type Metu struct{
  X, Y float64
}

// 方法，指针接受者
func (v *Metu) Scale(f float64){
  v.X = v.X * f
  v.Y = v.Y * f
}

// 函数
func ScaleFunc(v *Metu, f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}

func main() {
  v := Metu{3, 4}
  // 接受者可以为值
  v.Scale(2)
  // 函数指针参数
  ScaleFunc(&v, 10)

  // p为指针
  p := &Metu{3, 4}
  p.Scale(3)
  ScaleFunc(p, 10)

  fmt.Println(v, p)
}
