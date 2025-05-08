package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"os/exec"
	"strconv"
)

func main() {
	logger := slog.Default()

	if err := run(logger); err != nil {
		log.Fatal(err)
	}
}

func run(logger *slog.Logger) error {
	cmd := exec.Command("/usr/bin/pactl", "subscribe")
	r, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	defer func() {
		err := r.Close()
		if err != nil {
			logger.Error("error closing pipe", slog.String("error", err.Error()))
		}
	}()

	if err := cmd.Start(); err != nil {
		return err
	}

	// Event 'remove' on source-output #2180
	bufreader := bufio.NewReaderSize(r, 37)

	for {
		line, isPrefix, err := bufreader.ReadLine()
		if err != nil {
			break
		}

		// should check change on sink and differentiate from change on source
		if !bytes.Contains(line, []byte("change")) {
			continue
		}

		mute, err := CheckMute()
		if err != nil {
			return err
		}

		volume, err := CheckVolume()
		if err != nil {
			return err
		}

		fmt.Printf("mute: %t, volume: %d line: %s prefix: %t\n", mute, volume, string(line), isPrefix)
	}

	return nil
}

func CheckMute() (bool, error) {
	cmd := exec.Command("/usr/bin/pactl", "get-sink-mute", "@DEFAULT_SINK@")

	out, err := cmd.Output()
	if err != nil {
		return false, fmt.Errorf("error checking if sink is mute: %w", err)
	}

	return bytes.Contains(out, []byte("yes")), nil
}

func CheckVolume() (int, error) {
	cmd := exec.Command("/usr/bin/pactl", "get-sink-volume", "@DEFAULT_SINK@")

	out, err := cmd.Output()
	if err != nil {
		return 0, fmt.Errorf("error checking if sink is mute: %w", err)
	}

	idx := bytes.IndexRune(out, '%')
	if idx < 2 {
		return 0, errors.New("error parsing pactl get-sink-volume cannot find '%'")
	}

	volume, err := strconv.Atoi(string(out[idx-2 : idx]))
	if err != nil {
		return 0, fmt.Errorf("error parsing pactl get-sink-volume: %w", err)
	}

	return volume, nil
}
