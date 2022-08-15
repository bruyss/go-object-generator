package plc

import (
	"fmt"
	"os"
	"testing"

	"github.com/bruyss/go-object-generator/utils"
)

func TestMain(m *testing.M) {
	utils.InitializeDevLogger()
	fmt.Println("Before running tests...")
	exitVal := m.Run()
	fmt.Println("After running tests...")
	os.Exit(exitVal)
}
