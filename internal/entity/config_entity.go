package entity

import (
	"encoding/json"

	"github.com/answerdev/answer/pkg/converter"
)

// Config config
type Config struct {
	ID    int    `xorm:"not null pk autoincr INT(11) id"`
	Key   string `xorm:"unique VARCHAR(128) key"`
	Value string `xorm:"TEXT value"`
}

// TableName config table name
func (c *Config) TableName() string {
	return "config"
}

func (c *Config) BuildByJSON(data []byte) {
	cf := &Config{}
	_ = json.Unmarshal(data, cf)
	c.ID = cf.ID
	c.Key = cf.Key
	c.Value = cf.Value
}

func (c *Config) JsonString() string {
	data, _ := json.Marshal(c)
	return string(data)
}

// GetIntValue get int value
func (c *Config) GetIntValue() int {
	return converter.StringToInt(c.Value)
}

// GetArrayStringValue get array string value
func (c *Config) GetArrayStringValue() []string {
	var arr []string
	_ = json.Unmarshal([]byte(c.Value), &arr)
	return arr
}

// GetByteValue get byte value
func (c *Config) GetByteValue() []byte {
	return []byte(c.Value)
}
