package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	cmw "github.com/jimmale/concurrentmultiwriter/pkg"
	"io"
	"os"
	"time"
)

// SlowDiscarder discards anything written to it, but takes 1 second to do so
type SlowDiscarder struct {
}

// Write discards anything written to it
func (s SlowDiscarder) Write(p []byte) (n int, err error) {
	time.Sleep(1 * time.Second)
	return len(p), nil
}

func main() {

	// these are the five writers that we want one input to be written to
	md5hasher := md5.New()
	sha1hasher := sha1.New()
	sha256hasher := sha256.New()
	fastDiscarder := io.Discard
	slowDiscarder := SlowDiscarder{}

	// let's make a concurrent multiwriter
	mycmw := cmw.MultiWriter(md5hasher, sha1hasher, sha256hasher, fastDiscarder, slowDiscarder)

	// let's collect some input that we want to write to all of these Writers
	fmt.Print("Enter a line of text: ")
	bufferedInput := bufio.NewReader(os.Stdin)
	inputString, err := bufferedInput.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	// let's write that input to all five of the above writers
	_, err = mycmw.Write([]byte(inputString))
	if err != nil {
		fmt.Println("Error writing input to concurrent multiwriter:", err)
		os.Exit(1)
	}

	fmt.Println("\n================================================================================")
	fmt.Printf("md5sum:         %x\n", md5hasher.Sum(nil))
	fmt.Printf("sha1sum:        %x\n", sha1hasher.Sum(nil))
	fmt.Printf("sha256sum:      %x\n", sha256hasher.Sum(nil))
	fmt.Println("fastDiscarder:  <text discarded>")
	fmt.Println("slowDiscarder:  <text discarded>")
	fmt.Println("================================================================================")
}
