package config

type Config struct {
	Exchanges []ExchangeConfig
}

type ExchangeConfig struct {
	Name        string
	API         string
	Credentials string // json containing apikeys and other secrets.
}

func LoadConfiguration() {

}
