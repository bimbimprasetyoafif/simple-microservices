//go:build integration
// +build integration

package testing

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestE2ETestSuite(t *testing.T) {
	suite.Run(t, &e2eTestSuite{})
}
