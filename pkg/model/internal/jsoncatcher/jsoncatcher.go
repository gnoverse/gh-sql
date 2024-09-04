// Package jsoncatcher compares re-emitted, parsed JSON data with its raw
// counterpart to find fields which have not been correctly parsed.
package jsoncatcher

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/buger/jsonparser"
)

// CatchOptions contains options to modify the behaviour of Catch.
type CatchOptions struct {
	// Optional func, should return true if a mismatching key should be skipped
	// (as it is parsed otherwise).
	Skip func(key []byte) bool
	// Catch will write to Missed JSON objects containing the raw unparsed
	// object, and the fields it didn't parse, in a JSON stream (ie. one per line).
	Missed io.Writer
}

type missedLog struct {
	Missed []string        `json:"missed"`
	Raw    json.RawMessage `json:"raw"`
}

// ErrMissed is returned by [CatchOptions.Catch] for missed
// fields.
type ErrMissed []string

func (e ErrMissed) Error() string {
	return "missed JSON fields: " + strings.Join([]string(e), ", ")
}

// TODO: recursive, match values

func (c CatchOptions) Catch(rawOriginal []byte, parsed any) error {
	marshaled, err := json.Marshal(parsed)
	if err != nil {
		return err
	}

	var missed ErrMissed
	err = jsonparser.ObjectEach(rawOriginal, func(key, value []byte, valT jsonparser.ValueType, offset int) error {
		// Find the corresponding key also in marshaled.
		_, _, _, err := jsonparser.Get(marshaled, string(key))
		if err == jsonparser.KeyPathNotFoundError {
			if c.Skip != nil && !c.Skip(key) {
				missed = append(missed, string(key))
			}
			return nil
		}
		return err
	})
	if err != nil {
		return err
	}

	if len(missed) == 0 {
		// nothing missed, all good
		return nil
	}

	// some fields were missed
	if c.Missed != nil {
		// if we need to write to c.Missed, do so
		ml, err := json.Marshal(missedLog{
			Missed: []string(missed),
			Raw:    json.RawMessage(rawOriginal),
		})
		if err != nil {
			return fmt.Errorf("could not marshal missed log: %w", err)
		}
		_, err = c.Missed.Write(ml)
		if err != nil {
			log.Printf("warning: could not write to 'missed' buffer/file: %v", err)
		}
	}
	return missed
}
