package main

import (
	"github.com/jfrog/jfrog-cli/utils/tests"
	"testing"
)

var skipTests = map[string]bool{
	//Access
	"TestRefreshableAccessTokens": true,
	//Artifactory
	"TestArtifactoryProxy":                          true,
	"TestArtifactoryUploadAsArchiveWithIncludeDirs": true,
	"TestGitLfsCleanup":                             true,
	//GO
	"TestUnitTests": true,
	//Pip
	"TestPipInstallNativeSyntax":                        true,
	"TestPipInstallLegacy":                              true,
	"TestPipInstallNativeSyntax/setuppy":                true,
	"TestPipInstallLegacy/requirements-verbose":         true,
	"TestPipInstallLegacy/setuppy-with-module":          true,
	"TestPipInstallNativeSyntax/setuppy-with-module":    true,
	"TestPipInstallLegacy/requirements-use-cache":       true,
	"TestPipInstallLegacy/setuppy-verbose":              true,
	"TestPipInstallNativeSyntax/setuppy-verbose":        true,
	"TestPipInstallNativeSyntax/requirements-verbose":   true,
	"TestPipInstallLegacy/setuppy":                      true,
	"TestPipInstallNativeSyntax/requirements":           true,
	"TestPipInstallLegacy/requirements":                 true,
	"TestPipInstallNativeSyntax/requirements-use-cache": true,
	//Plugins
	"TestPluginInstallUninstallOfficialRegistry": true,
	//Xray
	"TestXrayAuditNugetMultiProject":                 true,
	"TestXrayBinaryScanWithBypassArchiveLimits":      true,
	"TestXrayBinaryScanJson":                         true,
	"TestXrayAuditMavenJson":                         true,
	"TestXrayAuditNugetSimpleJson":                   true,
	"TestXrayAuditNpmJson":                           true,
	"TestXrayAuditNpmSimpleJson":                     true,
	"TestXrayBinaryScanSimpleJsonWithProgress":       true,
	"TestXrayAuditNugetJson":                         true,
	"TestXrayAuditMavenSimpleJson":                   true,
	"TestXrayAuditDetectTech":                        true,
	"TestXrayBinaryScanJsonWithProgress":             true,
	"TestXrayBinaryScanSimpleJson":                   true,
	"TestXrayAuditYarnJson":                          true,
	"TestXrayAuditYarnSimpleJson":                    true,
	"TestXrayAuditMultiProjects":                     true,
	"TestXrayAuditPipenvJson":                        true,
	"TestXrayAuditPipenvSimpleJson":                  true,
	"TestXrayAuditPoetryJson":                        true,
	"TestXrayAuditPoetrySimpleJson":                  true,
	"TestXrayAuditPipJsonWithRequirementsFile":       true,
	"TestXrayAuditPipSimpleJsonWithRequirementsFile": true,
	"TestXrayAuditPipJson":                           true,
}

func skipTestIfMarkedTo(t *testing.T) {
	if !t.Skipped() && !*tests.TestMarkedToSkip && skipTests[t.Name()] {
		t.Skipf("Skipping '%v' test. To run the test remove it from 'skipTests.go(the map skipTests)'", t.Name())
	}
}
