/**
 *  Author: SongLee24
 *  Email: lisong.shine@qq.com
 *  Date: 2018-08-15
 *
 *
 *  prometheus.Desc是指标的描述符，用于实现对指标的管理
 *
 */

package collector

import (
        "bufio"
        "fmt"
        "io"
        "os"
        "strings"
	"github.com/prometheus/client_golang/prometheus"
        "strconv"
)

type Metrics struct {
	metrics map[string]*prometheus.Desc
}



func NewMetrics(a string) *Metrics {
	return &Metrics{
		metrics: map[string]*prometheus.Desc{
			"gauge_metric": prometheus.NewDesc(a+"_gauge_metric","The description of my_gauge_metric", []string{"path"}, nil),
		},
	}
}

func (c *Metrics) Describe(ch chan<- *prometheus.Desc) {
	for _, m := range c.metrics {
		ch <- m
	}
}

func (c *Metrics) Collect(ch chan<- prometheus.Metric) {

	Gaugedata := c.getdata()
	for name, data := range Gaugedata {
		ch <-prometheus.MustNewConstMetric(c.metrics["gauge_metric"], prometheus.GaugeValue, float64(data), name)
	}
}


func (c *Metrics) getdata() (Gaugedata map[string]int) {
        filesize, _ := getfilesize()
 	Gaugedata = map[string]int{
		"/install": filesize,
	}
	return
 }

func getfilesize() (int, error) {

        file, err := os.Open("/root/go/src/filesize_exporter/size.txt")
        if err != nil {
                fmt.Sprintf("error!")
        }
        defer file.Close()

        size, err := parsefile(file)
        if err != nil {
                fmt.Sprintf("error!")
        }
        filesize, err := strconv.Atoi(size)

	return filesize, nil
}


func parsefile(data io.Reader) (string, error) {
        scanner := bufio.NewScanner(data)
        var entries string
        for scanner.Scan() {
                columns := strings.Fields(scanner.Text())
                entries = columns[0]
        }

        if err := scanner.Err(); err != nil {
                fmt.Errorf("failed to parse file: %s", err)
        }

        return entries, nil
}


