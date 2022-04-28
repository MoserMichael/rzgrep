package rzgrep

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"strings"
)
import "os"

const CMD_GO2JAVA_DECOMPILE_CLASS = 2
const CMD_JAVA2GO_SEND_DECOMPILE_RESULT = 3

type ArchiveEntry struct {
	archiveRootDir string
	classFiles     []string
}

type JavaDecompiler struct {
	command        *exec.Cmd
	stdout         io.ReadCloser
	stdin          io.WriteCloser
	rootDir        string
	archiveEntries []ArchiveEntry
	activeDir      string
}

func NewJavaDecompiler() *JavaDecompiler {
	pathToExe, err := os.Executable()
	if err != nil {
		fmt.Println("Error: Can't get exe path", err)
		return nil
	}

	tempDir, err := ioutil.TempDir("", "rzgrep")

	pathToJar := pathToExe + ".jar"

	//fmt.Printf("pathToJar %s\n", pathToJar)
	cmd := exec.Command("java", "-jar", pathToJar)

	// Create stdout, stderr streams of type io.Reader
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("Error: Can't start java decompiler / no stdout pipe", err)
		return nil
	}

	stdin, err := cmd.StdinPipe()
	if err != nil {
		fmt.Println("Error: Can't start java decompiler / no stdin pipe", err)
		return nil
	}

	err = cmd.Start()
	if err != nil {
		fmt.Println("Error: Can't start java decompiler error:", err)
		return nil
	}

	//fmt.Printf("init JavaDecompiler\n")

	retVal := &JavaDecompiler{
		command: cmd,
		stdout:  stdout,
		stdin:   stdin,
		rootDir: tempDir,
	}

	retVal.InitArchive(tempDir)

	return retVal
}

func (ctx *JavaDecompiler) Close() {
	os.RemoveAll(ctx.rootDir)
	ctx.command.Process.Kill()
}

func (*JavaDecompiler) IsClassFile(fileName string) bool {
	return filepath.Ext(fileName) == ".class"
}

func (ctx *JavaDecompiler) InitArchive(dirName string) *JavaDecompiler {
	ctx.archiveEntries = append(ctx.archiveEntries, ArchiveEntry{archiveRootDir: dirName})
	//fmt.Printf("Push-entry: %s\n", ctx.archiveEntries)
	return ctx
}

func (ctx *JavaDecompiler) getActiveDir() string {
	if ctx.activeDir != "" {
		return ctx.activeDir
	}

	activeDir := ""
	for _, entry := range ctx.archiveEntries {
		if activeDir == "" {
			activeDir = entry.archiveRootDir
		} else {
			activeDir = filepath.Join(activeDir, entry.archiveRootDir)
		}
	}
	ctx.activeDir = activeDir
	return ctx.activeDir
}

func (ctx *JavaDecompiler) StoreClassFile(fileName string, reader io.Reader) error {
	filePath := filepath.Join(ctx.getActiveDir(), fileName)

	dirName := filepath.Dir(filePath)
	os.MkdirAll(dirName, os.ModePerm)

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	//fmt.Printf("fileName: %s\n", fileName)
	//fmt.Printf("Writing class file to: %s\n", filePath)
	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}

	base := filepath.Base(fileName)
	if strings.Index(base, "$") == -1 {
		lastEntry := &ctx.archiveEntries[len(ctx.archiveEntries)-1]
		lastEntry.classFiles = append(lastEntry.classFiles, fileName)
	}

	return nil
}

func (ctx *JavaDecompiler) CloseArchive(rzgrepCtx *Ctx) error {

	lastEntry := &ctx.archiveEntries[len(ctx.archiveEntries)-1]
	//fmt.Printf("close archive %s - %d\n", lastEntry, len(lastEntry.classFiles))
	for _, classFile := range lastEntry.classFiles {
		decompiledClass := ctx.decompileClass(classFile)
		//fmt.Printf("decompile: %s - %s\n", classFile, decompiledClass)
		reader := strings.NewReader(decompiledClass)
		rzgrepCtx.runOnClassReader(classFile, reader)
	}

	// pop stack
	ctx.archiveEntries = ctx.archiveEntries[:len(ctx.archiveEntries)-1]
	ctx.activeDir = ""
	return nil
}

func (ctx *JavaDecompiler) decompileClass(fileName string) string {
	buf := make([]byte, 4)

	className := strings.Replace(fileName, "/", ".", -1)

	//fmt.Printf("sending decompile command..\n")
	binary.BigEndian.PutUint32(buf, CMD_GO2JAVA_DECOMPILE_CLASS)
	ctx.stdin.Write(buf)

	ctx.WriteUTF([]byte(ctx.getActiveDir()))
	ctx.WriteUTF([]byte(className))

	//fmt.Printf("decompile command sent!\n")

	for {
		numRead, err := io.ReadFull(ctx.stdout, buf)
		if err != nil || numRead < 4 {
			//fmt.Printf("numRead %d read error %s\n", numRead, err)
			return ""
		}
		cmd := binary.BigEndian.Uint32(buf)
		if cmd == CMD_JAVA2GO_SEND_DECOMPILE_RESULT {
			//fmt.Printf("got CMD_JAVA2GO_SEND_DECOMPILE_RESULT\n")
			decompiledClass, err := ctx.readString()
			if err != nil {
				fmt.Printf("Error reading decompiler response %s\n", err)
			}
			//fmt.Printf("got: %s\n", decompiledClass)
			return decompiledClass
		} else {
			fmt.Printf("Error: unexpected command received from java %d\n", cmd)
			return ""
		}
	}
}

func (ctx *JavaDecompiler) readString() (string, error) {
	buf := make([]byte, 4)

	_, err := io.ReadFull(ctx.stdout, buf)
	if err != nil {
		return "", err
	}
	len := binary.BigEndian.Uint32(buf)

	buffer := make([]byte, len)
	_, err = io.ReadFull(ctx.stdout, buffer)
	if err != nil {
		return "", err
	}

	return string(buffer), nil
}

func (ctx *JavaDecompiler) WriteUTF(value []byte) {
	var buff bytes.Buffer
	_ = binary.Write(&buff, binary.BigEndian, uint16(len(value)))
	_ = binary.Write(&buff, binary.BigEndian, value)

	ctx.stdin.Write(buff.Bytes())
}

func ReadUTF(reader *bytes.Reader) []byte {
	var length uint16
	err := binary.Read(reader, binary.BigEndian, &length)
	if err != nil {
		panic(err)
	}

	bytesString := make([]byte, length)
	_, err = io.ReadFull(reader, bytesString)
	if err != nil {
		panic(err)
	}

	return bytesString
}
