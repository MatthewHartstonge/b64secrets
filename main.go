// Copyright 2017 Matthew Hartstonge
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	// Standard Library Imports
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	// External Imports
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"

	// Internal Imports
	"github.com/MatthewHartstonge/b64secrets/models"
)

const (
	globPattern = "./*/*.yml"
)

func createSecretsFile(fp string, wpf string) error {
	logger := log.WithFields(log.Fields{
		"method": "createSecretsFile",
	})

	f, err := ioutil.ReadFile(fp)
	if err != nil {
		logger.WithFields(log.Fields{
			"error": err,
		}).Error("file read failure")
		return err
	}

	// Split on documents
	documents := strings.Split(string(f), "\n---")
	var parsedDocs [][]byte
	for _, doc := range documents {
		secretDefinition := models.Secret{}
		yaml.Unmarshal([]byte(doc), &secretDefinition)

		if secretDefinition.Kind != "Secret" {
			continue
		} else if secretDefinition.Type != "Opaque" {
			continue
		}

		for key, value := range secretDefinition.Data {
			secretDefinition.Data[key] = base64.StdEncoding.EncodeToString([]byte(value))
		}

		base64doc, err := yaml.Marshal(secretDefinition)
		if err != nil {
			logger.WithFields(log.Fields{
				"error": err,
			}).Error("yaml marshaling failure")
			return err
		}
		parsedDocs = append(parsedDocs, base64doc)
	}

	if len(parsedDocs) > 0 {
		base64SecretsFile, err := os.Create(wpf)
		if err != nil {
			logger.WithFields(log.Fields{
				"error": err,
			}).Error("file write failure")
			return err
		}
		defer base64SecretsFile.Close()

		for i := range parsedDocs {
			_, err := base64SecretsFile.WriteString("---\n")
			if err != nil {
				logger.WithFields(log.Fields{
					"error": err,
				}).Error("file write failure")
			}
			_, err = base64SecretsFile.Write(parsedDocs[i])
			if err != nil {
				logger.WithFields(log.Fields{
					"error": err,
				}).Error("file write failure")
			}
		}

		// Flush file to disk
		err = base64SecretsFile.Sync()
		if err != nil {
			logger.WithFields(log.Fields{
				"error": err,
			}).Error("file write failure")
		}

		logger.WithFields(log.Fields{
			"originalPath":  fp,
			"conformedPath": wpf,
		}).Info("Created conformed secrets file")
	}

	return nil
}

func main() {
	logger := log.WithFields(log.Fields{
		"method": "main",
	})

	logger.Info("b64secrets is converting secrets..")
	glob, err := filepath.Glob(globPattern)
	if err != nil {
		logger.WithFields(log.Fields{
			"error": err,
		}).Fatal("File globbing failed")
	}

	for _, readFilepath := range glob {
		// Don't create secrets from already converted secrets files
		if strings.Contains(readFilepath, ".base64.yml") {
			continue
		}

		writeFilepath := fmt.Sprintf("%s.base64.yml", strings.TrimRight(readFilepath, ".yml"))
		err := createSecretsFile(readFilepath, writeFilepath)
		if err != nil {
			logger.Error("error writing secrets file... Continuing...")
		}
	}

	logger.Info("b64secrets file conversions completed!")
}
