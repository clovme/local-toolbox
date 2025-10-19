package position

import (
	"sort"
	"toolbox/pkg/enums"
)

type Position int

const Name = "数据位置类型"

const (
	Body Position = iota + 0
	Sidebar
	Unknown
)

var (
	initiate = map[Position]enums.Enums{
		Body:    {Key: "body", Name: "主体", Desc: "主体位置"},
		Sidebar: {Key: "sidebar", Name: "侧边栏", Desc: "侧边栏位置"},
		Unknown: {Key: "Unknown", Name: "未知", Desc: "未知值"},
	}

	enumToValue = make(map[string]Position)
)

func init() {
	for code, meta := range initiate {
		enumToValue[meta.Key] = code
	}
}

// Int 获取枚举值
func (r Position) Int() int {
	return int(r)
}

// Key 获取enums.Key
func (r Position) Key() string {
	return initiate[r].Key
}

// Name 获取枚举名称
func (r Position) Name() string {
	return initiate[r].Name
}

// Desc 获取枚举描述
func (r Position) Desc() string {
	return initiate[r].Desc
}

// Enum 获取枚举值
func (r Position) Enum() int {
	return int(r)
}

// Is 比较枚举值
func (r Position) Is(v Position) bool {
	return v == r
}

// Code 获取Position
func Code(key string) Position {
	if enum, ok := enumToValue[key]; ok {
		return enum
	}
	return Unknown
}

// ValueList 获取所有枚举列表
func ValueList() []Position {
	values := make([]Position, 0, len(initiate))
	for k := range initiate {
		values = append(values, k)
	}
	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})
	return values
}

// ValueMap 获取所有枚举Map（安全副本）
func ValueMap() map[Position]enums.Enums {
	copyMap := make(map[Position]enums.Enums, len(initiate))
	for k, v := range initiate {
		copyMap[k] = v
	}
	return copyMap
}
