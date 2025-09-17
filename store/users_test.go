package store

import (
	"context"
	"testing"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stretchr/testify/require"
)

func TestUserStore(t *testing.T) {
	// _ = godotenv.Load()
	// if err := godotenv.Load(".env.test"); err != nil {
	//     t.Log("warning: could not load .env.test, relying on system env")
	// }
	// os.Setenv("ENV", string(config.Env_Test))
	// fmt.Println("env DATABASE_URL:", os.Getenv("DATABASE_URL"))
	// conf, err := config.New()
	// require.NoError(t, err)

	// fmt.Printf(">>> database_url: %s\n", conf.DatabaseUrl())

	// db, err := NewPostgresDb(conf)
	// require.NoError(t, err)
	// defer db.Close()

	// m, err := migrate.New(
	// 	fmt.Sprintf("file:///%s/migrations", conf.ProjectRoot),
	// 	conf.DatabaseUrl())
	// require.NoError(t, err)

	// if err := m.Up(); err != nil && err != migrate.ErrNoChange {
	// 	require.NoError(t, err)
	// }

	env := fixtures.NewTestEnv(t)
	cleanup := env.SetupDb(t)
	t.Cleanup(func() {
		cleanup(t)
	})

	userStore := NewUserStore(env.Db)
	user, err := userStore.CreateUser(context.Background(), "test@test.com", "testingpassword")
	require.NoError(t, err)

	require.Equal(t, "test@test.com", user.Email)
	require.NoError(t, user.ComparePassword("testingpassword"))
}
