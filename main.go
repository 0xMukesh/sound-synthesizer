package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"
	"time"
)

const (
	duration   = 5
	sampleRate = 44100
	freq       = 1000
)

func main() {
	// number of samples
	ns := duration * sampleRate
	// angle increment per sample
	angle := (math.Pi * 2.0) / float64(ns)
	// exponential decay; f(x) = (0.0001/x)^(1/ns)
	startAmplitude := 1.0
	endAmplitude := 1.0e-4
	decayFactor := math.Pow(endAmplitude/startAmplitude, 1.0/float64(ns))

	_, err := os.Stat("out.bin")
	if err == nil {
		if err := os.Remove("out.bin"); err != nil {
			panic(err.Error())
		}
	}

	f, err := os.Create("out.bin")
	if err != nil {
		panic(err.Error())
	}

	start := time.Now()

	for i := 0; i < ns; i++ {
		sample := math.Sin(angle * freq * float64(i))
		sample *= startAmplitude
		startAmplitude *= decayFactor

		var buf [8]byte
		binary.LittleEndian.PutUint32(buf[:], math.Float32bits(float32(sample)))

		_, err := f.Write(buf[:])
		if err != nil {
			panic(err.Error())
		}
	}

	fmt.Printf("done - %dms\n", time.Since(start).Milliseconds())
}
