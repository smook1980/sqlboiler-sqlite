# sqlboiler-sqlite

This package is a driver for sqlboiler. It requires cgo to build and
therefore is not in the main tree.

## Installation

This package depends on the `database/sql` sqlite driver https://github.com/mattn/go-sqlite,
which requires cgo and the sqlite .so/.dll installed. Refer to the installation
of the `github.com/mattn/go-sqlite` to complete this step.

Installation is simple, just use `go get`. Once the binary is in
your path `sqlboiler` will be able to use it if you run it with the
driver name `sqlite`.

```bash
# Note: You must run this outside of your Go module directory. This must be done
# in GOPATH mode to get the correct result. If you'd like to pin the version
# manually via Go modules you can attempt other installation instructions.

# Install sqlboiler sqlite driver
go get -u -t github.com/smook1980/sqlboiler-sqlite
# Generate models
sqlboiler sqlite
```

It's configuration keys in sqlboiler are simple:

```toml
# Absolute path is recommended since the location
# sqlite is being run can change.
# For example generation time and model test time.
[sqlite]
dbname = "/path/to/file"
```

## Development

This does use go-bindata to embed templates into the binary.
You can run `go-generate` in the driver folder to re-gen the bindata
after modifying templates. Other than that `go build` should be able to
be used to build the binary.
