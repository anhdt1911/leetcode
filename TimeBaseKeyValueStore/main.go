package main

import "fmt"

// key, value, timestamp, which one is unique?
// key is unique, each key will have a value in given timestamp
// so how to store value by timestamp
// can store by another map but couldn't get timestamp_prev
// store by an array of struct

type Value struct {
	value     string
	timestamp int
}

type TimeMap struct {
	timeMap map[string][]Value
}

func Constructor() TimeMap {
	return TimeMap{
		timeMap: make(map[string][]Value),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, ok := this.timeMap[key]; !ok {
		this.timeMap[key] = []Value{
			{value, timestamp},
		}
	} else {
		this.timeMap[key] = append(this.timeMap[key], Value{value, timestamp})
	}
}

func (this *TimeMap) Get(key string, timestamp int) string {
	keyArr, ok := this.timeMap[key]
	if !ok {
		return ""
	}
	if timestamp < keyArr[0].timestamp {
		return ""
	}
	left, right := 0, len(keyArr)-1
	for left < right {
		mid := (left + right) / 2
		if timestamp > keyArr[mid].timestamp {
			left = mid + 1
		} else if timestamp < keyArr[mid].timestamp {
			right = mid
		} else {
			return keyArr[mid].value
		}
	}
	if keyArr[left].timestamp > timestamp {
		return keyArr[left-1].value
	}
	return keyArr[left].value
}

func main() {
	tm := Constructor()
	tm.Set("love", "high", 10)
	tm.Set("love", "low", 20)
	fmt.Println(tm.Get("love", 5))
	fmt.Println(tm.Get("love", 10))
	fmt.Println(tm.Get("love", 15))
	fmt.Println(tm.Get("love", 20))
	fmt.Println(tm.Get("love", 25))
}
