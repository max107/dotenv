package dotenv

import (
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"github.com/matchsystems/werr"
)

var (
	_, b, _, _ = runtime.Caller(0)
	WorkDir    = filepath.Dir(b)
)

func dotenv(files []string) error {
	items := make([]string, 0, len(files))
	for _, f := range files {

		if st, err := os.Stat(f); os.IsNotExist(err) || st.IsDir() {
			continue
		}

		items = append(items, f)
	}

	if len(items) > 0 {
		return werr.Wrap(godotenv.Overload(items...))
	}

	return nil
}

func Load[T any]() (*T, error) {
	entity := new(T)

	files := []string{
		path.Join(WorkDir, "../../.env"),
		path.Join(WorkDir, "../../.env.local"),
	}

	if err := dotenv(files); err != nil {
		return nil, werr.Wrap(err)
	}

	if err := env.Parse(entity); err != nil {
		return nil, werr.Wrap(err)
	}

	return entity, nil
}

func LoadTest[T any]() (*T, error) {
	entity := new(T)

	files := []string{
		path.Join(WorkDir, "../../.env"),
		path.Join(WorkDir, "../../.env.test"),
	}

	if err := dotenv(files); err != nil {
		return nil, werr.Wrap(err)
	}

	if err := env.Parse(entity); err != nil {
		return nil, werr.Wrap(err)
	}

	return entity, nil
}

func MustLoad[T any]() *T {
	c, _ := Load[T]()

	return c
}

func MustLoadTest[T any]() *T {
	c, _ := LoadTest[T]()

	return c
}
