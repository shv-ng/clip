package clip

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
)

func RunBenchmark(file string) error {
	var s StatResult
	var data []float64

	f, err := os.Open(file)
	if err != nil {
		return fmt.Errorf(": %w", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		url := scanner.Text()
		shorted := Shortit(url)
		d := Diff(url, shorted)
		data = append(data, d)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	s.P50 = percentile(data, 50)
	s.P90 = percentile(data, 90)
	s.P99 = percentile(data, 99)

	summarize(&s, data)

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)

	// pretty print
	fmt.Println("=== Clip benchmark result ===")

	fmt.Fprintf(w, "File\t: %s\n", file)
	fmt.Fprintf(w, "Total URLs processed\t: %d\n", s.ProcessedUrl)
	w.Flush()

	fmt.Println("\nReduction Statistics:")

	w = tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintf(w, "\tAverage\t: %.2f%%\n", s.Avg)
	fmt.Fprintf(w, "\tMedian(p50)\t: %.2f%%\n", s.P50)
	fmt.Fprintf(w, "\tp90\t: %.2f%%\n", s.P90)
	fmt.Fprintf(w, "\tp99\t: %.2f%%\n", s.P99)
	fmt.Fprintf(w, "\tMin\t: %.2f%%\n", s.Min)
	fmt.Fprintf(w, "\tMax\t: %.2f%%\n", s.Max)
	w.Flush()

	fmt.Println()
	w = tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintf(w, "URLs that became longer\t: %.2f%%\n", s.LT0)
	fmt.Fprintf(w, "URLs shortened > 50%%\t: %.2f%%\n", s.GT50)
	w.Flush()

	return nil
}

func percentile(data []float64, p float64) float64 {
	if len(data) == 0 || p > 100 {
		return 0
	}

	sort.Float64s(data)

	index := p / 100 * float64(len(data)-1)

	i := int(index)
	frac := index - float64(i)

	if i+1 < len(data) {
		return data[i]*(1-frac) + data[i+1]*frac
	}

	return data[i]
}

type StatResult struct {
	ProcessedUrl int
	P50          float64
	P90          float64
	P99          float64
	Avg          float64
	Min          float64
	Max          float64
	GT50         float64
	LT0          float64
}

func summarize(stat *StatResult, data []float64) {
	n := len(data)
	if n == 0 {
		return
	}

	stat.ProcessedUrl = n

	stat.Min = data[0]
	stat.Max = data[0]

	var lt0, gt50 int

	var sum float64 = 0
	for _, d := range data {
		if d >= 50 {
			gt50 += 1
		}
		if d < 0 {
			lt0 += 1
		}

		if d < stat.Min {
			stat.Min = d
		}
		if d > stat.Max {
			stat.Max = d
		}

		sum += d
	}

	stat.Avg = sum / float64(n)
	stat.LT0 = 100 * float64(lt0) / float64(n)
	stat.GT50 = 100 * float64(gt50) / float64(n)
}
