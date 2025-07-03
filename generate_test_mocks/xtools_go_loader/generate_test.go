// Copyright 2015 Peter Goetz
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

package mockgen_test

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	"github.com/phildrip/pegomock/v4/pegomock/filehandling"
)

func TestMockGeneration(t *testing.T) {
	RunSpecs(t, "Generating mocks with golang.org/x/tools/go/packages")
}

var _ = It("Generate mocks", func() {
	filehandling.GenerateMockFile(
		[]string{"github.com/phildrip/pegomock/v4/test_interface", "Display"},
		"../../mock_display_test.go", "MockDisplay", "pegomock_test",
		"", false, os.Stdout)

	filehandling.GenerateMockFile(
		[]string{"github.com/phildrip/pegomock/v4/test_interface", "GenericDisplay"},
		"../../mock_generic_display_test.go", "MockGenericDisplay", "pegomock_test",
		"", false, os.Stdout)
})
