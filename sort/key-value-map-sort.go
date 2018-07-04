package main

import (
	"fmt"
	"sort"
)

func MapToKVs(m map[string]string) KVs {
	//	kvs := make(KVs, 0)
	//	for k, v := range m {
	//		kvs = append(kvs, KV{K: k, V: v})
	//	}
	kvs := make(KVs, len(m))
	i := 0
	for k, v := range m {
		kvs[i] = KV{K: k, V: v}
		i++
	}
	return kvs
}

// 独自の構造体
type KV struct {
	K string
	V string
}

type KVs []KV

type KVSorter struct {
	KVs
	Comp func(k1, k2 KV) bool
}

func (k KVSorter) Len() int {
	return len(k.KVs)
}

func (k KVSorter) Swap(i, j int) {
	k.KVs[i], k.KVs[j] = k.KVs[j], k.KVs[i]
}

func (k KVSorter) Less(i, j int) bool {
	return k.Comp(k.KVs[i], k.KVs[j])
}

//	key 昇順比較
func (k KVSorter) SetKeyAscComp() KVSorter {
	k.Comp = func(k1, k2 KV) bool {
		return k1.K < k2.K
	}
	return k
}

//	value 昇順比較
func (k KVSorter) SetValueAscComp() KVSorter {
	k.Comp = func(k1, k2 KV) bool {
		return k1.V < k2.V
	}
	return k
}

func main() {
	m := map[string]string{
		"Nanoha": "Raising Heart",
		"Hayate": "Schwertkreuz",
		"Fate":   "Bardiche",
	}

	kvs := MapToKVs(m)

	fmt.Println("before sort")
	fmt.Println(kvs)
	//	スライスの中身をコピーで入れ替えてると推測される
	sort.Sort((KVSorter{KVs: kvs}).SetKeyAscComp())
	fmt.Println("after sort [key asc]")
	fmt.Println(kvs)
	sort.Sort((KVSorter{KVs: kvs}).SetValueAscComp())
	fmt.Println("after sort [value asc]")
	fmt.Println(kvs)
}
