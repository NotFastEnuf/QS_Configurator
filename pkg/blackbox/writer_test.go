package blackbox_test

import (
	"os"
	"testing"

	"github.com/NotFastEnuf/QS_Configurator/pkg/blackbox"
	"github.com/NotFastEnuf/QS_Configurator/pkg/protocol/quic"
)

func TestStructTags(t *testing.T) {
	b := quic.BlackboxCompact{}
	w := blackbox.NewDefaultWriter(os.Stdout)

	w.WriteHeaders()
	w.WriteValue(&b)
}
