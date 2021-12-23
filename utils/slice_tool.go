// Package utils
// @Description: 切片操作
package utils

import "errors"

// IsContain
// @description 很遗憾 切片的遍历应该只能是O(n)了
// @param list 切片
// @param elem 元素
// @return bool
func IsContain(list []interface{}, elem interface{}) bool {
	for _, val := range list {
		if elem == val {
			return true
		}
	}
	return false
}

// IsContainInt
// @description 判断切片是否存在指定元素
func IsContainInt(list []int, elem int) bool {
	for _, val := range list {
		if elem == val {
			return true
		}
	}
	return false
}

// DeleteByElemInt
// @description 删除切片中指定元素(首次出现)
// @return []int 删除元素后的切片
func DeleteByElemInt(list []int, elem int) ([]int, error) {
	index := -1
	for i, val := range list {
		if val == elem {
			index = i
			break
		}
	}
	if index < 0 {
		return list, errors.New("element not in splice")
	}
	return append(list[:index], list[index+1:]...), nil

}

// DeleteAllByElemInt
// @description 删除切片中指定元素(所有)
// @return []int 删除元素后的切片
func DeleteAllByElemInt(list []int, elem int) ([]int, error) {
	ret := make([]int, 0)
	for _, val := range list {
		if val != elem {
			ret = append(ret, val)
		}
	}
	if len(ret) == len(list) {
		return ret, errors.New("element not in splice")
	}
	return ret, nil
}

// DeleteByIndexInt
// @description 删除切片中指定下标对应的元素
// @return []int 删除元素后的切片
func DeleteByIndexInt(list []int, index int) ([]int, error) {
	if index < 0 || index >= len(list) {
		return list, errors.New("index out of range")
	}
	return append(list[:index], list[index+1:]...), nil
}

// GetIndexInt
// @description 获取元素在切片中的下标
func GetIndexInt(list []int, elem int) (int, error) {
	for i, e := range list {
		if elem == e {
			return i, nil
		}
	}
	return -1, errors.New("index out of range")
}
