package rzgrep

import (
	"archive/tar"
	"archive/zip"
	"bufio"
	"bytes"
	"cbuf"
	"compress/bzip2"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	_ "log"
	"os"
	"path"
	"regexp"
	"strings"
)

type ColorOutput int32

const (
	NoColor       ColorOutput = 0
	ColorTerminal             = 1
	ColorTags                 = 2
)

type EntryType int32

var (
	colorTermStart = string([]byte{27, 91, 51, 49, 109})
	colorTermEnd   = string([]byte{27, 91, 48, 109})
)

const (
	RegularFileEntry EntryType = 1
	DirEntry                   = 2
	ZipFileEntry               = 4
	GzipFileEntry              = 8
	Bzip2FileEntry             = 16
	TarFileEntry               = 32
	InvalidEntry               = 1 << 16
)

type EType struct {
	eType EntryType
	file  string
}

type Ctx struct {
	recentLines *cbuf.CBuf[string]
	verbose     bool
	pathNam     string
	path        []EType
	regExp      *regexp.Regexp
	hasErrors   bool
	colorOutput ColorOutput
}

func NewCtx(verbose bool, regExp *regexp.Regexp, context int, color bool) *Ctx {
	var ctxBuf *cbuf.CBuf[string]

	if context != 0 {
		ctxBuf = cbuf.NewCBuf[string](context + 1)
	}

	colorOutput := NoColor
	if color {
		if isStdoutTerminal() {
			colorOutput = ColorTerminal
		} else {
			colorOutput = ColorTags
		}
	}
	return &Ctx{recentLines: ctxBuf, verbose: verbose, pathNam: "", path: nil, regExp: regExp, colorOutput: colorOutput}
}

func isStdoutTerminal() bool {
	if fileInfo, _ := os.Stdout.Stat(); (fileInfo.Mode() & os.ModeCharDevice) != 0 {
		return true
	}
	return false
}

func (ctx *Ctx) runOnDir(dir string) error {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("can't enumerate directory %v", err)
	}

	for _, file := range files {
		pathName := path.Join(path.Dir(dir), file.Name())
		err = ctx.runOnFile(pathName)
	}
	return err
}

func (ctx *Ctx) runOnFile(fName string) error {

	ty, err := ctx.classifyFile(fName)
	if err != nil {
		err = fmt.Errorf("Error %s : %s\n", fName, err)
	} else if ty == RegularFileEntry {
		err = ctx.runOnRegularFile(fName)
	} else if ty == DirEntry {
		err = ctx.runOnDir(fName)
	} else if ty == ZipFileEntry {
		err = ctx.runOnZipFile(fName)
	} else if ty&GzipFileEntry != 0 {
		err = ctx.runOnGzipFile(fName, ty)
	} else if ty&Bzip2FileEntry != 0 {
		err = ctx.runOnBzip2File(fName, ty)
	} else {
		err = fmt.Errorf("error, unsupported option %d", ty)
	}
	return err
}

func (ctx *Ctx) runOnZipFile(fName string) error {

	ctx.push(fName, ZipFileEntry)
	defer ctx.pop()

	archive, err := zip.OpenReader(fName)
	if err != nil {
		fmt.Printf("Can't open zip file: %s error: %v\n", fName, err)
		return err
	}
	defer archive.Close()

	for _, fileEntry := range archive.File {
		fileReader, err := fileEntry.Open()
		if err != nil {
			fmt.Printf("Error: can't open %s in zip file\n", fileEntry.Name)
			return err
		}

		ctx.runOnAnyReader(fileEntry.Name, fileReader, int64(fileEntry.UncompressedSize64))
		fileReader.Close()
	}
	return nil
}

func (ctx *Ctx) runOnZipReader(reader io.Reader, fileSize int64) error {

	// copy reader content to memory buffer, with objective to create ioReader from it.
	buff := bytes.NewBuffer([]byte{})
	_, err := io.Copy(buff, reader)
	if err != nil {
		fmt.Printf("Can't open zip archive: %s error: %v\n", ctx.getLoc(), err)
		return err
	}
	ioReader := bytes.NewReader(buff.Bytes())

	archive, err := zip.NewReader(ioReader, fileSize)
	if err != nil {
		fmt.Printf("Can't open zip archive: %s error: %v\n", ctx.getLoc(), err)
		return err
	}

	for _, fileEntry := range archive.File {
		fileReader, err := fileEntry.Open()
		if err != nil {
			fmt.Printf("Error: can't open %s in zip file\n", fileEntry.Name)
			return err
		}
		ctx.runOnAnyReader(fileEntry.Name, fileReader, int64(fileEntry.UncompressedSize64))
		fileReader.Close()
	}
	return nil

}

