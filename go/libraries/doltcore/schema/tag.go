// Copyright 2020 Liquidata, Inc.
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

package schema

import (
	"math"
	"math/rand"
	"time"
)

const (
	// ReservedTagMin is the start of a range of tags which the user should not be able to use in their schemas.
	//ReservedTagMin uint64 = 1 << 63
	ReservedTagMin uint64 = 1 << 50

	// InvalidTag is used as an invalid tag
	InvalidTag uint64 = math.MaxUint64
)


const (
	DocNameTag = ReservedTagMin
	DocTextTag = ReservedTagMin + 1
)

var randGen = rand.New(rand.NewSource(time.Now().UnixNano()))

func AutoGenerateTag(sch Schema) uint64 {
	var maxTagVal uint64 = 128 * 128

	allCols := sch.GetAllCols()
	for maxTagVal/2 < uint64(allCols.Size()) {
		if maxTagVal == ReservedTagMin-1 {
			panic("There is no way anyone should ever have this many columns.  You are a bad person if you hit this panic.")
		} else if maxTagVal*128 < maxTagVal {
			maxTagVal = ReservedTagMin - 1
		} else {
			maxTagVal = maxTagVal * 128
		}
	}

	var randTag uint64
	for {
		randTag = uint64(randGen.Int63n(int64(maxTagVal)))

		if _, ok := allCols.GetByTag(randTag); !ok {
			break
		}
	}

	return randTag
}


