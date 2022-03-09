package cryptocomp

import "fmt"

type uri string

func newUri() uri {
	return uri("https://min-api.cryptocompare.com/data/v2/histohour?")
}

//fsym=BTC&tsym=USD&toTs=1639591200&limit=2000&api_key=605154dc9ffd403a3f818ccb12405c0caded85d41fae631de6eb1d9c366299ab
func (u uri) ts(ts string) uri {
	return uri(fmt.Sprintf("%s&toTs=%s", u, ts))
}

func (u uri) apikey(key string) uri {
	return uri(fmt.Sprintf("%s&api_key=%s", u, key))
}

func (u uri) limit(limit int) uri {
	return uri(fmt.Sprintf("%s&limit=%d", u, limit))
}

func (u uri) symbol(symbol string) uri {
	return uri(fmt.Sprintf("%s&fsym=%s", u, symbol))
}

func (u uri) tsym(symbol string) uri {
	return uri(fmt.Sprintf("%s&tsym=%s", u, symbol))
}

func (u uri) build() string {
	return string(u)
}
