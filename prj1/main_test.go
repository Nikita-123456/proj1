package main
import (
	"testing"
	"time"
	"os"
)
func TestWorking(t *testing.T){
	start := time.Now()

	Working()
	timeElapsed := time.Since(start)

	if  timeElapsed > (2*time.Second){
        os.Exit(1)
	}

}
 