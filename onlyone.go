package onlyone

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/nightlyone/lockfile"
)

// Chk - check for a lock-file (for app already running)
func Chk(lockFileNames ...string) error {
	lockFileName, err := getLockFileName()
	if err != nil && len(lockFileNames) < 1 {
		return err
	}
	if lockFileName != "" {
		lockFileNames = append(lockFileNames, lockFileName)
	}

	var lock lockfile.Lockfile
	for _, lockFileName := range lockFileNames {
		lockFileName = strings.ReplaceAll(lockFileName, " ", "")
		os.MkdirAll(filepath.Dir(lockFileName), 0755)
		lock, err = lockfile.New(lockFileName)
		if err != nil {
			continue
		}
		err = lock.TryLock()
		if err == nil {
			return nil
		}
	}
	return err
}

// ChkFatal - as a Chk(), but call a panic() if error
func ChkFatal(lockFileNames ...string) {
	err := Chk(lockFileNames...)
	if err != nil {
		panic(err)
	}
}

func getLockFileName() (string, error) {
	exename, err := os.Executable()
	if err != nil {
		return "error", err
	}
	ext := filepath.Ext(exename)
	return strings.TrimSuffix(exename, ext) + ".pid", nil
}
