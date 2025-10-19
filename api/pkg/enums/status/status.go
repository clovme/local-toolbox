package status

import (
	"sort"
	"toolbox/pkg/enums"
)

type Status int

const Name = "status"

const (
	Disable Status = iota
	Enable
)

var (
	initiate = map[Status]enums.Enums{
		Enable:  {Key: "enable", Name: "启用", Desc: "启用"},
		Disable: {Key: "disable", Name: "禁用", Desc: "禁用"},
	}

	enumToValue = make(map[string]Status)
)

func init() {
	for code, meta := range initiate {
		enumToValue[meta.Key] = code
	}
}

// Key 获取enums.Key
func (r Status) Key() string {
	return initiate[r].Key
}

// Name 获取枚举名称
func (r Status) Name() string {
	return initiate[r].Name
}

// Desc 获取枚举描述
func (r Status) Desc() string {
	return initiate[r].Desc
}

// Int 获取枚举值
func (r Status) Int() int {
	return int(r)
}

// Is 比较枚举值
func (r Status) Is(v Status) bool {
	return v == r
}

// Code 获取Status
func Code(key string) Status {
	if enum, ok := enumToValue[key]; ok {
		return enum
	}
	return Disable
}

// ValueList 获取所有枚举列表
func ValueList() []Status {
	values := make([]Status, 0, len(initiate))
	for k := range initiate {
		values = append(values, k)
	}
	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})
	return values
}

// ValueMap 获取所有枚举Map（安全副本）
func ValueMap() map[Status]enums.Enums {
	copyMap := make(map[Status]enums.Enums, len(initiate))
	for k, v := range initiate {
		copyMap[k] = v
	}
	return copyMap
}
