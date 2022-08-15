package adapter

import (
	"fmt"
	"github.com/monopeelz/linear-avocado/pkg/scanner"
	"os"
	"strings"
)

var (
	PublicKeyDetect  = scanner.NewRule("G0001", "foo", "public_key detection", "LOW")
	PrivateKeyDetect = scanner.NewRule("G0002", "bar", "private_key detection", "LOW")
)

type simpleScanner struct {
}

func (c simpleScanner) Rules() []scanner.Rule {
	return []scanner.Rule{
		*PublicKeyDetect,
		*PrivateKeyDetect,
	}
}

func (c simpleScanner) ScanFile(path string) ([]scanner.Finding, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return c.scanByte(content)
}

func (c simpleScanner) scanByte(b []byte) ([]scanner.Finding, error) {
	line := 1
	slide := make([]byte, 0)
	detections := make([]scanner.Finding, 0)
	size := len(b)
	for idx, a := range b {
		if a == 32 || a == 10 || (idx == size-1 && len(slide) != 0) {
			if rule := c.scanSecret(slide); rule != nil {
				detections = append(detections, scanner.Finding{
					Type:     rule.Type,
					RuleId:   rule.ID,
					Location: scanner.Location{},
					Metadata: scanner.Metadata{
						Description: rule.Description,
						Severity:    rule.Severity,
					},
				})
			}
			slide = []byte{}
		} else {
			slide = append(slide, a)
		}
		if a == 10 {
			line += 1
			continue
		}

	}
	fmt.Println("detections", detections)
	return detections, nil
}

func (c simpleScanner) scanSecret(i []byte) *scanner.Rule {
	word := string(i)
	if strings.Contains(word, "private_key") {
		fmt.Println("private_key")
		return PrivateKeyDetect
	} else if strings.Contains(word, "public_key") {
		return PublicKeyDetect
	}
	return nil
}

func NewSimpleScanner() scanner.Scanner {
	return &simpleScanner{}
}
