package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	_version = "1.0"
)

var (
	line, text string
	parts      []string
	err        error
)

type CpuStat struct {
	user, nice, system         float64
	idle, iowait, irq, softirq float64
	steal, guest, guest_nice   float64
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func collectCpuStat() (cpustats CpuStat) {
	file, err := os.Open("/proc/stat")
	check(err)

	defer file.Close()

	reader := bufio.NewReader(file)
	// 0x0A for 10 newline char.
	text, err = reader.ReadString(10)
	check(err)
	parts = strings.Fields(text)
	parts_float64 := make([]float64, 11)
	for v := 1; v < len(parts); v++ {
		parts_float64[v], err = strconv.ParseFloat((parts[v]), 64)
		check(err)
	}
	cpustats = CpuStat{user: parts_float64[1], nice: parts_float64[2], system: parts_float64[3], idle: parts_float64[4], iowait: parts_float64[5], irq: parts_float64[6], softirq: parts_float64[7], steal: parts_float64[8], guest: parts_float64[9], guest_nice: parts_float64[10]}

	return cpustats
}

func calcCpuPerc() (cpu_perc, user, nice, system, idle float64) {
	cpustats_prev := collectCpuStat()
	time.Sleep(1 * time.Second)
	cpustats := collectCpuStat()

	prevIdle := (cpustats_prev.idle + cpustats_prev.iowait)
	Idle := (cpustats.idle + cpustats.iowait)
	prevNonIdle := (cpustats_prev.user + cpustats_prev.nice + cpustats_prev.system + cpustats_prev.irq + cpustats_prev.softirq + cpustats_prev.steal)
	nonIdle := (cpustats.user + cpustats.nice + cpustats.system + cpustats.irq + cpustats.softirq + cpustats.steal)

	prevTotal := prevIdle + prevNonIdle
	total := Idle + nonIdle

	totald := total - prevTotal
	idled := Idle - prevIdle

	cpu_perc = (totald - idled) * 100 / totald
	return cpu_perc, cpustats.user, cpustats.nice, cpustats.system, cpustats.idle
}

func help() {
	fmt.Printf("%s v%s\n", os.Args[0], _version)
	fmt.Println()
	fmt.Printf("Usage : ./%s -w %%WARNING -c %%CRITICAL\n", os.Args[0])
	fmt.Println()
	fmt.Println("WARNING and CRITICAL values are percentage values without %")
	fmt.Println()
	fmt.Println("2016 - Aydin Doyak <aydintd@gmail.com>")
	os.Exit(5)
}

func main() {
	var warn, crit float64
	if len(os.Args) != 5 {
		help()
		os.Exit(5)
	}

	if os.Args[1] == "-w" && os.Args[3] == "-c" {
		warn, err = strconv.ParseFloat(os.Args[2], 64)
		check(err)
		crit, err = strconv.ParseFloat(os.Args[4], 64)
		check(err)
		if warn >= crit {
			fmt.Println("WARNING value can not be bigger than CRITICAL value")
			os.Exit(5)
		} else if crit > 100 {
			fmt.Println("%CRITICAL value can not be bigger than %100")
			os.Exit(5)
		}
	} else {
		help()
	}

	cpu_perc, user, nice, system, idle := calcCpuPerc()

	switch {
	case 0 <= cpu_perc && cpu_perc < warn:
		fmt.Printf("CPU: OK - TotalAvg: %%%2f|avg=%%%2f;;;; user=%.0f;;;; nice=%.0f;;;; system=%.0f;;;; idle=%.0f;;;; \n", cpu_perc, cpu_perc, user, nice, system, idle)
		os.Exit(0)
	case warn <= cpu_perc && cpu_perc < crit:
		fmt.Printf("CPU: WARNING - TotalAvg: %%%2f|avg=%%%2f;;;; user=%.0f;;;; nice=%.0f;;;; system=%.0f;;;; idle=%.0f;;;; \n", cpu_perc, cpu_perc, user, nice, system, idle)
		os.Exit(1)
	case crit <= cpu_perc:
		fmt.Printf("CPU: CRITICAL - TotalAvg: %%%2f|avg=%%%2f;;;; user=%.0f;;;; nice=%.0f;;;; system=%.0f;;;; idle=%.0f;;;; \n", cpu_perc, cpu_perc, user, nice, system, idle)
		os.Exit(2)
	default:
		fmt.Printf("UNKNOWN value.\n")
		os.Exit(3)
	}
}
