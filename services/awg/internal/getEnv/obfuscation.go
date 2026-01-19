package getenv

import (
	"fmt"
	"os"
	"strconv"

	awgctrlgo "github.com/slipynil/awgctrl-go"
)

// NewObfuscation creates a new Obfuscation struct from environment variables.
// You can also use Obfuscation struct for initialization awg service
func NewObfuscation() (*awgctrlgo.Obfuscation, error) {
	jc := os.Getenv("JS")
	if len(jc) == 0 {
		return nil, fmt.Errorf("JS is not set")
	}
	jcInt, err := strconv.Atoi(jc)
	if err != nil {
		return nil, fmt.Errorf("JS is not a valid integer")
	}

	jmin := os.Getenv("JMIN")
	if len(jmin) == 0 {
		return nil, fmt.Errorf("JMIN is not set")
	}
	jminInt, err := strconv.Atoi(jmin)
	if err != nil {
		return nil, fmt.Errorf("JMIN is not a valid integer")
	}

	jmax := os.Getenv("JMAX")
	if len(jmax) == 0 {
		return nil, fmt.Errorf("JMAX is not set")
	}
	jmaxInt, err := strconv.Atoi(jmax)
	if err != nil {
		return nil, fmt.Errorf("JMAX is not a valid integer")
	}

	s1 := os.Getenv("S1")
	if len(s1) == 0 {
		return nil, fmt.Errorf("S1 is not set")
	}
	s1Int, err := strconv.Atoi(s1)
	if err != nil {
		return nil, fmt.Errorf("S1 is not a valid integer")
	}

	s2 := os.Getenv("S2")
	if len(s2) == 0 {
		return nil, fmt.Errorf("S2 is not set")
	}
	s2Int, err := strconv.Atoi(s2)
	if err != nil {
		return nil, fmt.Errorf("S2 is not a valid integer")
	}

	h1 := os.Getenv("H1")
	if len(h1) == 0 {
		return nil, fmt.Errorf("H1 is not set")
	}
	h1Int, err := strconv.Atoi(h1)
	if err != nil {
		return nil, fmt.Errorf("H1 is not a valid integer")
	}

	h2 := os.Getenv("H2")
	if len(h2) == 0 {
		return nil, fmt.Errorf("H2 is not set")
	}
	h2Int, err := strconv.Atoi(h2)
	if err != nil {
		return nil, fmt.Errorf("H2 is not a valid integer")
	}

	h3 := os.Getenv("H3")
	if len(h3) == 0 {
		return nil, fmt.Errorf("H3 is not set")
	}
	h3Int, err := strconv.Atoi(h3)
	if err != nil {
		return nil, fmt.Errorf("H3 is not a valid integer")
	}

	h4 := os.Getenv("H4")
	if len(h4) == 0 {
		return nil, fmt.Errorf("H4 is not set")
	}
	h4Int, err := strconv.Atoi(h4)
	if err != nil {
		return nil, fmt.Errorf("H4 is not a valid integer")
	}

	return &awgctrlgo.Obfuscation{
		Jc:   jcInt,
		Jmin: jminInt,
		Jmax: jmaxInt,
		S1:   s1Int,
		S2:   s2Int,
		H1:   h1Int,
		H2:   h2Int,
		H3:   h3Int,
		H4:   h4Int,
	}, nil
}
