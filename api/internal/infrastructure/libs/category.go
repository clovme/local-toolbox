package libs

import (
	"strconv"
	"toolbox/internal/schema/vo/category"
)

func GetCategoryWithChildren(all []category.ApiCategoryVO, id int64) ([]int64, error) {
	var result []int64

	// 建一个 map 加快查找
	childrenMap := make(map[int64][]int64)
	for _, c := range all {
		id_, _ := strconv.ParseInt(c.ID, 10, 64)
		pid_, _ := strconv.ParseInt(c.Pid, 10, 64)
		childrenMap[pid_] = append(childrenMap[pid_], id_)
	}

	// 递归收集节点
	var collect func(int64)
	collect = func(pid int64) {
		for _, child := range childrenMap[pid] {
			result = append(result, child)
			collect(child)
		}
	}

	result = append(result, id)

	// 再查子节点
	collect(id)

	return result, nil
}
