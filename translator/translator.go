package translator

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// The suffix of a mojo-gojo file.
const suffix = ".mg"

// TranslateDirectory processes all mojo-gojo files in the provided directory,
// translates them to go files, and saves the results.
// It returns a list of the created go files.
func TranslateDirectory(dir string) ([]string, error) {
	var generatedFiles []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), suffix) {
			data, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			translatedContent := translateEmojiToKeyword(string(data))

			tempFilePath := strings.TrimSuffix(path, suffix) + ".go"
			err = ioutil.WriteFile(tempFilePath, []byte(translatedContent), 0644)
			if err != nil {
				return err
			}

			generatedFiles = append(generatedFiles, tempFilePath)
		}
		return nil
	})

	return generatedFiles, err
}

// translateEmojiToKeyword translates the content of a mojo-gojo file to a go file.
func translateEmojiToKeyword(content string) string {
	for emoji, keyword := range emojiToKeyword {
		content = strings.ReplaceAll(content, emoji, keyword)
	}
	return content
}
