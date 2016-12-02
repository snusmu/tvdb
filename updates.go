package tvdb

import "net/url"

type SeriesUpdate struct {
	ID          uint64 `json:"id"`
	LastUpdated uint64 `json:"lastUpdated"`
}

// Updates retrieves list of all updated series withing last few weeks
func (c *Conn) Updates(fromTime string) ([]SeriesUpdate, error) {
	v := make(url.Values)
	v.Set("fromTime", fromTime)
	var r struct {
		Data  []SeriesUpdate `json:"data"`
		Error requestErrors  `json:"error"`
	}
	if err := c.get(makeURL("/updated/query", v.Encode()), &r); err != nil {
		return nil, err
	}
	return r.Data, nil
}
