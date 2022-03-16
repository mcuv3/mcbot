package cryptocomp

import (
	"fmt"
	"strings"
)

type uri struct {
	uri  string
	key  string
	lim  string
	sym  string
	tsy  string
	time string
}

func newUri() *uri {
	return &uri{
		uri: "https://min-api.cryptocompare.com/data/v2/histohour",
	}
}

//fsym=BTC&tsym=USD&toTs=1639591200&limit=2000&api_key=605154dc9ffd403a3f818ccb12405c0caded85d41fae631de6eb1d9c366299ab
func (u *uri) ts(ts string) *uri {
	u.time = fmt.Sprintf("toTs=%s", ts)
	return u
}

func (u *uri) apikey(key string) *uri {
	u.key = fmt.Sprintf("api_key=%s", key)
	return u
}

func (u *uri) limit(limit int) *uri {
	u.lim = fmt.Sprintf("limit=%d", limit)
	return u
}

func (u *uri) symbol(symbol string) *uri {
	u.sym = fmt.Sprintf("fsym=%s", symbol)
	return u
}

func (u *uri) tsym(symbol string) *uri {
	u.tsy = fmt.Sprintf("tsym=%s", symbol)
	return u
}

func (u uri) build() string {
	return fmt.Sprintf("%s?%s&%s", u.uri, u.key, strings.Join([]string{u.sym, u.tsy, u.time, u.lim}, "&"))
}
