package db

import (
	. "github.com/Frostman/aptomi/pkg/slinga/util"
	"os"
	"path/filepath"
)

// AptomiOject represents an aptomi entity, which gets stored in aptomi DB
type AptomiOject string

// AptomiCurrentRunDir is where results of the last run are stored
const AptomiCurrentRunDir = "last-run-results"

const (
	/*
		The following objects can be added to Aptomi
	*/

	// TypeCluster is k8s cluster or any other cluster
	TypeCluster AptomiOject = "cluster"

	// TypeService is service definitions
	TypeService AptomiOject = "service"

	// TypeContext is how service gets allocated
	TypeContext AptomiOject = "context"

	// TypeRules is global rules of the land
	TypeRules AptomiOject = "rules"

	// TypeDependencies is who requested what
	TypeDependencies AptomiOject = "dependencies"

	/*
		The following objects must be configured to point to external resources
	*/

	// TypeUsersFile is where users are stored (this is for file-based storage)
	TypeUsersFile AptomiOject = "users"

	// TypeUsersLDAP is where ldap configuration is stored\
	TypeUsersLDAP AptomiOject = "ldap"

	// TypeSecrets is where secret tokens are stored (later in Hashicorp Vault)
	TypeSecrets AptomiOject = "secrets"

	// TypeCharts is where binary charts/images are stored (later in external repo)
	TypeCharts AptomiOject = "charts"

	/*
		The following objects are generated by aptomi during or after dependency resolution via policy
	*/

	// TypeRevision holds revision number for the last successful aptomi run
	TypeRevision AptomiOject = "revision"

	// TypePolicyResolution holds usage data for components/dependencies
	TypePolicyResolution AptomiOject = "db"

	// TypeLogs contains debug logs
	TypeLogs AptomiOject = "logs"

	// TypeGraphics contains images generated by graphviz
	TypeGraphics AptomiOject = "graphics"
)

// Return aptomi DB directory
func getAptomiEnvVarAsDir(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		panic("Environment variable is not present. Must point to a directory")
	}
	if stat, err := os.Stat(value); err != nil || !stat.IsDir() {
		panic("Directory doesn't exist or error encountered")
	}
	return value
}

// GetAptomiBaseDir returns base directory, i.e. the value of APTOMI_DB environment variable
func GetAptomiBaseDir() string {
	return getAptomiEnvVarAsDir("APTOMI_DB")
}

// GetAptomiPolicyDir returns default aptomi policy dir
func GetAptomiPolicyDir() string {
	return filepath.Join(GetAptomiBaseDir(), "policy")
}

// GetAptomiDebugLogName returns filename for aptomi debug log (in db, but outside of current run)
func GetAptomiDebugLogName() string {
	dir := filepath.Join(GetAptomiBaseDir(), string(TypeLogs))
	if stat, err := os.Stat(dir); err != nil || !stat.IsDir() {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic("Directory can't be created or error encountered")
		}
	}
	if stat, err := os.Stat(dir); err != nil || !stat.IsDir() {
		panic("Directory can't be created or error encountered")
	}
	return filepath.Join(dir, "debug.log")
}

// GetAptomiObjectFilePatternYaml returns file pattern for aptomi objects (so they can be loaded from those files)
func GetAptomiObjectFilePatternYaml(baseDir string, aptomiObject AptomiOject) string {
	return filepath.Join(baseDir, "**", string(aptomiObject)+"*.yaml")
}

// GetAptomiObjectFilePatternTgz returns file pattern for tgz objects (so they can be loaded from those files)
func GetAptomiObjectFilePatternTgz(baseDir string, aptomiObject AptomiOject, chartName string) string {
	return filepath.Join(baseDir, "**", chartName+".tgz")
}

// GetAptomiObjectWriteFileGlobal returns file name for global aptomi objects (e.g. revision)
// It will place files into AptomiCurrentRunDir. It will create the corresponding directories if they don't exist
func GetAptomiObjectWriteFileGlobal(baseDir string, aptomiObject AptomiOject) string {
	return filepath.Join(baseDir, string(aptomiObject)+".yaml")
}

// GetAptomiObjectFileFromRun returns file name for aptomi objects (so they can be saved)
// It will place files into AptomiCurrentRunDir. It will create the corresponding directories if they don't exist
func GetAptomiObjectFileFromRun(baseDir string, runDir string, aptomiObject AptomiOject, fileName string) string {
	return filepath.Join(baseDir, runDir, string(aptomiObject), fileName)
}

// GetAptomiObjectWriteFileCurrentRun returns file name for aptomi objects (so they can be saved)
// It will place files into AptomiCurrentRunDir. It will create the corresponding directories if they don't exist
func GetAptomiObjectWriteFileCurrentRun(baseDir string, aptomiObject AptomiOject, fileName string) string {
	dir := filepath.Join(baseDir, AptomiCurrentRunDir, string(aptomiObject))
	if stat, err := os.Stat(dir); err != nil || !stat.IsDir() {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic("Directory can't be created or error encountered")
		}
	}
	if stat, err := os.Stat(dir); err != nil || !stat.IsDir() {
		panic("Directory can't be created or error encountered")
	}
	return filepath.Join(dir, fileName)
}

// CleanCurrentRunDirectory deletes contents of a "current run" directory
func CleanCurrentRunDirectory(baseDir string) {
	dir := filepath.Join(baseDir, AptomiCurrentRunDir)
	err := DeleteDirectoryContents(dir)
	if err != nil {
		panic("Directory contents can't be deleted")
	}
}
