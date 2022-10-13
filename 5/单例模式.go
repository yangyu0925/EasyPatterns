package main

import (
	"fmt"
	"sync"
)

//饿汉式
//type single struct {}
//
//var instance *single = new(single)
//
//func GetInstance() *single {
//	return instance
//}
//
//func (s *single) SomeThing()  {
//	fmt.Println("单例对象的某方法")
//}
//
//func main() {
//	s := GetInstance()
//	s.SomeThing()
//}

//懒汉式
//type single struct {}
//
//var instance *single
//
//func GetInstance() *single {
//	if instance == nil {
//		instance = new(single)
//		return instance
//	}
//
//	return instance
//}
//
//func (s *single) SomeThing()  {
//	fmt.Println("单例对象的某方法")
//}
//func main() {
//	s := GetInstance()
//	s.SomeThing()
//}

//定义锁
//var lock sync.Mutex
//
//type single struct {}
//
//var instance *single
//
//func GetInstance() *single {
//	lock.Lock()
//	defer lock.Unlock()
//
//	if instance == nil {
//		return new(single)
//	}
//
//	return instance
//}
//
//func (s *single) SomeThing()  {
//	fmt.Println("单例对象的某方法")
//}
//
//func main() {
//	s := GetInstance()
//	s.SomeThing()
//}


//标记
//var initialized uint32
//var lock sync.Mutex
//
//type single struct {}
//
//var instance *single
//
//func GetInstance() *single {
//	if atomic.LoadUint32(&initialized) == 1 {
//		return instance
//	}
//
//	lock.Lock()
//	defer lock.Unlock()
//
//	if initialized == 0 {
//		instance = new(single)
//
//		atomic.StoreUint32(&initialized, 1)
//	}
//
//	return instance
//}
//
//func (s *single) SomeThing()  {
//	fmt.Println("单例对象的某方法")
//}
//
//func main() {
//	s := GetInstance()
//	s.SomeThing()
//}


//Once

var once sync.Once

type single struct {}

var instance *single

func GetInstance() *single {

	once.Do(func() {
		instance = new(single)
	})

	return instance
}

func (s *single) SomeThing()  {
	fmt.Println("单例对象的某方法")
}

func main() {
	s := GetInstance()
	s.SomeThing()
}