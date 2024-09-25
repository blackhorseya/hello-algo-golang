package chapter_hashing

import (
	"fmt"
)

type pair struct {
	key int
	val string
}

type arrayHashMap struct {
	buckets []*pair
}

func newArrayHashMap() *arrayHashMap {
	return &arrayHashMap{
		buckets: make([]*pair, 100),
	}
}

/* 雜湊函式 */
func (a *arrayHashMap) hashFunc(key int) int {
	index := key % 100
	return index
}

/* 查詢操作 */
func (a *arrayHashMap) get(key int) string {
	index := a.hashFunc(key)
	e := a.buckets[index]
	if e == nil {
		return "Not Found"
	}
	return e.val
}

/* 新增操作 */
func (a *arrayHashMap) put(key int, val string) {
	pair := &pair{key: key, val: val}
	index := a.hashFunc(key)
	a.buckets[index] = pair
}

/* 刪除操作 */
func (a *arrayHashMap) remove(key int) {
	index := a.hashFunc(key)
	// 置為 nil ，代表刪除
	a.buckets[index] = nil
}

/* 獲取所有鍵對 */
func (a *arrayHashMap) pairSet() []*pair {
	var pairs []*pair
	for _, pair := range a.buckets {
		if pair != nil {
			pairs = append(pairs, pair)
		}
	}
	return pairs
}

/* 獲取所有鍵 */
func (a *arrayHashMap) keySet() []int {
	var keys []int
	for _, pair := range a.buckets {
		if pair != nil {
			keys = append(keys, pair.key)
		}
	}
	return keys
}

/* 獲取所有值 */
func (a *arrayHashMap) valueSet() []string {
	var values []string
	for _, pair := range a.buckets {
		if pair != nil {
			values = append(values, pair.val)
		}
	}
	return values
}

/* 列印雜湊表 */
func (a *arrayHashMap) print() {
	for _, pair := range a.buckets {
		if pair != nil {
			fmt.Println(pair.key, "->", pair.val)
		}
	}
}
