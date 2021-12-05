package constructor

import (
	"encoding/json"
	"time"
)

// 原型模式：通过对已有对象的克隆，创出新的对象
// 场景：创建一个新对象开销很大：复杂运算、网络磁盘io

type Keyword struct {
	word      string
	visit     int
	UpdatedAt *time.Time
}

// Clone 使用序列化与反序列化的方式深拷贝
func (k *Keyword) Clone() *Keyword {
	var newK Keyword
	b,_:=json.Marshal(k)
	json.Unmarshal(b,&newK)
	return &newK
}

// Keywords 关键字 map
type Keywords map[string]*Keyword

// Clone 复制一个新的 keywords
// 对map进行拷贝
// updatedWords: 需要更新的关键词列表，由于从数据库中获取数据常常是数组的方式
func (words Keywords) Clone(updatedWords []*Keyword) Keywords {
	newKeywords := Keywords{}

	for k, v := range words {
		// 这里是浅拷贝，直接拷贝了地址
		newKeywords[k] = v
	}

	// 替换掉需要更新的字段，这里用的是深拷贝
	for _, word := range updatedWords {
		newKeywords[word.word] = word.Clone()
	}

	return newKeywords
}