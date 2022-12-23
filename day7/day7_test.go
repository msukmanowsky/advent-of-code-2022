package day7

import (
	"bufio"
	"strings"
	"testing"
)

func Test_SplitCommand(t *testing.T) {
	var rawCmdInput = strings.TrimSpace(`
$ cd /
$ ls
233998 glh.fcb
184686 jzn
dir qcznqph
dir qtbprrq
299692 rbssdzm.ccn
dir vtb
`)
	scanner := bufio.NewScanner(strings.NewReader(rawCmdInput))
	scanner.Split(SplitCommand)
	i := 0
	for scanner.Scan() {
		cmd := scanner.Text()
		if i == 0 && cmd != "$ cd /" {
			t.Fatalf("wanted: \"$ cd /\", got: \"%s\"", cmd)
		}
		if i == 1 && !strings.HasPrefix(cmd, "$ ls") {
			t.Fatalf("wanted prefix: $ ls, got: \"%s\"", cmd)
		}
		i += 1
	}
}

type parseCommandTest struct {
	input string
	want  Command
}

var tests = []parseCommandTest{
	{
		input: "$ cd /",
		want: Command{
			Name:   "cd",
			Args:   []string{"/"},
			Output: "",
		},
	},
	{
		input: "$ ls\n" +
			"233998 glh.fcb\n" +
			"184686 jzn\n" +
			"dir qcznqph\n" +
			"dir qtbprrq\n" +
			"299692 rbssdzm.ccn\n" +
			"dir vtb",
		want: Command{
			Name: "ls",
			Args: []string{},
			Output: "233998 glh.fcb\n" +
				"184686 jzn\n" +
				"dir qcznqph\n" +
				"dir qtbprrq\n" +
				"299692 rbssdzm.ccn\n" +
				"dir vtb",
		},
	},
}

func areCommandsEqual(cmd1 Command, cmd2 Command) bool {
	return cmd1.Name == cmd2.Name && strings.Join(cmd1.Args, " ") == strings.Join(cmd2.Args, " ") && cmd1.Output == cmd2.Output
}

func Test_parseCommand(t *testing.T) {
	for _, test := range tests {
		cmd, _ := parseCommand(test.input)
		if !areCommandsEqual(*cmd, test.want) {
			t.Errorf("input: %s\ngot: %+v\nwanted: %+v", test.input, *cmd, test.want)
		}
	}
}

func Test_DirSize(t *testing.T) {
	rootDir := Dir{
		Name: "",
		Files: map[string]File{
			"file1": {
				Name: "file1",
				Size: 100,
			},
			"file2": {
				Name: "file2",
				Size: 100,
			},
		},
		Dirs: map[string]Dir{
			"1": {
				Name:  "1",
				Files: make(map[string]File),
				Dirs: map[string]Dir{
					"1.1": {
						Name: "1.1",
						Files: map[string]File{
							"file1.1-1": {
								Name: "file1.1-1",
								Size: 100,
							},
							"file1.1-2": {
								Name: "file1.1-2",
								Size: 100,
							},
						},
						Dirs: map[string]Dir{
							"1.1.1": {
								Name: "1.1.1",
								Files: map[string]File{
									"file1.1.1-1": {
										Name: "file1.1.1-1",
										Size: 100,
									},
								},
								Dirs: make(map[string]Dir),
							},
						},
					},
				},
			},
		},
	}

	want := uint(500)
	if size := rootDir.Size(); size != want {
		t.Fatalf("wanted: %d, got: %d", want, size)
	}
}
