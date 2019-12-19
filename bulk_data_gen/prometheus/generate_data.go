package prometheus

import (
	. "github.com/naivewong/influxdb-comparisons/bulk_data_gen/common"
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const MaxNumDataFiles = 170
var (
	EpochDuration = 10 * time.Second
	NHostSims = 505
)

type GeneralMeasurement struct {
	timestamp time.Time

	measurement []byte
	tagNames    [][]byte
	tagValues   [][]byte
}

func NewGeneralMeasurement(measurement []byte, start time.Time) *GeneralMeasurement {
	return &GeneralMeasurement{
		measurement: measurement,
		timestamp:   start,
	}
}

func (m *GeneralMeasurement) AddTags(name, value string) {
	m.tagNames = append(m.tagNames, []byte(name))
	m.tagValues = append(m.tagValues, []byte(value))
}

func (m *GeneralMeasurement) Tick(d time.Duration) {
	m.timestamp = m.timestamp.Add(d)
}

func (m *GeneralMeasurement) ToPoint(p *Point, value float64) bool {
	p.SetMeasurementName(m.measurement)
	p.SetTimestamp(&m.timestamp)

	for i := range m.tagNames {
		p.AppendTag(m.tagNames[i], m.tagValues[i])
	}

	p.AppendField([]byte("value"), value)
	return true
}

type PrometheusSimulator struct {
	madePoints int64
	madeValues int64
	maxPoints  int64

	simulatedMeasurementIndex int

	hostIndex int

	timestampNow   time.Time
	timestampStart time.Time
	timestampEnd   time.Time

	mode string // random, timeseries.

	measurements []*GeneralMeasurement
	hostCount    int64

	scanners []*bufio.Scanner
	files    []*os.File

	tmpData []float64
}

func (g *PrometheusSimulator) SeenPoints() int64 {
	return g.madePoints
}

func (g *PrometheusSimulator) SeenValues() int64 {
	return g.madeValues
}

func (g *PrometheusSimulator) Total() int64 {
	return g.maxPoints
}

func (g *PrometheusSimulator) Finished() bool {
	return g.madePoints >= g.maxPoints
}

func (g *PrometheusSimulator) Close() {
	for _, file := range g.files {
		file.Close()
	}
}

func (d *PrometheusSimulator) Next(p *Point) {
	// switch to the next host if needed
	if d.simulatedMeasurementIndex == len(d.measurements) {
		d.simulatedMeasurementIndex = 0
		d.hostIndex++
		if d.mode == "timeseries" {
			d.scanners[d.hostIndex].Scan()
			s := d.scanners[d.hostIndex].Text()
			data := strings.Split(s, ",")
			for i := range data {
				v, _ := strconv.ParseFloat(data[i], 64)
				d.tmpData[i] = v
			}
		}
	}

	if d.hostIndex == int(d.hostCount) {
		d.hostIndex = 0
		for _, m := range d.measurements {
			m.Tick(EpochDuration)
		}
	}

	p.AppendTag([]byte("instance"), []byte(fmt.Sprintf("pc9%05d:9100", d.hostIndex)))

	switch d.mode {
	case "random":
		d.measurements[d.simulatedMeasurementIndex].ToPoint(p, rand.Float64())
	case "timeseries":
		d.measurements[d.simulatedMeasurementIndex].ToPoint(p, d.tmpData[d.simulatedMeasurementIndex])
	}

	d.madePoints++
	d.simulatedMeasurementIndex++
	d.madeValues += int64(len(p.FieldValues))

	return
}

type PrometheusSimulatorConfig struct {
	Start time.Time
	End   time.Time

	HostCount  int64
	HostOffset int64

	DataDir   string
	SeriesDir string
	Mode      string
}

func (d *PrometheusSimulatorConfig) ToSimulator() *PrometheusSimulator {
	epochs := d.End.Sub(d.Start).Nanoseconds() / EpochDuration.Nanoseconds()
	maxPoints := epochs * (d.HostCount * int64(NHostSims))
	pg := &PrometheusSimulator{
		madePoints: 0,
		madeValues: 0,
		maxPoints:  maxPoints,

		simulatedMeasurementIndex: 0,

		hostIndex: 0,

		timestampNow:   d.Start,
		timestampStart: d.Start,
		timestampEnd:   d.End,

		hostCount: d.HostCount,
		mode:      d.Mode,
	}

	if d.Mode == "timeseries" {
		num := d.HostCount
		if num > MaxNumDataFiles {
			num = MaxNumDataFiles
		}
		for i := 0; i < int(num); i++ {
			file, err := os.Open(filepath.Join(d.DataDir, "data" + strconv.Itoa(i)))
			if err != nil {
				fmt.Println("Error open data file")
				os.Exit(1)
			}
			pg.files = append(pg.files, file)
			pg.scanners = append(pg.scanners, bufio.NewScanner(file))
		}
		for i := num; i < d.HostCount; i++ {
			file, err := os.Open(filepath.Join(d.DataDir, "data" + strconv.Itoa(int(i - num))))
			if err != nil {
				fmt.Println("Error open data file")
				os.Exit(1)
			}
			pg.files = append(pg.files, file)
			pg.scanners = append(pg.scanners, bufio.NewScanner(file))
		}
		pg.scanners[0].Scan()
		s := pg.scanners[0].Text()
		data := strings.Split(s, ",")
		for i := range data {
			v, _ := strconv.ParseFloat(data[i], 64)
			pg.tmpData = append(pg.tmpData, v)
		}
	}

	var temp *GeneralMeasurement

	file, err := os.Open(filepath.Join(d.SeriesDir, "node_exporter.json"))
	if err != nil {
		fmt.Println("Error open node_exporter.json")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() && i < NHostSims {
		m := make(map[string]string)
		err = json.Unmarshal([]byte(scanner.Text()), &m)
		if err != nil {
			fmt.Println("Error parsing Json")
			os.Exit(1)
		}
		temp = NewGeneralMeasurement([]byte(m["__name__"]), d.Start)
		for k, v := range m {
			if k == "__name__" || k == "instance" {
				continue
			}
			temp.AddTags(k, v)
		}
		pg.measurements = append(pg.measurements, temp)
		i++
	}
	file.Close()

	return pg
}