func (ctx *Ctx) runOnAnyReader(fName string, reader io.Reader, fileSize int64) {

	entryType := ctx.classifyFileName(fName)

	ctx.push(fName, entryType)
	defer ctx.pop()

	var err error

	if entryType&RegularFileEntry != 0 {
		ctx.runOnReader(reader)
	} else if entryType&ZipFileEntry != 0 {
		err = ctx.runOnZipReader(reader, fileSize)
	} else if entryType&GzipFileEntry != 0 {
		ctx.runOnGzipReader(reader, entryType)
	} else if entryType&Bzip2FileEntry != 0 {
		ctx.runOnBZip2Reader(reader, entryType)
	}

	if err != nil {
		fmt.Printf("Error: %s %v", ctx.getLoc(), err)
	}
}

func (ctx *Ctx) runOnGzipFile(fName string, entryType EntryType) error {
	ctx.push(fName, GzipFileEntry)
	defer ctx.pop()

	file, err := os.Open(fName)
	if err != nil {
		fmt.Printf("Error: Can't open gzip file %s, %v\n", fName, err)
	}
	defer file.Close()

	var reader io.Reader = file
	ctx.runOnGzipReader(reader, entryType&(^GzipFileEntry))
	return err
}

func (ctx *Ctx) runOnGzipReader(reader io.Reader, entryType EntryType) {

	gzf, err := gzip.NewReader(reader)
	if err != nil {
		fmt.Printf("Error: Can't open gzip reader %v\n", err)
	}
	defer gzf.Close()

	if entryType&TarFileEntry != 0 {
		ctx.runOnTarReader(gzf)
	} else {
		ctx.runOnReader(gzf)
	}
}

func (ctx *Ctx) runOnBzip2File(fName string, entryType EntryType) error {
	ctx.push(fName, Bzip2FileEntry)
	defer ctx.pop()

	file, err := os.Open(fName)
	if err != nil {
		fmt.Printf("%s: Error: Can't open gzip file %v\n", ctx.getLoc(), err)
		return err
	}
	defer file.Close()

	var reader io.Reader = file
	ctx.runOnBZip2Reader(reader, entryType&(^Bzip2FileEntry))
	return nil
}

func (ctx *Ctx) runOnBZip2Reader(reader io.Reader, entryType EntryType) {
	gzf := bzip2.NewReader(reader)
	if gzf == nil {
		fmt.Printf("Error: Can't open bzip2 reader\n")
	}

	if entryType&TarFileEntry != 0 {
		ctx.runOnTarReader(gzf)
	} else {
		ctx.runOnReader(gzf)
	}
}

func (ctx *Ctx) runOnTarReader(reader io.Reader) error {
	tarReader := tar.NewReader(reader)
	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
			return err
		}

		if header.Typeflag == tar.TypeReg {
			ctx.runOnAnyReader(header.Name, tarReader, header.Size)
		}
	}
	return nil
}

func (ctx *Ctx) runOnRegularFile(fName string) error {

	ctx.push(fName, RegularFileEntry)
	defer ctx.pop()

	file, err := os.Open(fName)
	if err != nil {
		fmt.Printf("Error: Can't open file %s, %v\n", fName, err)
	}
	defer file.Close()

	var reader io.Reader = file
	ctx.runOnReader(reader)

	return err
}

func (ctx *Ctx) runOnReader(reader io.Reader) {

	var scanner = bufio.NewScanner(reader)
	lineNo := 1

	if ctx.recentLines != nil {
		ctx.recentLines.Clear()
	}

	showLinesAfter := 0
	var machPositions [][]int
	var hasMatch bool

	for scanner.Scan() {
		line := scanner.Text()

		if ctx.colorOutput == NoColor {
			if ctx.regExp.FindStringIndex(line) != nil {
				hasMatch = true
			} else {
				hasMatch = false
			}
		} else {
			machPositions = ctx.regExp.FindAllStringIndex(line, -1)
			if machPositions != nil {
				hasMatch = true
			} else {
				hasMatch = false
			}
		}
		if hasMatch {

			// show lines/context before match (if requested)
			if ctx.recentLines != nil {
				for !ctx.recentLines.IsEmpty() {
					numEntries := ctx.recentLines.NumEntries()
					bufLine := ctx.recentLines.Pop()
					fmt.Printf("%s:(%d) %s\n", ctx.getLoc(), lineNo-numEntries, *bufLine)
				}
				showLinesAfter = ctx.recentLines.Size()
			}

			// show the match itself
			if ctx.colorOutput == NoColor {
				fmt.Printf("%s:(%d) %s\n", ctx.getLoc(), lineNo, line)
			} else {
				// show highlighted match
				fmt.Printf("%s:(%d) ", ctx.getLoc(), lineNo)
				var lastPos int = 0
				var lastEndMatch int = -1
				for _, oneMatch := range machPositions {
					startPos := oneMatch[0]
					endPos := oneMatch[1]
					prefixStr := line[lastPos:startPos]
					matchStr := line[startPos:endPos]
					fmt.Printf("%s", prefixStr)
					if ctx.colorOutput == ColorTags {
						fmt.Printf("<b>%s</b>\n", matchStr)
					}
					if ctx.colorOutput == ColorTerminal {
						fmt.Print(colorTermStart, matchStr, colorTermEnd)
					}
					lastEndMatch = endPos
				}
				if lastEndMatch != -1 {
					fmt.Printf("%s\n", line[lastEndMatch:])
				}
			}
		} else {
			if showLinesAfter != 0 {
				fmt.Printf("%s:(%d) %s\n", ctx.getLoc(), lineNo, line)
				showLinesAfter -= 1
				if showLinesAfter == 1 {
					fmt.Printf("--\n")
					showLinesAfter = 0
				}
			} else {
				if ctx.recentLines != nil {
					if ctx.recentLines.IsFull() {
						ctx.recentLines.Pop()
					}
					ctx.recentLines.Push(&line)
				}
			}
		}
		lineNo += 1
	}
}

