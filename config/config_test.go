package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ConfigSuite struct {
	suite.Suite
}

func TestConfigSuite(t *testing.T) {
	suite.Run(t, new(ConfigSuite))
}

func (s *ConfigSuite) TestGetFBAppToken_HasValue_ReturnValue() {
	os.Setenv(fbToken, "test_token")
	s.Equal("test_token", GetFBAppToken())
}

func (s *ConfigSuite) TestGetFBAppToken_DoesNotHaveValue_ReturnEmptyString() {
	os.Setenv(fbToken, "")
	s.Empty(GetFBAppToken())
}

func (s *ConfigSuite) TestGetPageID_HasValue_ReturnValue() {
	os.Setenv(pageID, "test_id")
	s.Equal("test_id", GetPageID())
}

func (s *ConfigSuite) TestGetPageID_DoesNotHaveValue_ReturnEmptyString() {
	os.Setenv(pageID, "")
	s.Empty(GetPageID())
}

func (s *ConfigSuite) TestGetExportDirectory_HasValue_ReturnValue() {
	os.Setenv(exportDirectory, "test_directory")
	s.Equal("test_directory", GetExportDirectory())
}

func (s *ConfigSuite) TestGetExportDirectory_DoesNotHaveValue_ReturnEmptyString() {
	os.Setenv(exportDirectory, "")
	s.Empty(GetExportDirectory())
}
