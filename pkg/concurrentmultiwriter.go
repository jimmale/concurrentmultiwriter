package concurrentmultiwriter

import (
	"errors"
	"io"
	"sync"
)

// ╔══════════════════════════════════════════════════════════════════════════╗
// ║                        Concurrent Multi Writer                           ║
// ╚══════════════════════════════════════════════════════════════════════════╝

// ConcurrentMultiWriter TODO write doc
type ConcurrentMultiWriter struct {
	writers []io.Writer
}

// Write TODO write docs
func (cmw *ConcurrentMultiWriter) Write(p []byte) (n int, err error) {

	wg := sync.WaitGroup{}
	errSlice := make([]error, len(cmw.writers))
	nSlice := make([]int, len(cmw.writers))

	for index, writer := range cmw.writers {
		wg.Add(1)
		go func(writer io.Writer, index int, p []byte) {

			n, err := writer.Write(p)

			errSlice[index] = err
			nSlice[index] = n

			wg.Done()

		}(writer, index, p)
	}
	wg.Wait()

	minWrittenChars := intMin(nSlice...)
	collectedErrors := wrapErrors(errSlice...)

	return minWrittenChars, collectedErrors
}

// MultiWriter creates a writer that duplicates its writes to all the provided writers, similar to the Unix tee(1)
// command.
//
// Each write is written to each listed writer concurrently, using goroutines and a waitgroup.
// If a listed writer returns an error, that overall write operation stops and returns the error; it does not continue
// down the list.
//
// If multiple listed writers return an error on a given write operation, the errors will be wrapped using errors.Join
func MultiWriter(writers ...io.Writer) *ConcurrentMultiWriter {
	return &ConcurrentMultiWriter{writers: writers}
}

func intMin(ints ...int) int {
	minNumber := ints[0]
	for _, v := range ints {
		if v < minNumber {
			minNumber = v
		}
	}
	return minNumber
}

// TODO this might not give useful feedback
func wrapErrors(errs ...error) error {
	var result error = nil
	var nonNilErrs []error

	for _, v := range errs {
		if v != nil {
			nonNilErrs = append(nonNilErrs, v)
		}
	}

	if len(nonNilErrs) > 0 {
		result = errors.Join(nonNilErrs...)
	}

	return result
}
