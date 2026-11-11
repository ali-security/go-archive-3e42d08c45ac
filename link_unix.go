//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd

package archive

import (
	"os"

	"golang.org/x/sys/unix"
)

// linkNoFollow creates newname as a hard link to the oldname file. Unlike
// [os.Link], it never follows oldname if it is a symbolic link (link(2)
// does on some platforms, such as macOS), so that a hardlink to a symlink in
// an archive creates a hardlink to the symlink itself, and not to the file it
// points to, which may be outside of the extraction directory.
func linkNoFollow(oldname, newname string) error {
	if err := unix.Linkat(unix.AT_FDCWD, oldname, unix.AT_FDCWD, newname, 0); err != nil {
		return &os.LinkError{Op: "link", Old: oldname, New: newname, Err: err}
	}
	return nil
}
