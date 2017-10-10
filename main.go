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
	"gopkg.in/yaml.v2"

	// Internal Imports
	"github.com/MatthewHartstonge/b64secrets/models"
)

func createSecretsFile(fp string, wpf string) {
	f, err := ioutil.ReadFile(fp)

	// Split on documents
	documents := strings.Split(string(f), "\n---")

	base64SecretsFile, err := os.Create(wpf)
	if err != nil {
		panic(err)
	}
	defer base64SecretsFile.Close()

	for _, doc := range documents {
		secretDefinition := models.Secret{}
		yaml.Unmarshal([]byte(doc), &secretDefinition)

		if secretDefinition.Kind != "Secret" {
			continue
		}

		for key, value := range secretDefinition.Data {
			secretDefinition.Data[key] = base64.StdEncoding.EncodeToString([]byte(value))
		}

		base64doc, err := yaml.Marshal(secretDefinition)
		if err != nil {
			panic(err)
		}

		base64SecretsFile.WriteString("---\n")
		base64SecretsFile.Write(base64doc)
	}

	// Flush file to disk
	base64SecretsFile.Sync()
}

func main() {
	glob, err := filepath.Glob("./secrets-*.yml")
	if err != nil {
		panic(err)
	}

	for _, readFilepath := range glob {
		writeFilepath := fmt.Sprintf("%s.base64.yml", strings.TrimRight(readFilepath, ".yml"))
		createSecretsFile(readFilepath, writeFilepath)
	}
}