func (ctx *Ctx) getLoc() string {

	if len(ctx.pathNam) == 0 {
		if ctx.verbose {
			log.Printf("dbg: building path... %d\n", len(ctx.path))
		}
		var res strings.Builder

		for idx, eType := range ctx.path {
			if idx > 0 {
				res.WriteString("|")
			}
			res.WriteString(eType.file)
		}
		ctx.pathNam = res.String()
	}
	return ctx.pathNam
}

func (ctx *Ctx) classifyFile(fName string) (EntryType, error) {

	stat, err := os.Stat(fName)
	if err != nil {
		return InvalidEntry, fmt.Errorf("failed to open %s, error: %w", fName, err)
	}
	if stat.IsDir() {
		return DirEntry, nil
	}

	return ctx.classifyFileName(fName), nil
}

// see: https://www.gnu.org/software/tar/manual/tar.html#Compression
func (ctx *Ctx) classifyFileName(fName string) EntryType {
	if strings.HasSuffix(fName, ".zip") ||
		strings.HasSuffix(fName, ".jar") ||
		strings.HasSuffix(fName, ".war") ||
		strings.HasSuffix(fName, ".ear") {
		return ZipFileEntry
	}
	if strings.HasSuffix(fName, ".tar") {
		return TarFileEntry
	}
	if strings.HasSuffix(fName, ".tgz") ||
		strings.HasSuffix(fName, ".taz") ||
		strings.HasSuffix(fName, ".tar.gz") {
		return TarFileEntry | GzipFileEntry
	}

	if strings.HasSuffix(fName, ".tbz2") ||
		strings.HasSuffix(fName, ".tbz") ||
		strings.HasSuffix(fName, ".tar.bz2") ||
		strings.HasSuffix(fName, ".tar.bz") {
		return TarFileEntry | Bzip2FileEntry
	}
	return RegularFileEntry
}

func (ctx *Ctx) push(file string, eType EntryType) {
	elm := EType{
		eType: eType,
		file:  file,
	}

	ctx.path = append(ctx.path, elm)

	if ctx.verbose {
		log.Printf("visit %p %s %s %d %v\n", ctx.path, file, entryTypeName(eType), len(ctx.path), ctx.path)
	}
}

func entryTypeName(eType EntryType) string {
	title := ""

	if eType == RegularFileEntry {
		title += "File"
	}
	if eType == DirEntry {
		title += "directory"
	}
	if eType&ZipFileEntry != 0 {
		title += "zip"
	}
	if eType&TarFileEntry != 0 {
		title += "tar"
	}
	return title
}

func (ctx *Ctx) pop() (*EType, error) {
	if ctx.verbose {
		log.Printf("dbg: pop... %d\n", len(ctx.path))
	}
	nSize := len(ctx.path) - 1
	if nSize == -1 {
		return nil, fmt.Errorf("stack underflow")
	}

	rVal := ctx.path[nSize]
	ctx.path = ctx.path[:nSize]

	return &rVal, nil
}

type CmdParams struct {
	verbose bool
	inFile  string
	regExp  *regexp.Regexp
	context int
	color   bool
}

func parseCmdLine() *CmdParams {
	inFile := flag.String("in", "", "file or directory to scan")
	regExp := flag.String("e", "", "regular expression to search for. Syntax defined here: https://github.com/google/re2/wiki/Syntax")
	verbose := flag.Bool("v", false, "debug option")
	context := flag.Int("C", 0, "display a number of lines around a matching line")
	color := flag.Bool("color", false, "color matches on terminal (otherwise mark with <b> </b> tags)")

	flag.Parse()

	if inFile == nil || *inFile == "" {
		fmt.Printf("Error: No input file present\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if regExp == nil || *regExp == "" {
		fmt.Printf("Error: No search expression present\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	cRegExp, err := regexp.Compile(*regExp)
	if err != nil {
		fmt.Printf("Error: the regular expression '%s' has an error %v", *regExp, err)
	}

	if *verbose {
		fmt.Printf("regexp. raw: %s compiled: %s context lines: %d\n", *regExp, cRegExp, *context)
	}

	return &CmdParams{verbose: *verbose, inFile: *inFile, regExp: cRegExp, context: *context, color: *color}
}

func RunMain() {
	cmdParams := parseCmdLine()

	ctx := NewCtx(cmdParams.verbose, cmdParams.regExp, cmdParams.context, cmdParams.color)

	err := ctx.runOnFile(cmdParams.inFile)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
}
