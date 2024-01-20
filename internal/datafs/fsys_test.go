package datafs

import (
	"context"
	"io/fs"
	"net/url"
	"os"
	"runtime"
	"testing"

	"github.com/hairyhenderson/go-fsimpl"
	"github.com/hairyhenderson/go-fsimpl/gitfs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFSysForPath(t *testing.T) {
	t.Run("no provider", func(t *testing.T) {
		ctx := ContextWithFSProvider(context.Background(), nil)
		_, err := FSysForPath(ctx, "foo")
		require.Error(t, err)

		_, err = FSysForPath(ctx, "foo://bar")
		require.Error(t, err)
	})

	t.Run("file url", func(t *testing.T) {
		fsp := fsimpl.FSProviderFunc(func(u *url.URL) (fs.FS, error) {
			if runtime.GOOS == "windows" {
				assert.Equal(t, "file:///C:/tmp/foo", u.String())
				return os.DirFS("C:/tmp/foo"), nil
			}

			assert.Equal(t, "file:///tmp/foo", u.String())
			return os.DirFS("/tmp/foo"), nil
		}, "file")

		ctx := ContextWithFSProvider(context.Background(), fsp)
		fsys, err := FSysForPath(ctx, "file:///tmp/foo")
		require.NoError(t, err)
		require.NotNil(t, fsys)
	})

	t.Run("git url", func(t *testing.T) {
		fsp := fsimpl.FSProviderFunc(func(u *url.URL) (fs.FS, error) {
			assert.Equal(t, "git://github.com/hairyhenderson/gomplate", u.String())
			return gitfs.New(u)
		}, "git")

		ctx := ContextWithFSProvider(context.Background(), fsp)

		fsys, err := FSysForPath(ctx, "git://github.com/hairyhenderson/gomplate//README.md")
		require.NoError(t, err)
		require.NotNil(t, fsys)
	})

	t.Run("git+file url", func(t *testing.T) {
		fsp := fsimpl.FSProviderFunc(func(u *url.URL) (fs.FS, error) {
			if runtime.GOOS == "windows" {
				assert.Equal(t, "git+file:///C:/tmp/repo", u.String())
			} else {
				assert.Equal(t, "git+file:///tmp/repo", u.String())
			}

			return gitfs.New(u)
		}, "git+file")

		ctx := ContextWithFSProvider(context.Background(), fsp)
		fsys, err := FSysForPath(ctx, "git+file:///tmp/repo//README.md")
		require.NoError(t, err)
		require.NotNil(t, fsys)
	})
}
