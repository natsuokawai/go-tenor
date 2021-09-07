package gotenor

import (
	"net/url"
	"strconv"
)

func (t *Tenor) GetRandom(query, filter, locale string, limit int) (*tenorData, error) {
	url := t._urlBuilder("random") + "&q=" + url.QueryEscape(query)

	if limit > 0 {
		url += "&limit=" + strconv.Itoa(limit)
	}

	if filter != "" {
		url += "&contentfilter=" + filter
	}

	if locale != "" {
		url += "&locale=" + locale
	}

	body, err := t._fetch(url)
	if err != nil {
		return nil, err
	}

	return t._parseData(body)
}
