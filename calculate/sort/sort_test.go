package sort_test

import (
	"sort"
	"testing"
)

func TestSortOpt(t *testing.T) {
	data := []int{4, 2, 9, 5, 1, 6, 3, 8, 7}
	t.Log("Before sorting:", data)

	// 排序
	sort.Ints(data)

	t.Log("After sorting:", data)
	if data[0] != 1 || data[len(data)-1] != 9 {
		t.Error("Expected sorted array [1, 2, 3, 4, 5, 6, 7, 8, 9], but got", data)
	}

	// 反序排序
	sort.Slice(data, func(i, j int) bool { return data[i] > data[j] })
	t.Log("After reverse sorting:", data)
	if data[0] != 9 || data[len(data)-1] != 1 {
		t.Error("Expected reversed sorted array [9, 8, 7, 6, 5, 4, 3, 2, 1], but got", data)
	}
	// 按字符串排序
	dataStr := []string{"apple", "banana", "cherry", "date", "elderberry"}
	t.Log("Before sorting strings:", dataStr)

	sort.Strings(dataStr)

	t.Log("After sorting strings:", dataStr)
	if dataStr[0] != "apple" || dataStr[len(dataStr)-1] != "elderberry" {
		t.Error("Expected sorted string array ['apple', 'banana', 'cherry', 'date', 'elderberry'], but got", dataStr)
	}

	// 按 map排序
	dataMap := map[string]int{"apple": 5, "banana": 3, "cherry": 4, "date": 1, "elderberry": 2}
	t.Log("Before sorting map:", dataMap)

	keys := make([]string,
		0, len(dataMap))
	for k := range dataMap {
		keys = append(keys, k)
	}

}
