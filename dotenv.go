package dotenv

import (
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
	"github.com/matchsystems/werr"
)

// WorkDir get path to project root. Work only with go mod vendor
func WorkDir() string {
	_, b, _, _ := runtime.Caller(0)

	return path.Join(filepath.Dir(b), "../../../../")
}

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

func Load[T any](workdir string) (*T, error) {
	entity := new(T)

	files := []string{
		path.Join(workdir, ".env"),
		path.Join(workdir, ".env.local"),
	}

	if err := dotenv(files); err != nil {
		return nil, werr.Wrap(err)
	}

	if err := env.Parse(entity); err != nil {
		return nil, werr.Wrap(err)
	}

	return entity, nil
}

func LoadTest[T any](workdir string) (*T, error) {
	entity := new(T)

	files := []string{
		path.Join(workdir, ".env"),
		path.Join(workdir, ".env.test"),
	}

	if err := dotenv(files); err != nil {
		return nil, werr.Wrap(err)
	}

	if err := env.Parse(entity); err != nil {
		return nil, werr.Wrap(err)
	}

	return entity, nil
}

func MustLoad[T any](workdir string) *T {
	c, _ := Load[T](workdir)

	return c
}

func MustLoadTest[T any](workdir string) *T {
	c, _ := LoadTest[T](workdir)

	return c
}
