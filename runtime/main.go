package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

var (
	err error
	numCPU int
	numCPUCtx int
	perCPU []float64
	perCPUctx []float64
	infoStat []cpu.InfoStat
	timeStat []cpu.TimesStat
)

func main() {
	fmt.Println("Tests using runtime for metrics")

	// NumCPU returns the number of logical CPUs usable by the current process.
	fmt.Println(runtime.NumCPU())

	v, _ := mem.VirtualMemory()

    // almost every return value is a struct
    fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

    // convert to JSON. String() is also implemented
    fmt.Println(v)

	// Counts - counts the number of a total CPU logical and physical if you pass true and only physical if you pass false.
	numCPU, err = cpu.Counts(true)

	if err != nil {
		fmt.Print("Error function cpu.Counts", err)
	}

	fmt.Printf("Nº CPUs %d\n", numCPU)

	// Percent calculates the percentage of cpu used either per CPU or comercial. 
	// If an interval of 0 is given it will compare the current cpu times against the last call. 
	// Returns one value per cpu, or a single value if percpu is set to false.
	time.Sleep(2 * time.Second)
	perCPU, err = cpu.Percent(0, true)
	if err != nil {
		fmt.Println("Error function cpu.Percent()", err)
	}

	fmt.Printf("Percent used by CPUs %v\n\n", perCPU)

	// CPUInfo on linux will return 1 item per physical thread.
	infoStat, err = cpu.Info()

	if err != nil {
		fmt.Println("Error func cpu.Info()", err)
	}

	fmt.Printf("CPU InfoStat %v\n\n", infoStat)

	// TimesStat contains the amounts of time the CPU has spent performing different kinds of work. 
	// Time units are in seconds. It is based on linux /proc/stat file.
	timeStat, err = cpu.Times(true)

	if err != nil {
		fmt.Println("Error func cpu.Times()", err)
	}

	fmt.Printf("CPU TimeStat %v\n", timeStat)

	for _, tmStat := range timeStat {
		fmt.Printf("CPU TimeStat %v\n", tmStat)
	}

}
