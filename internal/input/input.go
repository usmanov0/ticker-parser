package input

import (
	"bufio"
	"context"
	"os"
)

func WaitForStop(ctx context.Context, cancel context.CancelFunc) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "STOP" {
			cancel()
			return
		}
	}
}
