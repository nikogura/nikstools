package nikstools

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
	"os/exec"
)

const DEFAULT_PROFILE_DIR = ".nikstools"

// SetProfile updated user's shell profile to include kubernetes utility functions
func SetProfile(homeDir string) (profileTypeName string, err error) {
	shellFuncData, err := GetBashfunc()
	if err != nil {
		err = errors.Wrapf(err, "failed getting shell func data")
		return profileTypeName, err
	}

	profileSourceString := fmt.Sprintf("if [ -f ~/%s ]; then . ~/%s; fi", DEFAULT_PROFILE_DIR, DEFAULT_PROFILE_DIR)
	nikstoolsProfile := fmt.Sprintf("%s/%s", homeDir, DEFAULT_PROFILE_DIR)

	err = os.WriteFile(nikstoolsProfile, shellFuncData, 0644)
	if err != nil {
		err = errors.Wrapf(err, "failed writing nikstools profile to file %s", nikstoolsProfile)
		return profileTypeName, err
	}

	fmt.Printf("profile created successfully at %s\n", nikstoolsProfile)

	zshrc := fmt.Sprintf("%s/.zshrc", homeDir)
	file, zshRcErr := os.OpenFile(zshrc, os.O_APPEND|os.O_WRONLY, 0644)
	if zshRcErr != nil {
		fmt.Println("~/.zshrc not found")
	} else {
		profileTypeName = "zshrc"
		ProfileBackup("", zshrc)
		fmt.Fprintf(file, "\n%s\n", profileSourceString)
		fmt.Printf("%s updated to source ~/%s, reload the shell to start using the functions\n", "~/.zshrc", DEFAULT_PROFILE_DIR)
		return profileTypeName, err
	}

	bashrc := fmt.Sprintf("%s/.bashrc", homeDir)
	file, err = os.OpenFile(bashrc, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("~/.bashrc not found")
	} else {
		profileTypeName = "bashrc"
		ProfileBackup("", bashrc)
		fmt.Fprintf(file, "\n%s\n", profileSourceString)
		fmt.Printf("%s updated to source ~/%s, reload the shell to start using the functions\n", "~/.bashrc", DEFAULT_PROFILE_DIR)
		return profileTypeName, err
	}

	bashprofile := fmt.Sprintf("%s/.bash_profile", homeDir)
	file, err = os.OpenFile(bashprofile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("~/.bash_profile not found")
	} else {
		profileTypeName = "bash_profile"
		ProfileBackup("", bashprofile)
		fmt.Fprintf(file, "\n%s\n", profileSourceString)
		fmt.Printf("%s updated to source ~/%s, reload the shell to start using the functions\n", "~/.bash_profile", DEFAULT_PROFILE_DIR)
		return profileTypeName, err
	}

	profile := fmt.Sprintf("%s/.profile", homeDir)

	file, err = os.OpenFile(profile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("~/.profile not found")
	} else {
		profileTypeName = "profile"
		ProfileBackup("", profile)
		fmt.Fprintf(file, "\n%s\n", profileSourceString)
		fmt.Printf("%s updated to source ~/%s, reload the shell to start using the functions\n", "~/.profile", DEFAULT_PROFILE_DIR)
		return profileTypeName, err
	}

	defer file.Close()

	return profileTypeName, err
}

func ProfileBackup(tmpDir string, fileName string) {
	var err error
	if tmpDir == "" {
		tmpDir, err = os.MkdirTemp("", "backup")
		if err != nil {
			log.Fatal("Error creating temp dir\n")
		}
	}
	fmt.Printf("taking back up of existing config at: %s/kubeconfig.ProfileBackup\n", tmpDir)
	commandStr := fmt.Sprintf("cp %s %s/kubeconfig.ProfileBackup", fileName, tmpDir)
	cmd := exec.Command("bash", "-c", commandStr)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err = cmd.Run()
	if err != nil {
		fmt.Println("error unable to take backup of profile file:", err)
		defer os.RemoveAll(tmpDir)
	}
}

func GetBashfunc() (profilebytes []byte, err error) {

	// TODO get bash funcs from embedded fs
	return profilebytes, err
}
