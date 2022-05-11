# snoopy

This small utility copies its stdin to the clipboard on macOS while preventing properly written Clipboard Managers from accessing the copied data.
This is essentially a Go version of [pbsecret](https://github.com/roosto/pbsecret) and I only wrote it to experiment with Go and Objective-C.

The purpose of this program is to be used together with CLI Password Managers (like `lpass`, `op` or `pass`) to write automation scripts.

## Usage

You would generally pipe the output of some CLI Password Manager to `snoopy`, or run the following command as a test:

```
$ echo "password" | snoopy
```

## Credits

The original Objective-C code was written by Dustin Goldman <dusto@goldmans.net> for [pbsecret](https://github.com/roosto/pbsecret).
