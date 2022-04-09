package contracts

import (
	"reflect"
	"sort"
)

type Context interface {
	// Get 从上下文中检索数据
	// retrieves data from the context.
	Get(key string) interface{}

	// Set 在上下文中保存数据
	// saves data in the context.
	Set(key string, val interface{})
}

type Fields map[string]any

type Interface interface {
	reflect.Type
	GetType() reflect.Type

	IsSubClass(class interface{}) bool

	// ClassName 获取类名
	ClassName() string
}

type Class interface {
	Interface

	// New 通过 Fields
	New(fields Fields) interface{}

	NewByTag(fields Fields, tag string) interface{}
}

type FieldsProvider interface {
	Fields() Fields
}

type Json interface {
	ToJson() string
}

type Getter interface {
	GetString(key string) string
	GetInt64(key string) int64
	GetInt(key string) int
	GetFloat64(key string) float64
	GetFloat(key string) float32
	GetBool(key string) bool
	GetFields(key string) Fields
}

type OptionalGetter interface {
	StringOption(key string, defaultValue string) string
	Int64Option(key string, defaultValue int64) int64
	IntOption(key string, defaultValue int) int
	Float64Option(key string, defaultValue float64) float64
	FloatOption(key string, defaultValue float32) float32
	BoolOption(key string, defaultValue bool) bool
	FieldsOption(key string, defaultValue Fields) Fields
}

type Collection[T any] interface {
	Json
	// sort

	sort.Interface
	Sort(func(previous T, next T) bool) Collection[T]
	IsEmpty() bool

	// filter

	Map(func(item T, index int) T) Collection[T]       // 返回新的集合
	Each(func(item T, index int)) Collection[T]        // 纯遍历
	Filter(func(item T, index int) bool) Collection[T] // 返回 true 保留下来
	Skip(func(item T, index int) bool) Collection[T]   // 返回 true 会跳过

	Where(field string, args ...interface{}) Collection[T]
	WhereNil(field string) Collection[T]
	WhereNotNil(field string) Collection[T]
	WhereLt(field string, arg interface{}) Collection[T]
	WhereLte(field string, arg interface{}) Collection[T]
	WhereGt(field string, arg interface{}) Collection[T]
	WhereGte(field string, arg interface{}) Collection[T]
	WhereIn(field string, arg interface{}) Collection[T]
	WhereNotIn(field string, arg interface{}) Collection[T]

	// keys、values

	// Pluck 数据类型为 []map、[]struct 的时候起作用
	Pluck(key string) Fields
	// Only 数据类型为 []map、[]struct 的时候起作用
	Only(keys ...string) Collection[Fields]

	// First 获取首个元素, []struct或者[]map可以获取指定字段
	First() T

	// Last 获取最后一个元素, []struct或者[]map可以获取指定字段
	Last() T

	// FirstValue 第一个元素的指定字段
	FirstValue(field string) any
	// LastValue 最后一个元素的指定字段
	LastValue(field string) any

	// union、merge...

	// Prepend 从开头插入元素
	Prepend(item ...T) Collection[T]
	// Push 从最后插入元素
	Push(items ...T) Collection[T]
	// Pull 从尾部获取并移出一个元素
	Pull(defaultValue ...T) T
	// Shift 从头部获取并移出一个元素
	Shift(defaultValue ...T) T
	// Put 替换一个元素，如果 index 不存在会执行 Push，返回新集合
	Put(index int, item T) Collection[T]
	// Offset 替换一个元素，如果 index 不存在会执行 Push
	Offset(index int, item T) Collection[T]
	// Merge 合并其他集合
	Merge(collections ...Collection[T]) Collection[T]
	// Reverse 返回一个顺序翻转后的集合
	Reverse() Collection[T]
	// Chunk 分块，handler 返回 error 表示中断
	Chunk(size int, handler func(collection Collection[T], page int)) error
	// Random 随机返回n个元素，默认1个
	Random(size ...uint) Collection[T]

	// aggregate

	Sum(key ...string) float64
	Max(key ...string) float64
	Min(key ...string) float64
	Avg(key ...string) float64
	Count() int

	// convert

	ToIntArray() (results []int)
	ToInt64Array() (results []int64)
	ToInterfaceArray() []interface{}
	ToFloat64Array() (results []float64)
	ToFloatArray() (results []float32)
	ToBoolArray() (results []bool)
	ToStringArray() (results []string)
	ToFields() Fields
	ToArrayFields() []Fields
}
