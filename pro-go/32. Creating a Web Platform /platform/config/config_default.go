package config

import "strings"

type DefaultConfig struct {
	configData map[string]any
}

func (c *DefaultConfig) get(name string) (any, bool) {
	data := c.configData
	var result any
	var found bool
	for _, k := range strings.Split(name, ":") {
		result, found = data[k]
		if section, ok := result.(map[string]any); ok && found {
			data = section
		}
	}
	return result, found
}
func (c *DefaultConfig) GetString(name string) (value string, found bool) {
	if result, found := c.get(name); found {
		return result.(string), true
	}
	return "", found
}

func (c *DefaultConfig) GetInt(name string) (value int, found bool) {
	if result, found := c.get(name); found {
		return int(result.(float64)), true
	}
	return 0, found
}

func (c *DefaultConfig) GetBool(name string) (value bool, found bool) {
	if result, found := c.get(name); found {
		return result.(bool), true
	}
	return false, found
}

func (c *DefaultConfig) GetFloat(name string) (value float64, found bool) {
	if result, found := c.get(name); found {
		return result.(float64), true
	}
	return 0, found
}

func (c *DefaultConfig) GetStringDefault(name string, defaultValue string) (value string) {
	if result, found := c.GetString(name); found {
		return result
	}
	return defaultValue
}

func (c *DefaultConfig) GetIntDefault(name string, defaultValue int) (value int) {
	if result, found := c.GetInt(name); found {
		return result
	}
	return defaultValue
}

func (c *DefaultConfig) GetBoolDefault(name string, defaultValue bool) (value bool) {
	if result, found := c.GetBool(name); found {
		return result
	}
	return defaultValue
}

func (c *DefaultConfig) GetFloatDefault(name string, defaultValue float64) (value float64) {
	result, found := c.GetFloat(name)
	if found {
		return result
	}
	return defaultValue
}

func (c *DefaultConfig) GetSection(name string) (section Configuration, found bool) {
	if result, f := c.get(name); f {
		if s, ok := result.(map[string]any); ok {
			return &DefaultConfig{s}, true
		}
	}
	return nil, false
}
