package plc

import (
	"os"
	"testing"

	"github.com/bruyss/go-object-generator/utils"
)

func TestMain(m *testing.M) {
	utils.InitializeDevLogger()
	exitVal := m.Run()
	os.Exit(exitVal)
}
