# sqlboiler-sqlite

This package is a sqlboiler driver for sqlite without any c-go dependencies. 

## Installation

This package depends on the `database/sql` sqlite driver https://modernc.org/sqlite. 
Installation is simple, just use `go get`. Once the binary is in your path `sqlboiler` 
will be able to use it if you run it with the driver name `sqlite`.

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

