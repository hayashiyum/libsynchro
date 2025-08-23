package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Repository struct {
	RepositoryId  string `json:"repositoryId"`
	RepositoryUrl string `json:"repositoryUrl"`
	Username      string `json:"username"`
	Password      string `json:"password"`
}

type UploadFileData struct {
	ArtifactId     string `json:"artifactId"`
	GroupId        string `json:"groupId"`
	Version        string `json:"version"`
	UploadFileName string `json:"uploadFileName"` // Single file upload name
	Base64FileData string `json:"base64FileData"` // Single file Base64 data
	UploadFilePath string `json:"uploadFilePath"` // Multi file upload path
}

// ReadRepositoryListFile reads the content of the repository list file.
func (a *App) ReadRepositoryListFile() (string, error) {
	content, err := ioutil.ReadFile(REPOSITORY_LIST_FILE_PATH)
	if err != nil {
		// a.log("ERROR", "[ERROR] Failed to read repository list")
		return "", err
	}
	return string(content), nil
}

// WriteRepositoryListFile writes the given content to the repository list file.
func (a *App) WriteRepositoryListFile(content string) error {
	err := ioutil.WriteFile(REPOSITORY_LIST_FILE_PATH, []byte(content), 0644)
	if err != nil {
		errorMessage := fmt.Sprintf("[ERROR] Failed to write to file %s", REPOSITORY_LIST_FILE_PATH)
		a.log("ERROR", errorMessage)
		return err
	}
	return nil
}

// DecodeFile decodes a Base64 string and writes it to a temporary file.
func (a *App) DecodeFile(data UploadFileData) error {
	decodedData, err := base64.StdEncoding.DecodeString(data.Base64FileData)
	if err != nil {
		errorMessage := fmt.Sprintf("[ERROR] Failed to Base64 decode file %s", data.UploadFileName)
		a.log("ERROR", errorMessage)
		return err
	}
	outputPath := filepath.Join(TMP_DIR_PATH, data.UploadFileName)
	err = os.WriteFile(outputPath, decodedData, 0644)
	if err != nil {
		errorMessage := fmt.Sprintf("[ERROR] Failed to decode, failed write output file %s", data.UploadFileName)
		a.log("ERROR", errorMessage)
		return err
	}
	return nil
}

// AddServerXmlToFile adds a <server> XML entry to the settings file for Maven deployment.
func (a *App) AddServerXmlToFile(repository Repository) error {
	// Read jaruper-settings.xml
	content, err := os.ReadFile(JARUPER_SETTINGS_FILE_PATH)
	if err != nil {
		a.log("ERROR", "[ERROR] Failed to read jaruper-settings.xml")
		return err
	}

	// Generate the <server> XML string from the Repository object
	serverXml := fmt.Sprintf(
		"      <server>\n"+
			"        <id>%s</id>\n"+
			"        <username>%s</username>\n"+
			"        <password>%s</password>\n"+
			"      </server>",
		repository.RepositoryId,
		repository.Username,
		repository.Password)

	// Find the existing server settings and insert the new XML.
	// This example simply inserts it right before the </servers> tag.
	updatedContent := strings.Replace(string(content), "</servers>", fmt.Sprintf("%s\n</servers>", serverXml), 1)

	// Create and write to the new settings.xml file
	err = os.WriteFile(TMP_SETTINGS_FILE_PATH, []byte(updatedContent), 0644)
	if err != nil {
		a.log("ERROR", "[ERROR] Failed to write tmp settings.xml")
		return err
	}
	return nil
}

// DeleteTmpFiles deletes the temporary directory and its contents, then recreates it.
func (a *App) DeleteTmpFiles() error {
	// Check if the directory exists
	if _, err := os.Stat(TMP_DIR_PATH); os.IsNotExist(err) {
		// If the directory does not exist, exit successfully without doing anything
		return nil
	}
	// Delete the directory and all its contents
	err := os.RemoveAll(TMP_DIR_PATH)
	if err != nil {
		return fmt.Errorf("failed to delete temporary directory '%s': %w", TMP_DIR_PATH, err)
	}
	// Recreate the temporary directory
	if err := os.Mkdir(TMP_DIR_PATH, 0755); err != nil {
		return fmt.Errorf("failed to create directory '%s': %w", TMP_DIR_PATH, err)
	}
	return nil
}

