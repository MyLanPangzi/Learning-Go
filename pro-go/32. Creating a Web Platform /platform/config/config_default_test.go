package config

import (
	"reflect"
	"testing"
)

type args struct {
	name string
}

func TestDefaultConfig_get(t *testing.T) {
	tests := []struct {
		name      string
		args      args
		wantValue any
		wantFound bool
	}{
		{"get bool", args{"section:bool"}, true, true},
		{"get int", args{"section:int"}, 1.0, true},
		{"get float", args{"section:float"}, 1.0, true},
		{"get string", args{"section:string"}, "1", true},
		{"get section", args{"section"}, map[string]any{
			"int":    1.0,
			"float":  1.0,
			"string": "1",
			"bool":   true,
		}, true},
	}
	c, err := Load("testdata/config.json")
	if err != nil {
		t.Error(err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue, gotFound := c.get(tt.args.name)
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("get() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
			if gotFound != tt.wantFound {
				t.Errorf("get() gotFound = %v, want %v", gotFound, tt.wantFound)
			}
		})
	}
}

func TestDefaultConfig_GetBoolDefault(t *testing.T) {

	tests := []struct {
		name      string
		args      args
		wantValue bool
	}{
		{"GetBoolDefault", args{"bool"}, true},
	}
	c, err := Load("testdata/config.json")
	if err != nil {
		t.Error(err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue := c.GetBoolDefault(tt.args.name, true)
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("GetBoolDefault() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestDefaultConfig_GetFloatDefault(t *testing.T) {
	tests := []struct {
		name      string
		args      args
		wantValue float64
	}{
		{"GetBoolDefault", args{"float"}, 1},
	}
	c, err := Load("testdata/config.json")
	if err != nil {
		t.Error(err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue := c.GetFloatDefault(tt.args.name, 1)
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("GetFloatDefault() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestDefaultConfig_GetIntDefault(t *testing.T) {
	tests := []struct {
		name      string
		args      args
		wantValue int
	}{
		{"GetIntDefault", args{"int"}, 1},
	}
	c, err := Load("testdata/config.json")
	if err != nil {
		t.Error(err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue := c.GetIntDefault(tt.args.name, 1)
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("GetIntDefault() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestDefaultConfig_GetStringDefault(t *testing.T) {
	tests := []struct {
		name      string
		args      args
		wantValue string
	}{
		{"GetStringDefault", args{"string"}, "1"},
	}
	c, err := Load("testdata/config.json")
	if err != nil {
		t.Error(err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotValue := c.GetStringDefault(tt.args.name, "1")
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("GetStringDefault() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}

func TestLoad(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		want    *DefaultConfig
		wantErr bool
	}{
		{name: "load", args: args{"testdata/config.json"}, want: &DefaultConfig{map[string]any{
			"section": map[string]any{
				"int":    1.0,
				"float":  1.0,
				"string": "1",
				"bool":   true,
			},
		}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Load(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Load() got = %v, want %v", got, tt.want)
			}
		})
	}
}
