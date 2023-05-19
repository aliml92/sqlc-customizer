package util

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
)

func FileExists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}

func OverridePackageName(filePath, newName, oldName string) error {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	modifiedData := strings.Replace(string(fileData), oldName, newName, 1)
	err = os.WriteFile(filePath, []byte(modifiedData), 0644)
	if err != nil {
		return err
	}
	return nil
}

func MoveFile(source string, destination string) error {
	// Open original file
	originalFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer originalFile.Close()


	// Create new file
	directory := path.Dir(destination)
	// create the directory if it does not exist
	err = os.MkdirAll(directory, 0755)
	if err != nil {
		return err
	}
	newFile, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer newFile.Close()

	// Copy the bytes to destination from source
	bytesWritten, err := io.Copy(newFile, originalFile)
	if err != nil {
		return err
	}
	fmt.Printf("Copied %d bytes.\n", bytesWritten)

	// Commit the file contents
	// Flushes memory to disk
	err = newFile.Sync()
	if err != nil {
		return err
	}
	// delete the original file
	err = os.Remove(source)
	if err != nil {
		return err
	}
	return nil
}

func OverrideStuct(filePath, oldWord, newWord string) (int, error) {	
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var (
		modifiedLines []string
		overrideCount int
	)


	for scanner.Scan() {
		line := scanner.Text()
		firstIndex, lastIndex := -1, -1

		if strings.Contains(line, oldWord) {
			firstIndex = strings.Index(line, oldWord)
			if firstIndex > 0 && ( line[firstIndex-1] == ' ' || line[firstIndex-1] == '*' || line[firstIndex-1] == '(' ){
				lastIndex = firstIndex + len(oldWord)
				if lastIndex < len(line) && ( line[lastIndex] == ' ' || line[lastIndex] == '*' || line[lastIndex] == ')' || line[lastIndex] == ',' || line[lastIndex] == '\n' ){
					line = line[:firstIndex] + newWord + line[lastIndex:]
					modifiedLines = append(modifiedLines, line)
					overrideCount++
					continue
				} else if lastIndex == len(line) {
					line = line[:firstIndex] + newWord
					modifiedLines = append(modifiedLines, line)
					overrideCount++
					continue
				}

			}
			firstIndex = strings.LastIndex(line, oldWord)
			if firstIndex > 0 && ( line[firstIndex-1] == ' ' || line[firstIndex-1] == '*' || line[firstIndex-1] == '(' ){
				lastIndex = firstIndex + len(oldWord)
				if lastIndex < len(line) && ( line[lastIndex] == ' ' || line[lastIndex] == '*' || line[lastIndex] == ')' || line[lastIndex] == ',' || line[lastIndex] == '\n' ){
					line = line[:firstIndex] + newWord + line[lastIndex:]
					modifiedLines = append(modifiedLines, line)
					overrideCount++	
					continue
				}
			}  
		} 
		modifiedLines = append(modifiedLines, line)
		continue
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	// Truncate the file and write the modified contents back
	file.Truncate(0)
	file.Seek(0, 0)
	writer := bufio.NewWriter(file)
	for _, line := range modifiedLines {
		fmt.Fprintln(writer, line)
	}
	writer.Flush()

	return overrideCount, nil
}

func AppendImportEntry(filePath string, newImportPath string) error {
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	importStatement := fmt.Sprintf("\t\"%s\"\n", newImportPath)
	importBlockEnd := ")\n"
	importIndex := -1
	importBlockEndIndex := -1

	if strings.Contains(string(fileData), importStatement) {
		fmt.Printf("Import entry for %s already exists\n", newImportPath)
		return nil
	}

	importIndex = strings.Index(string(fileData), "import (")
	if importIndex == -1 {
		importIndex = strings.Index(string(fileData), "package")
		importBlockEndIndex = strings.Index(string(fileData[importIndex:]), "\n")
		importStatement = fmt.Sprintf("\n\nimport (\n\t\"%s\"\n)\n", newImportPath)
		insertIndex := importIndex + importBlockEndIndex
		fileData = insertSlice(fileData, []byte(importStatement), insertIndex)
	
		err = os.WriteFile(filePath, fileData, 0644)
		if err != nil {
			return err
		}
	
		return nil
	}
	importBlockEndIndex = strings.Index(string(fileData[importIndex:]), importBlockEnd)
	if importBlockEndIndex == -1 {
		return errors.New("import end block not found")
	}
	insertIndex := importIndex + importBlockEndIndex
	fileData = insertSlice(fileData, []byte(importStatement), insertIndex)
	err = os.WriteFile(filePath, fileData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func insertSlice(slice, insertion []byte, index int) []byte {
	return append(slice[:index], append(insertion, slice[index:]...)...)
}

func AppendOmitempty(filePath string) (int, error) {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var (
		modifiedLines []string
		overrideCount int
	)


	for scanner.Scan() {
		line := scanner.Text()
		lastIndex := -1
		if strings.Contains(line, "json:") {
			lastIndex = strings.LastIndex(line, "\"")
			if lastIndex > 0 {
				line = line[:lastIndex] + ",omitempty" + line[lastIndex:]
				modifiedLines = append(modifiedLines, line)
				overrideCount++
				continue
			}
		}
		modifiedLines = append(modifiedLines, line)
		continue
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	// Truncate the file and write the modified contents back
	file.Truncate(0)
	file.Seek(0, 0)
	writer := bufio.NewWriter(file)
	for _, line := range modifiedLines {
		fmt.Fprintln(writer, line)
	}
	writer.Flush()

	return overrideCount, nil
}