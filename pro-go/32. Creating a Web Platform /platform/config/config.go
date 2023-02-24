package config

type Configuration interface {
	GetString(name string) (value string, found bool)
	GetInt(name string) (value int, found bool)
	GetBool(name string) (value bool, found bool)
	GetFloat(name string) (value float64, found bool)
	GetStringDefault(name string, defaultValue string) (value string)
	GetIntDefault(name string, defaultValue int) (value int)
	GetBoolDefault(name string, defaultValue bool) (value bool)
	GetFloatDefault(name string, defaultValue float64) (value float64)
	GetSection(name string) (section Configuration, found bool)
}
