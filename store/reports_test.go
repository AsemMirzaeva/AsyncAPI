package store_test

import (
	"asyncapi/fixtures"
	"asyncapi/store"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestReportStore(t *testing.T) {
	env := fixtures.NewTestEnv(t)
	cleanup := env.SetupDb(t)
	t.Cleanup(func() {
		cleanup(t)
	})

	ctx := context.Background()
	reportStore := store.NewReportStore(env.Db)
	userStore := store.NewUserStore(env.Db)
	user, err :=userStore.CreateUser(ctx, "test@test.com", "secretpassword")
	require.NoError(t, err)

	now := time.Now()

	report, err := reportStore.Create(ctx, user.Id, "monsters")
	require.NoError(t, err)
	require.Equal(t, user.Id, report.UserId)
	require.Equal(t, "monsters", report.ReportType)
	require.Less(t, now.UnixNano(), report.CompletedAt.UnixNano())

	startedAt := report.CreatedAt.Add(time.Second)
	completedAt := report.CreatedAt.Add(2 *time.Second)
	failedAt := report.CreatedAt.Add(3*time.Second)
	errorMsg := "there was a failure"
	downloadUrl := "http://localhost:8080/reports"
	downloadPath := "s3://reports-test/reports"
	downloadUrlExpiresAt := report.CreatedAt.Add(4 * time.Second)

	report.ReportType = "food"
	report.StartedAt = &startedAt
	report.CompletedAt = &completedAt
	report.FailedAt = &failedAt
	report.ErrorMessage = &errorMsg
	report.DownloadUrl = &downloadUrl
	report.OutputFilePath = &downloadPath
	report.DownloadUrlExpiresAt = &downloadUrlExpiresAt
	

	report2, err := reportStore.Update(ctx, report)
	require.NoError(t, err)

	require.Equal(t, report.UserId, report2.UserId)
	require.Equal(t, report.UserId, report2.UserId)
	require.Equal(t, report.UserId, report2.UserId)
	require.Equal(t, report.UserId, report2.UserId)
	require.Equal(t, report.UserId, report2.UserId)
	require.Equal(t, report.UserId, report2.UserId)
	require.Equal(t, report.UserId, report2.UserId)
}