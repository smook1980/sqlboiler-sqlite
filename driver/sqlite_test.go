package driver

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/volatiletech/sqlboiler/v4/drivers"
	_ "modernc.org/sqlite"
)

var (
	flagOverwriteGolden = flag.Bool("overwrite-golden", false, "Overwrite the golden file with the current execution results")
)

func TestDriver(t *testing.T) {
	rand.Seed(time.Now().Unix())
	b, err := ioutil.ReadFile("testdatabase.sql")
	if err != nil {
		t.Fatal(err)
	}

	tmpName := filepath.Join(os.TempDir(), fmt.Sprintf("sqlboiler-sqlite-%d.sql", rand.Int()))

	out := &bytes.Buffer{}
	createDB := exec.Command("sqlite3", tmpName)
	createDB.Stdout = out
	createDB.Stderr = out
	createDB.Stdin = bytes.NewReader(b)

	t.Log("sqlite file:", tmpName)
	if err := createDB.Run(); err != nil {
		t.Logf("sqlite output:\n%s\n", out.Bytes())
		t.Fatal(err)
	}
	t.Logf("sqlite output:\n%s\n", out.Bytes())

	config := drivers.Config{
		"dbname": tmpName,
	}

	s := &SQLiteDriver{}
	info, err := s.Assemble(config)
	if err != nil {
		t.Fatal(err)
	}

	got, err := json.Marshal(info)
	if err != nil {
		t.Fatal(err)
	}

	if *flagOverwriteGolden {
		if err = ioutil.WriteFile("sqlite.golden.json", got, 0664); err != nil {
			t.Fatal(err)
		}
		t.Log("wrote:", string(got))
		return
	}

	want, err := ioutil.ReadFile("sqlite.golden.json")
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(want, got) {
		t.Errorf("want:\n%s\ngot:\n%s\n", want, got)
	}

	bg, err := json.MarshalIndent(info, "", " ")
	if err != nil {
		t.Fatal(err)
	}

	ioutil.WriteFile("../compare_have.json", bg, 0644)
	if err != nil {
		t.Fatal(err)
	}

	var wnt interface{}
	err = json.Unmarshal(want, &wnt)
	if err != nil {
		t.Fatal(err)
	}

	bw, err := json.MarshalIndent(wnt, "", " ")
	if err != nil {
		t.Fatal(err)
	}

	ioutil.WriteFile("../compare_want.json", bw, 0644)
	if err != nil {
		t.Fatal(err)
	}
}
