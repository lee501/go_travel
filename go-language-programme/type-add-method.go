// 定义类型Integer
type Integer int 

// 面向对象的实现，增加方法Less()
func (a Integer) Less(b Integer) bool{
  return a < b
}

// 面向过程的实现方法
func Integer_Less(a, b Integer) bool{
  return a < b
}

// go语音类型是基于值传递的，如果需要修改变量的值，只能传递指针
// 需要改变对象的时候，使用指针
func (a *Integer) Add(b Integer){
  *a += b
}
