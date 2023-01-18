package contracts

type ConfigProvider func(env Env) interface{}

type Config interface {
	Getter
	FieldsProvider

	// Load 根据给定的字段提供者加载配置
	// load configuration based on given field provider.
	Load(provider FieldsProvider)

	// Reload 根据给定的字段提供者加载配置
	// reload configuration based on given field provider.
	Reload()

	// Get 获取指定的配置值
	// get the specified configuration value.
	Get(key string, defaultValue ...interface{}) interface{}

	// Set 设置给定的配置值
	// set a given configuration value.
	Set(key string, value interface{})

	// Unset 销毁指定的配置值
	// Destroy the specified configuration value.
	Unset(key string)
}

type Env interface {
	Getter
	OptionalGetter

	FieldsProvider

	// Load 加载配置
	// load configuration.
	Load() Fields
}