// FileUpload uploads a file to a Maven repository using the `mvn deploy:deploy-file` command.
func (a *App) FileUpload(uploadType string, repository Repository, data UploadFileData) error {
	cmdArgs := []string{
		"deploy:deploy-file",
		fmt.Sprintf("-s%s", "config/tmp/settings.xml"),
		fmt.Sprintf("-DgroupId=%s", data.GroupId),
		fmt.Sprintf("-DartifactId=%s", data.ArtifactId),
		fmt.Sprintf("-Dversion=%s", data.Version),
		fmt.Sprintf("-Durl=%s", repository.RepositoryUrl),
		fmt.Sprintf("-DrepositoryId=%s", repository.RepositoryId),
	}

	if uploadType == "single" {
		cmdArgs = append(cmdArgs, fmt.Sprintf("-Dfile=%s/%s", TMP_DIR_PATH, data.UploadFileName))
		a.log("INFO", fmt.Sprintf("[INFO] Upload Information: GroupId=%s, ArtifactId=%s, Version=%s, File=%s, URL=%s",
			data.GroupId, data.ArtifactId, data.Version, data.UploadFileName, repository.RepositoryUrl))
	} else if uploadType == "multi" {
		cmdArgs = append(cmdArgs, fmt.Sprintf("-Dfile=%s", data.UploadFilePath))
		a.log("INFO", fmt.Sprintf("[INFO] Upload Information: GroupId=%s, ArtifactId=%s, Version=%s, File=%s, URL=%s",
			data.GroupId, data.ArtifactId, data.Version, data.UploadFilePath, repository.RepositoryUrl))
	} else {
		return fmt.Errorf("System error")
	}

	cmd := exec.Command("mvn", cmdArgs...)
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		errorMessage := fmt.Sprintf("[ERROR] Failed to get stdout pipe, %v", err)
		a.log("ERROR", errorMessage)
		return err
	}
	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		errorMessage := fmt.Sprintf("[ERROR] Failed to get stderr pipe, %v", err)
		a.log("ERROR", errorMessage)
		return err
	}

	// Start the command
	if err := cmd.Start(); err != nil {
		errorMessage := fmt.Sprintf("[ERROR] Failed to start Maven command, %v", err)
		a.log("ERROR", errorMessage)
		return err
	}

	// Read stdout and stderr concurrently and send to frontend
	go a.readPipe(stdoutPipe, "INFO")  // Maven's informational output
	go a.readPipe(stderrPipe, "ERROR") // Maven's error output

	// Wait for the command to finish
	if err := cmd.Wait(); err != nil {
		errorMessage := fmt.Sprintf("[ERROR] Maven command failed, %v", err)
		a.log("ERROR", errorMessage)
		return err
	}

	log.Printf("File upload process succeeded, %s:%s:%s", data.GroupId, data.ArtifactId, data.Version)
	successMessage := fmt.Sprintf("[INFO] Successfully uploaded, %s:%s:%s",
		data.GroupId, data.ArtifactId, data.Version)
	a.log("INFO", successMessage)
	return nil
}

// log emits a log message event to the frontend.
func (a *App) log(logType string, message string) {
	runtime.EventsEmit(a.ctx, "logMessage", map[string]string{"type": logType, "message": message})
}

// readPipe reads from an io.ReadCloser and emits each line as a log message to the frontend.
func (a *App) readPipe(pipe io.ReadCloser, logType string) {
	scanner := bufio.NewScanner(pipe)
	for scanner.Scan() {
		line := scanner.Text()
		runtime.EventsEmit(a.ctx, "logMessage", map[string]string{"type": logType, "message": line})
	}
}

// Initialize function for application setup.
// It checks for and creates necessary directories and files.
func (a *App) Initialize() error {
	// Check and create config/tmp directory
	if _, err := os.Stat(TMP_DIR_PATH); os.IsNotExist(err) {
		if err := os.MkdirAll(TMP_DIR_PATH, 0755); err != nil {
			errorMessage := fmt.Sprintf("[ERROR] Failed to create directory %s\n    %v", TMP_DIR_PATH, err)
			a.log("ERROR", errorMessage)
			return err
		}
	}

	// Check and create config/repository-list.json file
	if _, err := os.Stat(REPOSITORY_LIST_FILE_PATH); os.IsNotExist(err) {
		file, err := os.Create(REPOSITORY_LIST_FILE_PATH)
		if err != nil {
			errorMessage := fmt.Sprintf("[ERROR] Failed to create file %s\n    %v", REPOSITORY_LIST_FILE_PATH, err)
			a.log("ERROR", errorMessage)
			return err
		}
		defer file.Close()
		if _, err := file.WriteString("[]"); err != nil {
			errorMessage := fmt.Sprintf("[ERROR] Failed to write content to file  %s\n    %v", REPOSITORY_LIST_FILE_PATH, err)
			a.log("ERROR", errorMessage)
			return err
		}
	}
	return nil
}
