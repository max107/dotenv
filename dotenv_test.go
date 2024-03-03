package dotenv_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/max107/dotenv"
)

func TestDotenv(t *testing.T) {
	t.Parallel()

	type empty struct {
		Foo string `env:"FOO"`
	}

	wd, err := os.Getwd()
	require.NoError(t, err)

	t.Run("workdir", func(t *testing.T) {
		t.Parallel()

		require.Equal(t, wd, dotenv.WorkDir())
	})

	t.Run("empty", func(t *testing.T) {
		t.Parallel()

		_, err := dotenv.Load[empty](dotenv.WorkDir())
		require.NoError(t, err)
	})

	t.Run("unknown_dir", func(t *testing.T) {
		t.Parallel()

		_, err := dotenv.Load[empty]("unknown_dir")
		require.NoError(t, err)
	})

	t.Run("valid", func(t *testing.T) {
		t.Parallel()

		c, err := dotenv.Load[empty](dotenv.WorkDir())
		require.NoError(t, err)
		require.Equal(t, "bar", c.Foo)
	})
}
