// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package httpcheckreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/httpcheckreceiver"

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/config/confighttp"
	"go.uber.org/multierr"
)

func TestValidate(t *testing.T) {
	testCases := []struct {
		desc        string
		cfg         *Config
		expectedErr error
	}{
		{
			desc: "missing endpoint",
			cfg: &Config{
				Targets: []*targetConfig{
					{
						HTTPClientSettings: confighttp.HTTPClientSettings{},
					},
				},
			},
			expectedErr: multierr.Combine(
				errMissingEndpoint,
			),
		},
		{
			desc: "invalid endpoint",
			cfg: &Config{
				Targets: []*targetConfig{
					{
						HTTPClientSettings: confighttp.HTTPClientSettings{
							Endpoint: "invalid://endpoint:  12efg",
						},
					},
				},
			},
			expectedErr: multierr.Combine(
				fmt.Errorf("%w: %s", errInvalidEndpoint, `parse "invalid://endpoint:  12efg": invalid port ":  12efg" after host`),
			),
		},
		{
			desc: "invalid config with multiple targets",
			cfg: &Config{
				Targets: []*targetConfig{
					{
						HTTPClientSettings: confighttp.HTTPClientSettings{
							Endpoint: "https://localhost:80",
						},
					},
					{
						HTTPClientSettings: confighttp.HTTPClientSettings{
							Endpoint: "invalid://endpoint:  12efg",
						},
					},
				},
			},
			expectedErr: multierr.Combine(
				fmt.Errorf("%w: %s", errInvalidEndpoint, `parse "invalid://endpoint:  12efg": invalid port ":  12efg" after host`),
			),
		},
		{
			desc: "valid config",
			cfg: &Config{
				Targets: []*targetConfig{
					{
						HTTPClientSettings: confighttp.HTTPClientSettings{
							Endpoint: "https://opentelemetry.io",
						},
					},
				},
			},
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			actualErr := tc.cfg.Validate()
			if tc.expectedErr != nil {
				require.EqualError(t, actualErr, tc.expectedErr.Error())
			} else {
				require.NoError(t, actualErr)
			}

		})
	}
}
