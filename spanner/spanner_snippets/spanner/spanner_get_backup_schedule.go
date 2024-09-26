// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package spanner

// [START spanner_get_backup_schedule]

import (
	"context"
	"fmt"
	"io"

	database "cloud.google.com/go/spanner/admin/database/apiv1"
	"cloud.google.com/go/spanner/admin/database/apiv1/databasepb"
)

func getBackupSchedule(w io.Writer, dbName string, scheduleId string) error {
	ctx := context.Background()

	client, err := database.NewDatabaseAdminClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	req := databasepb.GetBackupScheduleRequest{
		Name: fmt.Sprintf("%s/backupSchedules/%s", dbName, scheduleId),
	}

	res, err := client.GetBackupSchedule(ctx, &req)
	if err != nil {
		return err
	}

	fmt.Fprintf(w, "Backup schedule: %s", res)
	return nil
}

// [END spanner_get_backup_schedule]
