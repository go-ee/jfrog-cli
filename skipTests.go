package main

import (
	"github.com/jfrog/jfrog-cli/utils/tests"
	"testing"
)

var skipTests = map[string]bool{
	//Access
	"TestRefreshableAccessTokens": true,
	//Artifactory
	"TestArtifactoryProxy":			                          	true,
	"TestArtifactoryUploadAsArchiveWithIncludeDirs": 			true,
	"TestGitLfsCleanup":							true,
	"TestArtifactoryDownloadByBuildUsingSimpleDownloadWithProject":      	true,
	"TestArtifactoryDownloadWithEnvProject":                             	true,
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
	"TestXrayAuditPipSimpleJson":                     true,
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
	"TestXrayAuditGradleJson":                        true,
	"TestXrayAuditGradleSimpleJson":                  true,
	//NPM
	"TestNpmNativeSyntax":                            true,
	"TestNpmLegacy":                                  true,
	"TestNpmConditionalUpload":                       true,
	"TestNpmWithGlobalConfig":                        true,
	"TestNpmPublishDetailedSummary":                  true,
	"TestNpmPublishWithDeploymentView":               true,
	"TestNpmNativeSyntax/npm_i_with_scoped_project":  true,
	"TestNpmNativeSyntax/npm_i_with_npmrc_project":   true,
	"TestNpmLegacy/npm_i_with_production":            true,
	"TestNpmNativeSyntax/npm_i_with_production":      true,
	"TestNpmLegacy/npm_i_with_npmrc_project":         true,
	"TestNpmNativeSyntax/npm_postinstall":            true,
	"TestNpmLegacy/npm_i_with_scoped_project":        true,
	"TestNpmLegacy/npm_postinstall":                  true,
	"TestNpmLegacy/npm_i_with_module":                true,
	"TestNpmNativeSyntax/npm_i_with_module":          true,
	"TestNpmLegacy/npm_ci":                           true,
	"TestNpmLegacy/npm_ci_with_module":               true,
	"TestNpmNativeSyntax/npm_ci_with_module":         true,
	"TestNpmNativeSyntax/npm_ci":                     true,
	"TestNpmPackInstall":                             true,
}

func skipTestIfMarkedTo(t *testing.T) {
	if !t.Skipped() && !*tests.TestMarkedToSkip && skipTests[t.Name()] {
		t.Skipf("Skipping '%v' test. To run the test remove it from 'skipTests.go(the map skipTests)'", t.Name())
	}
}
