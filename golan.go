package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/huandu/goroutine"
	"os"
	"os/exec"
	"runtime"
	"runtime/pprof"
	"strconv"
	"strings"
	"time"
)

type Process struct {
	pid int
	cpu float64
	mem float64
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Test1")
		panic(err)
	}
}

var t uint64

func cpuTTest() {
	cmd := exec.Command("ps", "aux")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		logrus.Fatal(err)
	}
	Processes := make([]*Process, 0)
	for {
		line, err := out.ReadString('\n')
		if err != nil {
			break
		}
		tokens := strings.Split(line, " ")
		ft := make([]string, 0)
		for _, t := range tokens {
			if t != "" && t != "\t" {
				ft = append(ft, t)
			}
		}
		//logrus.Println(len(ft), ft)
		pid, err := strconv.Atoi(ft[1])
		if err != nil {
			continue
		}
		cpu, err := strconv.ParseFloat(ft[2], 64)
		if err != nil {
			logrus.Fatal(err)
		}
		mem, err := strconv.ParseFloat(ft[3], 64)
		if err != nil {
			logrus.Fatal(err)
		}
		Processes = append(Processes, &Process{pid, cpu, mem})
	}
	for _, p := range Processes {
		if os.Getpid() == p.pid {
			logrus.Println("process", p.pid, "takes", p.cpu, "% of the CPU", p.mem, "% of the memery")
		}
	}
	fmt.Println("当前协程数量： ", runtime.NumGoroutine())
	stats := &runtime.MemStats{}
	runtime.ReadMemStats(stats)
	if (stats.LastGC - t) != 0 {
		fmt.Println("时间差：", stats.LastGC-t)
		t = stats.LastGC
	}
	fmt.Println(
		"已申请且仍在使用", stats.Alloc,
		"已申请的总", stats.TotalAlloc,
		"已申请且任在使用", stats.HeapAlloc, stats.NextGC,
		"从系统中获取", stats.HeapSys,
		"系统-使用:", stats.HeapSys-stats.HeapAlloc,
		stats.HeapSys-stats.HeapAlloc-stats.HeapIdle,
		"释放到系统", stats.HeapReleased,
		"闲置span中", stats.HeapIdle,
		"last time GC", stats.LastGC, stats.Frees)
}

func goTest() {
	id := goroutine.GoroutineId()
	logrus.Println("go id:", id)
}

type newTest struct {
	str string
	res int
	s   string
}

func f1() {
	for {
		str := make([]string, 5)
		str[0] = " "
	}
}
func xun() {
	fmt.Println("Test")
	for {
		stats := &runtime.MemStats{}
		runtime.ReadMemStats(stats)

		buf := make([]int, 10000)
		buf[0] = 0

		//fmt.Println("testing:", stats.HeapAlloc)
		time.Sleep(100 * time.Millisecond)
	}
	//fmt.Println(cur)
}

var ch1 = make(chan int)
var ch2 = make(chan int)

func timeTest() {
	fmt.Println("$$$$$$$$$$$")
	time.Sleep(20 * time.Second)
	fmt.Println("test*******")
	ch1 <- 1
}
func statsTest() {
	for {
		stats := &runtime.MemStats{}
		runtime.ReadMemStats(stats)
		if stats.HeapAlloc >= 4194304 {
			fmt.Println(stats.HeapAlloc)
			ch2 <- 2
		}
	}
}
func switchTest() {
	go statsTest()
	go timeTest()
	select {
	case <-ch1:
		logrus.Println(time.Now())
	case <-ch2:
		logrus.Println("stats:")
	}
}
func firstGc() {
	stats := &runtime.MemStats{}
	runtime.ReadMemStats(stats)
	fmt.Println("firstGc:", stats.HeapAlloc)
}

//打印GC回收的内存和时间
func print(heaps uint64) {
	for {
		stats := &runtime.MemStats{}
		runtime.ReadMemStats(stats)
		if (stats.LastGC - t) != 0 {
			logrus.Println("GC release memery:", stats.HeapAlloc-heaps)
			logrus.Println("last time GC", stats.LastGC)
			/*fmt.Println(
				"已申请且任在使用", stats.HeapAlloc, stats.NextGC,
				"last time GC", stats.LastGC)
			firstGc()*/
			t = stats.LastGC
		}
	}

	//time.Sleep(1*time.Second)
}

func cpuTest() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		logrus.Fatal(err)
	}
	//cpu使用情况记录
	if err := pprof.StartCPUProfile(f); err != nil {
		logrus.Println("startcpuprofile is err")
		f.Close()
		return
	}
	defer pprof.StopCPUProfile()
}
func blockTest() {
	f, err := os.Create("block.out")
	if err != nil {
		logrus.Println("create block.out is err!")
	}
	runtime.SetBlockProfileRate(1)
	if err := pprof.Lookup("block").WriteTo(f, 0); err != nil {
		logrus.Println("create block.out is err!")
	}
	f.Close()
}

func memTest() {
	runtime.MemProfileRate = 10
	f, err := os.Create("mem.out")
	if err != nil {
		logrus.Println("open mem.txt err")
		return
	}
	if err = pprof.WriteHeapProfile(f); err != nil {
		logrus.Println("writeHeapProfile is err!")
	}
	f.Close()
}

func main() {
	stats := &runtime.MemStats{}
	runtime.ReadMemStats(stats)
	heaps := stats.HeapAlloc
	fmt.Println(os.Getpid())
	cpuTrue := flag.Bool("cpu", false, "cpu is ture")
	fmt.Println("cpu:", *cpuTrue)
	flag.Parse()

	go xun()
	cpuTest()
	memTest()
	blockTest()
	fmt.Println("ok")

	print(heaps)

}
