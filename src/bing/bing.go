package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/generators"
	"github.com/faiface/beep/speaker"
)

func main() {
	// Define flags
	freqFlag := flag.Int("f", 116, "Frequency of the tone in Hz")
	timeFlag := flag.Int("t", 3, "Duration of the tone in seconds")
	flag.IntVar(freqFlag, "frequency", 116, "Frequency of the tone in Hz")
	flag.IntVar(timeFlag, "time", 3, "Duration of the tone in seconds")

	// Custom usage function to show only the executable name
	flag.Usage = func() {
		exeName := filepath.Base(os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [--frequency | -f] [--time | -t]\n", exeName)
		fmt.Fprintf(flag.CommandLine.Output(), "  --frequency | -f  int\n")
		fmt.Fprintf(flag.CommandLine.Output(), "        Frequency of the tone in Hz (default 116 [Bb])\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  --time | -t  int\n")
		fmt.Fprintf(flag.CommandLine.Output(), "        Duration of the tone in seconds (default 3)\n")
	}

	flag.Parse()

	// Convert duration to time.Duration
	duration := time.Duration(*timeFlag) * time.Second

	sampleRate := beep.SampleRate(44100)
	speaker.Init(sampleRate, sampleRate.N(time.Second/10))

	sineWave, err := generators.SinTone(sampleRate, *freqFlag)
	if err != nil {
		fmt.Println("Failed to generate sine wave:", err)
		return
	}
	streamer := beep.Take(sampleRate.N(duration), sineWave)

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	// Print "bong" while the sound is playing, inserting "o" in the middle each time
	go func() {
		bong := "bong"
		fmt.Printf("\r%s", bong)
		for {
			select {
			case <-done:
				return
			default:
				fmt.Printf("\r%s", bong)
				bong = bong[:len(bong)-2] + "ong"
				time.Sleep(33 * time.Millisecond)
			}
		}
	}()

	<-done
}
