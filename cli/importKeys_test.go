/*
Copyright 2022 Nethermind

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cli

import (
	"io"
	"path/filepath"
	"testing"

	"github.com/Melvillian/sedge/cli/actions"
	"github.com/Melvillian/sedge/configs"
	"github.com/Melvillian/sedge/internal/pkg/dependencies"
	sedge_mocks "github.com/Melvillian/sedge/mocks"
	"github.com/golang/mock/gomock"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestImportKeys_NumberOfArguments(t *testing.T) {
	// Silence logger
	log.SetOutput(io.Discard)

	tests := []struct {
		name string
		args []string
	}{
		{
			name: "no flags",
			args: []string{},
		},
		{
			name: "with flags",
			args: []string{"--network", "goerli"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := ImportKeysCmd(nil, nil)
			cmd.SetArgs(tt.args)
			cmd.SetOutput(io.Discard)
			err := cmd.Execute()
			assert.ErrorIs(t, err, ErrInvalidNumberOfArguments)
		})
	}
}

func TestImportKeys_ArgsAndFlags(t *testing.T) {
	// Silence logger
	log.SetOutput(io.Discard)

	tests := []struct {
		name              string
		args              []string
		expectedSetupOpts actions.SetupContainersOptions
		expectedOptions   actions.ImportValidatorKeysOptions
	}{
		{
			name: "no flags",
			args: []string{"lighthouse"},
			expectedSetupOpts: actions.SetupContainersOptions{
				GenerationPath: configs.DefaultAbsSedgeDataPath,
				Services:       []string{validator},
			},
			expectedOptions: actions.ImportValidatorKeysOptions{
				ValidatorClient: "lighthouse",
				Network:         "mainnet",
				GenerationPath:  configs.DefaultAbsSedgeDataPath,
				From:            filepath.Join(configs.DefaultAbsSedgeDataPath, "keystore"),
			},
		},
		{
			name: "with flags",
			args: []string{
				"prysm",
				"--network", "goerli",
				"--from", "/tmp/keystore",
				"--path", "/tmp/sedge",
				"--start-validator",
			},
			expectedSetupOpts: actions.SetupContainersOptions{
				GenerationPath: "/tmp/sedge",
				Services:       []string{validator},
			},
			expectedOptions: actions.ImportValidatorKeysOptions{
				ValidatorClient: "prysm",
				Network:         "goerli",
				StartValidator:  true,
				GenerationPath:  "/tmp/sedge",
				From:            "/tmp/keystore",
			},
		},
		{
			name: "with shorthand flags",
			args: []string{
				"teku",
				"-n", "goerli",
				"--from", "/tmp/keystore",
				"-p", "/tmp/sedge",
				"--stop-validator",
			},
			expectedSetupOpts: actions.SetupContainersOptions{
				GenerationPath: "/tmp/sedge",
				Services:       []string{validator},
			},
			expectedOptions: actions.ImportValidatorKeysOptions{
				ValidatorClient: "teku",
				Network:         "goerli",
				From:            "/tmp/keystore",
				GenerationPath:  "/tmp/sedge",
				StopValidator:   true,
			},
		},
		{
			name: "with custom configs",
			args: []string{
				"lighthouse",
				"--custom-config", "/tmp/config",
				"--custom-genesis", "/tmp/genesis",
				"--custom-deploy-block", "custom-deploy-block",
				"--container-tag", "test-tag",
			},
			expectedSetupOpts: actions.SetupContainersOptions{
				GenerationPath: configs.DefaultAbsSedgeDataPath,
				Services:       []string{validator},
			},
			expectedOptions: actions.ImportValidatorKeysOptions{
				ValidatorClient: "lighthouse",
				Network:         "mainnet",
				GenerationPath:  configs.DefaultAbsSedgeDataPath,
				From:            filepath.Join(configs.DefaultAbsSedgeDataPath, "keystore"),
				ContainerTag:    "test-tag",
				CustomConfig: actions.ImportValidatorKeysCustomOptions{
					NetworkConfigPath: "/tmp/config",
					GenesisPath:       "/tmp/genesis",
					DeployBlockPath:   "custom-deploy-block",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actions := sedge_mocks.NewMockSedgeActions(gomock.NewController(t))
			depsMgr := sedge_mocks.NewMockDependenciesManager(gomock.NewController(t))

			gomock.InOrder(
				depsMgr.EXPECT().Check([]string{dependencies.Docker}).Return([]string{dependencies.Docker}, nil).Times(1),
				depsMgr.EXPECT().DockerEngineIsOn().Return(nil).Times(1),
				depsMgr.EXPECT().DockerComposeIsInstalled().Return(nil).Times(1),
				actions.EXPECT().ValidateDockerComposeFile(filepath.Join(tt.expectedSetupOpts.GenerationPath, "docker-compose.yml")).Return(nil).Times(1),
				actions.EXPECT().SetupContainers(tt.expectedSetupOpts).Times(1),
				actions.EXPECT().ImportValidatorKeys(tt.expectedOptions).Times(1),
			)

			cmd := ImportKeysCmd(actions, depsMgr)
			cmd.SetArgs(tt.args)
			err := cmd.Execute()
			assert.NoError(t, err)
		})
	}
}
