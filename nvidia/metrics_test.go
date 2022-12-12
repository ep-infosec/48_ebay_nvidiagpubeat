//******************************************************************
//Copyright 2018 eBay Inc.
//Architect/Developer: Deepak Vasthimal

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at

// https://www.apache.org/licenses/LICENSE-2.0

//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//******************************************************************

package nvidia

import "testing"

func TestMetrics_NewMetrics(t *testing.T) {
	m := NewMetrics()
	output, _ := m.Get("test", "query")
	for _, o := range output {
		if o != nil {
			t.Errorf("output has to be nil.")
		}
	}
}
