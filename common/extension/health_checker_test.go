/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package extension

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

import (
	"github.com/apache/dubbo-go/cluster/router"
	"github.com/apache/dubbo-go/common"
	"github.com/apache/dubbo-go/protocol"
)

func TestGetHealthChecker(t *testing.T) {
	SetHealthChecker("mock", newMockHealthCheck)
	checker := GetHealthChecker("mock", common.NewURLWithOptions())
	assert.NotNil(t, checker)
}

type mockHealthChecker struct {
}

func (m mockHealthChecker) IsHealthy(invoker protocol.Invoker) bool {
	return true
}

func newMockHealthCheck(_ *common.URL) router.HealthChecker {
	return &mockHealthChecker{}
}
