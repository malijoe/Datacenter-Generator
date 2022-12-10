package hardware

type PortConfig struct {
	Name            string
	SocketType      string
	SupportedSpeeds []any
}

type Port struct {
	Config     PortConfig
	Group      string
	PortFormat string
	Index      int
}

type PortGroup struct {
	Name          string
	TotalQuantity int
	Members       []struct {
		Quantity        int
		PossibleConfigs []PortConfig
	}
}
