// Copyright 2014 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package driver

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

// newTempFile returns a new output file in dir with the provided prefix and suffix.
func newTempFile(dir, prefix, suffix string) (*os.File, error) {
	for index := 1; index < 10000; index++ {
		path := filepath.Join(dir, fmt.Sprintf("%s%03d%s", prefix, index, suffix))
		if _, err := os.Stat(path); err != nil {
			return os.Create(path)
		}
	}
	// Give up
	return nil, fmt.Errorf("could not create file of the form %s%03d%s", prefix, 1, suffix)
}

var tempFiles []string
var tempFilesMu = sync.Mutex{}

// deferDeleteTempFile marks a file to be deleted by next call to Cleanup()
func deferDeleteTempFile(path string) {
	tempFilesMu.Lock()
	tempFiles = append(tempFiles, path)
	tempFilesMu.Unlock()
}

// cleanupTempFiles removes any temporary files selected for deferred cleaning.
func cleanupTempFiles() {
	tempFilesMu.Lock()
	for _, f := range tempFiles {
		os.Remove(f)
	}
	tempFiles = nil
	tempFilesMu.Unlock()
}
