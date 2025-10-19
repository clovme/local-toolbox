package icon

import (
	"sort"
	"toolbox/pkg/enums"
)

type Icon int

const Name = "图标类型"

const (
	Font Icon = iota
	Image
	Unknown
)

var (
	initiate = map[Icon]enums.Enums{
		Unknown: {Key: "Unknown", Name: "未知", Desc: "未知图标"},
		Image:   {Key: "Image", Name: "图片", Desc: "图片图标"},
		Font:    {Key: "Font", Name: "字体", Desc: "字体图标"},
	}

	enumToValue = make(map[string]Icon)
)

func init() {
	for code, meta := range initiate {
		enumToValue[meta.Key] = code
	}
}

// Int 获取枚举值
func (r Icon) Int() int {
	return int(r)
}

// Key 获取enums.Key
func (r Icon) Key() string {
	return initiate[r].Key
}

// Name 获取枚举名称
func (r Icon) Name() string {
	return initiate[r].Name
}

// Desc 获取枚举描述
func (r Icon) Desc() string {
	return initiate[r].Desc
}

// Enum 获取枚举值
func (r Icon) Enum() int {
	return int(r)
}

// Is 比较枚举值
func (r Icon) Is(v Icon) bool {
	return v == r
}

// Code 获取Icon
func Code(key string) Icon {
	if enum, ok := enumToValue[key]; ok {
		return enum
	}
	return Unknown
}

// ValueList 获取所有枚举列表
func ValueList() []Icon {
	values := make([]Icon, 0, len(initiate))
	for k := range initiate {
		values = append(values, k)
	}
	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})
	return values
}

// ValueMap 获取所有枚举Map（安全副本）
func ValueMap() map[Icon]enums.Enums {
	copyMap := make(map[Icon]enums.Enums, len(initiate))
	for k, v := range initiate {
		copyMap[k] = v
	}
	return copyMap
}
