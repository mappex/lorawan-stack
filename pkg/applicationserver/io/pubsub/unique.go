// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
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

package pubsub

import (
	"fmt"
	"regexp"
)

var uniqueIDPattern = regexp.MustCompile("(.*)\\:(.*)")

// PubSubUID generates an unique ID based on an application unique ID and a pub/sub ID.
func PubSubUID(appUID, id string) string {
	return fmt.Sprintf("%s:%s", appUID, id)
}

// SplitPubSubUID parses a unique ID generated by `ID` and returns the application unique ID and the pub/sub ID.
func SplitPubSubUID(uid string) (string, string) {
	matches := uniqueIDPattern.FindStringSubmatch(uid)
	if len(matches) != 3 {
		panic(fmt.Sprintf("invalid uniqueID `%s` with matches %v", uid, matches))
	}
	return matches[1], matches[2]
}