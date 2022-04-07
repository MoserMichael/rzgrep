## rzgrep - grep for stuff in archives that are embedded within archives

This is a small utility, it greps through the contents of an archive file, it also greps through any archive files that are embedded within an archive file.
They like to do that in java, where they have jars nested in jars.

The following archives are supported right now: zip|jar|war|ear|tar|tgz|taz|tar.gz|tbz2|tbz|tar.bz2|tar.bz
I am trusting this link with the definition of these extensions [link - see --auto-compress option](https://www.gnu.org/software/tar/manual/tar.html#Compression)

The utility is written in golang, you can see how to use it from the test output:
The utility accepts a fixed search string, or a regular expression defined as [re2 regex syntax](https://github.com/google/re2/wiki/Syntax)


## What I learned from all this

I have managed to learn a bit of golang generics, a circular buffer based on generics is [here](https://github.com/MoserMichael/rzgrep/blob/master/src/cbuf/cbuf.go), very exiting to have generics, feels like the twenty first century!

I intended to remember/pick up some golang for my job.

Also there seems to be a limit, of what can be done with shell pipelines; this task of grepping a regular expression within a nested compressed archive seems to be right on the border of what is feasible with a command line.
Luckily this problem can be solved by introducing yet another command line tool...


## The Utility

The options of the utility

<pre>
Usage of ./rzgrep:
  -C int
    	display a number of lines around a matching line
  -color
    	color matches on terminal (otherwise mark with <b> </b> tags)
  -e string
    	regular expression to search for. Syntax defined here: https://github.com/google/re2/wiki/Syntax
  -in string
    	file or directory to scan
  -v	debug option
</pre>
The test output
<pre>
+ test_it
+ ./rzgrep -e Cl.se -in zip.jar
zip.jar|src/rzgrep.go:(156) 	defer archive.Close()
zip.jar|src/rzgrep.go:(166) 		fileReader.Close()
zip.jar|src/rzgrep.go:(195) 		fileReader.Close()
zip.jar|src/rzgrep.go:(233) 	defer file.Close()
zip.jar|src/rzgrep.go:(246) 	defer gzf.Close()
zip.jar|src/rzgrep.go:(264) 	defer file.Close()
zip.jar|src/rzgrep.go:(314) 	defer file.Close()
+ ./rzgrep -e Cl.se -in zip.ear
zip.ear|zip.jar|src/rzgrep.go:(156) 	defer archive.Close()
zip.ear|zip.jar|src/rzgrep.go:(166) 		fileReader.Close()
zip.ear|zip.jar|src/rzgrep.go:(195) 		fileReader.Close()
zip.ear|zip.jar|src/rzgrep.go:(233) 	defer file.Close()
zip.ear|zip.jar|src/rzgrep.go:(246) 	defer gzf.Close()
zip.ear|zip.jar|src/rzgrep.go:(264) 	defer file.Close()
zip.ear|zip.jar|src/rzgrep.go:(314) 	defer file.Close()
+ ./rzgrep -e Cl.se -in zip.tgz
zip.tgz|zip.jar|src/rzgrep.go:(156) 	defer archive.Close()
zip.tgz|zip.jar|src/rzgrep.go:(166) 		fileReader.Close()
zip.tgz|zip.jar|src/rzgrep.go:(195) 		fileReader.Close()
zip.tgz|zip.jar|src/rzgrep.go:(233) 	defer file.Close()
zip.tgz|zip.jar|src/rzgrep.go:(246) 	defer gzf.Close()
zip.tgz|zip.jar|src/rzgrep.go:(264) 	defer file.Close()
zip.tgz|zip.jar|src/rzgrep.go:(314) 	defer file.Close()
+ ./rzgrep -C 3 -e Cl.se -in zip.jar
zip.jar|src/rzgrep.go:(153) 		fmt.Printf("Can't open zip file: %s error: %v\n", fName, err)
zip.jar|src/rzgrep.go:(154) 		return err
zip.jar|src/rzgrep.go:(155) 	}
zip.jar|src/rzgrep.go:(156) 	defer archive.Close()
zip.jar|src/rzgrep.go:(157) 
zip.jar|src/rzgrep.go:(158) 	for _, fileEntry := range archive.File {
zip.jar|src/rzgrep.go:(159) 		fileReader, err := fileEntry.Open()
--
zip.jar|src/rzgrep.go:(163) 		}
zip.jar|src/rzgrep.go:(164) 
zip.jar|src/rzgrep.go:(165) 		ctx.runOnAnyReader(fileEntry.Name, fileReader, int64(fileEntry.UncompressedSize64))
zip.jar|src/rzgrep.go:(166) 		fileReader.Close()
zip.jar|src/rzgrep.go:(167) 	}
zip.jar|src/rzgrep.go:(168) 	return nil
zip.jar|src/rzgrep.go:(169) }
--
zip.jar|src/rzgrep.go:(192) 			return err
zip.jar|src/rzgrep.go:(193) 		}
zip.jar|src/rzgrep.go:(194) 		ctx.runOnAnyReader(fileEntry.Name, fileReader, int64(fileEntry.UncompressedSize64))
zip.jar|src/rzgrep.go:(195) 		fileReader.Close()
zip.jar|src/rzgrep.go:(196) 	}
zip.jar|src/rzgrep.go:(197) 	return nil
zip.jar|src/rzgrep.go:(198) 
--
zip.jar|src/rzgrep.go:(230) 	if err != nil {
zip.jar|src/rzgrep.go:(231) 		fmt.Printf("Error: Can't open gzip file %s, %v\n", fName, err)
zip.jar|src/rzgrep.go:(232) 	}
zip.jar|src/rzgrep.go:(233) 	defer file.Close()
zip.jar|src/rzgrep.go:(234) 
zip.jar|src/rzgrep.go:(235) 	var reader io.Reader = file
zip.jar|src/rzgrep.go:(236) 	ctx.runOnGzipReader(reader, entryType&(^GzipFileEntry))
--
zip.jar|src/rzgrep.go:(243) 	if err != nil {
zip.jar|src/rzgrep.go:(244) 		fmt.Printf("Error: Can't open gzip reader %v\n", err)
zip.jar|src/rzgrep.go:(245) 	}
zip.jar|src/rzgrep.go:(246) 	defer gzf.Close()
zip.jar|src/rzgrep.go:(247) 
zip.jar|src/rzgrep.go:(248) 	if entryType&TarFileEntry != 0 {
zip.jar|src/rzgrep.go:(249) 		ctx.runOnTarReader(gzf)
--
zip.jar|src/rzgrep.go:(261) 		fmt.Printf("%s: Error: Can't open gzip file %v\n", ctx.getLoc(), err)
zip.jar|src/rzgrep.go:(262) 		return err
zip.jar|src/rzgrep.go:(263) 	}
zip.jar|src/rzgrep.go:(264) 	defer file.Close()
zip.jar|src/rzgrep.go:(265) 
zip.jar|src/rzgrep.go:(266) 	var reader io.Reader = file
zip.jar|src/rzgrep.go:(267) 	ctx.runOnBZip2Reader(reader, entryType&(^Bzip2FileEntry))
--
zip.jar|src/rzgrep.go:(311) 	if err != nil {
zip.jar|src/rzgrep.go:(312) 		fmt.Printf("Error: Can't open file %s, %v\n", fName, err)
zip.jar|src/rzgrep.go:(313) 	}
zip.jar|src/rzgrep.go:(314) 	defer file.Close()
zip.jar|src/rzgrep.go:(315) 
zip.jar|src/rzgrep.go:(316) 	var reader io.Reader = file
zip.jar|src/rzgrep.go:(317) 	ctx.runOnReader(reader)
--
+ ./rzgrep -C 3 -e Cl.se -in zip.ear
zip.ear|zip.jar|src/rzgrep.go:(153) 		fmt.Printf("Can't open zip file: %s error: %v\n", fName, err)
zip.ear|zip.jar|src/rzgrep.go:(154) 		return err
zip.ear|zip.jar|src/rzgrep.go:(155) 	}
zip.ear|zip.jar|src/rzgrep.go:(156) 	defer archive.Close()
zip.ear|zip.jar|src/rzgrep.go:(157) 
zip.ear|zip.jar|src/rzgrep.go:(158) 	for _, fileEntry := range archive.File {
zip.ear|zip.jar|src/rzgrep.go:(159) 		fileReader, err := fileEntry.Open()
--
zip.ear|zip.jar|src/rzgrep.go:(163) 		}
zip.ear|zip.jar|src/rzgrep.go:(164) 
zip.ear|zip.jar|src/rzgrep.go:(165) 		ctx.runOnAnyReader(fileEntry.Name, fileReader, int64(fileEntry.UncompressedSize64))
zip.ear|zip.jar|src/rzgrep.go:(166) 		fileReader.Close()
zip.ear|zip.jar|src/rzgrep.go:(167) 	}
zip.ear|zip.jar|src/rzgrep.go:(168) 	return nil
zip.ear|zip.jar|src/rzgrep.go:(169) }
--
zip.ear|zip.jar|src/rzgrep.go:(192) 			return err
zip.ear|zip.jar|src/rzgrep.go:(193) 		}
zip.ear|zip.jar|src/rzgrep.go:(194) 		ctx.runOnAnyReader(fileEntry.Name, fileReader, int64(fileEntry.UncompressedSize64))
zip.ear|zip.jar|src/rzgrep.go:(195) 		fileReader.Close()
zip.ear|zip.jar|src/rzgrep.go:(196) 	}
zip.ear|zip.jar|src/rzgrep.go:(197) 	return nil
zip.ear|zip.jar|src/rzgrep.go:(198) 
--
zip.ear|zip.jar|src/rzgrep.go:(230) 	if err != nil {
zip.ear|zip.jar|src/rzgrep.go:(231) 		fmt.Printf("Error: Can't open gzip file %s, %v\n", fName, err)
zip.ear|zip.jar|src/rzgrep.go:(232) 	}
zip.ear|zip.jar|src/rzgrep.go:(233) 	defer file.Close()
zip.ear|zip.jar|src/rzgrep.go:(234) 
zip.ear|zip.jar|src/rzgrep.go:(235) 	var reader io.Reader = file
zip.ear|zip.jar|src/rzgrep.go:(236) 	ctx.runOnGzipReader(reader, entryType&(^GzipFileEntry))
--
zip.ear|zip.jar|src/rzgrep.go:(243) 	if err != nil {
zip.ear|zip.jar|src/rzgrep.go:(244) 		fmt.Printf("Error: Can't open gzip reader %v\n", err)
zip.ear|zip.jar|src/rzgrep.go:(245) 	}
zip.ear|zip.jar|src/rzgrep.go:(246) 	defer gzf.Close()
zip.ear|zip.jar|src/rzgrep.go:(247) 
zip.ear|zip.jar|src/rzgrep.go:(248) 	if entryType&TarFileEntry != 0 {
zip.ear|zip.jar|src/rzgrep.go:(249) 		ctx.runOnTarReader(gzf)
--
zip.ear|zip.jar|src/rzgrep.go:(261) 		fmt.Printf("%s: Error: Can't open gzip file %v\n", ctx.getLoc(), err)
zip.ear|zip.jar|src/rzgrep.go:(262) 		return err
zip.ear|zip.jar|src/rzgrep.go:(263) 	}
zip.ear|zip.jar|src/rzgrep.go:(264) 	defer file.Close()
zip.ear|zip.jar|src/rzgrep.go:(265) 
zip.ear|zip.jar|src/rzgrep.go:(266) 	var reader io.Reader = file
zip.ear|zip.jar|src/rzgrep.go:(267) 	ctx.runOnBZip2Reader(reader, entryType&(^Bzip2FileEntry))
--
zip.ear|zip.jar|src/rzgrep.go:(311) 	if err != nil {
zip.ear|zip.jar|src/rzgrep.go:(312) 		fmt.Printf("Error: Can't open file %s, %v\n", fName, err)
zip.ear|zip.jar|src/rzgrep.go:(313) 	}
zip.ear|zip.jar|src/rzgrep.go:(314) 	defer file.Close()
zip.ear|zip.jar|src/rzgrep.go:(315) 
zip.ear|zip.jar|src/rzgrep.go:(316) 	var reader io.Reader = file
zip.ear|zip.jar|src/rzgrep.go:(317) 	ctx.runOnReader(reader)
--
+ ./rzgrep -C 3 -e Cl.se -in zip.tgz
zip.tgz|zip.jar|src/rzgrep.go:(153) 		fmt.Printf("Can't open zip file: %s error: %v\n", fName, err)
zip.tgz|zip.jar|src/rzgrep.go:(154) 		return err
zip.tgz|zip.jar|src/rzgrep.go:(155) 	}
zip.tgz|zip.jar|src/rzgrep.go:(156) 	defer archive.Close()
zip.tgz|zip.jar|src/rzgrep.go:(157) 
zip.tgz|zip.jar|src/rzgrep.go:(158) 	for _, fileEntry := range archive.File {
zip.tgz|zip.jar|src/rzgrep.go:(159) 		fileReader, err := fileEntry.Open()
--
zip.tgz|zip.jar|src/rzgrep.go:(163) 		}
zip.tgz|zip.jar|src/rzgrep.go:(164) 
zip.tgz|zip.jar|src/rzgrep.go:(165) 		ctx.runOnAnyReader(fileEntry.Name, fileReader, int64(fileEntry.UncompressedSize64))
zip.tgz|zip.jar|src/rzgrep.go:(166) 		fileReader.Close()
zip.tgz|zip.jar|src/rzgrep.go:(167) 	}
zip.tgz|zip.jar|src/rzgrep.go:(168) 	return nil
zip.tgz|zip.jar|src/rzgrep.go:(169) }
--
zip.tgz|zip.jar|src/rzgrep.go:(192) 			return err
zip.tgz|zip.jar|src/rzgrep.go:(193) 		}
zip.tgz|zip.jar|src/rzgrep.go:(194) 		ctx.runOnAnyReader(fileEntry.Name, fileReader, int64(fileEntry.UncompressedSize64))
zip.tgz|zip.jar|src/rzgrep.go:(195) 		fileReader.Close()
zip.tgz|zip.jar|src/rzgrep.go:(196) 	}
zip.tgz|zip.jar|src/rzgrep.go:(197) 	return nil
zip.tgz|zip.jar|src/rzgrep.go:(198) 
--
zip.tgz|zip.jar|src/rzgrep.go:(230) 	if err != nil {
zip.tgz|zip.jar|src/rzgrep.go:(231) 		fmt.Printf("Error: Can't open gzip file %s, %v\n", fName, err)
zip.tgz|zip.jar|src/rzgrep.go:(232) 	}
zip.tgz|zip.jar|src/rzgrep.go:(233) 	defer file.Close()
zip.tgz|zip.jar|src/rzgrep.go:(234) 
zip.tgz|zip.jar|src/rzgrep.go:(235) 	var reader io.Reader = file
zip.tgz|zip.jar|src/rzgrep.go:(236) 	ctx.runOnGzipReader(reader, entryType&(^GzipFileEntry))
--
zip.tgz|zip.jar|src/rzgrep.go:(243) 	if err != nil {
zip.tgz|zip.jar|src/rzgrep.go:(244) 		fmt.Printf("Error: Can't open gzip reader %v\n", err)
zip.tgz|zip.jar|src/rzgrep.go:(245) 	}
zip.tgz|zip.jar|src/rzgrep.go:(246) 	defer gzf.Close()
zip.tgz|zip.jar|src/rzgrep.go:(247) 
zip.tgz|zip.jar|src/rzgrep.go:(248) 	if entryType&TarFileEntry != 0 {
zip.tgz|zip.jar|src/rzgrep.go:(249) 		ctx.runOnTarReader(gzf)
--
zip.tgz|zip.jar|src/rzgrep.go:(261) 		fmt.Printf("%s: Error: Can't open gzip file %v\n", ctx.getLoc(), err)
zip.tgz|zip.jar|src/rzgrep.go:(262) 		return err
zip.tgz|zip.jar|src/rzgrep.go:(263) 	}
zip.tgz|zip.jar|src/rzgrep.go:(264) 	defer file.Close()
zip.tgz|zip.jar|src/rzgrep.go:(265) 
zip.tgz|zip.jar|src/rzgrep.go:(266) 	var reader io.Reader = file
zip.tgz|zip.jar|src/rzgrep.go:(267) 	ctx.runOnBZip2Reader(reader, entryType&(^Bzip2FileEntry))
--
zip.tgz|zip.jar|src/rzgrep.go:(311) 	if err != nil {
zip.tgz|zip.jar|src/rzgrep.go:(312) 		fmt.Printf("Error: Can't open file %s, %v\n", fName, err)
zip.tgz|zip.jar|src/rzgrep.go:(313) 	}
zip.tgz|zip.jar|src/rzgrep.go:(314) 	defer file.Close()
zip.tgz|zip.jar|src/rzgrep.go:(315) 
zip.tgz|zip.jar|src/rzgrep.go:(316) 	var reader io.Reader = file
zip.tgz|zip.jar|src/rzgrep.go:(317) 	ctx.runOnReader(reader)
--
+ echo '*** Highlight search results***'
*** Highlight search results***
+ ./rzgrep -color -C 3 -e Cl.se -in zip.jar
zip.jar|src/rzgrep.go:(153) 		fmt.Printf("Can't open zip file: %s error: %v\n", fName, err)
zip.jar|src/rzgrep.go:(154) 		return err
zip.jar|src/rzgrep.go:(155) 	}
zip.jar|src/rzgrep.go:(156) 	defer archive.<b>Close</b>
()
zip.jar|src/rzgrep.go:(157) 
zip.jar|src/rzgrep.go:(158) 	for _, fileEntry := range archive.File {
zip.jar|src/rzgrep.go:(159) 		fileReader, err := fileEntry.Open()
--
zip.jar|src/rzgrep.go:(163) 		}
zip.jar|src/rzgrep.go:(164) 
zip.jar|src/rzgrep.go:(165) 		ctx.runOnAnyReader(fileEntry.Name, fileReader, int64(fileEntry.UncompressedSize64))
zip.jar|src/rzgrep.go:(166) 		fileReader.<b>Close</b>
()
zip.jar|src/rzgrep.go:(167) 	}
zip.jar|src/rzgrep.go:(168) 	return nil
zip.jar|src/rzgrep.go:(169) }
--
zip.jar|src/rzgrep.go:(192) 			return err
zip.jar|src/rzgrep.go:(193) 		}
zip.jar|src/rzgrep.go:(194) 		ctx.runOnAnyReader(fileEntry.Name, fileReader, int64(fileEntry.UncompressedSize64))
zip.jar|src/rzgrep.go:(195) 		fileReader.<b>Close</b>
()
zip.jar|src/rzgrep.go:(196) 	}
zip.jar|src/rzgrep.go:(197) 	return nil
zip.jar|src/rzgrep.go:(198) 
--
zip.jar|src/rzgrep.go:(230) 	if err != nil {
zip.jar|src/rzgrep.go:(231) 		fmt.Printf("Error: Can't open gzip file %s, %v\n", fName, err)
zip.jar|src/rzgrep.go:(232) 	}
zip.jar|src/rzgrep.go:(233) 	defer file.<b>Close</b>
()
zip.jar|src/rzgrep.go:(234) 
zip.jar|src/rzgrep.go:(235) 	var reader io.Reader = file
zip.jar|src/rzgrep.go:(236) 	ctx.runOnGzipReader(reader, entryType&(^GzipFileEntry))
--
zip.jar|src/rzgrep.go:(243) 	if err != nil {
zip.jar|src/rzgrep.go:(244) 		fmt.Printf("Error: Can't open gzip reader %v\n", err)
zip.jar|src/rzgrep.go:(245) 	}
zip.jar|src/rzgrep.go:(246) 	defer gzf.<b>Close</b>
()
zip.jar|src/rzgrep.go:(247) 
zip.jar|src/rzgrep.go:(248) 	if entryType&TarFileEntry != 0 {
zip.jar|src/rzgrep.go:(249) 		ctx.runOnTarReader(gzf)
--
zip.jar|src/rzgrep.go:(261) 		fmt.Printf("%s: Error: Can't open gzip file %v\n", ctx.getLoc(), err)
zip.jar|src/rzgrep.go:(262) 		return err
zip.jar|src/rzgrep.go:(263) 	}
zip.jar|src/rzgrep.go:(264) 	defer file.<b>Close</b>
()
zip.jar|src/rzgrep.go:(265) 
zip.jar|src/rzgrep.go:(266) 	var reader io.Reader = file
zip.jar|src/rzgrep.go:(267) 	ctx.runOnBZip2Reader(reader, entryType&(^Bzip2FileEntry))
--
zip.jar|src/rzgrep.go:(311) 	if err != nil {
zip.jar|src/rzgrep.go:(312) 		fmt.Printf("Error: Can't open file %s, %v\n", fName, err)
zip.jar|src/rzgrep.go:(313) 	}
zip.jar|src/rzgrep.go:(314) 	defer file.<b>Close</b>
()
zip.jar|src/rzgrep.go:(315) 
zip.jar|src/rzgrep.go:(316) 	var reader io.Reader = file
zip.jar|src/rzgrep.go:(317) 	ctx.runOnReader(reader)
--
+ ./rzgrep -color -C 3 -e Cl.se -in zip.ear
zip.ear|zip.jar|src/rzgrep.go:(153) 		fmt.Printf("Can't open zip file: %s error: %v\n", fName, err)
zip.ear|zip.jar|src/rzgrep.go:(154) 		return err
zip.ear|zip.jar|src/rzgrep.go:(155) 	}
zip.ear|zip.jar|src/rzgrep.go:(156) 	defer archive.<b>Close</b>
()
zip.ear|zip.jar|src/rzgrep.go:(157) 
zip.ear|zip.jar|src/rzgrep.go:(158) 	for _, fileEntry := range archive.File {
zip.ear|zip.jar|src/rzgrep.go:(159) 		fileReader, err := fileEntry.Open()
--
zip.ear|zip.jar|src/rzgrep.go:(163) 		}
zip.ear|zip.jar|src/rzgrep.go:(164) 
zip.ear|zip.jar|src/rzgrep.go:(165) 		ctx.runOnAnyReader(fileEntry.Name, fileReader, int64(fileEntry.UncompressedSize64))
zip.ear|zip.jar|src/rzgrep.go:(166) 		fileReader.<b>Close</b>
()
zip.ear|zip.jar|src/rzgrep.go:(167) 	}
zip.ear|zip.jar|src/rzgrep.go:(168) 	return nil
zip.ear|zip.jar|src/rzgrep.go:(169) }
--
zip.ear|zip.jar|src/rzgrep.go:(192) 			return err
zip.ear|zip.jar|src/rzgrep.go:(193) 		}
zip.ear|zip.jar|src/rzgrep.go:(194) 		ctx.runOnAnyReader(fileEntry.Name, fileReader, int64(fileEntry.UncompressedSize64))
zip.ear|zip.jar|src/rzgrep.go:(195) 		fileReader.<b>Close</b>
()
zip.ear|zip.jar|src/rzgrep.go:(196) 	}
zip.ear|zip.jar|src/rzgrep.go:(197) 	return nil
zip.ear|zip.jar|src/rzgrep.go:(198) 
--
zip.ear|zip.jar|src/rzgrep.go:(230) 	if err != nil {
zip.ear|zip.jar|src/rzgrep.go:(231) 		fmt.Printf("Error: Can't open gzip file %s, %v\n", fName, err)
zip.ear|zip.jar|src/rzgrep.go:(232) 	}
zip.ear|zip.jar|src/rzgrep.go:(233) 	defer file.<b>Close</b>
()
zip.ear|zip.jar|src/rzgrep.go:(234) 
zip.ear|zip.jar|src/rzgrep.go:(235) 	var reader io.Reader = file
zip.ear|zip.jar|src/rzgrep.go:(236) 	ctx.runOnGzipReader(reader, entryType&(^GzipFileEntry))
--
zip.ear|zip.jar|src/rzgrep.go:(243) 	if err != nil {
zip.ear|zip.jar|src/rzgrep.go:(244) 		fmt.Printf("Error: Can't open gzip reader %v\n", err)
zip.ear|zip.jar|src/rzgrep.go:(245) 	}
zip.ear|zip.jar|src/rzgrep.go:(246) 	defer gzf.<b>Close</b>
()
zip.ear|zip.jar|src/rzgrep.go:(247) 
zip.ear|zip.jar|src/rzgrep.go:(248) 	if entryType&TarFileEntry != 0 {
zip.ear|zip.jar|src/rzgrep.go:(249) 		ctx.runOnTarReader(gzf)
--
zip.ear|zip.jar|src/rzgrep.go:(261) 		fmt.Printf("%s: Error: Can't open gzip file %v\n", ctx.getLoc(), err)
zip.ear|zip.jar|src/rzgrep.go:(262) 		return err
zip.ear|zip.jar|src/rzgrep.go:(263) 	}
zip.ear|zip.jar|src/rzgrep.go:(264) 	defer file.<b>Close</b>
()
zip.ear|zip.jar|src/rzgrep.go:(265) 
zip.ear|zip.jar|src/rzgrep.go:(266) 	var reader io.Reader = file
zip.ear|zip.jar|src/rzgrep.go:(267) 	ctx.runOnBZip2Reader(reader, entryType&(^Bzip2FileEntry))
--
zip.ear|zip.jar|src/rzgrep.go:(311) 	if err != nil {
zip.ear|zip.jar|src/rzgrep.go:(312) 		fmt.Printf("Error: Can't open file %s, %v\n", fName, err)
zip.ear|zip.jar|src/rzgrep.go:(313) 	}
zip.ear|zip.jar|src/rzgrep.go:(314) 	defer file.<b>Close</b>
()
zip.ear|zip.jar|src/rzgrep.go:(315) 
zip.ear|zip.jar|src/rzgrep.go:(316) 	var reader io.Reader = file
zip.ear|zip.jar|src/rzgrep.go:(317) 	ctx.runOnReader(reader)
--
+ ./rzgrep -color -C 3 -e Cl.se -in zip.tgz
zip.tgz|zip.jar|src/rzgrep.go:(153) 		fmt.Printf("Can't open zip file: %s error: %v\n", fName, err)
zip.tgz|zip.jar|src/rzgrep.go:(154) 		return err
zip.tgz|zip.jar|src/rzgrep.go:(155) 	}
zip.tgz|zip.jar|src/rzgrep.go:(156) 	defer archive.<b>Close</b>
()
zip.tgz|zip.jar|src/rzgrep.go:(157) 
zip.tgz|zip.jar|src/rzgrep.go:(158) 	for _, fileEntry := range archive.File {
zip.tgz|zip.jar|src/rzgrep.go:(159) 		fileReader, err := fileEntry.Open()
--
zip.tgz|zip.jar|src/rzgrep.go:(163) 		}
zip.tgz|zip.jar|src/rzgrep.go:(164) 
zip.tgz|zip.jar|src/rzgrep.go:(165) 		ctx.runOnAnyReader(fileEntry.Name, fileReader, int64(fileEntry.UncompressedSize64))
zip.tgz|zip.jar|src/rzgrep.go:(166) 		fileReader.<b>Close</b>
()
zip.tgz|zip.jar|src/rzgrep.go:(167) 	}
zip.tgz|zip.jar|src/rzgrep.go:(168) 	return nil
zip.tgz|zip.jar|src/rzgrep.go:(169) }
--
zip.tgz|zip.jar|src/rzgrep.go:(192) 			return err
zip.tgz|zip.jar|src/rzgrep.go:(193) 		}
zip.tgz|zip.jar|src/rzgrep.go:(194) 		ctx.runOnAnyReader(fileEntry.Name, fileReader, int64(fileEntry.UncompressedSize64))
zip.tgz|zip.jar|src/rzgrep.go:(195) 		fileReader.<b>Close</b>
()
zip.tgz|zip.jar|src/rzgrep.go:(196) 	}
zip.tgz|zip.jar|src/rzgrep.go:(197) 	return nil
zip.tgz|zip.jar|src/rzgrep.go:(198) 
--
zip.tgz|zip.jar|src/rzgrep.go:(230) 	if err != nil {
zip.tgz|zip.jar|src/rzgrep.go:(231) 		fmt.Printf("Error: Can't open gzip file %s, %v\n", fName, err)
zip.tgz|zip.jar|src/rzgrep.go:(232) 	}
zip.tgz|zip.jar|src/rzgrep.go:(233) 	defer file.<b>Close</b>
()
zip.tgz|zip.jar|src/rzgrep.go:(234) 
zip.tgz|zip.jar|src/rzgrep.go:(235) 	var reader io.Reader = file
zip.tgz|zip.jar|src/rzgrep.go:(236) 	ctx.runOnGzipReader(reader, entryType&(^GzipFileEntry))
--
zip.tgz|zip.jar|src/rzgrep.go:(243) 	if err != nil {
zip.tgz|zip.jar|src/rzgrep.go:(244) 		fmt.Printf("Error: Can't open gzip reader %v\n", err)
zip.tgz|zip.jar|src/rzgrep.go:(245) 	}
zip.tgz|zip.jar|src/rzgrep.go:(246) 	defer gzf.<b>Close</b>
()
zip.tgz|zip.jar|src/rzgrep.go:(247) 
zip.tgz|zip.jar|src/rzgrep.go:(248) 	if entryType&TarFileEntry != 0 {
zip.tgz|zip.jar|src/rzgrep.go:(249) 		ctx.runOnTarReader(gzf)
--
zip.tgz|zip.jar|src/rzgrep.go:(261) 		fmt.Printf("%s: Error: Can't open gzip file %v\n", ctx.getLoc(), err)
zip.tgz|zip.jar|src/rzgrep.go:(262) 		return err
zip.tgz|zip.jar|src/rzgrep.go:(263) 	}
zip.tgz|zip.jar|src/rzgrep.go:(264) 	defer file.<b>Close</b>
()
zip.tgz|zip.jar|src/rzgrep.go:(265) 
zip.tgz|zip.jar|src/rzgrep.go:(266) 	var reader io.Reader = file
zip.tgz|zip.jar|src/rzgrep.go:(267) 	ctx.runOnBZip2Reader(reader, entryType&(^Bzip2FileEntry))
--
zip.tgz|zip.jar|src/rzgrep.go:(311) 	if err != nil {
zip.tgz|zip.jar|src/rzgrep.go:(312) 		fmt.Printf("Error: Can't open file %s, %v\n", fName, err)
zip.tgz|zip.jar|src/rzgrep.go:(313) 	}
zip.tgz|zip.jar|src/rzgrep.go:(314) 	defer file.<b>Close</b>
()
zip.tgz|zip.jar|src/rzgrep.go:(315) 
zip.tgz|zip.jar|src/rzgrep.go:(316) 	var reader io.Reader = file
zip.tgz|zip.jar|src/rzgrep.go:(317) 	ctx.runOnReader(reader)
--
+ echo '*** eof test ***'
*** eof test ***
