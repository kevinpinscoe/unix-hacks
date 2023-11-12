# isitmounted is a shell tool to result if a directory is a mounted filesystem

isitmounted returns a value to the bash variable “$?”. If the value is 0, then the specified directory is a mountpoint. If the value is 1, then the directory is anything else: not a mountpoint, not a directory or even does not exist. No other signaling is performed.

There are other commands out there to determine this such as: mount, mountpoint and findmnt but this one I find easier to use in an shell if statement.

## Prerequisites

* A Linux PC. MacOS and Windows would probably work in cygwin or WSL I didn't try.
* Install the golang compiler and make tools using your package manager.

## Install

* Clone this repo (git clone)
* go mod tidy
* run make
* run make install (installs to your home directory under the bin subdirectory). You can edit Makefile to change the location of the install. Become sudo if using a system location.

## Run

```
if ! ~/bin/isitmounted /backup
then
   echo "error: /backup is not mounted"
   exit 1
fi
```
