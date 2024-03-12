// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Package tests contains test cases. To run the tests go to tests directory and run:
// RUN_TESTBED=1 go test -v

package tests

// This file contains Test functions which initiate the tests. The tests can be either
// coded in this file or use scenarios from perf_scenarios.go.

import (
	"testing"

	"github.com/open-telemetry/opentelemetry-collector-contrib/testbed/testbed"
)

// TestMain is used to initiate setup, execution and tear down of testbed.
func TestMain(m *testing.M) {
	testbed.DoTestMain(m, performanceResultsSummary)
}

// 头部采样 100%
func TestTraceBallastWithoutTailSampling(t *testing.T) {
	processors := map[string]string{
		"batch": `
  batch:
    send_batch_size: 1_000
    send_batch_max_size: 5_000
    timeout: 1s
`,
	}
	ScenarioSPSWithAttrs(
		t,
		[]string{},
		[]TestCase{
			{
				SPS:            5000,
				attrCount:      10,
				attrSizeByte:   10,
				expectedMaxCPU: 200,
				expectedMaxRAM: 2048,
				resultsSummary: performanceResultsSummary,
			},
			{
				SPS:            10000,
				attrCount:      10,
				attrSizeByte:   10,
				expectedMaxCPU: 200,
				expectedMaxRAM: 2048,
				resultsSummary: performanceResultsSummary,
			},
			{
				SPS:            15000,
				attrCount:      10,
				attrSizeByte:   10,
				expectedMaxCPU: 200,
				expectedMaxRAM: 2048,
				resultsSummary: performanceResultsSummary,
			},
			{
				SPS:            20000,
				attrCount:      10,
				attrSizeByte:   10,
				expectedMaxCPU: 200,
				expectedMaxRAM: 2048,
				resultsSummary: performanceResultsSummary,
			},
			{
				SPS:            40000,
				attrCount:      10,
				attrSizeByte:   10,
				expectedMaxCPU: 200,
				expectedMaxRAM: 2048,
				resultsSummary: performanceResultsSummary,
			},
		},
		processors,
		nil,
	)
}

func TestTraceBallastWithTailSampling(t *testing.T) {
	processors := map[string]string{
		"batch": `
  batch:
    send_batch_size: 1_000
    send_batch_max_size: 5_000
    timeout: 1s
`,
		"tail_sampling": `
  tail_sampling:
    decision_wait: 5s
    policies: [
      {
        name: test-policy,
        type: always_sample
      }
    ]`,
	}
	ScenarioSPSWithAttrs(
		t,
		[]string{},
		[]TestCase{
			{
				SPS:            5000,
				attrCount:      10,
				attrSizeByte:   10,
				expectedMaxCPU: 200,
				expectedMaxRAM: 2048,
				resultsSummary: performanceResultsSummary,
			},
			{
				SPS:            10000,
				attrCount:      10,
				attrSizeByte:   10,
				expectedMaxCPU: 200,
				expectedMaxRAM: 2048,
				resultsSummary: performanceResultsSummary,
			},
			{
				SPS:            15000,
				attrCount:      10,
				attrSizeByte:   10,
				expectedMaxCPU: 200,
				expectedMaxRAM: 2048,
				resultsSummary: performanceResultsSummary,
			},
			{
				SPS:            20000,
				attrCount:      10,
				attrSizeByte:   10,
				expectedMaxCPU: 200,
				expectedMaxRAM: 2048,
				resultsSummary: performanceResultsSummary,
			},
			{
				SPS:            40000,
				attrCount:      10,
				attrSizeByte:   10,
				expectedMaxCPU: 200,
				expectedMaxRAM: 2048,
				resultsSummary: performanceResultsSummary,
			},
		},
		processors,
		nil,
	)
}
