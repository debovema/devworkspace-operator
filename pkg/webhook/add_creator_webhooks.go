//
// Copyright (c) 2019-2020 Red Hat, Inc.
// This program and the accompanying materials are made
// available under the terms of the Eclipse Public License 2.0
// which is available at https://www.eclipse.org/legal/epl-2.0/
//
// SPDX-License-Identifier: EPL-2.0
//
// Contributors:
//   Red Hat, Inc. - initial API and implementation
package webhook

import "github.com/che-incubator/che-workspace-operator/pkg/webhook/workspace"

func init() {
	configureWebhookTasks = append(configureWebhookTasks, workspace.Configure)
}
