// Copyright 2016-2018, Pulumi Corporation.
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

package integration

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pulumi/pulumi/sdk/v2/go/common/util/contract"
)

func TestPrefixer(t *testing.T) {
	byts := make([]byte, 0, 1000)
	buf := bytes.NewBuffer(byts)
	prefixer := newPrefixer(buf, "OK: ")
	_, err := prefixer.Write([]byte("\nsadsada\n\nasdsadsa\nasdsadsa\n"))
	contract.AssertNoError(err)
	assert.Equal(t, []byte("OK: \nOK: sadsada\nOK: \nOK: asdsadsa\nOK: asdsadsa\n"), buf.Bytes())
}

// Test that RunCommand writes the command's output to a log file.
