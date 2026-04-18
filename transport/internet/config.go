package internet

import (
	"github.com/xtls/xray-core/common/errors"
	"github.com/xtls/xray-core/common/serial"
)

type ConfigCreator func() interface{}

var (
	globalTransportConfigCreatorCache = make(map[string]ConfigCreator)
)

const unknownProtocol = "unknown"

func RegisterProtocolConfigCreator(name string, creator ConfigCreator) error {
	if _, found := globalTransportConfigCreatorCache[name]; found {
		return errors.New("protocol ", name, " is already registered").AtError()
	}
	globalTransportConfigCreatorCache[name] = creator
	return nil
}

// Note: Each new transport needs to add init() func in transport/internet/xxx/config.go
// Otherwise, it will cause #3244
func CreateTransportConfig(name string) (interface{}, error) {
	creator, ok := globalTransportConfigCreatorCache[name]
	if !ok {
		return nil, errors.New("unknown transport protocol: ", name)
	}
	return creator(), nil
}

func (c *TransportConfig) GetTypedSettings() (interface{}, error) {
	return c.Settings.GetInstance()
}

func (c *TransportConfig) GetUnifiedProtocolName() string {
	return c.ProtocolName
}

func (c *StreamConfig) GetEffectiveProtocol() string {
	if c == nil || len(c.ProtocolName) == 0 {
		return "tcp"
	}

	return c.ProtocolName
}

func (c *StreamConfig) GetEffectiveTransportSettings() (interface{}, error) {
	protocol := c.GetEffectiveProtocol()
	return c.GetTransportSettingsFor(protocol)
}

func (c *StreamConfig) GetTransportSettingsFor(protocol string) (interface{}, error) {
	if c != nil {
		for _, settings := range c.TransportSettings {
			if settings.GetUnifiedProtocolName() == protocol {
				return settings.GetTypedSettings()
			}
		}
	}

	return CreateTransportConfig(protocol)
}

func (c *StreamConfig) GetEffectiveSecuritySettings() (interface{}, error) {
	for _, settings := range c.SecuritySettings {
		if settings.Type == c.SecurityType {
			return settings.GetInstance()
		}
	}
	return serial.GetInstance(c.SecurityType)
}

func (c *StreamConfig) HasSecuritySettings() bool {
	return len(c.SecuritySettings) > 0
}

func (c *ProxyConfig) HasTag() bool {
	return c != nil && len(c.Tag) > 0
}
