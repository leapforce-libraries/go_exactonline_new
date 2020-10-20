package exactonline

import (
	bigquerytools "github.com/Leapforce-nl/go_bigquerytools"
	http "github.com/Leapforce-nl/go_exactonline_new/http"
	salesorder "github.com/Leapforce-nl/go_exactonline_new/salesorder"
)

// ExactOnline stores ExactOnline configuration
//
type ExactOnline struct {
	SalesOrderClient *salesorder.Client
}

// methods
//
func NewExactOnline(division int, clientID string, clientSecret string, bigQuery *bigquerytools.BigQuery, isLive bool) (*ExactOnline, error) {
	eo := ExactOnline{}

	http, err := http.NewHttp(division, clientID, clientSecret, bigQuery, isLive)
	if err != nil {
		return nil, err
	}
	eo.SalesOrderClient = salesorder.NewClient(http)

	return &eo, nil
}
