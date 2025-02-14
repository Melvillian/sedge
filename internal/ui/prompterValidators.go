/*
Copyright 2022 Nethermind

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package ui

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/Melvillian/sedge/configs"
	"github.com/Melvillian/sedge/internal/utils"
)

var (
	ErrInvalidEthereumAddress = errors.New("invalid ethereum address")
	ErrInvalidInt64String     = errors.New("invalid int64 string")
	ErrInvalidFileExtension   = errors.New("invalid file extension")
	ErrInvalidURL             = errors.New("invalid URL")
	ErrInvalidDigitString     = errors.New("invalid number")
)

var digitsString = regexp.MustCompile("^0$|^[1-9][0-9]*$")

func EthAddressValidator(address string, allowEmpty bool) error {
	if allowEmpty && address == "" {
		return nil
	}
	if !utils.IsAddress(address) {
		return ErrInvalidEthereumAddress
	}
	return nil
}

func FilePathValidator(ans interface{}) error {
	if str, ok := ans.(string); ok {
		fileInfo, err := os.Stat(str)
		if err != nil {
			return err
		}
		if fileInfo.IsDir() {
			return errors.New("is a directory not a file")
		}
	}
	return nil
}

func DirPathValidator(ans string) error {
	fileInfo, err := os.Stat(ans)
	if err != nil {
		return err
	}
	if !fileInfo.IsDir() {
		return errors.New("is a file not a directory")
	}
	return nil
}

func URLValidator(ans interface{}) error {
	if str, ok := ans.(string); ok {
		if str == "" {
			return nil
		}
		if invalidURL, ok := utils.UriValidator([]string{str}); !ok {
			return fmt.Errorf("%w: %s", ErrInvalidURL, invalidURL)
		}
	}
	return nil
}

var int64Regex = regexp.MustCompile("^[0-9]{1,15}$")

func Int64Validator(ans interface{}) error {
	if str, ok := ans.(string); ok {
		if !int64Regex.MatchString(str) {
			return ErrInvalidInt64String
		}
		if _, err := strconv.ParseInt(str, 10, 64); err != nil {
			return ErrInvalidInt64String
		}
	}
	return nil
}

func fileExtensionValidator(extensions []string) func(ans interface{}) error {
	return func(ans interface{}) error {
		if input, ok := ans.(string); ok {
			input := strings.ToLower(input)
			for _, ext := range extensions {
				if filepath.Ext(input) == ext {
					return nil
				}
			}
		}
		return fmt.Errorf("%w: %s, must be one of these: %s", ErrInvalidFileExtension, ans, strings.Join(extensions, ", "))
	}
}

// DigitsStringValidator validates that the input is a valid string with only digits
func DigitsStringValidator(ttd string) error {
	if !digitsString.MatchString(ttd) {
		return ErrInvalidDigitString
	}
	return nil
}

func GraffitiValidator(str string) error {
	if len(str) > 16 {
		return fmt.Errorf(configs.ErrGraffitiLength, str, len(str))
	}
	return nil
}
