package constructor

import "sync"

type Singleton struct {

}

// 饿汉式
var singleton *Singleton

func init(){
	singleton = &Singleton{}
}

func GetInstance()*Singleton{
	return singleton
}

// 懒汉式
// 在获取时才创建对象，通过once保证对象创建唯一
var (
	lazySingleton *Singleton
	once sync.Once
)

func GetLazyInstance()*Singleton{
	once.Do(func() {
		lazySingleton = &Singleton{}
	})
	return lazySingleton
}