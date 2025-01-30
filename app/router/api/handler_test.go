package api

import (
	"fmt"
	"testing"
	"time"
)

// ALTER SYSTEM SET timezone = 'America/Caracas';
func TestNow(t *testing.T) {
	t.Log("Test Now")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}
