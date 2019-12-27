package influxdb

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"time"
	bulkQuerygen "github.com/naivewong/influxdb-comparisons/bulk_query_gen"
)

// InfluxDevopsSingleHost produces Influx-specific queries for the devops single-host case.
type InfluxPrometheusType1 struct {
	db          string
	interval    bulkQuerygen.TimeInterval
	currentTime time.Time
	currentIdx  int
	mode        string
	org         string
	predicate   string
}

func NewFluxPrometheusType1(dbConfig bulkQuerygen.DatabaseConfig, interval bulkQuerygen.TimeInterval, duration time.Duration, scaleVar int) bulkQuerygen.QueryGenerator {
	g := &InfluxPrometheusType1{
		db:          dbConfig[bulkQuerygen.DatabaseName],
		interval:    interval,
		currentTime: interval.Start.Add(interval.End.Sub(interval.Start)/2),
		mode:        dbConfig["mode"],
		org:         dbConfig["org"],
	}
	switch dbConfig["mode"] {
	case "origin":
		file, err := os.Open("./type1.txt")
		if err != nil {
			fmt.Println("Error open type1.txt", err)
			os.Exit(1)
		}
		s := bufio.NewScanner(file)
		for s.Scan() {
			g.predicate = g.predicate + "r._measurement==\"" + s.Text() + "\" or "
		}
	case "same-host":
		for i := 0; i < 30; i++ {
			g.predicate = fmt.Sprintf("%sr.instance==\"pc9%05d:9100\" or ", g.predicate, i)
		}
	}
	if len(g.predicate) > 4 {
		g.predicate = g.predicate[: len(g.predicate) - 4]
	}
	return g
}

func (d *InfluxPrometheusType1) Dispatch(i int) bulkQuerygen.Query {
	q := bulkQuerygen.NewHTTPQuery() // from pool
	query := fmt.Sprintf(`from(bucket:"%s") `+
			`|> range(start:%s, stop:%s) `+
			`|> filter(fn:(r) => %s)`,
			d.db,
			d.currentTime.Format(time.RFC3339),
			d.currentTime.Add(5 * time.Minute).Format(time.RFC3339),
			d.predicate)
	getValues := url.Values{}
	getValues.Set("org", d.org)
	q.Method = []byte("POST")
	q.Path = []byte(fmt.Sprintf("/query?%s", getValues.Encode()))
	q.Body = []byte(query)
	switch d.mode {
	case "origin":
		q.HumanLabel = []byte("Prometheus type1 origin")
	case "same-host":
		q.HumanLabel = []byte("Prometheus type1 same host")

	}
	q.HumanDescription = []byte(fmt.Sprintf("%s: %s", q.HumanLabel, d.currentTime.Format(time.RFC3339)))
	
	d.currentTime = d.currentTime.Add(5 * time.Minute)
	return q
}
