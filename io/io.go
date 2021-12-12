package io

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/atotto/clipboard"
)

// ReadAll reads input from a piped stdin or clipboard if no stdin was passed.
func ReadAll() ([]byte, error) {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return nil, err
	}

	if fi.Mode()&os.ModeNamedPipe == 0 {
		clip, err := clipboard.ReadAll()
		if err != nil {
			return nil, err
		}

		return []byte(clip), nil
	}

	return ioutil.ReadAll(os.Stdin)
}

// WriteAll writes to clipboard and stdout
func WriteAll(data []byte) error {
	if err := clipboard.WriteAll(string(data)); err != nil {
		return err
	}
	_, err := fmt.Printf("%s", string(data))
	return err
}
