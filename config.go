package nflog

type Config struct {
	Groups []uint16
    PayloadLength uint32
	Return struct {
		Errors bool
	}
}

func NewConfig() *Config {
    cfg := &Config{}
    cfg.PayloadLength = 0x40000000

    return cfg
}

func (c Config) Validate() error {
	if len(c.Groups) == 0 {
		return ConfigurationError("No groups defined")
	}

	if len(c.Groups) > 32 {
		return ConfigurationError("Number of groups should be <= 32")
	}

	return nil
}
