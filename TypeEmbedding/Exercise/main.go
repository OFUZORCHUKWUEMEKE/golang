package main

import "fmt"

type Bytes int
type Celsius float32

type BandwidthUsage struct {
	amount []Bytes
}
type CpuTemp struct {
	temp []Celsius
}

type MemoryUsage struct {
	amount []Bytes
}

func (b *BandwidthUsage) AverageBandWidth() int {
	sum := 0
	for i := 0; i < len(b.amount); i++ {
		sum += int(b.amount[i])
	}
	return sum / len(b.amount)
}
func (b *CpuTemp) AverageCputWidth() int {
	sum := 0
	for i := 0; i < len(b.temp); i++ {
		sum += int(b.temp[i])
	}
	return sum / len(b.temp)
}
func (m *MemoryUsage) AverageMemoryUsage() int {
	sum := 0
	for i := 0; i < len(m.amount); i++ {
		sum += int(m.amount[i])
	}
	return sum / len(m.amount)
}

type Dashboard struct {
	BandwidthUsage
	CpuTemp
	MemoryUsage
}

func main() {
	bandwidth := BandwidthUsage{[]Bytes{1000, 2000, 3000, 4000, 5000}}
	temp := CpuTemp{[]Celsius{50, 51, 52, 53, 54}}
	memory := MemoryUsage{[]Bytes{80000, 60000, 40000, 50000}}

	dash := Dashboard{
		BandwidthUsage: bandwidth,
		CpuTemp:        temp,
		MemoryUsage:    memory,
	}
	fmt.Printf("Average bandwidth usage: %v\n", dash.AverageBandWidth())
	fmt.Printf("Average bandwidth usage: %v\n", dash.AverageMemoryUsage())
	fmt.Printf("Average bandwidth usage: %v\n", dash.AverageCputWidth())
}
