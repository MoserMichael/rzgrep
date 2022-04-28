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
  -j	use java decompiler for .class files
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
zip.jar|src/rzgrep.go:(156) 	defer archive.<b>Close</b>()
zip.jar|src/rzgrep.go:(157) 
zip.jar|src/rzgrep.go:(158) 	for _, fileEntry := range archive.File {
zip.jar|src/rzgrep.go:(159) 		fileReader, err := fileEntry.Open()
--
zip.jar|src/rzgrep.go:(163) 		}
zip.jar|src/rzgrep.go:(164) 
zip.jar|src/rzgrep.go:(165) 		ctx.runOnAnyReader(fileEntry.Name, fileReader, int64(fileEntry.UncompressedSize64))
zip.jar|src/rzgrep.go:(166) 		fileReader.<b>Close</b>()
zip.jar|src/rzgrep.go:(167) 	}
zip.jar|src/rzgrep.go:(168) 	return nil
zip.jar|src/rzgrep.go:(169) }
--
zip.jar|src/rzgrep.go:(192) 			return err
zip.jar|src/rzgrep.go:(193) 		}
zip.jar|src/rzgrep.go:(194) 		ctx.runOnAnyReader(fileEntry.Name, fileReader, int64(fileEntry.UncompressedSize64))
zip.jar|src/rzgrep.go:(195) 		fileReader.<b>Close</b>()
zip.jar|src/rzgrep.go:(196) 	}
zip.jar|src/rzgrep.go:(197) 	return nil
zip.jar|src/rzgrep.go:(198) 
--
zip.jar|src/rzgrep.go:(230) 	if err != nil {
zip.jar|src/rzgrep.go:(231) 		fmt.Printf("Error: Can't open gzip file %s, %v\n", fName, err)
zip.jar|src/rzgrep.go:(232) 	}
zip.jar|src/rzgrep.go:(233) 	defer file.<b>Close</b>()
zip.jar|src/rzgrep.go:(234) 
zip.jar|src/rzgrep.go:(235) 	var reader io.Reader = file
zip.jar|src/rzgrep.go:(236) 	ctx.runOnGzipReader(reader, entryType&(^GzipFileEntry))
--
zip.jar|src/rzgrep.go:(243) 	if err != nil {
zip.jar|src/rzgrep.go:(244) 		fmt.Printf("Error: Can't open gzip reader %v\n", err)
zip.jar|src/rzgrep.go:(245) 	}
zip.jar|src/rzgrep.go:(246) 	defer gzf.<b>Close</b>()
zip.jar|src/rzgrep.go:(247) 
zip.jar|src/rzgrep.go:(248) 	if entryType&TarFileEntry != 0 {
zip.jar|src/rzgrep.go:(249) 		ctx.runOnTarReader(gzf)
--
zip.jar|src/rzgrep.go:(261) 		fmt.Printf("%s: Error: Can't open gzip file %v\n", ctx.getLoc(), err)
zip.jar|src/rzgrep.go:(262) 		return err
zip.jar|src/rzgrep.go:(263) 	}
zip.jar|src/rzgrep.go:(264) 	defer file.<b>Close</b>()
zip.jar|src/rzgrep.go:(265) 
zip.jar|src/rzgrep.go:(266) 	var reader io.Reader = file
zip.jar|src/rzgrep.go:(267) 	ctx.runOnBZip2Reader(reader, entryType&(^Bzip2FileEntry))
--
zip.jar|src/rzgrep.go:(311) 	if err != nil {
zip.jar|src/rzgrep.go:(312) 		fmt.Printf("Error: Can't open file %s, %v\n", fName, err)
zip.jar|src/rzgrep.go:(313) 	}
zip.jar|src/rzgrep.go:(314) 	defer file.<b>Close</b>()
zip.jar|src/rzgrep.go:(315) 
zip.jar|src/rzgrep.go:(316) 	var reader io.Reader = file
zip.jar|src/rzgrep.go:(317) 	ctx.runOnReader(reader)
--
+ ./rzgrep -color -C 3 -e Cl.se -in zip.ear
zip.ear|zip.jar|src/rzgrep.go:(153) 		fmt.Printf("Can't open zip file: %s error: %v\n", fName, err)
zip.ear|zip.jar|src/rzgrep.go:(154) 		return err
zip.ear|zip.jar|src/rzgrep.go:(155) 	}
zip.ear|zip.jar|src/rzgrep.go:(156) 	defer archive.<b>Close</b>()
zip.ear|zip.jar|src/rzgrep.go:(157) 
zip.ear|zip.jar|src/rzgrep.go:(158) 	for _, fileEntry := range archive.File {
zip.ear|zip.jar|src/rzgrep.go:(159) 		fileReader, err := fileEntry.Open()
--
zip.ear|zip.jar|src/rzgrep.go:(163) 		}
zip.ear|zip.jar|src/rzgrep.go:(164) 
zip.ear|zip.jar|src/rzgrep.go:(165) 		ctx.runOnAnyReader(fileEntry.Name, fileReader, int64(fileEntry.UncompressedSize64))
zip.ear|zip.jar|src/rzgrep.go:(166) 		fileReader.<b>Close</b>()
zip.ear|zip.jar|src/rzgrep.go:(167) 	}
zip.ear|zip.jar|src/rzgrep.go:(168) 	return nil
zip.ear|zip.jar|src/rzgrep.go:(169) }
--
zip.ear|zip.jar|src/rzgrep.go:(192) 			return err
zip.ear|zip.jar|src/rzgrep.go:(193) 		}
zip.ear|zip.jar|src/rzgrep.go:(194) 		ctx.runOnAnyReader(fileEntry.Name, fileReader, int64(fileEntry.UncompressedSize64))
zip.ear|zip.jar|src/rzgrep.go:(195) 		fileReader.<b>Close</b>()
zip.ear|zip.jar|src/rzgrep.go:(196) 	}
zip.ear|zip.jar|src/rzgrep.go:(197) 	return nil
zip.ear|zip.jar|src/rzgrep.go:(198) 
--
zip.ear|zip.jar|src/rzgrep.go:(230) 	if err != nil {
zip.ear|zip.jar|src/rzgrep.go:(231) 		fmt.Printf("Error: Can't open gzip file %s, %v\n", fName, err)
zip.ear|zip.jar|src/rzgrep.go:(232) 	}
zip.ear|zip.jar|src/rzgrep.go:(233) 	defer file.<b>Close</b>()
zip.ear|zip.jar|src/rzgrep.go:(234) 
zip.ear|zip.jar|src/rzgrep.go:(235) 	var reader io.Reader = file
zip.ear|zip.jar|src/rzgrep.go:(236) 	ctx.runOnGzipReader(reader, entryType&(^GzipFileEntry))
--
zip.ear|zip.jar|src/rzgrep.go:(243) 	if err != nil {
zip.ear|zip.jar|src/rzgrep.go:(244) 		fmt.Printf("Error: Can't open gzip reader %v\n", err)
zip.ear|zip.jar|src/rzgrep.go:(245) 	}
zip.ear|zip.jar|src/rzgrep.go:(246) 	defer gzf.<b>Close</b>()
zip.ear|zip.jar|src/rzgrep.go:(247) 
zip.ear|zip.jar|src/rzgrep.go:(248) 	if entryType&TarFileEntry != 0 {
zip.ear|zip.jar|src/rzgrep.go:(249) 		ctx.runOnTarReader(gzf)
--
zip.ear|zip.jar|src/rzgrep.go:(261) 		fmt.Printf("%s: Error: Can't open gzip file %v\n", ctx.getLoc(), err)
zip.ear|zip.jar|src/rzgrep.go:(262) 		return err
zip.ear|zip.jar|src/rzgrep.go:(263) 	}
zip.ear|zip.jar|src/rzgrep.go:(264) 	defer file.<b>Close</b>()
zip.ear|zip.jar|src/rzgrep.go:(265) 
zip.ear|zip.jar|src/rzgrep.go:(266) 	var reader io.Reader = file
zip.ear|zip.jar|src/rzgrep.go:(267) 	ctx.runOnBZip2Reader(reader, entryType&(^Bzip2FileEntry))
--
zip.ear|zip.jar|src/rzgrep.go:(311) 	if err != nil {
zip.ear|zip.jar|src/rzgrep.go:(312) 		fmt.Printf("Error: Can't open file %s, %v\n", fName, err)
zip.ear|zip.jar|src/rzgrep.go:(313) 	}
zip.ear|zip.jar|src/rzgrep.go:(314) 	defer file.<b>Close</b>()
zip.ear|zip.jar|src/rzgrep.go:(315) 
zip.ear|zip.jar|src/rzgrep.go:(316) 	var reader io.Reader = file
zip.ear|zip.jar|src/rzgrep.go:(317) 	ctx.runOnReader(reader)
--
+ ./rzgrep -color -C 3 -e Cl.se -in zip.tgz
zip.tgz|zip.jar|src/rzgrep.go:(153) 		fmt.Printf("Can't open zip file: %s error: %v\n", fName, err)
zip.tgz|zip.jar|src/rzgrep.go:(154) 		return err
zip.tgz|zip.jar|src/rzgrep.go:(155) 	}
zip.tgz|zip.jar|src/rzgrep.go:(156) 	defer archive.<b>Close</b>()
zip.tgz|zip.jar|src/rzgrep.go:(157) 
zip.tgz|zip.jar|src/rzgrep.go:(158) 	for _, fileEntry := range archive.File {
zip.tgz|zip.jar|src/rzgrep.go:(159) 		fileReader, err := fileEntry.Open()
--
zip.tgz|zip.jar|src/rzgrep.go:(163) 		}
zip.tgz|zip.jar|src/rzgrep.go:(164) 
zip.tgz|zip.jar|src/rzgrep.go:(165) 		ctx.runOnAnyReader(fileEntry.Name, fileReader, int64(fileEntry.UncompressedSize64))
zip.tgz|zip.jar|src/rzgrep.go:(166) 		fileReader.<b>Close</b>()
zip.tgz|zip.jar|src/rzgrep.go:(167) 	}
zip.tgz|zip.jar|src/rzgrep.go:(168) 	return nil
zip.tgz|zip.jar|src/rzgrep.go:(169) }
--
zip.tgz|zip.jar|src/rzgrep.go:(192) 			return err
zip.tgz|zip.jar|src/rzgrep.go:(193) 		}
zip.tgz|zip.jar|src/rzgrep.go:(194) 		ctx.runOnAnyReader(fileEntry.Name, fileReader, int64(fileEntry.UncompressedSize64))
zip.tgz|zip.jar|src/rzgrep.go:(195) 		fileReader.<b>Close</b>()
zip.tgz|zip.jar|src/rzgrep.go:(196) 	}
zip.tgz|zip.jar|src/rzgrep.go:(197) 	return nil
zip.tgz|zip.jar|src/rzgrep.go:(198) 
--
zip.tgz|zip.jar|src/rzgrep.go:(230) 	if err != nil {
zip.tgz|zip.jar|src/rzgrep.go:(231) 		fmt.Printf("Error: Can't open gzip file %s, %v\n", fName, err)
zip.tgz|zip.jar|src/rzgrep.go:(232) 	}
zip.tgz|zip.jar|src/rzgrep.go:(233) 	defer file.<b>Close</b>()
zip.tgz|zip.jar|src/rzgrep.go:(234) 
zip.tgz|zip.jar|src/rzgrep.go:(235) 	var reader io.Reader = file
zip.tgz|zip.jar|src/rzgrep.go:(236) 	ctx.runOnGzipReader(reader, entryType&(^GzipFileEntry))
--
zip.tgz|zip.jar|src/rzgrep.go:(243) 	if err != nil {
zip.tgz|zip.jar|src/rzgrep.go:(244) 		fmt.Printf("Error: Can't open gzip reader %v\n", err)
zip.tgz|zip.jar|src/rzgrep.go:(245) 	}
zip.tgz|zip.jar|src/rzgrep.go:(246) 	defer gzf.<b>Close</b>()
zip.tgz|zip.jar|src/rzgrep.go:(247) 
zip.tgz|zip.jar|src/rzgrep.go:(248) 	if entryType&TarFileEntry != 0 {
zip.tgz|zip.jar|src/rzgrep.go:(249) 		ctx.runOnTarReader(gzf)
--
zip.tgz|zip.jar|src/rzgrep.go:(261) 		fmt.Printf("%s: Error: Can't open gzip file %v\n", ctx.getLoc(), err)
zip.tgz|zip.jar|src/rzgrep.go:(262) 		return err
zip.tgz|zip.jar|src/rzgrep.go:(263) 	}
zip.tgz|zip.jar|src/rzgrep.go:(264) 	defer file.<b>Close</b>()
zip.tgz|zip.jar|src/rzgrep.go:(265) 
zip.tgz|zip.jar|src/rzgrep.go:(266) 	var reader io.Reader = file
zip.tgz|zip.jar|src/rzgrep.go:(267) 	ctx.runOnBZip2Reader(reader, entryType&(^Bzip2FileEntry))
--
zip.tgz|zip.jar|src/rzgrep.go:(311) 	if err != nil {
zip.tgz|zip.jar|src/rzgrep.go:(312) 		fmt.Printf("Error: Can't open file %s, %v\n", fName, err)
zip.tgz|zip.jar|src/rzgrep.go:(313) 	}
zip.tgz|zip.jar|src/rzgrep.go:(314) 	defer file.<b>Close</b>()
zip.tgz|zip.jar|src/rzgrep.go:(315) 
zip.tgz|zip.jar|src/rzgrep.go:(316) 	var reader io.Reader = file
zip.tgz|zip.jar|src/rzgrep.go:(317) 	ctx.runOnReader(reader)
--
+ echo '*** Java decompiler: search in compiled classes ***'
*** Java decompiler: search in compiled classes ***
+ ./rzgrep -color -C 3 -e for -in rzgrep.jar -j
decomp/PipeDecompiler.class:(36)     public void run() {
decomp/PipeDecompiler.class:(37)         try {
decomp/PipeDecompiler.class:(38)             int cmd;
decomp/PipeDecompiler.class:(39)             log("waiting <b>for</b> commands...");
decomp/PipeDecompiler.class:(40)             while (true) {
decomp/PipeDecompiler.class:(41)                 cmd = this.inData.readInt();
decomp/PipeDecompiler.class:(42)                 if (cmd == CMD_GO2JAVA_DECOMPILE_CLASS) {
--
org/jd/core/v1/model/fragment/FlexibleFragment.class:(51)         return this.label;
org/jd/core/v1/model/fragment/FlexibleFragment.class:(52)     }
org/jd/core/v1/model/fragment/FlexibleFragment.class:(53)     
org/jd/core/v1/model/fragment/FlexibleFragment.class:(54)     public boolean incLineCount(boolean <b>for</b>ce) {
org/jd/core/v1/model/fragment/FlexibleFragment.class:(55)         if (this.lineCount < this.maximalLineCount) {
org/jd/core/v1/model/fragment/FlexibleFragment.class:(56)             this.lineCount++;
org/jd/core/v1/model/fragment/FlexibleFragment.class:(57)             return true;
--
org/jd/core/v1/model/fragment/FlexibleFragment.class:(59)         return false;
org/jd/core/v1/model/fragment/FlexibleFragment.class:(60)     }
org/jd/core/v1/model/fragment/FlexibleFragment.class:(61)     
org/jd/core/v1/model/fragment/FlexibleFragment.class:(62)     public boolean decLineCount(boolean <b>for</b>ce) {
org/jd/core/v1/model/fragment/FlexibleFragment.class:(63)         if (this.lineCount > this.minimalLineCount) {
org/jd/core/v1/model/fragment/FlexibleFragment.class:(64)             this.lineCount--;
org/jd/core/v1/model/fragment/FlexibleFragment.class:(65)             return true;
--
org/jd/core/v1/model/javafragment/EndBlockFragment.class:(21)         return this.start;
org/jd/core/v1/model/javafragment/EndBlockFragment.class:(22)     }
org/jd/core/v1/model/javafragment/EndBlockFragment.class:(23)     
org/jd/core/v1/model/javafragment/EndBlockFragment.class:(24)     public boolean incLineCount(boolean <b>for</b>ce) {
org/jd/core/v1/model/javafragment/EndBlockFragment.class:(25)         if (this.lineCount < this.maximalLineCount) {
org/jd/core/v1/model/javafragment/EndBlockFragment.class:(26)             this.lineCount++;
org/jd/core/v1/model/javafragment/EndBlockFragment.class:(27)             return true;
--
org/jd/core/v1/model/javafragment/EndBlockFragment.class:(29)         return false;
org/jd/core/v1/model/javafragment/EndBlockFragment.class:(30)     }
org/jd/core/v1/model/javafragment/EndBlockFragment.class:(31)     
org/jd/core/v1/model/javafragment/EndBlockFragment.class:(32)     public boolean decLineCount(boolean <b>for</b>ce) {
org/jd/core/v1/model/javafragment/EndBlockFragment.class:(33)         if (this.lineCount > this.minimalLineCount) {
org/jd/core/v1/model/javafragment/EndBlockFragment.class:(34)             this.lineCount--;
org/jd/core/v1/model/javafragment/EndBlockFragment.class:(35)             return true;
--
org/jd/core/v1/model/javafragment/EndBodyFragment.class:(21)         return this.start;
org/jd/core/v1/model/javafragment/EndBodyFragment.class:(22)     }
org/jd/core/v1/model/javafragment/EndBodyFragment.class:(23)     
org/jd/core/v1/model/javafragment/EndBodyFragment.class:(24)     public boolean incLineCount(boolean <b>for</b>ce) {
org/jd/core/v1/model/javafragment/EndBodyFragment.class:(25)         if (this.lineCount < this.maximalLineCount) {
org/jd/core/v1/model/javafragment/EndBodyFragment.class:(26)             this.lineCount++;
org/jd/core/v1/model/javafragment/EndBodyFragment.class:(27)             if (!<b>for</b>ce)
org/jd/core/v1/model/javafragment/EndBodyFragment.class:(28)                 if (this.lineCount == 1 && this.start.getLineCount() == 0)
org/jd/core/v1/model/javafragment/EndBodyFragment.class:(29)                     this.start.setLineCount(this.lineCount);  
org/jd/core/v1/model/javafragment/EndBodyFragment.class:(30)             return true;
--
org/jd/core/v1/model/javafragment/EndBodyFragment.class:(32)         return false;
org/jd/core/v1/model/javafragment/EndBodyFragment.class:(33)     }
org/jd/core/v1/model/javafragment/EndBodyFragment.class:(34)     
org/jd/core/v1/model/javafragment/EndBodyFragment.class:(35)     public boolean decLineCount(boolean <b>for</b>ce) {
org/jd/core/v1/model/javafragment/EndBodyFragment.class:(36)         if (this.lineCount > this.minimalLineCount) {
org/jd/core/v1/model/javafragment/EndBodyFragment.class:(37)             this.lineCount--;
org/jd/core/v1/model/javafragment/EndBodyFragment.class:(38)             if (!<b>for</b>ce && 
org/jd/core/v1/model/javafragment/EndBodyFragment.class:(39)                 this.lineCount == 0)
org/jd/core/v1/model/javafragment/EndBodyFragment.class:(40)                 this.start.setLineCount(this.lineCount); 
org/jd/core/v1/model/javafragment/EndBodyFragment.class:(41)             return true;
--
org/jd/core/v1/model/javafragment/EndBodyInParameterFragment.class:(9)         super(minimalLineCount, lineCount, maximalLineCount, weight, label, startBodyFragment);
org/jd/core/v1/model/javafragment/EndBodyInParameterFragment.class:(10)     }
org/jd/core/v1/model/javafragment/EndBodyInParameterFragment.class:(11)     
org/jd/core/v1/model/javafragment/EndBodyInParameterFragment.class:(12)     public boolean incLineCount(boolean <b>for</b>ce) {
org/jd/core/v1/model/javafragment/EndBodyInParameterFragment.class:(13)         if (this.lineCount < this.maximalLineCount) {
org/jd/core/v1/model/javafragment/EndBodyInParameterFragment.class:(14)             this.lineCount++;
org/jd/core/v1/model/javafragment/EndBodyInParameterFragment.class:(15)             return true;
--
org/jd/core/v1/model/javafragment/EndBodyInParameterFragment.class:(17)         return false;
org/jd/core/v1/model/javafragment/EndBodyInParameterFragment.class:(18)     }
org/jd/core/v1/model/javafragment/EndBodyInParameterFragment.class:(19)     
org/jd/core/v1/model/javafragment/EndBodyInParameterFragment.class:(20)     public boolean decLineCount(boolean <b>for</b>ce) {
org/jd/core/v1/model/javafragment/EndBodyInParameterFragment.class:(21)         if (this.lineCount > this.minimalLineCount) {
org/jd/core/v1/model/javafragment/EndBodyInParameterFragment.class:(22)             this.lineCount--;
org/jd/core/v1/model/javafragment/EndBodyInParameterFragment.class:(23)             return true;
--
org/jd/core/v1/model/javafragment/EndSingleStatementBlockFragment.class:(21)         return this.start;
org/jd/core/v1/model/javafragment/EndSingleStatementBlockFragment.class:(22)     }
org/jd/core/v1/model/javafragment/EndSingleStatementBlockFragment.class:(23)     
org/jd/core/v1/model/javafragment/EndSingleStatementBlockFragment.class:(24)     public boolean incLineCount(boolean <b>for</b>ce) {
org/jd/core/v1/model/javafragment/EndSingleStatementBlockFragment.class:(25)         if (this.lineCount < this.maximalLineCount) {
org/jd/core/v1/model/javafragment/EndSingleStatementBlockFragment.class:(26)             this.lineCount++;
org/jd/core/v1/model/javafragment/EndSingleStatementBlockFragment.class:(27)             if (!<b>for</b>ce)
org/jd/core/v1/model/javafragment/EndSingleStatementBlockFragment.class:(28)                 if (this.start.getLineCount() == 0)
org/jd/core/v1/model/javafragment/EndSingleStatementBlockFragment.class:(29)                     this.start.setLineCount(1);  
org/jd/core/v1/model/javafragment/EndSingleStatementBlockFragment.class:(30)             return true;
--
org/jd/core/v1/model/javafragment/EndSingleStatementBlockFragment.class:(32)         return false;
org/jd/core/v1/model/javafragment/EndSingleStatementBlockFragment.class:(33)     }
org/jd/core/v1/model/javafragment/EndSingleStatementBlockFragment.class:(34)     
org/jd/core/v1/model/javafragment/EndSingleStatementBlockFragment.class:(35)     public boolean decLineCount(boolean <b>for</b>ce) {
org/jd/core/v1/model/javafragment/EndSingleStatementBlockFragment.class:(36)         if (this.lineCount > this.minimalLineCount) {
org/jd/core/v1/model/javafragment/EndSingleStatementBlockFragment.class:(37)             this.lineCount--;
org/jd/core/v1/model/javafragment/EndSingleStatementBlockFragment.class:(38)             if (!<b>for</b>ce && 
org/jd/core/v1/model/javafragment/EndSingleStatementBlockFragment.class:(39)                 this.lineCount == 0)
org/jd/core/v1/model/javafragment/EndSingleStatementBlockFragment.class:(40)                 this.start.setLineCount(0); 
org/jd/core/v1/model/javafragment/EndSingleStatementBlockFragment.class:(41)             return true;
--
org/jd/core/v1/model/javafragment/ImportsFragment.class:(37)     }
org/jd/core/v1/model/javafragment/ImportsFragment.class:(38)     
org/jd/core/v1/model/javafragment/ImportsFragment.class:(39)     public int getLineCount() {
org/jd/core/v1/model/javafragment/ImportsFragment.class:(40)         assert this.lineCount != -1 : "Call initLineCounts() be<b>for</b>e";
org/jd/core/v1/model/javafragment/ImportsFragment.class:(41)         return this.lineCount;
org/jd/core/v1/model/javafragment/ImportsFragment.class:(42)     }
org/jd/core/v1/model/javafragment/ImportsFragment.class:(43)     
--
org/jd/core/v1/model/javafragment/LineNumberTokensFragment.class:(25)     
org/jd/core/v1/model/javafragment/LineNumberTokensFragment.class:(26)     protected static int searchFirstLineNumber(List<Token> tokens) {
org/jd/core/v1/model/javafragment/LineNumberTokensFragment.class:(27)         SearchLineNumberVisitor visitor = new SearchLineNumberVisitor();
org/jd/core/v1/model/javafragment/LineNumberTokensFragment.class:(28)         <b>for</b> (Token token : tokens) {
org/jd/core/v1/model/javafragment/LineNumberTokensFragment.class:(29)             token.accept(visitor);
org/jd/core/v1/model/javafragment/LineNumberTokensFragment.class:(30)             if (visitor.lineNumber != 0)
org/jd/core/v1/model/javafragment/LineNumberTokensFragment.class:(31)                 return visitor.lineNumber - visitor.newLineCounter; 
--
org/jd/core/v1/model/javafragment/StartBlockFragment.class:(23)         this.lineCount = lineCount;
org/jd/core/v1/model/javafragment/StartBlockFragment.class:(24)     }
org/jd/core/v1/model/javafragment/StartBlockFragment.class:(25)     
org/jd/core/v1/model/javafragment/StartBlockFragment.class:(26)     public boolean incLineCount(boolean <b>for</b>ce) {
org/jd/core/v1/model/javafragment/StartBlockFragment.class:(27)         if (this.lineCount < this.maximalLineCount) {
org/jd/core/v1/model/javafragment/StartBlockFragment.class:(28)             this.lineCount++;
org/jd/core/v1/model/javafragment/StartBlockFragment.class:(29)             if (!<b>for</b>ce)
org/jd/core/v1/model/javafragment/StartBlockFragment.class:(30)                 if (this.lineCount == 1 && this.end.getLineCount() == 0)
org/jd/core/v1/model/javafragment/StartBlockFragment.class:(31)                     this.end.setLineCount(this.lineCount);  
org/jd/core/v1/model/javafragment/StartBlockFragment.class:(32)             return true;
--
org/jd/core/v1/model/javafragment/StartBlockFragment.class:(34)         return false;
org/jd/core/v1/model/javafragment/StartBlockFragment.class:(35)     }
org/jd/core/v1/model/javafragment/StartBlockFragment.class:(36)     
org/jd/core/v1/model/javafragment/StartBlockFragment.class:(37)     public boolean decLineCount(boolean <b>for</b>ce) {
org/jd/core/v1/model/javafragment/StartBlockFragment.class:(38)         if (this.lineCount > this.minimalLineCount) {
org/jd/core/v1/model/javafragment/StartBlockFragment.class:(39)             this.lineCount--;
org/jd/core/v1/model/javafragment/StartBlockFragment.class:(40)             if (!<b>for</b>ce)
org/jd/core/v1/model/javafragment/StartBlockFragment.class:(41)                 if (this.lineCount == 1)
org/jd/core/v1/model/javafragment/StartBlockFragment.class:(42)                     this.end.setLineCount(this.lineCount);  
org/jd/core/v1/model/javafragment/StartBlockFragment.class:(43)             return true;
--
org/jd/core/v1/model/javafragment/StartBodyFragment.class:(23)         this.lineCount = lineCount;
org/jd/core/v1/model/javafragment/StartBodyFragment.class:(24)     }
org/jd/core/v1/model/javafragment/StartBodyFragment.class:(25)     
org/jd/core/v1/model/javafragment/StartBodyFragment.class:(26)     public boolean incLineCount(boolean <b>for</b>ce) {
org/jd/core/v1/model/javafragment/StartBodyFragment.class:(27)         if (this.lineCount < this.maximalLineCount) {
org/jd/core/v1/model/javafragment/StartBodyFragment.class:(28)             this.lineCount++;
org/jd/core/v1/model/javafragment/StartBodyFragment.class:(29)             if (!<b>for</b>ce)
org/jd/core/v1/model/javafragment/StartBodyFragment.class:(30)                 if (this.lineCount == 1 && this.end.getLineCount() == 0)
org/jd/core/v1/model/javafragment/StartBodyFragment.class:(31)                     this.end.setLineCount(this.lineCount);  
org/jd/core/v1/model/javafragment/StartBodyFragment.class:(32)             return true;
--
org/jd/core/v1/model/javafragment/StartBodyFragment.class:(34)         return false;
org/jd/core/v1/model/javafragment/StartBodyFragment.class:(35)     }
org/jd/core/v1/model/javafragment/StartBodyFragment.class:(36)     
org/jd/core/v1/model/javafragment/StartBodyFragment.class:(37)     public boolean decLineCount(boolean <b>for</b>ce) {
org/jd/core/v1/model/javafragment/StartBodyFragment.class:(38)         if (this.lineCount > this.minimalLineCount) {
org/jd/core/v1/model/javafragment/StartBodyFragment.class:(39)             this.lineCount--;
org/jd/core/v1/model/javafragment/StartBodyFragment.class:(40)             if (!<b>for</b>ce)
org/jd/core/v1/model/javafragment/StartBodyFragment.class:(41)                 if (this.lineCount == 1)
org/jd/core/v1/model/javafragment/StartBodyFragment.class:(42)                     this.end.setLineCount(this.lineCount);  
org/jd/core/v1/model/javafragment/StartBodyFragment.class:(43)             return true;
--
org/jd/core/v1/model/javafragment/StartSingleStatementBlockFragment.class:(23)         this.end = end;
org/jd/core/v1/model/javafragment/StartSingleStatementBlockFragment.class:(24)     }
org/jd/core/v1/model/javafragment/StartSingleStatementBlockFragment.class:(25)     
org/jd/core/v1/model/javafragment/StartSingleStatementBlockFragment.class:(26)     public boolean incLineCount(boolean <b>for</b>ce) {
org/jd/core/v1/model/javafragment/StartSingleStatementBlockFragment.class:(27)         if (this.lineCount < this.maximalLineCount) {
org/jd/core/v1/model/javafragment/StartSingleStatementBlockFragment.class:(28)             this.lineCount++;
org/jd/core/v1/model/javafragment/StartSingleStatementBlockFragment.class:(29)             if (!<b>for</b>ce)
org/jd/core/v1/model/javafragment/StartSingleStatementBlockFragment.class:(30)                 if (this.end.getLineCount() == 0)
org/jd/core/v1/model/javafragment/StartSingleStatementBlockFragment.class:(31)                     this.end.setLineCount(1);  
org/jd/core/v1/model/javafragment/StartSingleStatementBlockFragment.class:(32)             return true;
--
org/jd/core/v1/model/javafragment/StartSingleStatementBlockFragment.class:(34)         return false;
org/jd/core/v1/model/javafragment/StartSingleStatementBlockFragment.class:(35)     }
org/jd/core/v1/model/javafragment/StartSingleStatementBlockFragment.class:(36)     
org/jd/core/v1/model/javafragment/StartSingleStatementBlockFragment.class:(37)     public boolean decLineCount(boolean <b>for</b>ce) {
org/jd/core/v1/model/javafragment/StartSingleStatementBlockFragment.class:(38)         if (this.lineCount > this.minimalLineCount) {
org/jd/core/v1/model/javafragment/StartSingleStatementBlockFragment.class:(39)             this.lineCount--;
org/jd/core/v1/model/javafragment/StartSingleStatementBlockFragment.class:(40)             if (!<b>for</b>ce)
org/jd/core/v1/model/javafragment/StartSingleStatementBlockFragment.class:(41)                 if (this.lineCount == 1)
org/jd/core/v1/model/javafragment/StartSingleStatementBlockFragment.class:(42)                     this.end.setLineCount(1);  
org/jd/core/v1/model/javafragment/StartSingleStatementBlockFragment.class:(43)             return true;
--
org/jd/core/v1/model/javafragment/TokensFragment.class:(44)     
org/jd/core/v1/model/javafragment/TokensFragment.class:(45)     protected static int getLineCount(List<Token> tokens) {
org/jd/core/v1/model/javafragment/TokensFragment.class:(46)         LineCountVisitor visitor = new LineCountVisitor();
org/jd/core/v1/model/javafragment/TokensFragment.class:(47)         <b>for</b> (Token token : tokens)
org/jd/core/v1/model/javafragment/TokensFragment.class:(48)             token.accept(visitor); 
org/jd/core/v1/model/javafragment/TokensFragment.class:(49)         return visitor.lineCount;
org/jd/core/v1/model/javafragment/TokensFragment.class:(50)     }
--
org/jd/core/v1/model/javasyntax/AbstractJavaSyntaxVisitor.class:(605)     }
org/jd/core/v1/model/javasyntax/AbstractJavaSyntaxVisitor.class:(606)     
org/jd/core/v1/model/javasyntax/AbstractJavaSyntaxVisitor.class:(607)     protected void acceptListDeclaration(List<? extends Declaration> list) {
org/jd/core/v1/model/javasyntax/AbstractJavaSyntaxVisitor.class:(608)         <b>for</b> (Declaration declaration : list)
org/jd/core/v1/model/javasyntax/AbstractJavaSyntaxVisitor.class:(609)             declaration.accept(this); 
org/jd/core/v1/model/javasyntax/AbstractJavaSyntaxVisitor.class:(610)     }
org/jd/core/v1/model/javasyntax/AbstractJavaSyntaxVisitor.class:(611)     
--
org/jd/core/v1/model/javasyntax/AbstractJavaSyntaxVisitor.class:(612)     protected void acceptListExpression(List<? extends Expression> list) {
org/jd/core/v1/model/javasyntax/AbstractJavaSyntaxVisitor.class:(613)         <b>for</b> (Expression expression : list)
org/jd/core/v1/model/javasyntax/AbstractJavaSyntaxVisitor.class:(614)             expression.accept(this); 
org/jd/core/v1/model/javasyntax/AbstractJavaSyntaxVisitor.class:(615)     }
org/jd/core/v1/model/javasyntax/AbstractJavaSyntaxVisitor.class:(616)     
--
org/jd/core/v1/model/javasyntax/AbstractJavaSyntaxVisitor.class:(617)     protected void acceptListReference(List<? extends Reference> list) {
org/jd/core/v1/model/javasyntax/AbstractJavaSyntaxVisitor.class:(618)         <b>for</b> (Reference reference : list)
org/jd/core/v1/model/javasyntax/AbstractJavaSyntaxVisitor.class:(619)             reference.accept(this); 
org/jd/core/v1/model/javasyntax/AbstractJavaSyntaxVisitor.class:(620)     }
org/jd/core/v1/model/javasyntax/AbstractJavaSyntaxVisitor.class:(621)     
--
org/jd/core/v1/model/javasyntax/AbstractJavaSyntaxVisitor.class:(622)     protected void acceptListStatement(List<? extends Statement> list) {
org/jd/core/v1/model/javasyntax/AbstractJavaSyntaxVisitor.class:(623)         <b>for</b> (Statement statement : list)
org/jd/core/v1/model/javasyntax/AbstractJavaSyntaxVisitor.class:(624)             statement.accept(this); 
org/jd/core/v1/model/javasyntax/AbstractJavaSyntaxVisitor.class:(625)     }
org/jd/core/v1/model/javasyntax/AbstractJavaSyntaxVisitor.class:(626)     
--
org/jd/core/v1/model/javasyntax/AbstractJavaSyntaxVisitor.class:(656)     
org/jd/core/v1/model/javasyntax/AbstractJavaSyntaxVisitor.class:(657)     protected void safeAcceptListDeclaration(List<? extends Declaration> list) {
org/jd/core/v1/model/javasyntax/AbstractJavaSyntaxVisitor.class:(658)         if (list != null)
org/jd/core/v1/model/javasyntax/AbstractJavaSyntaxVisitor.class:(659)             <b>for</b> (Declaration declaration : list)
org/jd/core/v1/model/javasyntax/AbstractJavaSyntaxVisitor.class:(660)                 declaration.accept(this);  
org/jd/core/v1/model/javasyntax/AbstractJavaSyntaxVisitor.class:(661)     }
org/jd/core/v1/model/javasyntax/AbstractJavaSyntaxVisitor.class:(662)     
--
org/jd/core/v1/model/javasyntax/AbstractJavaSyntaxVisitor.class:(663)     protected void safeAcceptListStatement(List<? extends Statement> list) {
org/jd/core/v1/model/javasyntax/AbstractJavaSyntaxVisitor.class:(664)         if (list != null)
org/jd/core/v1/model/javasyntax/AbstractJavaSyntaxVisitor.class:(665)             <b>for</b> (Statement statement : list)
org/jd/core/v1/model/javasyntax/AbstractJavaSyntaxVisitor.class:(666)                 statement.accept(this);  
org/jd/core/v1/model/javasyntax/AbstractJavaSyntaxVisitor.class:(667)     }
org/jd/core/v1/model/javasyntax/AbstractJavaSyntaxVisitor.class:(668) }
--
org/jd/core/v1/model/javasyntax/declaration/ConstructorDeclaration.class:(14)     
org/jd/core/v1/model/javasyntax/declaration/ConstructorDeclaration.class:(15)     protected BaseTypeParameter typeParameters;
org/jd/core/v1/model/javasyntax/declaration/ConstructorDeclaration.class:(16)     
org/jd/core/v1/model/javasyntax/declaration/ConstructorDeclaration.class:(17)     protected BaseFormalParameter <b>for</b>malParameters;
org/jd/core/v1/model/javasyntax/declaration/ConstructorDeclaration.class:(18)     
org/jd/core/v1/model/javasyntax/declaration/ConstructorDeclaration.class:(19)     protected BaseType exceptionTypes;
org/jd/core/v1/model/javasyntax/declaration/ConstructorDeclaration.class:(20)     
--
org/jd/core/v1/model/javasyntax/declaration/ConstructorDeclaration.class:(22)     
org/jd/core/v1/model/javasyntax/declaration/ConstructorDeclaration.class:(23)     protected BaseStatement statements;
org/jd/core/v1/model/javasyntax/declaration/ConstructorDeclaration.class:(24)     
org/jd/core/v1/model/javasyntax/declaration/ConstructorDeclaration.class:(25)     public ConstructorDeclaration(int flags, BaseFormalParameter <b>for</b>malParameters, String descriptor, BaseStatement statements) {
org/jd/core/v1/model/javasyntax/declaration/ConstructorDeclaration.class:(26)         this.flags = flags;
org/jd/core/v1/model/javasyntax/declaration/ConstructorDeclaration.class:(27)         this.<b>for</b>        this.formalParameters = <b>for</b>malParameters;
org/jd/core/v1/model/javasyntax/declaration/ConstructorDeclaration.class:(28)         this.descriptor = descriptor;
org/jd/core/v1/model/javasyntax/declaration/ConstructorDeclaration.class:(29)         this.statements = statements;
org/jd/core/v1/model/javasyntax/declaration/ConstructorDeclaration.class:(30)     }
--
org/jd/core/v1/model/javasyntax/declaration/ConstructorDeclaration.class:(31)     
org/jd/core/v1/model/javasyntax/declaration/ConstructorDeclaration.class:(32)     public ConstructorDeclaration(BaseAnnotationReference annotationReferences, int flags, BaseTypeParameter typeParameters, BaseFormalParameter <b>for</b>malParameters, BaseType exceptionTypes, String descriptor, BaseStatement statements) {
org/jd/core/v1/model/javasyntax/declaration/ConstructorDeclaration.class:(33)         this.annotationReferences = annotationReferences;
org/jd/core/v1/model/javasyntax/declaration/ConstructorDeclaration.class:(34)         this.flags = flags;
org/jd/core/v1/model/javasyntax/declaration/ConstructorDeclaration.class:(35)         this.typeParameters = typeParameters;
--
org/jd/core/v1/model/javasyntax/declaration/ConstructorDeclaration.class:(36)         this.<b>for</b>        this.formalParameters = <b>for</b>malParameters;
org/jd/core/v1/model/javasyntax/declaration/ConstructorDeclaration.class:(37)         this.exceptionTypes = exceptionTypes;
org/jd/core/v1/model/javasyntax/declaration/ConstructorDeclaration.class:(38)         this.descriptor = descriptor;
org/jd/core/v1/model/javasyntax/declaration/ConstructorDeclaration.class:(39)         this.statements = statements;
--
org/jd/core/v1/model/javasyntax/declaration/ConstructorDeclaration.class:(52)     }
org/jd/core/v1/model/javasyntax/declaration/ConstructorDeclaration.class:(53)     
org/jd/core/v1/model/javasyntax/declaration/ConstructorDeclaration.class:(54)     public BaseFormalParameter getFormalParameters() {
org/jd/core/v1/model/javasyntax/declaration/ConstructorDeclaration.class:(55)         return this.<b>for</b>malParameters;
org/jd/core/v1/model/javasyntax/declaration/ConstructorDeclaration.class:(56)     }
org/jd/core/v1/model/javasyntax/declaration/ConstructorDeclaration.class:(57)     
org/jd/core/v1/model/javasyntax/declaration/ConstructorDeclaration.class:(58)     public BaseType getExceptionTypes() {
--
org/jd/core/v1/model/javasyntax/declaration/FieldDeclarators.class:(24)     }
org/jd/core/v1/model/javasyntax/declaration/FieldDeclarators.class:(25)     
org/jd/core/v1/model/javasyntax/declaration/FieldDeclarators.class:(26)     public void setFieldDeclaration(FieldDeclaration fieldDeclaration) {
org/jd/core/v1/model/javasyntax/declaration/FieldDeclarators.class:(27)         <b>for</b> (FieldDeclarator fieldDeclarator : this)
org/jd/core/v1/model/javasyntax/declaration/FieldDeclarators.class:(28)             fieldDeclarator.setFieldDeclaration(fieldDeclaration); 
org/jd/core/v1/model/javasyntax/declaration/FieldDeclarators.class:(29)     }
org/jd/core/v1/model/javasyntax/declaration/FieldDeclarators.class:(30)     
--
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(20)     
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(21)     protected Type returnedType;
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(22)     
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(23)     protected BaseFormalParameter <b>for</b>malParameters;
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(24)     
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(25)     protected BaseType exceptionTypes;
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(26)     
--
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(53)         this.defaultAnnotationValue = defaultAnnotationValue;
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(54)     }
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(55)     
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(56)     public MethodDeclaration(int flags, String name, Type returnedType, BaseFormalParameter <b>for</b>malParameters, String descriptor, BaseStatement statements) {
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(57)         this.flags = flags;
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(58)         this.name = name;
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(59)         this.returnedType = returnedType;
--
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(60)         this.<b>for</b>        this.formalParameters = <b>for</b>malParameters;
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(61)         this.descriptor = descriptor;
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(62)         this.statements = statements;
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(63)     }
--
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(64)     
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(65)     public MethodDeclaration(int flags, String name, Type returnedType, BaseFormalParameter <b>for</b>malParameters, String descriptor, ElementValue defaultAnnotationValue) {
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(66)         this.flags = flags;
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(67)         this.name = name;
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(68)         this.returnedType = returnedType;
--
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(69)         this.<b>for</b>        this.formalParameters = <b>for</b>malParameters;
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(70)         this.descriptor = descriptor;
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(71)         this.defaultAnnotationValue = defaultAnnotationValue;
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(72)     }
--
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(73)     
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(74)     public MethodDeclaration(BaseAnnotationReference annotationReferences, int flags, String name, BaseTypeParameter typeParameters, Type returnedType, BaseFormalParameter <b>for</b>malParameters, BaseType exceptionTypes, String descriptor, BaseStatement statements, ElementValue defaultAnnotationValue) {
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(75)         this.annotationReferences = annotationReferences;
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(76)         this.flags = flags;
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(77)         this.name = name;
--
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(78)         this.typeParameters = typeParameters;
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(79)         this.returnedType = returnedType;
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(80)         this.<b>for</b>        this.formalParameters = <b>for</b>malParameters;
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(81)         this.exceptionTypes = exceptionTypes;
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(82)         this.descriptor = descriptor;
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(83)         this.statements = statements;
--
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(105)     }
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(106)     
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(107)     public BaseFormalParameter getFormalParameters() {
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(108)         return this.<b>for</b>malParameters;
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(109)     }
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(110)     
org/jd/core/v1/model/javasyntax/declaration/MethodDeclaration.class:(111)     public BaseType getExceptionTypes() {
--
org/jd/core/v1/model/javasyntax/type/AbstractTypeArgumentVisitor.class:(14) 
org/jd/core/v1/model/javasyntax/type/AbstractTypeArgumentVisitor.class:(15) public abstract class AbstractTypeArgumentVisitor implements TypeArgumentVisitor {
org/jd/core/v1/model/javasyntax/type/AbstractTypeArgumentVisitor.class:(16)     public void visit(TypeArguments arguments) {
org/jd/core/v1/model/javasyntax/type/AbstractTypeArgumentVisitor.class:(17)         <b>for</b> (TypeArgument typeArgument : arguments)
org/jd/core/v1/model/javasyntax/type/AbstractTypeArgumentVisitor.class:(18)             typeArgument.accept(this); 
org/jd/core/v1/model/javasyntax/type/AbstractTypeArgumentVisitor.class:(19)     }
org/jd/core/v1/model/javasyntax/type/AbstractTypeArgumentVisitor.class:(20)     
--
org/jd/core/v1/model/javasyntax/type/ObjectType.class:(231)             GenericType gt = (GenericType)typeArgument;
org/jd/core/v1/model/javasyntax/type/ObjectType.class:(232)             BaseType bt = typeBounds.get(gt.getName());
org/jd/core/v1/model/javasyntax/type/ObjectType.class:(233)             if (bt != null)
org/jd/core/v1/model/javasyntax/type/ObjectType.class:(234)                 <b>for</b> (Type type : bt) {
org/jd/core/v1/model/javasyntax/type/ObjectType.class:(235)                     if (this.dimension == type.getDimension()) {
org/jd/core/v1/model/javasyntax/type/ObjectType.class:(236)                         Class<?> typeClass = type.getClass();
org/jd/core/v1/model/javasyntax/type/ObjectType.class:(237)                         if (typeClass == ObjectType.class || typeClass == InnerObjectType.class) {
--
org/jd/core/v1/model/javasyntax/type/TypeParameters.class:(21)         super(types.length + 1);
org/jd/core/v1/model/javasyntax/type/TypeParameters.class:(22)         assert types != null && types.length > 0 : "Uses 'TypeParameter' instead";
org/jd/core/v1/model/javasyntax/type/TypeParameters.class:(23)         add(type);
org/jd/core/v1/model/javasyntax/type/TypeParameters.class:(24)         <b>for</b> (TypeParameter t : types)
org/jd/core/v1/model/javasyntax/type/TypeParameters.class:(25)             add(t); 
org/jd/core/v1/model/javasyntax/type/TypeParameters.class:(26)     }
org/jd/core/v1/model/javasyntax/type/TypeParameters.class:(27)     
--
org/jd/core/v1/model/javasyntax/type/TypeParameters.class:(32)     public String toString() {
org/jd/core/v1/model/javasyntax/type/TypeParameters.class:(33)         StringBuilder sb = new StringBuilder();
org/jd/core/v1/model/javasyntax/type/TypeParameters.class:(34)         sb.append(get(0).toString());
org/jd/core/v1/model/javasyntax/type/TypeParameters.class:(35)         <b>for</b> (int i = 1; i < size(); i++) {
org/jd/core/v1/model/javasyntax/type/TypeParameters.class:(36)             sb.append(" & ");
org/jd/core/v1/model/javasyntax/type/TypeParameters.class:(37)             sb.append(get(i).toString());
org/jd/core/v1/model/javasyntax/type/TypeParameters.class:(38)         } 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/model/cfg/BasicBlock.class:(274)             return true; 
org/jd/core/v1/service/converter/classfiletojavasyntax/model/cfg/BasicBlock.class:(275)         if (this.branch == basicBlock)
org/jd/core/v1/service/converter/classfiletojavasyntax/model/cfg/BasicBlock.class:(276)             return true; 
org/jd/core/v1/service/converter/classfiletojavasyntax/model/cfg/BasicBlock.class:(277)         <b>for</b> (ExceptionHandler exceptionHandler : this.exceptionHandlers) {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/cfg/BasicBlock.class:(278)             if (exceptionHandler.getBasicBlock() == basicBlock)
org/jd/core/v1/service/converter/classfiletojavasyntax/model/cfg/BasicBlock.class:(279)                 return true; 
org/jd/core/v1/service/converter/classfiletojavasyntax/model/cfg/BasicBlock.class:(280)         } 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/model/cfg/BasicBlock.class:(281)         <b>for</b> (SwitchCase switchCase : this.switchCases) {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/cfg/BasicBlock.class:(282)             if (switchCase.getBasicBlock() == basicBlock)
org/jd/core/v1/service/converter/classfiletojavasyntax/model/cfg/BasicBlock.class:(283)                 return true; 
org/jd/core/v1/service/converter/classfiletojavasyntax/model/cfg/BasicBlock.class:(284)         } 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/model/cfg/BasicBlock.class:(294)             this.next = nevv; 
org/jd/core/v1/service/converter/classfiletojavasyntax/model/cfg/BasicBlock.class:(295)         if (this.branch == old)
org/jd/core/v1/service/converter/classfiletojavasyntax/model/cfg/BasicBlock.class:(296)             this.branch = nevv; 
org/jd/core/v1/service/converter/classfiletojavasyntax/model/cfg/BasicBlock.class:(297)         <b>for</b> (ExceptionHandler exceptionHandler : this.exceptionHandlers)
org/jd/core/v1/service/converter/classfiletojavasyntax/model/cfg/BasicBlock.class:(298)             exceptionHandler.replace(old, nevv); 
org/jd/core/v1/service/converter/classfiletojavasyntax/model/cfg/BasicBlock.class:(299)         <b>for</b> (SwitchCase switchCase : this.switchCases)
org/jd/core/v1/service/converter/classfiletojavasyntax/model/cfg/BasicBlock.class:(300)             switchCase.replace(old, nevv); 
org/jd/core/v1/service/converter/classfiletojavasyntax/model/cfg/BasicBlock.class:(301)         if (this.sub1 == old)
org/jd/core/v1/service/converter/classfiletojavasyntax/model/cfg/BasicBlock.class:(302)             this.sub1 = nevv; 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/model/cfg/BasicBlock.class:(314)             this.next = nevv; 
org/jd/core/v1/service/converter/classfiletojavasyntax/model/cfg/BasicBlock.class:(315)         if (olds.contains(this.branch))
org/jd/core/v1/service/converter/classfiletojavasyntax/model/cfg/BasicBlock.class:(316)             this.branch = nevv; 
org/jd/core/v1/service/converter/classfiletojavasyntax/model/cfg/BasicBlock.class:(317)         <b>for</b> (ExceptionHandler exceptionHandler : this.exceptionHandlers)
org/jd/core/v1/service/converter/classfiletojavasyntax/model/cfg/BasicBlock.class:(318)             exceptionHandler.replace(olds, nevv); 
org/jd/core/v1/service/converter/classfiletojavasyntax/model/cfg/BasicBlock.class:(319)         <b>for</b> (SwitchCase switchCase : this.switchCases)
org/jd/core/v1/service/converter/classfiletojavasyntax/model/cfg/BasicBlock.class:(320)             switchCase.replace(olds, nevv); 
org/jd/core/v1/service/converter/classfiletojavasyntax/model/cfg/BasicBlock.class:(321)         if (olds.contains(this.sub1))
org/jd/core/v1/service/converter/classfiletojavasyntax/model/cfg/BasicBlock.class:(322)             this.sub1 = nevv; 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/model/cfg/BasicBlock.class:(331)             this.exceptionHandlers = new DefaultList<>();
org/jd/core/v1/service/converter/classfiletojavasyntax/model/cfg/BasicBlock.class:(332)             this.exceptionHandlers.add(new ExceptionHandler(internalThrowableName, basicBlock));
org/jd/core/v1/service/converter/classfiletojavasyntax/model/cfg/BasicBlock.class:(333)         } else {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/cfg/BasicBlock.class:(334)             <b>for</b> (ExceptionHandler exceptionHandler : this.exceptionHandlers) {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/cfg/BasicBlock.class:(335)                 if (exceptionHandler.getBasicBlock() == basicBlock) {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/cfg/BasicBlock.class:(336)                     exceptionHandler.addInternalThrowableName(internalThrowableName);
org/jd/core/v1/service/converter/classfiletojavasyntax/model/cfg/BasicBlock.class:(337)                     return;
--
org/jd/core/v1/service/converter/classfiletojavasyntax/model/javasyntax/declaration/ClassFileBodyDeclaration.class:(72)         if (innerTypeDeclarations != null) {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/javasyntax/declaration/ClassFileBodyDeclaration.class:(73)             updateFirstLineNumber((List)(this.innerTypeDeclarations = innerTypeDeclarations));
org/jd/core/v1/service/converter/classfiletojavasyntax/model/javasyntax/declaration/ClassFileBodyDeclaration.class:(74)             this.innerTypeMap = new HashMap<>();
org/jd/core/v1/service/converter/classfiletojavasyntax/model/javasyntax/declaration/ClassFileBodyDeclaration.class:(75)             <b>for</b> (ClassFileTypeDeclaration innerType : innerTypeDeclarations)
org/jd/core/v1/service/converter/classfiletojavasyntax/model/javasyntax/declaration/ClassFileBodyDeclaration.class:(76)                 this.innerTypeMap.put(innerType.getInternalTypeName(), innerType); 
org/jd/core/v1/service/converter/classfiletojavasyntax/model/javasyntax/declaration/ClassFileBodyDeclaration.class:(77)         } 
org/jd/core/v1/service/converter/classfiletojavasyntax/model/javasyntax/declaration/ClassFileBodyDeclaration.class:(78)     }
--
org/jd/core/v1/service/converter/classfiletojavasyntax/model/javasyntax/declaration/ClassFileBodyDeclaration.class:(91)     }
org/jd/core/v1/service/converter/classfiletojavasyntax/model/javasyntax/declaration/ClassFileBodyDeclaration.class:(92)     
org/jd/core/v1/service/converter/classfiletojavasyntax/model/javasyntax/declaration/ClassFileBodyDeclaration.class:(93)     protected void updateFirstLineNumber(List<? extends ClassFileMemberDeclaration> members) {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/javasyntax/declaration/ClassFileBodyDeclaration.class:(94)         <b>for</b> (ClassFileMemberDeclaration member : members) {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/javasyntax/declaration/ClassFileBodyDeclaration.class:(95)             int lineNumber = member.getFirstLineNumber();
org/jd/core/v1/service/converter/classfiletojavasyntax/model/javasyntax/declaration/ClassFileBodyDeclaration.class:(96)             if (lineNumber > 0) {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/javasyntax/declaration/ClassFileBodyDeclaration.class:(97)                 if (this.firstLineNumber == 0) {
--
org/jd/core/v1/service/converter/classfiletojavasyntax/model/javasyntax/declaration/ClassFileConstructorDeclaration.class:(43)         this.flags = flags;
org/jd/core/v1/service/converter/classfiletojavasyntax/model/javasyntax/declaration/ClassFileConstructorDeclaration.class:(44)     }
org/jd/core/v1/service/converter/classfiletojavasyntax/model/javasyntax/declaration/ClassFileConstructorDeclaration.class:(45)     
org/jd/core/v1/service/converter/classfiletojavasyntax/model/javasyntax/declaration/ClassFileConstructorDeclaration.class:(46)     public void setFormalParameters(BaseFormalParameter <b>for</b>malParameters) {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/javasyntax/declaration/ClassFileConstructorDeclaration.class:(47)         this.<b>for</b>        this.formalParameters = <b>for</b>malParameters;
org/jd/core/v1/service/converter/classfiletojavasyntax/model/javasyntax/declaration/ClassFileConstructorDeclaration.class:(48)     }
org/jd/core/v1/service/converter/classfiletojavasyntax/model/javasyntax/declaration/ClassFileConstructorDeclaration.class:(49)     
org/jd/core/v1/service/converter/classfiletojavasyntax/model/javasyntax/declaration/ClassFileConstructorDeclaration.class:(50)     public void setStatements(BaseStatement statements) {
--
org/jd/core/v1/service/converter/classfiletojavasyntax/model/javasyntax/declaration/ClassFileMethodDeclaration.class:(65)         this.flags = flags;
org/jd/core/v1/service/converter/classfiletojavasyntax/model/javasyntax/declaration/ClassFileMethodDeclaration.class:(66)     }
org/jd/core/v1/service/converter/classfiletojavasyntax/model/javasyntax/declaration/ClassFileMethodDeclaration.class:(67)     
org/jd/core/v1/service/converter/classfiletojavasyntax/model/javasyntax/declaration/ClassFileMethodDeclaration.class:(68)     public void setFormalParameters(BaseFormalParameter <b>for</b>malParameters) {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/javasyntax/declaration/ClassFileMethodDeclaration.class:(69)         this.<b>for</b>        this.formalParameters = <b>for</b>malParameters;
org/jd/core/v1/service/converter/classfiletojavasyntax/model/javasyntax/declaration/ClassFileMethodDeclaration.class:(70)     }
org/jd/core/v1/service/converter/classfiletojavasyntax/model/javasyntax/declaration/ClassFileMethodDeclaration.class:(71)     
org/jd/core/v1/service/converter/classfiletojavasyntax/model/javasyntax/declaration/ClassFileMethodDeclaration.class:(72)     public void setStatements(BaseStatement statements) {
--
org/jd/core/v1/service/converter/classfiletojavasyntax/model/javasyntax/declaration/ClassFileStaticInitializerDeclaration.class:(51)     
org/jd/core/v1/service/converter/classfiletojavasyntax/model/javasyntax/declaration/ClassFileStaticInitializerDeclaration.class:(52)     public void setFlags(int flags) {}
org/jd/core/v1/service/converter/classfiletojavasyntax/model/javasyntax/declaration/ClassFileStaticInitializerDeclaration.class:(53)     
org/jd/core/v1/service/converter/classfiletojavasyntax/model/javasyntax/declaration/ClassFileStaticInitializerDeclaration.class:(54)     public void setFormalParameters(BaseFormalParameter <b>for</b>malParameters) {}
org/jd/core/v1/service/converter/classfiletojavasyntax/model/javasyntax/declaration/ClassFileStaticInitializerDeclaration.class:(55)     
org/jd/core/v1/service/converter/classfiletojavasyntax/model/javasyntax/declaration/ClassFileStaticInitializerDeclaration.class:(56)     public void setStatements(BaseStatement statements) {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/javasyntax/declaration/ClassFileStaticInitializerDeclaration.class:(57)         this.statements = statements;
--
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/AbstractLocalVariable.class:(125)     
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/AbstractLocalVariable.class:(126)     protected void fireChangeEvent(Map<String, BaseType> typeBounds) {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/AbstractLocalVariable.class:(127)         if (this.variablesOnLeft != null)
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/AbstractLocalVariable.class:(128)             <b>for</b> (AbstractLocalVariable v : this.variablesOnLeft)
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/AbstractLocalVariable.class:(129)                 v.variableOnRight(typeBounds, this);  
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/AbstractLocalVariable.class:(130)         if (this.variablesOnRight != null)
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/AbstractLocalVariable.class:(131)             <b>for</b> (AbstractLocalVariable v : this.variablesOnRight)
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/AbstractLocalVariable.class:(132)                 v.variableOnLeft(typeBounds, this);  
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/AbstractLocalVariable.class:(133)     }
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/AbstractLocalVariable.class:(134)     
--
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(112)             }  
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(113)         if (alvToMerge == null) {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(114)             if (this.children != null)
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(115)                 <b>for</b> (Frame frame : this.children)
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(116)                     frame.mergeLocalVariable(typeBounds, localVariableMaker, lv);  
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(117)         } else if (lv != alvToMerge) {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(118)             <b>for</b> (LocalVariableReference reference : alvToMerge.getReferences())
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(119)                 reference.setLocalVariable(lv); 
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(120)             lv.getReferences().addAll(alvToMerge.getReferences());
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(121)             lv.setFromOffset(alvToMerge.getFromOffset());
--
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(152)         } 
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(153)         if (alvToRemove == null) {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(154)             if (this.children != null)
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(155)                 <b>for</b> (Frame frame : this.children)
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(156)                     frame.removeLocalVariable(lv);  
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(157)         } else {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(158)             this.localVariableArray[index] = alvToRemove.getNext();
--
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(168)     
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(169)     public void close() {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(170)         if (this.newExpressions != null)
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(171)             <b>for</b> (Map.Entry<NewExpression, AbstractLocalVariable> entry : this.newExpressions.entrySet()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(172)                 ObjectType ot1 = ((NewExpression)entry.getKey()).getObjectType();
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(173)                 ObjectType ot2 = (ObjectType)((AbstractLocalVariable)entry.getValue()).getType();
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(174)                 if (ot1.getTypeArguments() == null && ot2.getTypeArguments() != null)
--
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(180)         HashSet<String> names = new HashSet<>(parentNames);
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(181)         HashMap<Type, Boolean> types = new HashMap<>();
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(182)         int length = this.localVariableArray.length;
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(183)         <b>for</b> (int i = 0; i < length; i++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(184)             AbstractLocalVariable lv = this.localVariableArray[i];
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(185)             while (lv != null) {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(186)                 if (types.containsKey(lv.getType())) {
--
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(206)             }  
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(207)         if (!types.isEmpty()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(208)             GenerateLocalVariableNameVisitor visitor = new GenerateLocalVariableNameVisitor(names, types);
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(209)             <b>for</b> (int j = 0; j < length; j++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(210)                 AbstractLocalVariable lv = this.localVariableArray[j];
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(211)                 while (lv != null) {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(212)                     if (lv.name == null) {
--
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(222)             } 
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(223)         } 
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(224)         if (this.children != null)
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(225)             <b>for</b> (Frame child : this.children)
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(226)                 child.createNames(names);  
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(227)     }
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(228)     
--
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(232)         if (containsLineNumber)
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(233)             mergeDeclarations(); 
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(234)         if (this.children != null)
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(235)             <b>for</b> (Frame child : this.children)
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(236)                 child.createDeclarations();  
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(237)     }
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(238)     
--
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(241)         HashMap<Frame, HashSet<AbstractLocalVariable>> map = createMapForInlineDeclarations();
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(242)         if (!map.isEmpty()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(243)             SearchUndeclaredLocalVariableVisitor visitor = new SearchUndeclaredLocalVariableVisitor();
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(244)             <b>for</b> (Map.Entry<Frame, HashSet<AbstractLocalVariable>> entry : map.entrySet()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(245)                 Statements statements = ((Frame)entry.getKey()).statements;
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(246)                 ListIterator<Statement> iterator = statements.listIterator();
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(247)                 HashSet<AbstractLocalVariable> undeclaredLocalVariables = entry.getValue();
--
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(266)                                 iterator.previous(); 
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(267)                             DefaultList<AbstractLocalVariable> sorted = new DefaultList<>(undeclaredLocalVariablesInStatement);
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(268)                             sorted.sort(ABSTRACT_LOCAL_VARIABLE_COMPARATOR);
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(269)                             <b>for</b> (AbstractLocalVariable lv : sorted) {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(270)                                 iterator.add(new LocalVariableDeclarationStatement(lv.getType(), new LocalVariableDeclarator(lv.getName())));
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(271)                                 lv.setDeclared(true);
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(272)                                 undeclaredLocalVariables.remove(lv);
--
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(312)                 Expressions expressions = new Expressions();
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(313)                 splitMultiAssignment(2147483647, undeclaredLocalVariablesInStatement, expressions, boe);
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(314)                 iterator.remove();
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(315)                 <b>for</b> (Expression exp : expressions)
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(316)                     iterator.add(newDeclarationStatement(undeclaredLocalVariables, undeclaredLocalVariablesInStatement, (BinaryOperatorExpression)exp)); 
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(317)                 if (expressions.isEmpty())
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(318)                     iterator.add(es); 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(368)             Expressions expressions = new Expressions();
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(369)             int toOffset = fs.getToOffset();
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(370)             if (init.isList()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(371)                 <b>for</b> (Expression exp : init) {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(372)                     splitMultiAssignment(toOffset, undeclaredLocalVariablesInStatement, expressions, exp);
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(373)                     if (expressions.isEmpty())
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(374)                         expressions.add(exp); 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(386)         } 
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(387)     }
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(388)     
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(389)     protected void updateForStatement(HashSet<AbstractLocalVariable> undeclaredLocalVariables, HashSet<AbstractLocalVariable> undeclaredLocalVariablesInStatement, ClassFileForStatement <b>for</b>Statement, Expression init) {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(390)         if (init.getClass() != BinaryOperatorExpression.class)
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(391)             return; 
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(392)         BinaryOperatorExpression boe = (BinaryOperatorExpression)init;
--
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(394)             return; 
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(395)         ClassFileLocalVariableReferenceExpression reference = (ClassFileLocalVariableReferenceExpression)boe.getLeftExpression();
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(396)         AbstractLocalVariable localVariable = reference.getLocalVariable();
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(397)         if (localVariable.isDeclared() || localVariable.getToOffset() > <b>for</b>Statement.getToOffset())
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(398)             return; 
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(399)         undeclaredLocalVariables.remove(localVariable);
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(400)         undeclaredLocalVariablesInStatement.remove(localVariable);
--
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(401)         localVariable.setDeclared(true);
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(402)         VariableInitializer variableInitializer = (boe.getRightExpression().getClass() == NewInitializedArray.class) ? ((NewInitializedArray)boe.getRightExpression()).getArrayInitializer() : new ExpressionVariableInitializer(boe.getRightExpression());
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(403)         <b>for</b>Statement.setDeclaration(new LocalVariableDeclaration(localVariable.getType(), new LocalVariableDeclarator(boe.getLineNumber(), reference.getName(), variableInitializer)));
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(404)         <b>for</b>Statement.setInit(null);
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(405)     }
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(406)     
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(407)     protected void updateForStatement(HashSet<AbstractLocalVariable> variablesToDeclare, HashSet<AbstractLocalVariable> foundVariables, ClassFileForStatement <b>for</b>Statement, Expressions init) {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(408)         DefaultList<BinaryOperatorExpression> boes = new DefaultList<>();
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(409)         DefaultList<AbstractLocalVariable> localVariables = new DefaultList<>();
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(410)         Type type0 = null, type1 = null;
--
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(411)         int minDimension = 0, maxDimension = 0;
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(412)         <b>for</b> (Expression expression : init) {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(413)             if (expression.getClass() != BinaryOperatorExpression.class)
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(414)                 return; 
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(415)             BinaryOperatorExpression boe = (BinaryOperatorExpression)expression;
--
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(416)             if (boe.getLeftExpression().getClass() != ClassFileLocalVariableReferenceExpression.class)
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(417)                 return; 
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(418)             AbstractLocalVariable localVariable = ((ClassFileLocalVariableReferenceExpression)boe.getLeftExpression()).getLocalVariable();
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(419)             if (localVariable.isDeclared() || localVariable.getToOffset() > <b>for</b>Statement.getToOffset())
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(420)                 return; 
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(421)             if (type1 == null) {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(422)                 type1 = localVariable.getType();
--
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(435)             localVariables.add(localVariable);
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(436)             boes.add(boe);
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(437)         } 
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(438)         <b>for</b> (AbstractLocalVariable lv : localVariables) {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(439)             variablesToDeclare.remove(lv);
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(440)             foundVariables.remove(lv);
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(441)             lv.setDeclared(true);
--
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(442)         } 
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(443)         if (minDimension == maxDimension) {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(444)             <b>for</b>Statement.setDeclaration(new LocalVariableDeclaration(type1, createDeclarators1(boes, false)));
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(445)         } else {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(446)             <b>for</b>Statement.setDeclaration(new LocalVariableDeclaration(type0, createDeclarators1(boes, true)));
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(447)         } 
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(448)         <b>for</b>Statement.setInit(null);
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(449)     }
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(450)     
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(451)     protected LocalVariableDeclarators createDeclarators1(DefaultList<BinaryOperatorExpression> boes, boolean setDimension) {
--
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(452)         LocalVariableDeclarators declarators = new LocalVariableDeclarators(boes.size());
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(453)         <b>for</b> (BinaryOperatorExpression boe : boes) {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(454)             ClassFileLocalVariableReferenceExpression reference = (ClassFileLocalVariableReferenceExpression)boe.getLeftExpression();
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(455)             VariableInitializer variableInitializer = (boe.getRightExpression().getClass() == NewInitializedArray.class) ? ((NewInitializedArray)boe.getRightExpression()).getArrayInitializer() : new ExpressionVariableInitializer(boe.getRightExpression());
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(456)             LocalVariableDeclarator declarator = new LocalVariableDeclarator(boe.getLineNumber(), reference.getName(), variableInitializer);
--
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(565)     
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(566)     protected LocalVariableDeclarators createDeclarators2(DefaultList<LocalVariableDeclarationStatement> declarations, boolean setDimension) {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(567)         LocalVariableDeclarators declarators = new LocalVariableDeclarators(declarations.size());
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(568)         <b>for</b> (LocalVariableDeclarationStatement declaration : declarations) {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(569)             LocalVariableDeclarator declarator = (LocalVariableDeclarator)declaration.getLocalVariableDeclarators();
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(570)             if (setDimension)
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/Frame.class:(571)                 declarator.setDimension(declaration.getType().getDimension()); 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/LocalVariableSet.class:(135)     
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/LocalVariableSet.class:(136)     public AbstractLocalVariable[] initialize(Frame rootFrame) {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/LocalVariableSet.class:(137)         AbstractLocalVariable[] cache = new AbstractLocalVariable[this.array.length];
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/LocalVariableSet.class:(138)         <b>for</b> (int index = this.array.length - 1; index >= 0; index--) {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/LocalVariableSet.class:(139)             AbstractLocalVariable lv = this.array[index];
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/LocalVariableSet.class:(140)             if (lv != null) {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/LocalVariableSet.class:(141)                 AbstractLocalVariable previous = null;
--
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/RootFrame.class:(15)     
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/RootFrame.class:(16)     public void createDeclarations() {
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/RootFrame.class:(17)         if (this.children != null)
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/RootFrame.class:(18)             <b>for</b> (Frame child : this.children)
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/RootFrame.class:(19)                 child.createDeclarations();  
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/RootFrame.class:(20)     }
org/jd/core/v1/service/converter/classfiletojavasyntax/model/localvariable/RootFrame.class:(21) }
--
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(144)         if (fields == null)
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(145)             return null; 
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(146)         DefaultList<ClassFileFieldDeclaration> list = new DefaultList<>(fields.length);
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(147)         <b>for</b> (Field field : fields) {
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(148)             BaseAnnotationReference annotationReferences = convertAnnotationReferences(converter, field);
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(149)             Type typeField = parser.parseFieldSignature(classFile, field);
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(150)             ExpressionVariableInitializer variableInitializer = convertFieldInitializer(field, typeField);
--
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(159)         if (methods == null)
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(160)             return null; 
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(161)         DefaultList<ClassFileConstructorOrMethodDeclaration> list = new DefaultList<>(methods.length);
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(162)         <b>for</b> (Method method : methods) {
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(163)             Map<String, TypeArgument> bindings;
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(164)             Map<String, BaseType> typeBounds;
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(165)             String name = method.getName();
--
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(207)         if (innerClassFiles == null)
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(208)             return null; 
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(209)         DefaultList<ClassFileTypeDeclaration> list = new DefaultList<>(innerClassFiles.size());
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(210)         <b>for</b> (ClassFile innerClassFile : innerClassFiles) {
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(211)             ClassFileTypeDeclaration innerTypeDeclaration;
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(212)             int flags = innerClassFile.getAccessFlags();
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(213)             if ((flags & 0x4000) != 0) {
--
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(284)         if (moduleInfos == null || moduleInfos.length == 0)
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(285)             return null; 
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(286)         DefaultList<ModuleDeclaration.ModuleInfo> list = new DefaultList<>(moduleInfos.length);
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(287)         <b>for</b> (ModuleInfo moduleInfo : moduleInfos)
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(288)             list.add(new ModuleDeclaration.ModuleInfo(moduleInfo.getName(), moduleInfo.getFlags(), moduleInfo.getVersion())); 
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(289)         return list;
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(290)     }
--
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(293)         if (packageInfos == null || packageInfos.length == 0)
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(294)             return null; 
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(295)         DefaultList<ModuleDeclaration.PackageInfo> list = new DefaultList<>(packageInfos.length);
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(296)         <b>for</b> (PackageInfo packageInfo : packageInfos) {
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(297)             DefaultList<String> moduleInfoNames = (packageInfo.getModuleInfoNames() == null) ? null : new DefaultList<>(packageInfo.getModuleInfoNames());
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(298)             list.add(new ModuleDeclaration.PackageInfo(packageInfo.getInternalName(), packageInfo.getFlags(), moduleInfoNames));
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(299)         } 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(304)         if (serviceInfos == null || serviceInfos.length == 0)
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(305)             return null; 
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(306)         DefaultList<ModuleDeclaration.ServiceInfo> list = new DefaultList<>(serviceInfos.length);
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(307)         <b>for</b> (ServiceInfo serviceInfo : serviceInfos) {
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(308)             DefaultList<String> implementationTypeNames = (serviceInfo.getImplementationTypeNames() == null) ? null : new DefaultList<>(serviceInfo.getImplementationTypeNames());
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(309)             list.add(new ModuleDeclaration.ServiceInfo(serviceInfo.getInterfaceTypeName(), implementationTypeNames));
org/jd/core/v1/service/converter/classfiletojavasyntax/processor/ConvertClassFileProcessor.class:(310)         } 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AggregateFieldsUtil.class:(11)             if (size > 1) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AggregateFieldsUtil.class:(12)                 int firstIndex = 0, lastIndex = 0;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AggregateFieldsUtil.class:(13)                 ClassFileFieldDeclaration firstField = fields.get(0);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AggregateFieldsUtil.class:(14)                 <b>for</b> (int index = 1; index < size; index++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AggregateFieldsUtil.class:(15)                     ClassFileFieldDeclaration field = fields.get(index);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AggregateFieldsUtil.class:(16)                     if (firstField.getFirstLineNumber() == 0 || firstField.getFlags() != field.getFlags() || !firstField.getType().equals(field.getType())) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AggregateFieldsUtil.class:(17)                         firstField = field;
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AggregateFieldsUtil.class:(47)             } else {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AggregateFieldsUtil.class:(48)                 declarators.add(bfd.getFirst());
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AggregateFieldsUtil.class:(49)             } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AggregateFieldsUtil.class:(50)             <b>for</b> (ClassFileFieldDeclaration f : sublist) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AggregateFieldsUtil.class:(51)                 bfd = f.getFieldDeclarators();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AggregateFieldsUtil.class:(52)                 if (bfd.isList()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AggregateFieldsUtil.class:(53)                     declarators.addAll(bfd.getList());
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AnnotationConverter.class:(54)         if (invisibles == null)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AnnotationConverter.class:(55)             return convert(visibles); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AnnotationConverter.class:(56)         AnnotationReferences<AnnotationReference> aral = new AnnotationReferences<>();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AnnotationConverter.class:(57)         <b>for</b> (Annotation a : visibles.getAnnotations())
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AnnotationConverter.class:(58)             aral.add(convert(a)); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AnnotationConverter.class:(59)         <b>for</b> (Annotation a : invisibles.getAnnotations())
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AnnotationConverter.class:(60)             aral.add(convert(a)); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AnnotationConverter.class:(61)         return aral;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AnnotationConverter.class:(62)     }
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AnnotationConverter.class:(66)         if (as.length == 1)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AnnotationConverter.class:(67)             return convert(as[0]); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AnnotationConverter.class:(68)         AnnotationReferences<AnnotationReference> aral = new AnnotationReferences<>(as.length);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AnnotationConverter.class:(69)         <b>for</b> (Annotation a : as)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AnnotationConverter.class:(70)             aral.add(convert(a)); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AnnotationConverter.class:(71)         return aral;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AnnotationConverter.class:(72)     }
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AnnotationConverter.class:(89)                         convert(elementValue)));
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AnnotationConverter.class:(90)         } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AnnotationConverter.class:(91)         ElementValuePairs list = new ElementValuePairs(elementValuePairs.length);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AnnotationConverter.class:(92)         <b>for</b> (ElementValuePair elementValuePair : elementValuePairs) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AnnotationConverter.class:(93)             String elementName = elementValuePair.getElementName();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AnnotationConverter.class:(94)             ElementValue elementValue = elementValuePair.getElementValue();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AnnotationConverter.class:(95)             list.add(new ElementValuePair(elementName, convert(elementValue)));
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AnnotationConverter.class:(164)             this.elementValue = new ElementValueArrayInitializerElementValue(this.elementValue);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AnnotationConverter.class:(165)         } else {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AnnotationConverter.class:(166)             ElementValues list = new ElementValues(values.length);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AnnotationConverter.class:(167)             <b>for</b> (ElementValue value : values) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AnnotationConverter.class:(168)                 value.accept(this);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AnnotationConverter.class:(169)                 list.add(this.elementValue);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/AnnotationConverter.class:(170)             } 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(130)         ConstantPool constants = method.getConstants();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(131)         byte[] code = ((AttributeCode)method.<AttributeCode>getAttribute("Code")).getCode();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(132)         boolean syntheticFlag = ((method.getAccessFlags() & 0x1000) != 0);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(133)         <b>for</b> (int offset = fromOffset; offset < toOffset; offset++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(134)             Expression indexRef, arrayRef, valueRef, expression1, expression2, expression3;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(135)             Type type1, type2, type3;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(136)             ConstantMemberRef constantMemberRef;
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1007)         } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1008)         Expressions parameters = new Expressions(parameterTypes.size());
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1009)         int count = parameterTypes.size() - 1;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1010)         <b>for</b> (int i = count; i >= 0; i--) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1011)             parameter = stack.pop();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1012)             if (parameter.getClass() == NewArray.class)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1013)                 parameter = NewArrayMaker.make(statements, (NewArray)parameter); 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1423)         String name1 = constants.getConstantUtf8(cnat1.getNameIndex());
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1424)         String descriptor1 = constants.getConstantUtf8(cnat1.getDescriptorIndex());
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1425)         if (typeName.equals(this.internalTypeName))
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1426)             <b>for</b> (ClassFileConstructorOrMethodDeclaration methodDeclaration : this.bodyDeclaration.getMethodDeclarations()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1427)                 if ((methodDeclaration.getFlags() & 0x1002) == 4098 && methodDeclaration.getMethod().getName().equals(name1) && methodDeclaration.getMethod().getDescriptor().equals(descriptor1)) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1428)                     ClassFileMethodDeclaration cfmd = (ClassFileMethodDeclaration)methodDeclaration;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1429)                     stack.push(new LambdaIdentifiersExpression(lineNumber, indyMethodTypes.returnedType, indyMethodTypes.returnedType, 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1445)         stack.push(new MethodReferenceExpression(lineNumber, indyMethodTypes.returnedType, (Expression)indyParameters, typeName, name1, descriptor1));
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1446)     }
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1447)     
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1448)     private static List<String> prepareLambdaParameters(BaseFormalParameter <b>for</b>malParameters, int parameterCount) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1449)         if (<b>for</b>malParameters == null || parameterCount == 0)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1450)             return null; 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1451)         LambdaParameterNamesVisitor lambdaParameterNamesVisitor = new LambdaParameterNamesVisitor(null);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1452)         <b>for</b>malParameters.accept(lambdaParameterNamesVisitor);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1453)         List<String> names = lambdaParameterNamesVisitor.getNames();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1454)         assert names.size() >= parameterCount;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1455)         if (names.size() == parameterCount)
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1902)             String internalName = ((ObjectType)expression.getType()).getInternalName();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1903)             if (!ot.getInternalName().equals(internalName)) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1904)                 this.memberVisitor.init(name, null);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1905)                 <b>for</b> (ClassFileFieldDeclaration field : this.bodyDeclaration.getFieldDeclarations()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1906)                     field.getFieldDeclarators().accept(this.memberVisitor);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1907)                     if (this.memberVisitor.found())
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1908)                         return new SuperExpression(expression.getLineNumber(), expression.getType()); 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1917)             String internalName = ((ObjectType)expression.getType()).getInternalName();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1918)             if (!ot.getInternalName().equals(internalName)) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1919)                 this.memberVisitor.init(name, descriptor);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1920)                 <b>for</b> (ClassFileConstructorOrMethodDeclaration member : this.bodyDeclaration.getMethodDeclarations()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1921)                     member.accept(this.memberVisitor);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1922)                     if (this.memberVisitor.found())
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1923)                         return new SuperExpression(expression.getLineNumber(), expression.getType()); 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1995)         int toOffset = basicBlock.getToOffset();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1996)         if (toOffset > maxOffset)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1997)             toOffset = maxOffset; 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1998)         <b>for</b> (; offset < toOffset; offset++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(1999)             int deltaOffset, low, high, count, opcode = code[offset] & 0xFF;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(2000)             switch (opcode) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(2001)                 case 16:
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(2105)         if (offset >= toOffset)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(2106)             return 0; 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(2107)         int lastOffset = offset;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(2108)         <b>for</b> (; offset < toOffset; offset++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(2109)             int deltaOffset, low, high, count, opcode = code[offset] & 0xFF;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(2110)             lastOffset = offset;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(2111)             switch (opcode) {
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(2217)     
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(2218)     public static int evalStackDepth(ConstantPool constants, byte[] code, BasicBlock bb) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(2219)         int depth = 0;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(2220)         <b>for</b> (int offset = bb.getFromOffset(), toOffset = bb.getToOffset(); offset < toOffset; offset++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(2221)             ConstantMemberRef constantMemberRef;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(2222)             ConstantNameAndType constantNameAndType;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeParser.class:(2223)             String descriptor;
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(53)     }
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(54)     
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(55)     protected static void writeByteCode(String linePrefix, StringBuilder sb, ConstantPool constants, byte[] code, int fromOffset, int toOffset) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(56)         <b>for</b> (int offset = fromOffset; offset < toOffset; offset++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(57)             int i, low, high, value, npairs, j;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(58)             ConstantMemberRef constantMemberRef;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(59)             int k;
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(115)                     sb.append(" default").append(" -> ").append(offset + ((code[i++] & 0xFF) << 24 | (code[i++] & 0xFF) << 16 | (code[i++] & 0xFF) << 8 | code[i++] & 0xFF));
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(116)                     low = (code[i++] & 0xFF) << 24 | (code[i++] & 0xFF) << 16 | (code[i++] & 0xFF) << 8 | code[i++] & 0xFF;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(117)                     high = (code[i++] & 0xFF) << 24 | (code[i++] & 0xFF) << 16 | (code[i++] & 0xFF) << 8 | code[i++] & 0xFF;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(118)                     <b>for</b> (value = low; value <= high; value++)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(119)                         sb.append(", ").append(value).append(" -> ").append(offset + ((code[i++] & 0xFF) << 24 | (code[i++] & 0xFF) << 16 | (code[i++] & 0xFF) << 8 | code[i++] & 0xFF)); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(120)                     offset = i - 1;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(121)                     break;
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(123)                     i = offset + 4 & 0xFFFC;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(124)                     sb.append(" default").append(" -> ").append(offset + ((code[i++] & 0xFF) << 24 | (code[i++] & 0xFF) << 16 | (code[i++] & 0xFF) << 8 | code[i++] & 0xFF));
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(125)                     npairs = (code[i++] & 0xFF) << 24 | (code[i++] & 0xFF) << 16 | (code[i++] & 0xFF) << 8 | code[i++] & 0xFF;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(126)                     <b>for</b> (j = 0, k = 0; k < npairs; k++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(127)                         sb.append(", ").append((code[i++] & 0xFF) << 24 | (code[i++] & 0xFF) << 16 | (code[i++] & 0xFF) << 8 | code[i++] & 0xFF);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(128)                         sb.append(" -> ").append(offset + ((code[i++] & 0xFF) << 24 | (code[i++] & 0xFF) << 16 | (code[i++] & 0xFF) << 8 | code[i++] & 0xFF));
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(129)                     } 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(278)                 sb.append(" '");
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(279)                 stringIndex = ((ConstantString)constant).getStringIndex();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(280)                 str = constants.getConstantUtf8(stringIndex);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(281)                 <b>for</b> (char c : str.toCharArray()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(282)                     switch (c) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(283)                         case '\b':
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(284)                             sb.append("\\\\b");
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(310)         if (lineNumberTable != null) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(311)             sb.append(linePrefix).append("Line number table:\n");
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(312)             sb.append(linePrefix).append("  Java source line number -> byte code offset\n");
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(313)             <b>for</b> (LineNumber lineNumber : lineNumberTable.getLineNumberTable()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(314)                 sb.append(linePrefix).append("  #");
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(315)                 sb.append(lineNumber.getLineNumber()).append("\t-> ");
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(316)                 sb.append(lineNumber.getStartPc()).append('\n');
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(323)         if (localVariableTable != null) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(324)             sb.append(linePrefix).append("Local variable table:\n");
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(325)             sb.append(linePrefix).append("  start\tlength\tslot\tname\tdescriptor\n");
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(326)             <b>for</b> (LocalVariable localVariable : localVariableTable.getLocalVariableTable()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(327)                 sb.append(linePrefix).append("  ");
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(328)                 sb.append(localVariable.getStartPc()).append('\t');
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(329)                 sb.append(localVariable.getLength()).append('\t');
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(336)         if (localVariableTypeTable != null) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(337)             sb.append(linePrefix).append("Local variable type table:\n");
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(338)             sb.append(linePrefix).append("  start\tlength\tslot\tname\tsignature\n");
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(339)             <b>for</b> (LocalVariableType localVariable : localVariableTypeTable.getLocalVariableTypeTable()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(340)                 sb.append(linePrefix).append("  ");
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(341)                 sb.append(localVariable.getStartPc()).append('\t');
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(342)                 sb.append(localVariable.getLength()).append('\t');
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(352)         if (codeExceptions != null) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(353)             sb.append(linePrefix).append("Exception table:\n");
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(354)             sb.append(linePrefix).append("  from\tto\ttarget\ttype\n");
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(355)             <b>for</b> (CodeException codeException : codeExceptions) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(356)                 sb.append(linePrefix).append("  ");
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(357)                 sb.append(codeException.getStartPc()).append('\t');
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ByteCodeWriter.class:(358)                 sb.append(codeException.getEndPc()).append('\t');
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphGotoReducer.class:(4) 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphGotoReducer.class:(5) public class ControlFlowGraphGotoReducer {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphGotoReducer.class:(6)     public static void reduce(ControlFlowGraph cfg) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphGotoReducer.class:(7)         <b>for</b> (BasicBlock basicBlock : cfg.getBasicBlocks()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphGotoReducer.class:(8)             if (basicBlock.getType() == 67108864) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphGotoReducer.class:(9)                 BasicBlock successor = basicBlock.getNext();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphGotoReducer.class:(10)                 if (basicBlock == successor) {
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphGotoReducer.class:(14)                 } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphGotoReducer.class:(15)                 Set<BasicBlock> successorPredecessors = successor.getPredecessors();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphGotoReducer.class:(16)                 successorPredecessors.remove(basicBlock);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphGotoReducer.class:(17)                 <b>for</b> (BasicBlock predecessor : basicBlock.getPredecessors()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphGotoReducer.class:(18)                     predecessor.replace(basicBlock, successor);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphGotoReducer.class:(19)                     successorPredecessors.add(predecessor);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphGotoReducer.class:(20)                 } 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(21)         initial.set(0);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(22)         arrayOfDominatorIndexes[0] = initial;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(23)         int i;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(24)         <b>for</b> (i = 0; i < length; i++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(25)             initial = new BitSet(length);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(26)             initial.flip(0, length);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(27)             arrayOfDominatorIndexes[i] = initial;
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(31)         initial.set(0);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(32)         do {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(33)             boolean change = false;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(34)             <b>for</b> (BasicBlock basicBlock : list) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(35)                 int index = basicBlock.getIndex();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(36)                 BitSet dominatorIndexes = arrayOfDominatorIndexes[index];
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(37)                 initial = (BitSet)dominatorIndexes.clone();
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(38)                 <b>for</b> (BasicBlock predecessorBB : basicBlock.getPredecessors())
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(39)                     dominatorIndexes.and(arrayOfDominatorIndexes[predecessorBB.getIndex()]); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(40)                 dominatorIndexes.set(index);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(41)                 i = change | (!initial.equals(dominatorIndexes) ? 1 : 0);
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(49)         int length = list.size();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(50)         BitSet[] arrayOfMemberIndexes = new BitSet[length];
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(51)         int i;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(52)         <b>for</b> (i = 0; i < length; i++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(53)             int index;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(54)             BasicBlock current = list.get(i);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(55)             BitSet dominatorIndexes = arrayOfDominatorIndexes[i];
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(65)                         arrayOfMemberIndexes[index] = searchLoopMemberIndexes(length, arrayOfMemberIndexes[index], current, current.getNext()); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(66)                     break;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(67)                 case 64:
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(68)                     <b>for</b> (BasicBlock.SwitchCase switchCase : current.getSwitchCases()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(69)                         index = switchCase.getBasicBlock().getIndex();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(70)                         if (index >= 0 && dominatorIndexes.get(index))
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(71)                             arrayOfMemberIndexes[index] = searchLoopMemberIndexes(length, arrayOfMemberIndexes[index], current, switchCase.getBasicBlock()); 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(73)                     break;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(74)             } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(75)         } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(76)         <b>for</b> (i = 0; i < length; i++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(77)             if (arrayOfMemberIndexes[i] != null) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(78)                 BitSet memberIndexes = arrayOfMemberIndexes[i];
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(79)                 int maxOffset = -1;
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(80)                 <b>for</b> (int k = 0; k < length; k++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(81)                     if (memberIndexes.get(k)) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(82)                         int offset = ((BasicBlock)list.get(k)).getFromOffset();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(83)                         if (maxOffset < offset)
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(105)             } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(106)         } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(107)         DefaultList<Loop> loops = new DefaultList<>();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(108)         <b>for</b> (int j = 0; j < length; j++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(109)             if (arrayOfMemberIndexes[j] != null) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(110)                 BitSet memberIndexes = arrayOfMemberIndexes[j];
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(111)                 BasicBlock start = list.get(j);
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(161)         if (!visited.get(current.getIndex())) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(162)             visited.set(current.getIndex());
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(163)             if (current != start)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(164)                 <b>for</b> (BasicBlock predecessor : current.getPredecessors())
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(165)                     recursiveBackwardSearchLoopMemberIndexes(visited, predecessor, start);  
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(166)         } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(167)     }
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(169)     protected static Loop makeLoop(List<BasicBlock> list, BasicBlock start, BitSet searchZoneIndexes, BitSet memberIndexes) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(170)         int length = list.size();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(171)         int maxOffset = -1;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(172)         <b>for</b> (int i = 0; i < length; i++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(173)             if (memberIndexes.get(i)) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(174)                 int offset = checkMaxOffset(list.get(i));
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(175)                 if (maxOffset < offset)
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(179)         memberIndexes.clear();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(180)         recursiveForwardSearchLoopMemberIndexes(memberIndexes, searchZoneIndexes, start, maxOffset);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(181)         HashSet<BasicBlock> members = new HashSet<>(memberIndexes.cardinality());
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(182)         <b>for</b> (int j = 0; j < length; j++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(183)             if (memberIndexes.get(j))
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(184)                 members.add(list.get(j)); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(185)         } 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(202)                 HashSet<BasicBlock> set = new HashSet<>();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(203)                 if (recursiveForwardSearchLastLoopMemberIndexes(members, searchZoneIndexes, set, end, null)) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(204)                     members.addAll(set);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(205)                     <b>for</b> (BasicBlock member : set) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(206)                         if (member.getIndex() >= 0)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(207)                             memberIndexes.set(member.getIndex()); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(208)                     } 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(213)         if (end != BasicBlock.END) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(214)             HashSet<BasicBlock> m = new HashSet<>(members);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(215)             HashSet<BasicBlock> set = new HashSet<>();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(216)             <b>for</b> (BasicBlock member : m) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(217)                 if (member.getType() == 32768 && member != start) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(218)                     set.clear();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(219)                     if (recursiveForwardSearchLastLoopMemberIndexes(members, searchZoneIndexes, set, member.getNext(), end))
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(229)     
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(230)     private static BasicBlock searchEndBasicBlock(BitSet memberIndexes, int maxOffset, Set<BasicBlock> members) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(231)         BasicBlock end = BasicBlock.END;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(232)         <b>for</b> (BasicBlock member : members) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(233)             BasicBlock bb;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(234)             switch (member.getType()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(235)                 case 32768:
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(246)                         maxOffset = bb.getFromOffset();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(247)                     } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(248)                 case 64:
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(249)                     <b>for</b> (BasicBlock.SwitchCase switchCase : member.getSwitchCases()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(250)                         bb = switchCase.getBasicBlock();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(251)                         if (!memberIndexes.get(bb.getIndex()) && maxOffset < bb.getFromOffset()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(252)                             end = bb;
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(259)                         end = bb;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(260)                         maxOffset = bb.getFromOffset();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(261)                     } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(262)                     <b>for</b> (BasicBlock.ExceptionHandler exceptionHandler : member.getExceptionHandlers()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(263)                         bb = exceptionHandler.getBasicBlock();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(264)                         if (!memberIndexes.get(bb.getIndex()) && maxOffset < bb.getFromOffset()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(265)                             end = bb;
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(274)     private static int checkMaxOffset(BasicBlock basicBlock) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(275)         int maxOffset = basicBlock.getFromOffset();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(276)         if (basicBlock.getType() == 512) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(277)             <b>for</b> (BasicBlock.ExceptionHandler exceptionHandler : basicBlock.getExceptionHandlers()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(278)                 int offset;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(279)                 if (exceptionHandler.getInternalThrowableName() == null) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(280)                     offset = checkThrowBlockOffset(exceptionHandler.getBasicBlock());
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(287)         } else if (basicBlock.getType() == 64) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(288)             BasicBlock lastBB = null;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(289)             BasicBlock previousBB = null;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(290)             <b>for</b> (BasicBlock.SwitchCase switchCase : basicBlock.getSwitchCases()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(291)                 BasicBlock bb = switchCase.getBasicBlock();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(292)                 if (lastBB == null || lastBB.getFromOffset() < bb.getFromOffset()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(293)                     previousBB = lastBB;
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(327)             if (current != target) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(328)                 recursiveForwardSearchLoopMemberIndexes(visited, searchZoneIndexes, current.getNext(), target);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(329)                 recursiveForwardSearchLoopMemberIndexes(visited, searchZoneIndexes, current.getBranch(), target);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(330)                 <b>for</b> (BasicBlock.SwitchCase switchCase : current.getSwitchCases())
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(331)                     recursiveForwardSearchLoopMemberIndexes(visited, searchZoneIndexes, switchCase.getBasicBlock(), target); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(332)                 <b>for</b> (BasicBlock.ExceptionHandler exceptionHandler : current.getExceptionHandlers())
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(333)                     recursiveForwardSearchLoopMemberIndexes(visited, searchZoneIndexes, exceptionHandler.getBasicBlock(), target); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(334)                 if (current.getType() == 268435456)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(335)                     visited.set(current.getNext().getIndex()); 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(345)             visited.set(current.getIndex());
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(346)             recursiveForwardSearchLoopMemberIndexes(visited, searchZoneIndexes, current.getNext(), maxOffset);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(347)             recursiveForwardSearchLoopMemberIndexes(visited, searchZoneIndexes, current.getBranch(), maxOffset);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(348)             <b>for</b> (BasicBlock.SwitchCase switchCase : current.getSwitchCases())
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(349)                 recursiveForwardSearchLoopMemberIndexes(visited, searchZoneIndexes, switchCase.getBasicBlock(), maxOffset); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(350)             <b>for</b> (BasicBlock.ExceptionHandler exceptionHandler : current.getExceptionHandlers())
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(351)                 recursiveForwardSearchLoopMemberIndexes(visited, searchZoneIndexes, exceptionHandler.getBasicBlock(), maxOffset); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(352)             if (current.getType() == 268435456)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(353)                 visited.set(current.getNext().getIndex()); 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(388)     
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(389)     protected static boolean predecessorsInSearchZone(BasicBlock basicBlock, BitSet searchZoneIndexes) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(390)         Set<BasicBlock> predecessors = basicBlock.getPredecessors();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(391)         <b>for</b> (BasicBlock predecessor : predecessors) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(392)             if (!inSearchZone(predecessor, searchZoneIndexes))
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(393)                 return false; 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(394)         } 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(401)     
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(402)     protected static BasicBlock recheckEndBlock(Set<BasicBlock> members, BasicBlock end) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(403)         boolean flag = false;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(404)         <b>for</b> (BasicBlock predecessor : end.getPredecessors()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(405)             if (!members.contains(predecessor)) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(406)                 flag = true;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(407)                 break;
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(409)         } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(410)         if (!flag) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(411)             BasicBlock newEnd = null;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(412)             <b>for</b> (BasicBlock member : members) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(413)                 if (member.matchType(876822149)) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(414)                     BasicBlock bb = member.getNext();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(415)                     if (bb != end && !members.contains(bb)) {
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(460)             } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(461)         } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(462)         loopBB.setSub1(start);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(463)         <b>for</b> (BasicBlock member : members) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(464)             if (member.matchType(876822149)) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(465)                 BasicBlock bb = member.getNext();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(466)                 if (bb == start) {
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(488)                     member.setBranch(newJumpBasicBlock(member, bb));
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(489)                 } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(490)             } else if (member.getType() == 64) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(491)                 <b>for</b> (BasicBlock.SwitchCase switchCase : member.getSwitchCases()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(492)                     BasicBlock bb = switchCase.getBasicBlock();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(493)                     if (bb == start) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(494)                         switchCase.setBasicBlock(BasicBlock.LOOP_START);
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(524)     public static void reduce(ControlFlowGraph cfg) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(525)         BitSet[] arrayOfDominatorIndexes = buildDominatorIndexes(cfg);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(526)         List<Loop> loops = identifyNaturalLoops(cfg, arrayOfDominatorIndexes);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(527)         <b>for</b> (int i = 0, loopsLength = loops.size(); i < loopsLength; i++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(528)             Loop loop = loops.get(i);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(529)             BasicBlock startBB = loop.getStart();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(530)             BasicBlock loopBB = reduceLoop(loop);
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(531)             <b>for</b> (int j = loopsLength - 1; j > i; j--) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(532)                 Loop otherLoop = loops.get(j);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(533)                 if (otherLoop.getStart() == startBB)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphLoopReducer.class:(534)                     otherLoop.setStart(loopBB); 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(38)         map[0] = MARK;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(39)         int lastOffset = 0;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(40)         int lastStatementOffset = -1;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(41)         <b>for</b> (int offset = 0; offset < length; offset++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(42)             ConstantMemberRef constantMemberRef;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(43)             ConstantNameAndType constantNameAndType;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(44)             String descriptor;
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(205)                     values = new int[high - low + 2];
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(206)                     offsets = new int[high - low + 2];
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(207)                     offsets[0] = defaultOffset;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(208)                     <b>for</b> (m = 1, len = high - low + 2; m < len; m++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(209)                         values[m] = low + m - 1;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(210)                         branchOffset = offsets[m] = offset + ((code[k++] & 0xFF) << 24 | (code[k++] & 0xFF) << 16 | (code[k++] & 0xFF) << 8 | code[k++] & 0xFF);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(211)                         map[branchOffset] = MARK;
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(224)                     values = new int[npairs + 1];
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(225)                     offsets = new int[npairs + 1];
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(226)                     offsets[0] = defaultOffset;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(227)                     <b>for</b> (n = 1; n <= npairs; n++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(228)                         values[n] = (code[k++] & 0xFF) << 24 | (code[k++] & 0xFF) << 16 | (code[k++] & 0xFF) << 8 | code[k++] & 0xFF;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(229)                         branchOffset = offsets[n] = offset + ((code[k++] & 0xFF) << 24 | (code[k++] & 0xFF) << 16 | (code[k++] & 0xFF) << 8 | code[k++] & 0xFF);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(230)                         map[branchOffset] = MARK;
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(315)         nextOffsets[lastOffset] = length;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(316)         CodeException[] codeExceptions = attributeCode.getExceptionTable();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(317)         if (codeExceptions != null)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(318)             <b>for</b> (CodeException codeException : codeExceptions) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(319)                 map[codeException.getStartPc()] = MARK;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(320)                 map[codeException.getHandlerPc()] = MARK;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(321)             }  
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(326)             int[] offsetToLineNumbers = new int[length];
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(327)             int k = 0;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(328)             int lineNumber = lineNumberTable[0].getLineNumber();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(329)             <b>for</b> (int m = 1, len = lineNumberTable.length; m < len; m++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(330)                 LineNumber lineNumberEntry = lineNumberTable[m];
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(331)                 int toIndex = lineNumberEntry.getStartPc();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(332)                 <b>for</b> (; k < toIndex; offsetToLineNumbers[k++] = lineNumber);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(333)                 if (lineNumber > lineNumberEntry.getLineNumber())
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(334)                     map[k] = MARK; 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(335)                 lineNumber = lineNumberEntry.getLineNumber();
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(336)             } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(337)             <b>for</b> (; k < length; offsetToLineNumbers[k++] = lineNumber);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(338)             cfg.setOffsetToLineNumbers(offsetToLineNumbers);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(339)         } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(340)         lastOffset = 0;
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(341)         BasicBlock startBasicBlock = cfg.newBasicBlock(1, 0, 0);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(342)         int j;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(343)         <b>for</b> (j = nextOffsets[0]; j < length; j = nextOffsets[j]) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(344)             if (map[j] != null) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(345)                 map[lastOffset] = cfg.newBasicBlock(lastOffset, j);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(346)                 lastOffset = j;
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(352)         BasicBlock successor = list.get(1);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(353)         startBasicBlock.setNext(successor);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(354)         successor.getPredecessors().add(startBasicBlock);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(355)         <b>for</b> (int i = 1, basicBlockLength = list.size(); i < basicBlockLength; i++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(356)             int[] values, offsets;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(357)             DefaultList<BasicBlock.SwitchCase> switchCases;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(358)             int defaultOffset;
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(399)                     bb = map[defaultOffset];
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(400)                     switchCases.add(new BasicBlock.SwitchCase(bb));
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(401)                     bb.getPredecessors().add(basicBlock);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(402)                     <b>for</b> (k = 1, len = offsets.length; k < len; k++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(403)                         int m = offsets[k];
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(404)                         if (m != defaultOffset) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(405)                             bb = map[m];
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(441)             int[] handlePcToStartPc = branchOffsets;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(442)             char[] handlePcMarks = types;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(443)             Arrays.sort(codeExceptions, CODE_EXCEPTION_COMPARATOR);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(444)             <b>for</b> (CodeException codeException : codeExceptions) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(445)                 int startPc = codeException.getStartPc();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(446)                 int handlerPc = codeException.getHandlerPc();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(447)                 if (startPc != handlerPc && (
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(477)                 } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(478)             } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(479)         } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(480)         <b>for</b> (BasicBlock bb : basicBlocks) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(481)             BasicBlock next = bb.getNext();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(482)             if (bb.getType() == 4 && next.getPredecessors().size() == 1) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphMaker.class:(483)                 if (next.getType() == 67108864 && ByteCodeParser.evalStackDepth(constants, code, bb) > 0) {
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(402)         BasicBlock.SwitchCase defaultSC = null;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(403)         BasicBlock.SwitchCase lastSC = null;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(404)         int maxOffset = -1;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(405)         <b>for</b> (BasicBlock.SwitchCase switchCase : basicBlock.getSwitchCases()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(406)             if (maxOffset < switchCase.getOffset())
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(407)                 maxOffset = switchCase.getOffset(); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(408)             if (switchCase.isDefaultCase()) {
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(416)         BasicBlock lastSwitchCaseBasicBlock = null;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(417)         BitSet v = new BitSet();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(418)         HashSet<BasicBlock> ends = new HashSet<>();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(419)         <b>for</b> (BasicBlock.SwitchCase switchCase : basicBlock.getSwitchCases()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(420)             BasicBlock bb = switchCase.getBasicBlock();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(421)             if (switchCase.getOffset() == maxOffset) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(422)                 lastSwitchCaseBasicBlock = bb;
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(425)             visit(v, bb, maxOffset, ends);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(426)         } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(427)         BasicBlock end = BasicBlock.END;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(428)         <b>for</b> (BasicBlock bb : ends) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(429)             if (!bb.matchType(1266696506) && (
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(430)                 end == BasicBlock.END || end.getFromOffset() < bb.getFromOffset()))
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(431)                 end = bb; 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(456)                     iterator.remove(); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(457)             } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(458)         } else {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(459)             <b>for</b> (BasicBlock.SwitchCase switchCase : basicBlock.getSwitchCases()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(460)                 if (switchCase.getBasicBlock() == end)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(461)                     switchCase.setBasicBlock(BasicBlock.SWITCH_BREAK); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(462)             } 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(463)         } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(464)         boolean reduced = true;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(465)         <b>for</b> (BasicBlock.SwitchCase switchCase : basicBlock.getSwitchCases())
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(466)             reduced &= reduce(visited, switchCase.getBasicBlock(), jsrTargets); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(467)         <b>for</b> (BasicBlock.SwitchCase switchCase : basicBlock.getSwitchCases()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(468)             BasicBlock bb = switchCase.getBasicBlock();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(469)             assert bb != end;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(470)             Set<BasicBlock> predecessors = bb.getPredecessors();
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(487)     
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(488)     protected static boolean searchLoopStart(BasicBlock basicBlock, int maxOffset) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(489)         WatchDog watchdog = new WatchDog();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(490)         <b>for</b> (BasicBlock.SwitchCase switchCase : basicBlock.getSwitchCases()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(491)             BasicBlock bb = switchCase.getBasicBlock();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(492)             watchdog.clear();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(493)             while (bb.getFromOffset() < maxOffset) {
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(502)                     next = bb.getBranch();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(503)                 } else if (bb.getType() == 64) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(504)                     int max = bb.getFromOffset();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(505)                     <b>for</b> (BasicBlock.SwitchCase sc : bb.getSwitchCases()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(506)                         if (max < sc.getBasicBlock().getFromOffset()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(507)                             next = sc.getBasicBlock();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(508)                             max = next.getFromOffset();
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(521)     protected static boolean reduceTryDeclaration(BitSet visited, BasicBlock basicBlock, BitSet jsrTargets) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(522)         boolean reduced = true;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(523)         BasicBlock finallyBB = null;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(524)         <b>for</b> (BasicBlock.ExceptionHandler exceptionHandler : basicBlock.getExceptionHandlers()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(525)             if (exceptionHandler.getInternalThrowableName() == null) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(526)                 reduced = reduce(visited, exceptionHandler.getBasicBlock(), jsrTargets);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(527)                 finallyBB = exceptionHandler.getBasicBlock();
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(536)         int maxOffset = basicBlock.getFromOffset();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(537)         boolean tryWithResourcesFlag = true;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(538)         BasicBlock tryWithResourcesBB = null;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(539)         <b>for</b> (BasicBlock.ExceptionHandler exceptionHandler : basicBlock.getExceptionHandlers()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(540)             if (exceptionHandler.getInternalThrowableName() != null)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(541)                 reduced &= reduce(visited, exceptionHandler.getBasicBlock(), jsrTargets); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(542)             BasicBlock bb = exceptionHandler.getBasicBlock();
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(552)                 } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(553)                 assert predecessors.size() == 2;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(554)                 if (tryWithResourcesBB == null) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(555)                     label99: <b>for</b> (BasicBlock predecessor : predecessors) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(556)                         if (predecessor != basicBlock) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(557)                             assert false;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(558)                             tryWithResourcesBB = predecessor;
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(565)             } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(566)         } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(567)         if (tryWithResourcesFlag) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(568)             <b>for</b> (BasicBlock.ExceptionHandler exceptionHandler : basicBlock.getExceptionHandlers())
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(569)                 exceptionHandler.getBasicBlock().getPredecessors().remove(basicBlock); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(570)             <b>for</b> (BasicBlock predecessor : basicBlock.getPredecessors()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(571)                 predecessor.replace(basicBlock, tryBB);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(572)                 tryBB.replace(basicBlock, predecessor);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(573)             } 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(577)             updateBlock(tryBB, end, maxOffset);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(578)             if (finallyBB != null && basicBlock.getExceptionHandlers().size() == 1 && tryBB.getType() == 1024 && tryBB.getNext() == BasicBlock.END && basicBlock.getFromOffset() == tryBB.getFromOffset() && !containsFinally(tryBB)) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(579)                 basicBlock.getExceptionHandlers().addAll(0, tryBB.getExceptionHandlers());
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(580)                 <b>for</b> (BasicBlock.ExceptionHandler exceptionHandler : tryBB.getExceptionHandlers()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(581)                     Set<BasicBlock> set = exceptionHandler.getBasicBlock().getPredecessors();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(582)                     set.clear();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(583)                     set.add(basicBlock);
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(589)                 predecessors.add(basicBlock);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(590)             } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(591)             int toOffset = maxOffset;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(592)             <b>for</b> (BasicBlock.ExceptionHandler exceptionHandler : basicBlock.getExceptionHandlers()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(593)                 BasicBlock bb = exceptionHandler.getBasicBlock();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(594)                 if (bb == end) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(595)                     exceptionHandler.setBasicBlock(BasicBlock.END);
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(621)     }
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(622)     
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(623)     protected static boolean containsFinally(BasicBlock basicBlock) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(624)         <b>for</b> (BasicBlock.ExceptionHandler exceptionHandler : basicBlock.getExceptionHandlers()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(625)             if (exceptionHandler.getInternalThrowableName() == null)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(626)                 return true; 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(627)         } 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(644)     }
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(645)     
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(646)     protected static BasicBlock searchJsrTarget(BasicBlock basicBlock, BitSet jsrTargets) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(647)         <b>for</b> (BasicBlock.ExceptionHandler exceptionHandler : basicBlock.getExceptionHandlers()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(648)             if (exceptionHandler.getInternalThrowableName() == null) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(649)                 BasicBlock bb = exceptionHandler.getBasicBlock();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(650)                 if (bb.getType() == 4) {
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(669)                 return next; 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(670)             end = next;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(671)         } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(672)         <b>for</b> (BasicBlock.ExceptionHandler exceptionHandler : basicBlock.getExceptionHandlers()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(673)             BasicBlock bb = exceptionHandler.getBasicBlock();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(674)             if (bb.getFromOffset() < maxOffset) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(675)                 last = splitSequence(bb, maxOffset);
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(749)         if (basicBlock.getExceptionHandlers().size() == 1) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(750)             BasicBlock subTry = basicBlock.getSub1();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(751)             if (subTry.matchType(7168)) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(752)                 <b>for</b> (BasicBlock.ExceptionHandler exceptionHandler : subTry.getExceptionHandlers()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(753)                     if (exceptionHandler.getInternalThrowableName() == null)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(754)                         return; 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(755)                 } 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(756)                 <b>for</b> (BasicBlock.ExceptionHandler exceptionHandler : subTry.getExceptionHandlers()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(757)                     BasicBlock bb = exceptionHandler.getBasicBlock();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(758)                     basicBlock.addExceptionHandler(exceptionHandler.getInternalThrowableName(), bb);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(759)                     bb.replace(subTry, basicBlock);
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(788)             branch.getPredecessors().remove(basicBlock);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(789)             Set<BasicBlock> nextPredecessors = basicBlock.getNext().getPredecessors();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(790)             nextPredecessors.remove(basicBlock);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(791)             <b>for</b> (BasicBlock predecessor : basicBlock.getPredecessors()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(792)                 predecessor.replace(basicBlock, basicBlock.getNext());
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(793)                 nextPredecessors.add(predecessor);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(794)             } 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(800)             while (iterator.hasNext()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(801)                 BasicBlock predecessor = iterator.next();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(802)                 if (predecessor != basicBlock && predecessor.getType() == 8192 && predecessor.getNext() == next) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(803)                     <b>for</b> (BasicBlock predecessorPredecessor : predecessor.getPredecessors()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(804)                         predecessorPredecessor.replace(predecessor, basicBlock);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(805)                         basicBlock.getPredecessors().add(predecessorPredecessor);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(806)                     } 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(838)                     changeEndLoopToJump(visitedMembers, basicBlock.getNext(), basicBlock.getSub1());
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(839)                     BasicBlock newLoopBB = basicBlock.getControlFlowGraph().newBasicBlock(basicBlock);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(840)                     Set<BasicBlock> predecessors = conditionalBranch.getPredecessors();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(841)                     <b>for</b> (BasicBlock predecessor : predecessors)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(842)                         predecessor.replace(conditionalBranch, BasicBlock.LOOP_END); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(843)                     newLoopBB.setNext(conditionalBranch);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(844)                     predecessors.clear();
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(909)                     case 4096:
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(910)                         visit(visited, basicBlock.getSub1(), maxOffset, ends);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(911)                     case 512:
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(912)                         <b>for</b> (BasicBlock.ExceptionHandler exceptionHandler : basicBlock.getExceptionHandlers())
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(913)                             visit(visited, exceptionHandler.getBasicBlock(), maxOffset, ends); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(914)                         visit(visited, basicBlock.getNext(), maxOffset, ends);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(915)                         break;
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(928)                     case 128:
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(929)                         visit(visited, basicBlock.getNext(), maxOffset, ends);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(930)                     case 64:
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(931)                         <b>for</b> (BasicBlock.SwitchCase switchCase : basicBlock.getSwitchCases())
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(932)                             visit(visited, switchCase.getBasicBlock(), maxOffset, ends); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(933)                         break;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(934)                 } 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(956)                 case 4096:
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(957)                     replaceLoopStartWithSwitchBreak(visited, basicBlock.getSub1());
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(958)                 case 512:
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(959)                     <b>for</b> (BasicBlock.ExceptionHandler exceptionHandler : basicBlock.getExceptionHandlers())
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(960)                         replaceLoopStartWithSwitchBreak(visited, exceptionHandler.getBasicBlock()); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(961)                     break;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(962)                 case 131072:
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(974)                 case 128:
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(975)                     replaceLoopStartWithSwitchBreak(visited, basicBlock.getNext());
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(976)                 case 64:
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(977)                     <b>for</b> (BasicBlock.SwitchCase switchCase : basicBlock.getSwitchCases())
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(978)                         replaceLoopStartWithSwitchBreak(visited, switchCase.getBasicBlock()); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(979)                     break;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(980)             } 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(1004)                 case 4096:
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(1005)                     updateBasicBlock = searchUpdateBlockAndCreateContinueLoop(visited, basicBlock, basicBlock.getSub1());
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(1006)                 case 512:
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(1007)                     <b>for</b> (BasicBlock.ExceptionHandler exceptionHandler : basicBlock.getExceptionHandlers()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(1008)                         if (updateBasicBlock == null)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(1009)                             updateBasicBlock = searchUpdateBlockAndCreateContinueLoop(visited, basicBlock, exceptionHandler.getBasicBlock()); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(1010)                     } 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(1029)                 case 128:
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(1030)                     updateBasicBlock = searchUpdateBlockAndCreateContinueLoop(visited, basicBlock, basicBlock.getNext());
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(1031)                 case 64:
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(1032)                     <b>for</b> (BasicBlock.SwitchCase switchCase : basicBlock.getSwitchCases()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(1033)                         if (updateBasicBlock == null)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(1034)                             updateBasicBlock = searchUpdateBlockAndCreateContinueLoop(visited, basicBlock, switchCase.getBasicBlock()); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(1035)                     } 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(1103)                         changeEndLoopToJump(visited, target, basicBlock.getSub1());
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(1104)                     } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(1105)                 case 512:
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(1106)                     <b>for</b> (BasicBlock.ExceptionHandler exceptionHandler : basicBlock.getExceptionHandlers()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(1107)                         if (exceptionHandler.getBasicBlock() == BasicBlock.LOOP_END) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(1108)                             exceptionHandler.setBasicBlock(newJumpBasicBlock(basicBlock, target));
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(1109)                             continue;
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(1150)                         changeEndLoopToJump(visited, target, basicBlock.getNext());
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(1151)                     } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(1152)                 case 64:
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(1153)                     <b>for</b> (BasicBlock.SwitchCase switchCase : basicBlock.getSwitchCases()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(1154)                         if (switchCase.getBasicBlock() == BasicBlock.LOOP_END) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(1155)                             switchCase.setBasicBlock(newJumpBasicBlock(basicBlock, target));
org/jd/core/v1/service/converter/classfiletojavasyntax/util/ControlFlowGraphReducer.class:(1156)                             continue;
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(56)     
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(57)     protected Map<String, BaseType> typeBounds;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(58)     
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(59)     protected FormalParameters <b>for</b>malParameters;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(60)     
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(61)     protected PopulateBlackListNamesVisitor populateBlackListNamesVisitor = new PopulateBlackListNamesVisitor(this.blackListNames);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(62)     
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(74)         this.createParameterVisitor = new CreateParameterVisitor(typeMaker);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(75)         this.createLocalVariableVisitor = new CreateLocalVariableVisitor(typeMaker);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(76)         if (classFile.getFields() != null)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(77)             <b>for</b> (Field field : classFile.getFields()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(78)                 String descriptor = field.getDescriptor();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(79)                 if (descriptor.charAt(descriptor.length() - 1) == ';')
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(80)                     typeMaker.makeFromDescriptor(descriptor).accept(this.populateBlackListNamesVisitor); 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(83)         if (classFile.getSuperTypeName() != null)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(84)             typeMaker.makeFromInternalTypeName(classFile.getSuperTypeName()).accept(this.populateBlackListNamesVisitor); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(85)         if (classFile.getInterfaceTypeNames() != null)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(86)             <b>for</b> (String interfaceTypeName : classFile.getInterfaceTypeNames())
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(87)                 typeMaker.makeFromInternalTypeName(interfaceTypeName).accept(this.populateBlackListNamesVisitor);  
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(88)         if (parameterTypes != null)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(89)             if (parameterTypes.isList()) {
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(90)                 <b>for</b> (Type type : parameterTypes)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(91)                     type.accept(this.populateBlackListNamesVisitor); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(92)             } else {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(93)                 parameterTypes.getFirst().accept(this.populateBlackListNamesVisitor);
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(113)             int lastParameterIndex = parameterTypes.size() - 1;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(114)             boolean varargs = ((method.getAccessFlags() & 0x80) != 0);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(115)             initLocalVariablesFromParameterTypes(classFile, parameterTypes, varargs, firstVariableIndex, lastParameterIndex);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(116)             this.<b>for</b>malParameters = new FormalParameters();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(117)             AttributeParameterAnnotations rvpa = method.<AttributeParameterAnnotations>getAttribute("RuntimeVisibleParameterAnnotations");
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(118)             AttributeParameterAnnotations ripa = method.<AttributeParameterAnnotations>getAttribute("RuntimeInvisibleParameterAnnotations");
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(119)             if (rvpa == null && ripa == null) {
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(120)                 <b>for</b> (int parameterIndex = 0, variableIndex = firstVariableIndex; parameterIndex <= lastParameterIndex; parameterIndex++, variableIndex++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(121)                     AbstractLocalVariable lv = this.localVariableSet.root(variableIndex);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(122)                     this.<b>for</b>malParameters.add(new ClassFileFormalParameter(lv, (varargs && parameterIndex == lastParameterIndex)));
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(123)                     if (PrimitiveType.TYPE_LONG.equals(lv.getType()) || PrimitiveType.TYPE_DOUBLE.equals(lv.getType()))
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(124)                         variableIndex++; 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(125)                 } 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(127)                 Annotations[] visiblesArray = (rvpa == null) ? null : rvpa.getParameterAnnotations();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(128)                 Annotations[] invisiblesArray = (ripa == null) ? null : ripa.getParameterAnnotations();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(129)                 AnnotationConverter annotationConverter = new AnnotationConverter(typeMaker);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(130)                 <b>for</b> (int parameterIndex = 0, variableIndex = firstVariableIndex; parameterIndex <= lastParameterIndex; parameterIndex++, variableIndex++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(131)                     AbstractLocalVariable lv = this.localVariableSet.root(variableIndex);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(132)                     Annotations visibles = (visiblesArray == null || visiblesArray.length <= parameterIndex) ? null : visiblesArray[parameterIndex];
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(133)                     Annotations invisibles = (invisiblesArray == null || invisiblesArray.length <= parameterIndex) ? null : invisiblesArray[parameterIndex];
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(134)                     BaseAnnotationReference annotationReferences = annotationConverter.convert(visibles, invisibles);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(135)                     this.<b>for</b>malParameters.add(new ClassFileFormalParameter(annotationReferences, lv, (varargs && parameterIndex == lastParameterIndex)));
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(136)                     if (PrimitiveType.TYPE_LONG.equals(lv.getType()) || PrimitiveType.TYPE_DOUBLE.equals(lv.getType()))
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(137)                         variableIndex++; 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(138)                 } 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(147)             AttributeLocalVariableTable localVariableTable = code.<AttributeLocalVariableTable>getAttribute("LocalVariableTable");
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(148)             if (localVariableTable != null) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(149)                 boolean staticFlag = ((method.getAccessFlags() & 0x8) != 0);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(150)                 <b>for</b> (LocalVariable localVariable : localVariableTable.getLocalVariableTable()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(151)                     AbstractLocalVariable lv;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(152)                     int index = localVariable.getIndex();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(153)                     int startPc = (!staticFlag && index == 0) ? 0 : localVariable.getStartPc();
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(170)             AttributeLocalVariableTypeTable localVariableTypeTable = code.<AttributeLocalVariableTypeTable>getAttribute("LocalVariableTypeTable");
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(171)             if (localVariableTypeTable != null) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(172)                 UpdateTypeVisitor updateTypeVisitor = new UpdateTypeVisitor(this.localVariableSet);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(173)                 <b>for</b> (LocalVariableType lv : localVariableTypeTable.getLocalVariableTypeTable()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(174)                     updateTypeVisitor.setLocalVariableType(lv);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(175)                     this.typeMaker.makeFromSignature(lv.getSignature()).accept(updateTypeVisitor);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(176)                 } 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(181)     protected void initLocalVariablesFromParameterTypes(ClassFile classFile, BaseType parameterTypes, boolean varargs, int firstVariableIndex, int lastParameterIndex) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(182)         HashMap<Type, Boolean> typeMap = new HashMap<>();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(183)         DefaultList<Type> t = parameterTypes.getList();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(184)         <b>for</b> (int parameterIndex = 0; parameterIndex <= lastParameterIndex; parameterIndex++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(185)             Type type = t.get(parameterIndex);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(186)             typeMap.put(type, Boolean.valueOf(typeMap.containsKey(type)));
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(187)         } 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(197)         } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(198)         StringBuilder sb = new StringBuilder();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(199)         GenerateParameterSuffixNameVisitor generateParameterSuffixNameVisitor = new GenerateParameterSuffixNameVisitor();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(200)         <b>for</b> (int i = 0, variableIndex = firstVariableIndex; i <= lastParameterIndex; i++, variableIndex++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(201)             Type type = t.get(i);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(202)             AbstractLocalVariable lv = this.localVariableSet.root(variableIndex);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(203)             if (lv == null) {
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(392)     }
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(393)     
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(394)     public BaseFormalParameter getFormalParameters() {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(395)         return this.<b>for</b>malParameters;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(396)     }
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(397)     
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LocalVariableMaker.class:(398)     public void pushFrame(Statements statements) {
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LoopStatementMaker.class:(403)             return null; 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LoopStatementMaker.class:(404)         SearchLocalVariableReferenceVisitor visitor1 = new SearchLocalVariableReferenceVisitor();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LoopStatementMaker.class:(405)         visitor1.init(syntheticIterator.getIndex());
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LoopStatementMaker.class:(406)         <b>for</b> (int i = 1, len = subStatements.size(); i < len; i++)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LoopStatementMaker.class:(407)             subStatements.get(i).accept(visitor1); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LoopStatementMaker.class:(408)         if (visitor1.containsReference())
org/jd/core/v1/service/converter/classfiletojavasyntax/util/LoopStatementMaker.class:(409)             return null; 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/MergeMembersUtil.class:(65)     protected static void sort(List<? extends ClassFileMemberDeclaration> members) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/MergeMembersUtil.class:(66)         int order = 0;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/MergeMembersUtil.class:(67)         int lastLineNumber = 0;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/MergeMembersUtil.class:(68)         <b>for</b> (ClassFileMemberDeclaration member : members) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/MergeMembersUtil.class:(69)             int lineNumber = member.getFirstLineNumber();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/MergeMembersUtil.class:(70)             if (lineNumber > 0 && lineNumber != lastLineNumber) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/MergeMembersUtil.class:(71)                 if (lastLineNumber > 0)
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(305)         List<SwitchStatement.Block> blocks = switchStatement.getBlocks();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(306)         DefaultStack<Expression> localStack = new DefaultStack<>(this.stack);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(307)         switchCases.sort(SWITCH_CASE_COMPARATOR);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(308)         <b>for</b> (int i = 0, len = switchCases.size(); i < len; i++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(309)             BasicBlock.SwitchCase sc = switchCases.get(i);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(310)             BasicBlock bb = sc.getBasicBlock();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(311)             int j = i + 1;
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(322)                 blocks.add(new SwitchStatement.LabelBlock(label, subStatements));
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(323)             } else {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(324)                 DefaultList<SwitchStatement.Label> labels = new DefaultList<>(j - i);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(325)                 <b>for</b> (; i < j; i++)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(326)                     labels.add(new SwitchStatement.ExpressionLabel(new IntegerConstantExpression(conditionType, ((BasicBlock.SwitchCase)switchCases.get(i)).getValue()))); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(327)                 blocks.add(new SwitchStatement.MultiLabelsBlock(labels, subStatements));
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(328)                 i--;
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(343)         Statements finallyStatements = null;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(344)         int assertStackSize = this.stack.size();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(345)         Statements tryStatements = makeSubStatements(watchdog, basicBlock.getSub1(), statements, jumps);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(346)         <b>for</b> (BasicBlock.ExceptionHandler exceptionHandler : basicBlock.getExceptionHandlers()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(347)             assert this.stack.size() == assertStackSize : "parseTry : problem with stack";
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(348)             if (exceptionHandler.getInternalThrowableName() == null) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(349)                 this.stack.push(FINALLY_EXCEPTION_EXPRESSION);
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(402)             replacePreOperatorWithPostOperator(catchStatements);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(403)             ClassFileTryStatement.CatchClause cc = new ClassFileTryStatement.CatchClause(lineNumber, ot, exception, catchStatements);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(404)             if (exceptionHandler.getOtherInternalThrowableNames() != null)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(405)                 <b>for</b> (String name : exceptionHandler.getOtherInternalThrowableNames())
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(406)                     cc.addType(this.typeMaker.makeFromInternalTypeName(name));  
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(407)             catchClauses.add(cc);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(408)         } 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(550)             if (sub1.getType() == 4194304 && sub1.getNext() == last && countStartLoop(sub1.getSub1()) == 0) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(551)                 changeEndLoopToStartLoop(new BitSet(), sub1.getSub1());
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(552)                 subStatements = makeSubStatements(watchdog, sub1.getSub1(), statements, jumps, updateStatements);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(553)                 assert subStatements.getLast() == ContinueStatement.CONTINUE : "StatementMaker.parseLoop(...) : unexpected basic block <b>for</b> create a do-while loop";
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(554)                 subStatements.removeLast();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(555)             } else {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(556)                 createDoWhileContinue(last);
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(569)         while (bb.matchType(876822149)) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(570)             switch (bb.getType()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(571)                 case 128:
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(572)                     <b>for</b> (BasicBlock.SwitchCase switchCase : bb.getSwitchCases())
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(573)                         count += countStartLoop(switchCase.getBasicBlock()); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(574)                     break;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(575)                 case 1024:
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(576)                 case 2048:
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(577)                 case 4096:
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(578)                     count += countStartLoop(bb.getSub1());
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(579)                     <b>for</b> (BasicBlock.ExceptionHandler exceptionHandler : bb.getExceptionHandlers())
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(580)                         count += countStartLoop(exceptionHandler.getBasicBlock()); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(581)                     break;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(582)                 case 131072:
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(597)         boolean change;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(598)         do {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(599)             change = false;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(600)             <b>for</b> (BasicBlock predecessor : last.getPredecessors()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(601)                 if (predecessor.getType() == 65536) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(602)                     BasicBlock l = predecessor.getSub1();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(603)                     if (l.matchType(876822149)) {
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(648)                 case 1024:
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(649)                 case 2048:
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(650)                 case 4096:
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(651)                     <b>for</b> (BasicBlock.ExceptionHandler exceptionHandler : basicBlock.getExceptionHandlers()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(652)                         if (exceptionHandler.getBasicBlock() == BasicBlock.LOOP_END) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(653)                             exceptionHandler.setBasicBlock(BasicBlock.LOOP_START);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(654)                             continue;
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(677)                     break;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(678)                 case 64:
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(679)                 case 128:
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(680)                     <b>for</b> (BasicBlock.SwitchCase switchCase : basicBlock.getSwitchCases()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(681)                         if (switchCase.getBasicBlock() == BasicBlock.LOOP_END) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(682)                             switchCase.setBasicBlock(BasicBlock.LOOP_START);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(683)                             continue;
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(775)     
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(776)     protected Expression createObjectTypeReferenceDotClassExpression(int lineNumber, String fieldName, MethodInvocationExpression mie) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(777)         this.memberVisitor.init(fieldName);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(778)         <b>for</b> (ClassFileFieldDeclaration field : this.bodyDeclaration.getFieldDeclarations()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(779)             field.getFieldDeclarators().accept(this.memberVisitor);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(780)             if (this.memberVisitor.found()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(781)                 field.setFlags(field.getFlags() | 0x1000);
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(783)             } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(784)         } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(785)         this.memberVisitor.init("class$");
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(786)         <b>for</b> (ClassFileConstructorOrMethodDeclaration member : this.bodyDeclaration.getMethodDeclarations()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(787)             member.accept(this.memberVisitor);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(788)             if (this.memberVisitor.found()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/StatementMaker.class:(789)                 member.setFlags(member.getFlags() | 0x1000);
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(61)                                                         MethodInvocationExpression mie = (MethodInvocationExpression)previousSwitchStatement.getCondition();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(62)                                                         if (mie.getName().equals("hashCode") && mie.getDescriptor().equals("()I")) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(63)                                                             HashMap<Integer, String> map = new HashMap<>();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(64)                                                             <b>for</b> (SwitchStatement.Block block : previousSwitchStatement.getBlocks()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(65)                                                                 BaseStatement stmts = block.getStatements();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(66)                                                                 assert stmts != null && stmts.getClass() == Statements.class && !((Statements)stmts).isEmpty();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(67)                                                                 <b>for</b> (Statement stmt : stmts) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(68)                                                                     if (stmt.getClass() != IfStatement.class)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(69)                                                                         break; 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(70)                                                                     IfStatement is = (IfStatement)stmt;
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(88)                                                                     map.put(index, string);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(89)                                                                 } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(90)                                                             } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(91)                                                             <b>for</b> (SwitchStatement.Block block : switchStatement.getBlocks()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(92)                                                                 if (block.getClass() == SwitchStatement.LabelBlock.class) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(93)                                                                     SwitchStatement.LabelBlock lb = (SwitchStatement.LabelBlock)block;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(94)                                                                     if (lb.getLabel() != SwitchStatement.DEFAULT_LABEL) {
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(100)                                                                 } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(101)                                                                 if (block.getClass() == SwitchStatement.MultiLabelsBlock.class) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(102)                                                                     SwitchStatement.MultiLabelsBlock lmb = (SwitchStatement.MultiLabelsBlock)block;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(103)                                                                     <b>for</b> (SwitchStatement.Label label : lmb.getLabels()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(104)                                                                         if (label != SwitchStatement.DEFAULT_LABEL) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(105)                                                                             SwitchStatement.ExpressionLabel el = (SwitchStatement.ExpressionLabel)label;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(106)                                                                             IntegerConstantExpression nce = (IntegerConstantExpression)el.getExpression();
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(151)             MethodInvocationExpression mie = (MethodInvocationExpression)ae.getExpression();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(152)             String methodName = mie.getName();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(153)             if (mie.getDescriptor().equals("()[I") && methodName.startsWith("$SWITCH_TABLE$"))
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(154)                 <b>for</b> (ClassFileConstructorOrMethodDeclaration declaration : bodyDeclaration.getMethodDeclarations()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(155)                     if (declaration.getMethod().getName().equals(methodName)) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(156)                         DefaultList<Statement> statements = declaration.getStatements().getList();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(157)                         updateSwitchStatement(switchStatement, statements.listIterator(3));
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(196)         ArrayExpression ae = (ArrayExpression)switchStatement.getCondition();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(197)         Expression expression = ((MethodReferenceExpression)ae.getIndex()).getExpression();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(198)         ObjectType type = (ObjectType)expression.getType();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(199)         <b>for</b> (SwitchStatement.Block block : switchStatement.getBlocks()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(200)             if (block.getClass() == SwitchStatement.LabelBlock.class) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(201)                 SwitchStatement.LabelBlock lb = (SwitchStatement.LabelBlock)block;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(202)                 if (lb.getLabel() != SwitchStatement.DEFAULT_LABEL) {
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(208)             } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(209)             if (block.getClass() == SwitchStatement.MultiLabelsBlock.class) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(210)                 SwitchStatement.MultiLabelsBlock lmb = (SwitchStatement.MultiLabelsBlock)block;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(211)                 <b>for</b> (SwitchStatement.Label label : lmb.getLabels()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(212)                     if (label != SwitchStatement.DEFAULT_LABEL) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(213)                         SwitchStatement.ExpressionLabel el = (SwitchStatement.ExpressionLabel)label;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/SwitchStatementMaker.class:(214)                         IntegerConstantExpression nce = (IntegerConstantExpression)el.getExpression();
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(120)                     typeTypes.interfaces = makeFromInternalTypeName(interfaceTypeNames[0]);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(121)                 } else {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(122)                     UnmodifiableTypes list = new UnmodifiableTypes(length);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(123)                     <b>for</b> (String interfaceTypeName : interfaceTypeNames)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(124)                         list.add(makeFromInternalTypeName(interfaceTypeName)); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(125)                     typeTypes.interfaces = list;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(126)                 } 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(197)     
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(198)     public static int countDimension(String descriptor) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(199)         int count = 0;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(200)         <b>for</b> (int i = 0, len = descriptor.length(); i < len && descriptor.charAt(i) == '['; i++)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(201)             count++; 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(202)         return count;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(203)     }
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(235)         boolean containsThrowsSignature = (signature.indexOf('^') != -1);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(236)         if (!containsThrowsSignature && exceptionTypeNames != null) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(237)             StringBuilder sb = new StringBuilder(signature);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(238)             <b>for</b> (String exceptionTypeName : exceptionTypeNames)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(239)                 sb.append("^L").append(exceptionTypeName).append(';'); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(240)             cacheKey = sb.toString();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(241)         } 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(269)                         methodTypes.exceptionTypes = makeFromInternalTypeName(exceptionTypeNames[0]);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(270)                     } else {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(271)                         UnmodifiableTypes list = new UnmodifiableTypes(exceptionTypeNames.length);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(272)                         <b>for</b> (String exceptionTypeName : exceptionTypeNames)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(273)                             list.add(makeFromInternalTypeName(exceptionTypeName)); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(274)                         methodTypes.exceptionTypes = list;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(275)                     }  
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(678)                     } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(679)                 } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(680)                 if (rightTypeTypes.interfaces != null)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(681)                     <b>for</b> (Type interfaze : rightTypeTypes.interfaces) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(682)                         ObjectType ot = searchSuperParameterizedType(superHashCode, superInternalTypeName, ((ObjectType)interfaze).createType((BaseTypeArgument)null));
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(683)                         if (ot != null) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(684)                             this.superParameterizedObjectTypes.put(key, ot);
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(708)                     } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(709)                 } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(710)                 if (rightTypeTypes.interfaces != null)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(711)                     <b>for</b> (Type interfaze : rightTypeTypes.interfaces) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(712)                         bindTypesToTypesVisitor.init();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(713)                         interfaze.accept(bindTypesToTypesVisitor);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(714)                         ObjectType ot = (ObjectType)bindTypesToTypesVisitor.getType();
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(747)             superClassAndInterfaceNames = this.hierarchy.get(rightInternalName);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(748)         } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(749)         if (superClassAndInterfaceNames != null) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(750)             <b>for</b> (String name : superClassAndInterfaceNames) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(751)                 if (leftInternalName.equals(name)) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(752)                     this.assignableRawTypes.put(key, Boolean.TRUE);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(753)                     return true;
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(754)                 } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(755)             } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(756)             <b>for</b> (String name : superClassAndInterfaceNames) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(757)                 if (isRawTypeAssignable(leftHashCode, leftInternalName, name)) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(758)                     this.assignableRawTypes.put(key, Boolean.TRUE);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(759)                     return true;
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(789)         skipMembers(reader);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(790)         String signature = null;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(791)         int count = reader.readUnsignedShort();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(792)         <b>for</b> (int j = 0; j < count; j++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(793)             int attributeNameIndex = reader.readUnsignedShort();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(794)             int attributeLength = reader.readInt();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(795)             if ("Signature".equals(constants[attributeNameIndex])) {
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(813)             } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(814)             int length = superClassAndInterfaceNames.length;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(815)             UnmodifiableTypes list = new UnmodifiableTypes(length - 1);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(816)             <b>for</b> (int i = 1; i < length; i++)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(817)                 list.add(makeFromInternalTypeName(superClassAndInterfaceNames[i])); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(818)             typeTypes.interfaces = list;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(819)         } 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(858)                             type = loadFieldType(typeTypes.superType, fieldName, descriptor); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(859)                         if (type == null && typeTypes.interfaces != null)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(860)                             if (typeTypes.interfaces.isList()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(861)                                 <b>for</b> (Type interfaze : typeTypes.interfaces) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(862)                                     type = loadFieldType((ObjectType)interfaze, fieldName, descriptor);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(863)                                     if (type != null)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(864)                                         break; 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(926)                             methodTypes = loadMethodTypes(typeTypes.superType, methodName, descriptor); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(927)                         if (methodTypes == null && typeTypes.interfaces != null)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(928)                             if (typeTypes.interfaces.isList()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(929)                                 <b>for</b> (Type interfaze : typeTypes.interfaces) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(930)                                     methodTypes = loadMethodTypes((ObjectType)interfaze, methodName, descriptor);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(931)                                     if (methodTypes != null)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(932)                                         break; 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1008)         String outerTypeName = null;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1009)         ObjectType outerObjectType = null;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1010)         int count = reader.readUnsignedShort();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1011)         <b>for</b> (int i = 0; i < count; i++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1012)             int attributeNameIndex = reader.readUnsignedShort();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1013)             int attributeLength = reader.readInt();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1014)             if ("InnerClasses".equals(constants[attributeNameIndex])) {
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1015)                 int innerClassesCount = reader.readUnsignedShort();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1016)                 <b>for</b> (int j = 0; j < innerClassesCount; j++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1017)                     int innerTypeIndex = reader.readUnsignedShort();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1018)                     int outerTypeIndex = reader.readUnsignedShort();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1019)                     reader.skip(4);
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1078)             Object[] constants = loadClassFile(internalTypeName, reader);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1079)             int count = reader.readUnsignedShort();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1080)             int i;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1081)             <b>for</b> (i = 0; i < count; i++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1082)                 reader.skip(2);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1083)                 int nameIndex = reader.readUnsignedShort();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1084)                 int descriptorIndex = reader.readUnsignedShort();
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1085)                 String signature = null;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1086)                 int attributeCount = reader.readUnsignedShort();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1087)                 <b>for</b> (int j = 0; j < attributeCount; j++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1088)                     int attributeNameIndex = reader.readUnsignedShort();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1089)                     int attributeLength = reader.readInt();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1090)                     if ("Signature".equals(constants[attributeNameIndex])) {
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1103)                 } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1104)             } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1105)             count = reader.readUnsignedShort();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1106)             <b>for</b> (i = 0; i < count; i++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1107)                 MethodTypes methodTypes;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1108)                 reader.skip(2);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1109)                 int nameIndex = reader.readUnsignedShort();
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1111)                 String signature = null;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1112)                 String[] exceptionTypeNames = null;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1113)                 int attributeCount = reader.readUnsignedShort();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1114)                 <b>for</b> (int j = 0; j < attributeCount; j++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1115)                     int exceptionCount, attributeNameIndex = reader.readUnsignedShort();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1116)                     int attributeLength = reader.readInt();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1117)                     String str = (String)constants[attributeNameIndex];
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1123)                             exceptionCount = reader.readUnsignedShort();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1124)                             if (exceptionCount > 0) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1125)                                 exceptionTypeNames = new String[exceptionCount];
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1126)                                 <b>for</b> (int k = 0; k < exceptionCount; k++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1127)                                     int exceptionClassIndex = reader.readUnsignedShort();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1128)                                     Integer cc = (Integer)constants[exceptionClassIndex];
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1129)                                     exceptionTypeNames[k] = (String)constants[cc.intValue()];
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1173)         int count = reader.readUnsignedShort();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1174)         String[] superClassAndInterfaceNames = new String[count + 1];
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1175)         superClassAndInterfaceNames[0] = superClassName;
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1176)         <b>for</b> (int i = 1; i <= count; i++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1177)             int interfaceIndex = reader.readUnsignedShort();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1178)             Integer cc = (Integer)constants[interfaceIndex];
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1179)             superClassAndInterfaceNames[i] = (String)constants[cc.intValue()];
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1184)     
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1185)     private static void skipMembers(ClassFileReader reader) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1186)         int count = reader.readUnsignedShort();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1187)         <b>for</b> (int i = 0; i < count; i++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1188)             reader.skip(6);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1189)             skipAttributes(reader);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1190)         } 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1195)         if (count == 0)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1196)             return null; 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1197)         Object[] constants = new Object[count];
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1198)         <b>for</b> (int i = 1; i < count; i++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1199)             int tag = reader.readByte();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1200)             switch (tag) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1201)                 case 1:
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1237)     
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1238)     private static void skipAttributes(ClassFileReader reader) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1239)         int count = reader.readUnsignedShort();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1240)         <b>for</b> (int i = 0; i < count; i++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1241)             reader.skip(2);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1242)             int attributeLength = reader.readInt();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1243)             reader.skip(attributeLength);
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1263)                             bool = multipleMethods(typeTypes.superType.getInternalName(), suffixKey); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1264)                         if (bool == null && typeTypes.interfaces != null)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1265)                             if (typeTypes.interfaces.isList()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1266)                                 <b>for</b> (Type interfaze : typeTypes.interfaces) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1267)                                     bool = multipleMethods(((ObjectType)interfaze).getInternalName(), suffixKey);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1268)                                     if (bool != null)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeMaker.class:(1269)                                         break; 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeParametersToTypeArgumentsBinder.class:(207)                             expressionType = expressionObjectType.createType((BaseTypeArgument)null);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeParametersToTypeArgumentsBinder.class:(208)                         } else if (typeParameters.isList()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeParametersToTypeArgumentsBinder.class:(209)                             TypeArguments tas = new TypeArguments(typeParameters.size());
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeParametersToTypeArgumentsBinder.class:(210)                             <b>for</b> (TypeParameter typeParameter : typeParameters)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeParametersToTypeArgumentsBinder.class:(211)                                 tas.add(bindings.get(typeParameter.getIdentifier())); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeParametersToTypeArgumentsBinder.class:(212)                             expressionType = expressionObjectType.createType(tas);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeParametersToTypeArgumentsBinder.class:(213)                         } else {
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeParametersToTypeArgumentsBinder.class:(234)                 if (typeParameters != null && typeArguments == null)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeParametersToTypeArgumentsBinder.class:(235)                     if (typeParameters.isList()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeParametersToTypeArgumentsBinder.class:(236)                         TypeArguments tas = new TypeArguments(typeParameters.size());
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeParametersToTypeArgumentsBinder.class:(237)                         <b>for</b> (TypeParameter typeParameter : typeParameters)
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeParametersToTypeArgumentsBinder.class:(238)                             tas.add(new GenericType(typeParameter.getIdentifier())); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeParametersToTypeArgumentsBinder.class:(239)                         neObjectType = neObjectType.createType(tas);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeParametersToTypeArgumentsBinder.class:(240)                     } else {
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeParametersToTypeArgumentsBinder.class:(247)                 } 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeParametersToTypeArgumentsBinder.class:(248)                 Map<String, TypeArgument> bindings = createBindings(null, typeParameters, typeArguments, null, type, t, parameterTypes, parameters);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeParametersToTypeArgumentsBinder.class:(249)                 ne.setParameterTypes(parameterTypes = bind(bindings, parameterTypes));
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeParametersToTypeArgumentsBinder.class:(250)                 <b>for</b> (Map.Entry<String, TypeArgument> entry : bindings.entrySet()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeParametersToTypeArgumentsBinder.class:(251)                     this.typeArgumentToTypeVisitor.init();
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeParametersToTypeArgumentsBinder.class:(252)                     ((TypeArgument)entry.getValue()).accept(this.typeArgumentToTypeVisitor);
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeParametersToTypeArgumentsBinder.class:(253)                     entry.setValue(this.typeArgumentToTypeVisitor.getType());
--
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeParametersToTypeArgumentsBinder.class:(325)             }  
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeParametersToTypeArgumentsBinder.class:(326)         if (bindings.containsValue(null))
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeParametersToTypeArgumentsBinder.class:(327)             if (eraseTypeArguments(expression, typeParameters, typeArguments)) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeParametersToTypeArgumentsBinder.class:(328)                 <b>for</b> (Map.Entry<String, TypeArgument> entry : bindings.entrySet())
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeParametersToTypeArgumentsBinder.class:(329)                     entry.setValue(null); 
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeParametersToTypeArgumentsBinder.class:(330)             } else {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeParametersToTypeArgumentsBinder.class:(331)                 <b>for</b> (Map.Entry<String, TypeArgument> entry : bindings.entrySet()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeParametersToTypeArgumentsBinder.class:(332)                     if (entry.getValue() == null) {
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeParametersToTypeArgumentsBinder.class:(333)                         BaseType baseType = typeBounds.get(entry.getKey());
org/jd/core/v1/service/converter/classfiletojavasyntax/util/TypeParametersToTypeArgumentsBinder.class:(334)                         if (baseType == null) {
--
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(226)     public void visit(SuperConstructorInvocationExpression expression) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(227)         BaseExpression parameters = expression.getParameters();
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(228)         if (parameters != null) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(229)             boolean <b>for</b>ce = (parameters.size() > 0 && this.typeMaker.multipleMethods(expression.getObjectType().getInternalName(), "<init>", parameters.size()));
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(230)             expression.setParameters(updateExpressions(((ClassFileSuperConstructorInvocationExpression)expression).getParameterTypes(), parameters, <b>for</b>ce));
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(231)         } 
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(232)     }
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(233)     
--
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(234)     public void visit(ConstructorInvocationExpression expression) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(235)         BaseExpression parameters = expression.getParameters();
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(236)         if (parameters != null) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(237)             boolean <b>for</b>ce = (parameters.size() > 0 && this.typeMaker.multipleMethods(expression.getObjectType().getInternalName(), "<init>", parameters.size()));
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(238)             expression.setParameters(updateExpressions(((ClassFileConstructorInvocationExpression)expression).getParameterTypes(), parameters, <b>for</b>ce));
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(239)         } 
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(240)     }
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(241)     
--
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(242)     public void visit(MethodInvocationExpression expression) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(243)         BaseExpression parameters = expression.getParameters();
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(244)         if (parameters != null) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(245)             boolean <b>for</b>ce = (parameters.size() > 0 && this.typeMaker.multipleMethods(expression.getInternalTypeName(), expression.getName(), parameters.size()));
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(246)             expression.setParameters(updateExpressions(((ClassFileMethodInvocationExpression)expression).getParameterTypes(), parameters, <b>for</b>ce));
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(247)         } 
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(248)         expression.getExpression().accept(this);
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(249)     }
--
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(297)         expression.setExpressionFalse(updateExpression(expressionType, expression.getExpressionFalse(), false));
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(298)     }
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(299)     
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(300)     protected BaseExpression updateExpressions(BaseType types, BaseExpression expressions, boolean <b>for</b>ce) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(301)         if (expressions != null)
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(302)             if (expressions.isList()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(303)                 DefaultList<Type> t = types.getList();
--
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(304)                 DefaultList<Expression> e = expressions.getList();
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(305)                 <b>for</b> (int i = e.size() - 1; i >= 0; i--)
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(306)                     e.set(i, updateExpression(t.get(i), e.get(i), <b>for</b>ce)); 
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(307)             } else {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(308)                 expressions = updateExpression(types.getFirst(), (Expression)expressions, <b>for</b>ce);
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(309)             }  
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(310)         return expressions;
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(311)     }
--
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(312)     
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(313)     private Expression updateExpression(Type type, Expression expression, boolean <b>for</b>ce) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(314)         Class<?> expressionClass = expression.getClass();
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(315)         if (expressionClass == NullExpression.class) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(316)             if (<b>for</b>ce) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(317)                 this.searchFirstLineNumberVisitor.init();
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(318)                 expression.accept(this.searchFirstLineNumberVisitor);
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/AddCastExpressionVisitor.class:(319)                 expression = new CastExpression(this.searchFirstLineNumberVisitor.getLineNumber(), type, expression);
--
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/BindTypeArgumentsToTypeArgumentsVisitor.class:(40)     public void visit(TypeArguments arguments) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/BindTypeArgumentsToTypeArgumentsVisitor.class:(41)         int size = arguments.size();
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/BindTypeArgumentsToTypeArgumentsVisitor.class:(42)         int i;
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/BindTypeArgumentsToTypeArgumentsVisitor.class:(43)         <b>for</b> (i = 0; i < size; i++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/BindTypeArgumentsToTypeArgumentsVisitor.class:(44)             TypeArgument ta = arguments.get(i);
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/BindTypeArgumentsToTypeArgumentsVisitor.class:(45)             ta.accept(this);
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/BindTypeArgumentsToTypeArgumentsVisitor.class:(46)             if (this.result != ta)
--
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/BindTypeArgumentsToTypeArgumentsVisitor.class:(53)                 TypeArguments newTypes = new TypeArguments(size);
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/BindTypeArgumentsToTypeArgumentsVisitor.class:(54)                 newTypes.addAll(arguments.subList(0, i));
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/BindTypeArgumentsToTypeArgumentsVisitor.class:(55)                 newTypes.add((TypeArgument)this.result);
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/BindTypeArgumentsToTypeArgumentsVisitor.class:(56)                 <b>for</b> (; ++i < size; i++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/BindTypeArgumentsToTypeArgumentsVisitor.class:(57)                     TypeArgument ta = arguments.get(i);
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/BindTypeArgumentsToTypeArgumentsVisitor.class:(58)                     ta.accept(this);
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/BindTypeArgumentsToTypeArgumentsVisitor.class:(59)                     if (this.result == null)
--
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/BindTypeParametersToNonWildcardTypeArgumentsVisitor.class:(46)     public void visit(TypeParameters parameters) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/BindTypeParametersToNonWildcardTypeArgumentsVisitor.class:(47)         int size = parameters.size();
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/BindTypeParametersToNonWildcardTypeArgumentsVisitor.class:(48)         TypeArguments arguments = new TypeArguments(size);
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/BindTypeParametersToNonWildcardTypeArgumentsVisitor.class:(49)         <b>for</b> (TypeParameter parameter : parameters) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/BindTypeParametersToNonWildcardTypeArgumentsVisitor.class:(50)             parameter.accept(this);
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/BindTypeParametersToNonWildcardTypeArgumentsVisitor.class:(51)             if (this.result == null)
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/BindTypeParametersToNonWildcardTypeArgumentsVisitor.class:(52)                 return; 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/BindTypesToTypesVisitor.class:(101)     public void visit(Types types) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/BindTypesToTypesVisitor.class:(102)         int size = types.size();
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/BindTypesToTypesVisitor.class:(103)         int i;
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/BindTypesToTypesVisitor.class:(104)         <b>for</b> (i = 0; i < size; i++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/BindTypesToTypesVisitor.class:(105)             Type t = types.get(i);
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/BindTypesToTypesVisitor.class:(106)             t.accept(this);
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/BindTypesToTypesVisitor.class:(107)             if (this.result != t)
--
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/BindTypesToTypesVisitor.class:(113)             Types newTypes = new Types(size);
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/BindTypesToTypesVisitor.class:(114)             newTypes.addAll(types.subList(0, i));
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/BindTypesToTypesVisitor.class:(115)             newTypes.add((Type)this.result);
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/BindTypesToTypesVisitor.class:(116)             <b>for</b> (; ++i < size; i++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/BindTypesToTypesVisitor.class:(117)                 Type t = types.get(i);
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/BindTypesToTypesVisitor.class:(118)                 t.accept(this);
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/BindTypesToTypesVisitor.class:(119)                 newTypes.add((Type)this.result);
--
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/CreateInstructionsVisitor.class:(42)         ClassFileBodyDeclaration bodyDeclaration = (ClassFileBodyDeclaration)declaration;
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/CreateInstructionsVisitor.class:(43)         List<ClassFileConstructorOrMethodDeclaration> methods = bodyDeclaration.getMethodDeclarations();
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/CreateInstructionsVisitor.class:(44)         if (methods != null) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/CreateInstructionsVisitor.class:(45)             <b>for</b> (ClassFileConstructorOrMethodDeclaration method : methods) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/CreateInstructionsVisitor.class:(46)                 if ((method.getFlags() & 0x1040) != 0) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/CreateInstructionsVisitor.class:(47)                     method.accept(this);
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/CreateInstructionsVisitor.class:(48)                     continue;
--
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/CreateInstructionsVisitor.class:(56)                 } 
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/CreateInstructionsVisitor.class:(57)                 if (method.getParameterTypes() != null) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/CreateInstructionsVisitor.class:(58)                     if (method.getParameterTypes().isList()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/CreateInstructionsVisitor.class:(59)                         <b>for</b> (Type type1 : method.getParameterTypes()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/CreateInstructionsVisitor.class:(60)                             if (type1.isObject() && type1.getName() == null) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/CreateInstructionsVisitor.class:(61)                                 method.setFlags(method.getFlags() | 0x1000);
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/CreateInstructionsVisitor.class:(62)                                 method.accept(this);
--
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/CreateInstructionsVisitor.class:(72)                     } 
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/CreateInstructionsVisitor.class:(73)                 } 
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/CreateInstructionsVisitor.class:(74)             } 
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/CreateInstructionsVisitor.class:(75)             <b>for</b> (ClassFileConstructorOrMethodDeclaration method : methods) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/CreateInstructionsVisitor.class:(76)                 if ((method.getFlags() & 0x1040) == 0)
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/CreateInstructionsVisitor.class:(77)                     method.accept(this); 
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/CreateInstructionsVisitor.class:(78)             } 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/InitInstanceFieldVisitor.class:(223)     protected void updateFieldsAndConstructors() {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/InitInstanceFieldVisitor.class:(224)         int count = this.putFields.size();
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/InitInstanceFieldVisitor.class:(225)         if (count > 0) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/InitInstanceFieldVisitor.class:(226)             <b>for</b> (BinaryOperatorExpression putField : this.putFields) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/InitInstanceFieldVisitor.class:(227)                 FieldReferenceExpression fre = (FieldReferenceExpression)putField.getLeftExpression();
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/InitInstanceFieldVisitor.class:(228)                 FieldDeclarator declaration = this.fieldDeclarators.get(fre.getName());
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/InitInstanceFieldVisitor.class:(229)                 if (declaration != null) {
--
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/InitInstanceFieldVisitor.class:(232)                     ((ClassFileFieldDeclaration)declaration.getFieldDeclaration()).setFirstLineNumber(expression.getLineNumber());
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/InitInstanceFieldVisitor.class:(233)                 } 
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/InitInstanceFieldVisitor.class:(234)             } 
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/InitInstanceFieldVisitor.class:(235)             <b>for</b> (Data data : this.datas) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/InitInstanceFieldVisitor.class:(236)                 data.statements.subList(data.index, data.index + count).clear();
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/InitInstanceFieldVisitor.class:(237)                 if (data.statements.isEmpty()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/InitInstanceFieldVisitor.class:(238)                     data.declaration.setStatements((BaseStatement)null);
--
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/InitStaticFieldVisitor.class:(75)             this.methods = bodyDeclaration.getMethodDeclarations();
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/InitStaticFieldVisitor.class:(76)             if (this.methods != null) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/InitStaticFieldVisitor.class:(77)                 this.deleteStaticDeclaration = null;
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/InitStaticFieldVisitor.class:(78)                 <b>for</b> (int i = 0, len = this.methods.size(); i < len; i++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/InitStaticFieldVisitor.class:(79)                     ((ClassFileConstructorOrMethodDeclaration)this.methods.get(i)).accept(this);
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/InitStaticFieldVisitor.class:(80)                     if (this.deleteStaticDeclaration != null) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/InitStaticFieldVisitor.class:(81)                         if (this.deleteStaticDeclaration.booleanValue())
--
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/InitStaticFieldVisitor.class:(108)                 if (statements.size() > 0) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/InitStaticFieldVisitor.class:(109)                     DefaultList<Statement> list = statements.getList();
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/InitStaticFieldVisitor.class:(110)                     Iterator<FieldDeclarator> fieldDeclaratorIterator = this.fields.iterator();
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/InitStaticFieldVisitor.class:(111)                     <b>for</b> (int i = 0, len = list.size(); i < len; i++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/InitStaticFieldVisitor.class:(112)                         Statement statement = list.get(i);
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/InitStaticFieldVisitor.class:(113)                         if (setStaticFieldInitializer(statement, fieldDeclaratorIterator)) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/InitStaticFieldVisitor.class:(114)                             if (i > 0) {
--
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/MergeTryWithResourcesStatementVisitor.class:(35)     }
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/MergeTryWithResourcesStatementVisitor.class:(36)     
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/MergeTryWithResourcesStatementVisitor.class:(37)     public void visit(SwitchStatement statement) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/MergeTryWithResourcesStatementVisitor.class:(38)         <b>for</b> (SwitchStatement.Block block : statement.getBlocks())
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/MergeTryWithResourcesStatementVisitor.class:(39)             block.getStatements().accept(this); 
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/MergeTryWithResourcesStatementVisitor.class:(40)     }
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/MergeTryWithResourcesStatementVisitor.class:(41)     
--
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/MergeTryWithResourcesStatementVisitor.class:(136)     }
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/MergeTryWithResourcesStatementVisitor.class:(137)     
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/MergeTryWithResourcesStatementVisitor.class:(138)     protected void acceptListStatement(List<? extends Statement> list) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/MergeTryWithResourcesStatementVisitor.class:(139)         <b>for</b> (Statement statement : list)
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/MergeTryWithResourcesStatementVisitor.class:(140)             statement.accept(this); 
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/MergeTryWithResourcesStatementVisitor.class:(141)     }
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/MergeTryWithResourcesStatementVisitor.class:(142)     
--
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/MergeTryWithResourcesStatementVisitor.class:(143)     protected void safeAcceptListStatement(List<? extends Statement> list) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/MergeTryWithResourcesStatementVisitor.class:(144)         if (list != null)
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/MergeTryWithResourcesStatementVisitor.class:(145)             <b>for</b> (Statement statement : list)
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/MergeTryWithResourcesStatementVisitor.class:(146)                 statement.accept(this);  
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/MergeTryWithResourcesStatementVisitor.class:(147)     }
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/MergeTryWithResourcesStatementVisitor.class:(148) }
--
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/PopulateBindingsWithTypeParameterVisitor.class:(27)     }
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/PopulateBindingsWithTypeParameterVisitor.class:(28)     
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/PopulateBindingsWithTypeParameterVisitor.class:(29)     public void visit(TypeParameters parameters) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/PopulateBindingsWithTypeParameterVisitor.class:(30)         <b>for</b> (TypeParameter parameter : parameters)
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/PopulateBindingsWithTypeParameterVisitor.class:(31)             parameter.accept(this); 
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/PopulateBindingsWithTypeParameterVisitor.class:(32)     }
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/PopulateBindingsWithTypeParameterVisitor.class:(33) }
--
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(78)             if (this.statementCountToRemove > 0)
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(79)                 if (i > this.statementCountToRemove) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(80)                     List<Statement> list = statements.subList(i - this.statementCountToRemove, i);
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(81)                     <b>for</b> (Statement statement : list)
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(82)                         statement.accept(this.declaredSyntheticLocalVariableVisitor); 
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(83)                     lastStatement.accept(this.declaredSyntheticLocalVariableVisitor);
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(84)                     list.clear();
--
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(86)                     this.statementCountToRemove = 0;
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(87)                 } else {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(88)                     List<Statement> list = statements;
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(89)                     <b>for</b> (Statement statement : list)
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(90)                         statement.accept(this.declaredSyntheticLocalVariableVisitor); 
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(91)                     list.clear();
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(92)                     if (i < size)
--
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(126)     }
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(127)     
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(128)     public void visit(SwitchStatement statement) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(129)         <b>for</b> (SwitchStatement.Block block : statement.getBlocks())
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(130)             block.getStatements().accept(this); 
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(131)     }
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(132)     
--
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(146)             removeFinallyStatements(tryStatements);
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(147)             this.statementCountToRemove = finallyStatementsSize;
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(148)             if (catchClauses != null)
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(149)                 <b>for</b> (TryStatement.CatchClause cc : catchClauses)
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(150)                     removeFinallyStatements((Statements)cc.getStatements());  
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(151)             this.statementCountInFinally = oldStatementCountInFinally;
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(152)             if (statement.getResources() != null)
--
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(160)             this.statementCountToRemove += finallyStatementsSize;
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(161)             removeFinallyStatements(tryStatements);
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(162)             if (catchClauses != null)
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(163)                 <b>for</b> (TryStatement.CatchClause cc : catchClauses)
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(164)                     removeFinallyStatements((Statements)cc.getStatements());  
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(165)             this.statementCountInFinally = oldStatementCountInFinally;
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(166)             this.statementCountToRemove = oldStatementCountToRemove;
--
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(256)     
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(257)     protected void safeAcceptListStatement(List<? extends Statement> list) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(258)         if (list != null)
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(259)             <b>for</b> (Statement statement : list)
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(260)                 statement.accept(this);  
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(261)     }
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/RemoveFinallyStatementsVisitor.class:(262) }
--
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/SearchFirstLineNumberVisitor.class:(65)     public void visit(Statements statements) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/SearchFirstLineNumberVisitor.class:(66)         if (this.lineNumber == -1) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/SearchFirstLineNumberVisitor.class:(67)             List<Statement> list = statements;
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/SearchFirstLineNumberVisitor.class:(68)             <b>for</b> (Statement statement : list) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/SearchFirstLineNumberVisitor.class:(69)                 statement.accept(this);
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/SearchFirstLineNumberVisitor.class:(70)                 if (this.lineNumber != -1)
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/SearchFirstLineNumberVisitor.class:(71)                     break; 
--
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/UpdateClassTypeArgumentsVisitor.class:(93)     public void visit(TypeArguments arguments) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/UpdateClassTypeArgumentsVisitor.class:(94)         int size = arguments.size();
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/UpdateClassTypeArgumentsVisitor.class:(95)         int i;
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/UpdateClassTypeArgumentsVisitor.class:(96)         <b>for</b> (i = 0; i < size; i++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/UpdateClassTypeArgumentsVisitor.class:(97)             TypeArgument ta = arguments.get(i);
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/UpdateClassTypeArgumentsVisitor.class:(98)             ta.accept(this);
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/UpdateClassTypeArgumentsVisitor.class:(99)             if (this.result != ta)
--
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/UpdateClassTypeArgumentsVisitor.class:(106)                 TypeArguments newTypes = new TypeArguments(size);
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/UpdateClassTypeArgumentsVisitor.class:(107)                 newTypes.addAll(arguments.subList(0, i));
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/UpdateClassTypeArgumentsVisitor.class:(108)                 newTypes.add((TypeArgument)this.result);
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/UpdateClassTypeArgumentsVisitor.class:(109)                 <b>for</b> (; ++i < size; i++) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/UpdateClassTypeArgumentsVisitor.class:(110)                     TypeArgument ta = arguments.get(i);
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/UpdateClassTypeArgumentsVisitor.class:(111)                     ta.accept(this);
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/UpdateClassTypeArgumentsVisitor.class:(112)                     newTypes.add((TypeArgument)this.result);
--
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/UpdateIntegerConstantTypeVisitor.class:(299)         if (expressions.isList()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/UpdateIntegerConstantTypeVisitor.class:(300)             DefaultList<Type> t = types.getList();
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/UpdateIntegerConstantTypeVisitor.class:(301)             DefaultList<Expression> e = expressions.getList();
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/UpdateIntegerConstantTypeVisitor.class:(302)             <b>for</b> (int i = e.size() - 1; i >= 0; i--) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/UpdateIntegerConstantTypeVisitor.class:(303)                 Type type = t.get(i);
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/UpdateIntegerConstantTypeVisitor.class:(304)                 if (type.getDimension() == 0 && type.isPrimitive()) {
org/jd/core/v1/service/converter/classfiletojavasyntax/visitor/UpdateIntegerConstantTypeVisitor.class:(305)                     Expression parameter = e.get(i);
--
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(79)         if (aic != null) {
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(80)             DefaultList<ClassFile> innerClassFiles = new DefaultList<>();
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(81)             String innerTypePrefix = internalTypeName + '$';
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(82)             <b>for</b> (InnerClass ic : aic.getInnerClasses()) {
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(83)                 if (!internalTypeName.equals(ic.getInnerTypeName()) && (
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(84)                     internalTypeName.equals(ic.getOuterTypeName()) || ic.getInnerTypeName().startsWith(innerTypePrefix))) {
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(85)                     int length;
--
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(129)         if (count == 0)
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(130)             return null; 
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(131)         Constant[] constants = new Constant[count];
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(132)         <b>for</b> (int i = 1; i < count; i++) {
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(133)             int tag = reader.readByte();
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(134)             switch (tag) {
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(135)                 case 1:
--
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(183)         if (count == 0)
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(184)             return null; 
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(185)         String[] interfaceTypeNames = new String[count];
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(186)         <b>for</b> (int i = 0; i < count; i++) {
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(187)             int index = reader.readUnsignedShort();
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(188)             interfaceTypeNames[i] = constants.getConstantTypeName(index);
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(189)         } 
--
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(195)         if (count == 0)
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(196)             return null; 
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(197)         Field[] fields = new Field[count];
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(198)         <b>for</b> (int i = 0; i < count; i++) {
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(199)             int accessFlags = reader.readUnsignedShort();
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(200)             int nameIndex = reader.readUnsignedShort();
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(201)             int descriptorIndex = reader.readUnsignedShort();
--
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(212)         if (count == 0)
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(213)             return null; 
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(214)         Method[] methods = new Method[count];
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(215)         <b>for</b> (int i = 0; i < count; i++) {
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(216)             int accessFlags = reader.readUnsignedShort();
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(217)             int nameIndex = reader.readUnsignedShort();
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(218)             int descriptorIndex = reader.readUnsignedShort();
--
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(229)         if (count == 0)
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(230)             return null; 
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(231)         HashMap<String, Attribute> attributes = new HashMap<>();
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(232)         <b>for</b> (int i = 0; i < count; i++) {
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(233)             int attributeNameIndex = reader.readUnsignedShort();
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(234)             int attributeLength = reader.readInt();
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(235)             Constant constant = constants.getConstant(attributeNameIndex);
--
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(381)         if (count == 0)
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(382)             return null; 
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(383)         ElementValuePair[] pairs = new ElementValuePair[count];
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(384)         <b>for</b> (int i = 0; i < count; i++) {
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(385)             int elementNameIndex = reader.readUnsignedShort();
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(386)             String elementName = constants.getConstantUtf8(elementNameIndex);
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(387)             pairs[i] = new ElementValuePair(elementName, loadElementValue(reader, constants));
--
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(394)         if (count == 0)
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(395)             return null; 
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(396)         ElementValue[] values = new ElementValue[count];
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(397)         <b>for</b> (int i = 0; i < count; i++)
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(398)             values[i] = loadElementValue(reader, constants); 
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(399)         return values;
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(400)     }
--
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(404)         if (count == 0)
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(405)             return null; 
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(406)         BootstrapMethod[] values = new BootstrapMethod[count];
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(407)         <b>for</b> (int i = 0; i < count; i++) {
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(408)             int bootstrapArguments[], bootstrapMethodRef = reader.readUnsignedShort();
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(409)             int numBootstrapArguments = reader.readUnsignedShort();
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(410)             if (numBootstrapArguments == 0) {
--
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(411)                 bootstrapArguments = EMPTY_INT_ARRAY;
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(412)             } else {
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(413)                 bootstrapArguments = new int[numBootstrapArguments];
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(414)                 <b>for</b> (int j = 0; j < numBootstrapArguments; j++)
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(415)                     bootstrapArguments[j] = reader.readUnsignedShort(); 
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(416)             } 
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(417)             values[i] = new BootstrapMethod(bootstrapMethodRef, bootstrapArguments);
--
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(433)         if (count == 0)
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(434)             return null; 
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(435)         CodeException[] codeExceptions = new CodeException[count];
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(436)         <b>for</b> (int i = 0; i < count; i++)
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(437)             codeExceptions[i] = new CodeException(i, reader
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(438)                     .readUnsignedShort(), reader
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(439)                     .readUnsignedShort(), reader
--
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(452)         if (count == 0)
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(453)             return null; 
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(454)         String[] exceptionTypeNames = new String[count];
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(455)         <b>for</b> (int i = 0; i < count; i++) {
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(456)             int exceptionClassIndex = reader.readUnsignedShort();
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(457)             exceptionTypeNames[i] = constants.getConstantTypeName(exceptionClassIndex);
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(458)         } 
--
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(464)         if (count == 0)
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(465)             return null; 
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(466)         InnerClass[] innerClasses = new InnerClass[count];
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(467)         <b>for</b> (int i = 0; i < count; i++) {
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(468)             int innerTypeIndex = reader.readUnsignedShort();
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(469)             int outerTypeIndex = reader.readUnsignedShort();
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(470)             int innerNameIndex = reader.readUnsignedShort();
--
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(482)         if (count == 0)
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(483)             return null; 
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(484)         LocalVariable[] localVariables = new LocalVariable[count];
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(485)         <b>for</b> (int i = 0; i < count; i++) {
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(486)             int startPc = reader.readUnsignedShort();
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(487)             int length = reader.readUnsignedShort();
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(488)             int nameIndex = reader.readUnsignedShort();
--
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(500)         if (count == 0)
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(501)             return null; 
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(502)         LocalVariableType[] localVariables = new LocalVariableType[count];
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(503)         <b>for</b> (int i = 0; i < count; i++) {
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(504)             int startPc = reader.readUnsignedShort();
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(505)             int length = reader.readUnsignedShort();
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(506)             int nameIndex = reader.readUnsignedShort();
--
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(518)         if (count == 0)
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(519)             return null; 
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(520)         LineNumber[] lineNumbers = new LineNumber[count];
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(521)         <b>for</b> (int i = 0; i < count; i++)
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(522)             lineNumbers[i] = new LineNumber(reader.readUnsignedShort(), reader.readUnsignedShort()); 
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(523)         return lineNumbers;
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(524)     }
--
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(528)         if (count == 0)
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(529)             return null; 
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(530)         MethodParameter[] parameters = new MethodParameter[count];
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(531)         <b>for</b> (int i = 0; i < count; i++) {
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(532)             int nameIndex = reader.readUnsignedShort();
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(533)             String name = constants.getConstantUtf8(nameIndex);
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(534)             parameters[i] = new MethodParameter(name, reader.readUnsignedShort());
--
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(541)         if (count == 0)
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(542)             return null; 
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(543)         ModuleInfo[] moduleInfos = new ModuleInfo[count];
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(544)         <b>for</b> (int i = 0; i < count; i++) {
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(545)             int moduleInfoIndex = reader.readUnsignedShort();
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(546)             int moduleFlag = reader.readUnsignedShort();
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(547)             int moduleVersionIndex = reader.readUnsignedShort();
--
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(557)         if (count == 0)
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(558)             return null; 
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(559)         PackageInfo[] packageInfos = new PackageInfo[count];
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(560)         <b>for</b> (int i = 0; i < count; i++) {
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(561)             int packageInfoIndex = reader.readUnsignedShort();
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(562)             int packageFlag = reader.readUnsignedShort();
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(563)             String packageInfoName = constants.getConstantTypeName(packageInfoIndex);
--
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(571)         if (count == 0)
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(572)             return null; 
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(573)         String[] names = new String[count];
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(574)         <b>for</b> (int i = 0; i < count; i++)
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(575)             names[i] = constants.getConstantTypeName(reader.readUnsignedShort()); 
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(576)         return names;
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(577)     }
--
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(581)         if (count == 0)
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(582)             return null; 
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(583)         ServiceInfo[] services = new ServiceInfo[count];
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(584)         <b>for</b> (int i = 0; i < count; i++)
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(585)             services[i] = new ServiceInfo(constants
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(586)                     .getConstantTypeName(reader.readUnsignedShort()), 
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(587)                     loadConstantClassNames(reader, constants)); 
--
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(593)         if (count == 0)
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(594)             return null; 
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(595)         Annotation[] annotations = new Annotation[count];
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(596)         <b>for</b> (int i = 0; i < count; i++) {
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(597)             int descriptorIndex = reader.readUnsignedShort();
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(598)             String descriptor = constants.getConstantUtf8(descriptorIndex);
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(599)             annotations[i] = new Annotation(descriptor, loadElementValuePairs(reader, constants));
--
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(606)         if (count == 0)
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(607)             return null; 
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(608)         Annotations[] parameterAnnotations = new Annotations[count];
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(609)         <b>for</b> (int i = 0; i < count; i++) {
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(610)             Annotation[] annotations = loadAnnotations(reader, constants);
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(611)             if (annotations != null)
org/jd/core/v1/service/deserializer/classfile/ClassFileDeserializer.class:(612)                 parameterAnnotations[i] = new Annotations(annotations); 
--
org/jd/core/v1/service/deserializer/classfile/ClassFileReader.class:(73)                 case 12:
org/jd/core/v1/service/deserializer/classfile/ClassFileReader.class:(74)                 case 13:
org/jd/core/v1/service/deserializer/classfile/ClassFileReader.class:(75)                     if (this.offset + 1 > maxOffset)
org/jd/core/v1/service/deserializer/classfile/ClassFileReader.class:(76)                         throw new UTFDataFormatException("mal<b>for</b>med input: partial character at end"); 
org/jd/core/v1/service/deserializer/classfile/ClassFileReader.class:(77)                     char2 = this.data[this.offset++];
org/jd/core/v1/service/deserializer/classfile/ClassFileReader.class:(78)                     if ((char2 & 0xC0) != 128)
org/jd/core/v1/service/deserializer/classfile/ClassFileReader.class:(79)                         throw new UTFDataFormatException("mal<b>for</b>med input around byte " + this.offset); 
org/jd/core/v1/service/deserializer/classfile/ClassFileReader.class:(80)                     charArray[charArrayOffset++] = (char)((c & 0x1F) << 6 | char2 & 0x3F);
org/jd/core/v1/service/deserializer/classfile/ClassFileReader.class:(81)                     continue;
org/jd/core/v1/service/deserializer/classfile/ClassFileReader.class:(82)                 case 14:
--
org/jd/core/v1/service/deserializer/classfile/ClassFileReader.class:(83)                     if (this.offset + 2 > maxOffset)
org/jd/core/v1/service/deserializer/classfile/ClassFileReader.class:(84)                         throw new UTFDataFormatException("mal<b>for</b>med input: partial character at end"); 
org/jd/core/v1/service/deserializer/classfile/ClassFileReader.class:(85)                     char2 = this.data[this.offset++];
org/jd/core/v1/service/deserializer/classfile/ClassFileReader.class:(86)                     char3 = this.data[this.offset++];
org/jd/core/v1/service/deserializer/classfile/ClassFileReader.class:(87)                     if ((char2 & 0xC0) != 128 || (char3 & 0xC0) != 128)
--
org/jd/core/v1/service/deserializer/classfile/ClassFileReader.class:(88)                         throw new UTFDataFormatException("mal<b>for</b>med input around byte " + (this.offset - 1)); 
org/jd/core/v1/service/deserializer/classfile/ClassFileReader.class:(89)                     charArray[charArrayOffset++] = (char)((c & 0xF) << 12 | (char2 & 0x3F) << 6 | (char3 & 0x3F) << 0);
org/jd/core/v1/service/deserializer/classfile/ClassFileReader.class:(90)                     continue;
org/jd/core/v1/service/deserializer/classfile/ClassFileReader.class:(91)             } 
--
org/jd/core/v1/service/deserializer/classfile/ClassFileReader.class:(92)             throw new UTFDataFormatException("mal<b>for</b>med input around byte " + this.offset);
org/jd/core/v1/service/deserializer/classfile/ClassFileReader.class:(93)         } 
org/jd/core/v1/service/deserializer/classfile/ClassFileReader.class:(94)         return new String(charArray, 0, charArrayOffset);
org/jd/core/v1/service/deserializer/classfile/ClassFileReader.class:(95)     }
--
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/util/JavaFragmentFactory.class:(28)         fragments.add(new SpacerFragment(0, 1, 1, 1, "Spacer after imports"));
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/util/JavaFragmentFactory.class:(29)     }
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/util/JavaFragmentFactory.class:(30)     
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/util/JavaFragmentFactory.class:(31)     public static void addSpacerBe<b>for</b>eMainDeclaration(List<Fragment> fragments) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/util/JavaFragmentFactory.class:(32)         fragments.add(new SpacerFragment(0, 0, 2147483647, 5, "Spacer be<b>for</b>e main declaration"));
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/util/JavaFragmentFactory.class:(33)     }
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/util/JavaFragmentFactory.class:(34)     
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/util/JavaFragmentFactory.class:(35)     public static void addEndArrayInitializerInParameter(List<Fragment> fragments, StartBlockFragment start) {
--
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/util/JavaFragmentFactory.class:(94)         fragments.add(new SpaceSpacerFragment(0, 1, 1, 16, "Spacer between switch label"));
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/util/JavaFragmentFactory.class:(95)     }
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/util/JavaFragmentFactory.class:(96)     
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/util/JavaFragmentFactory.class:(97)     public static void addSpacerBe<b>for</b>eExtends(List<Fragment> fragments) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/util/JavaFragmentFactory.class:(98)         fragments.add(new SpaceSpacerFragment(0, 0, 1, 2, "Spacer be<b>for</b>e extends"));
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/util/JavaFragmentFactory.class:(99)     }
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/util/JavaFragmentFactory.class:(100)     
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/util/JavaFragmentFactory.class:(101)     public static void addSpacerBe<b>for</b>eImplements(List<Fragment> fragments) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/util/JavaFragmentFactory.class:(102)         fragments.add(new SpaceSpacerFragment(0, 0, 1, 2, "Spacer be<b>for</b>e implements"));
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/util/JavaFragmentFactory.class:(103)     }
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/util/JavaFragmentFactory.class:(104)     
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/util/JavaFragmentFactory.class:(105)     public static void addSpacerBetweenEnumValues(List<Fragment> fragments, int preferredLineCount) {
--
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/util/StringUtil.class:(1) public class StringUtil {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/util/StringUtil.class:(2)     public static String escapeString(String s) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/util/StringUtil.class:(3)         int length = s.length();
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/util/StringUtil.class:(4)         <b>for</b> (int i = 0; i < length; i++) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/util/StringUtil.class:(5)             char c = s.charAt(i);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/util/StringUtil.class:(6)             if (c == '\\' || c == '"' || c < ' ') {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/util/StringUtil.class:(7)                 StringBuilder sb = new StringBuilder(length * 2);
--
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/util/StringUtil.class:(8)                 sb.append(s.substring(0, i));
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/util/StringUtil.class:(9)                 <b>for</b> (; i < length; i++) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/util/StringUtil.class:(10)                     c = s.charAt(i);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/util/StringUtil.class:(11)                     switch (c) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/util/StringUtil.class:(12)                         case '\\':
--
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(180)         if (size > 0) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(181)             Iterator<AnnotationReference> iterator = list.iterator();
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(182)             ((AnnotationReference)iterator.next()).accept(this);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(183)             <b>for</b> (int i = 1; i < size; i++) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(184)                 this.tokens.add(TextToken.SPACE);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(185)                 ((AnnotationReference)iterator.next()).accept(this);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(186)             } 
--
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(196)                 JavaFragmentFactory.addNewLineBetweenArrayInitializerBlock(this.fragments); 
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(197)             this.tokens = new TypeVisitor.Tokens(this);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(198)             declaration.get(0).accept(this);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(199)             <b>for</b> (int i = 1; i < size; i++) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(200)                 if (this.tokens.isEmpty()) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(201)                     JavaFragmentFactory.addSpacerBetweenArrayInitializerBlock(this.fragments);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(202)                     if (size > 10 && i % 10 == 0)
--
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(235)             BaseType superType = declaration.getSuperType();
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(236)             if (superType != null && !superType.equals(ObjectType.TYPE_OBJECT)) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(237)                 this.fragments.addTokensFragment(this.tokens);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(238)                 JavaFragmentFactory.addSpacerBe<b>for</b>eExtends(this.fragments);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(239)                 this.tokens = new TypeVisitor.Tokens(this);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(240)                 this.tokens.add(EXTENDS);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(241)                 this.tokens.add(TextToken.SPACE);
--
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(247)             if (interfaces != null) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(248)                 if (!this.tokens.isEmpty())
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(249)                     this.fragments.addTokensFragment(this.tokens); 
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(250)                 JavaFragmentFactory.addSpacerBe<b>for</b>eImplements(this.fragments);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(251)                 this.tokens = new TypeVisitor.Tokens(this);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(252)                 this.tokens.add(IMPLEMENTS);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(253)                 this.tokens.add(TextToken.SPACE);
--
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(301)             if (!this.fragments.isEmpty())
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(302)                 JavaFragmentFactory.addSpacerAfterImports(this.fragments); 
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(303)         } 
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(304)         JavaFragmentFactory.addSpacerBe<b>for</b>eMainDeclaration(this.fragments);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(305)         super.visit(compilationUnit);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(306)     }
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(307)     
--
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(327)             this.tokens.add(new DeclarationToken(4, this.currentInternalTypeName, this.currentTypeName, declaration.getDescriptor()));
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(328)             storeContext();
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(329)             this.currentMethodParamNames.clear();
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(330)             BaseFormalParameter <b>for</b>malParameters = declaration.getFormalParameters();
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(331)             if (<b>for</b>malParameters == null) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(332)                 this.tokens.add(TextToken.LEFTRIGHTROUNDBRACKETS);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(333)             } else {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(334)                 this.tokens.add(StartBlockToken.START_PARAMETERS_BLOCK);
--
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(335)                 this.fragments.addTokensFragment(this.tokens);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(336)                 <b>for</b>malParameters.accept(this);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(337)                 this.tokens = new TypeVisitor.Tokens(this);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(338)                 this.tokens.add(EndBlockToken.END_PARAMETERS_BLOCK);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(339)             } 
--
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(419)             if (interfaces != null) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(420)                 this.tokens.add(StartBlockToken.START_DECLARATION_OR_STATEMENT_BLOCK);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(421)                 this.fragments.addTokensFragment(this.tokens);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(422)                 JavaFragmentFactory.addSpacerBe<b>for</b>eImplements(this.fragments);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(423)                 this.tokens = new TypeVisitor.Tokens(this);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(424)                 this.tokens.add(IMPLEMENTS);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(425)                 this.tokens.add(TextToken.SPACE);
--
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(436)             List<EnumDeclaration.Constant> constants = declaration.getConstants();
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(437)             if (constants != null && !constants.isEmpty()) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(438)                 int preferredLineNumber = 0;
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(439)                 <b>for</b> (EnumDeclaration.Constant constant : constants) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(440)                     if (constant.getArguments() != null || constant.getBodyDeclaration() != null) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(441)                         preferredLineNumber = 1;
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(442)                         break;
--
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(445)                 this.fragments.add(StartMovableJavaBlockFragment.START_MOVABLE_FIELD_BLOCK);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(446)                 ((EnumDeclaration.Constant)constants.get(0)).accept(this);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(447)                 int size = constants.size();
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(448)                 <b>for</b> (int i = 1; i < size; i++) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(449)                     JavaFragmentFactory.addSpacerBetweenEnumValues(this.fragments, preferredLineNumber);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(450)                     ((EnumDeclaration.Constant)constants.get(i)).accept(this);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(451)                 } 
--
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(561)         if (size > 0) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(562)             Iterator<FieldDeclarator> iterator = declarators.iterator();
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(563)             ((FieldDeclarator)iterator.next()).accept(this);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(564)             <b>for</b> (int i = 1; i < size; i++) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(565)                 JavaFragmentFactory.addSpacerBetweenFieldDeclarators(this.fragments);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(566)                 ((FieldDeclarator)iterator.next()).accept(this);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(567)             } 
--
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(598)         if (size > 0) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(599)             Iterator<FormalParameter> iterator = declarations.iterator();
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(600)             ((FormalParameter)iterator.next()).accept(this);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(601)             <b>for</b> (int i = 1; i < size; i++) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(602)                 this.tokens.add(TextToken.COMMA_SPACE);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(603)                 ((FormalParameter)iterator.next()).accept(this);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(604)             } 
--
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(627)             BaseType interfaces = declaration.getInterfaces();
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(628)             if (interfaces != null) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(629)                 this.fragments.addTokensFragment(this.tokens);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(630)                 JavaFragmentFactory.addSpacerBe<b>for</b>eImplements(this.fragments);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(631)                 this.tokens = new TypeVisitor.Tokens(this);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(632)                 this.tokens.add(EXTENDS);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(633)                 this.tokens.add(TextToken.SPACE);
--
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(839)         if (size > 0) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(840)             Iterator<LocalVariableDeclarator> iterator = declarators.iterator();
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(841)             ((LocalVariableDeclarator)iterator.next()).accept(this);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(842)             <b>for</b> (int i = 1; i < size; i++) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(843)                 this.tokens.add(TextToken.COMMA_SPACE);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(844)                 ((LocalVariableDeclarator)iterator.next()).accept(this);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(845)             } 
--
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(854)             ((MemberDeclaration)iterator.next()).accept(this);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(855)             if (size > 1) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(856)                 int fragmentCount1 = -1;
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(857)                 <b>for</b> (int i = 1; i < size; i++) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(858)                     if (fragmentCount2 < this.fragments.size()) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(859)                         fragmentCount1 = this.fragments.size();
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(860)                         JavaFragmentFactory.addSpacerBetweenMembers(this.fragments);
--
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(893)             this.tokens.add(new DeclarationToken(3, this.currentInternalTypeName, declaration.getName(), declaration.getDescriptor()));
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(894)             storeContext();
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(895)             this.currentMethodParamNames.clear();
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(896)             BaseFormalParameter <b>for</b>malParameters = declaration.getFormalParameters();
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(897)             if (<b>for</b>malParameters == null) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(898)                 this.tokens.add(TextToken.LEFTRIGHTROUNDBRACKETS);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(899)             } else {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(900)                 this.tokens.add(StartBlockToken.START_PARAMETERS_BLOCK);
--
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(901)                 this.fragments.addTokensFragment(this.tokens);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(902)                 <b>for</b>malParameters.accept(this);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(903)                 this.tokens = new TypeVisitor.Tokens(this);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(904)                 this.tokens.add(EndBlockToken.END_PARAMETERS_BLOCK);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/CompilationUnitVisitor.class:(905)             } 
--
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/ExpressionVisitor.class:(199)             int size = list.size();
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/ExpressionVisitor.class:(200)             if (size > 0) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/ExpressionVisitor.class:(201)                 Iterator<Expression> iterator = list.iterator();
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/ExpressionVisitor.class:(202)                 <b>for</b> (int i = size - 1; i > 0; i--) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/ExpressionVisitor.class:(203)                     this.inExpressionFlag = true;
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/ExpressionVisitor.class:(204)                     ((Expression)iterator.next()).accept(this);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/ExpressionVisitor.class:(205)                     if (!this.tokens.isEmpty())
--
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/ExpressionVisitor.class:(300)                 default:
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/ExpressionVisitor.class:(301)                     this.tokens.add(TextToken.LEFTROUNDBRACKET);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/ExpressionVisitor.class:(302)                     this.tokens.add(newTextToken(parameters.get(0)));
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/ExpressionVisitor.class:(303)                     <b>for</b> (i = 1; i < size; i++) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/ExpressionVisitor.class:(304)                         this.tokens.add(TextToken.COMMA_SPACE);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/ExpressionVisitor.class:(305)                         this.tokens.add(newTextToken(parameters.get(i)));
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/ExpressionVisitor.class:(306)                     } 
--
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/StatementVisitor.class:(67)     
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/StatementVisitor.class:(68)     public static final KeywordToken FINALLY = new KeywordToken("finally");
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/StatementVisitor.class:(69)     
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/StatementVisitor.class:(70)     public static final KeywordToken FOR = new KeywordToken("<b>for</b>");
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/StatementVisitor.class:(71)     
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/StatementVisitor.class:(72)     public static final KeywordToken IF = new KeywordToken("if");
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/StatementVisitor.class:(73)     
--
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/StatementVisitor.class:(392)         if (size > 0) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/StatementVisitor.class:(393)             Iterator<Statement> iterator = list.iterator();
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/StatementVisitor.class:(394)             ((Statement)iterator.next()).accept(this);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/StatementVisitor.class:(395)             <b>for</b> (int i = 1; i < size; i++) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/StatementVisitor.class:(396)                 JavaFragmentFactory.addSpacerBetweenStatements(this.fragments);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/StatementVisitor.class:(397)                 ((Statement)iterator.next()).accept(this);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/StatementVisitor.class:(398)             } 
--
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/StatementVisitor.class:(502)                 this.tokens.add(TextToken.SPACE); 
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/StatementVisitor.class:(503)             this.tokens.add(StartBlockToken.START_RESOURCES_BLOCK);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/StatementVisitor.class:(504)             ((TryStatement.Resource)resources.get(0)).accept(this);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/StatementVisitor.class:(505)             <b>for</b> (int i = 1; i < size; i++) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/StatementVisitor.class:(506)                 this.tokens.add(TextToken.SEMICOLON_SPACE);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/StatementVisitor.class:(507)                 ((TryStatement.Resource)resources.get(i)).accept(this);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/StatementVisitor.class:(508)             } 
--
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/StatementVisitor.class:(528)         int fragmentCount1 = this.fragments.size(), fragmentCount2 = fragmentCount1;
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/StatementVisitor.class:(529)         statement.getTryStatements().accept(this);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/StatementVisitor.class:(530)         if (statement.getCatchClauses() != null)
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/StatementVisitor.class:(531)             <b>for</b> (TryStatement.CatchClause cc : statement.getCatchClauses()) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/StatementVisitor.class:(532)                 JavaFragmentFactory.addEndStatementsBlock(this.fragments, group);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/StatementVisitor.class:(533)                 ObjectType type = cc.getType();
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/StatementVisitor.class:(534)                 this.tokens = new TypeVisitor.Tokens(this);
--
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/StatementVisitor.class:(536)                 this.tokens.add(TextToken.SPACE_LEFTROUNDBRACKET);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/StatementVisitor.class:(537)                 this.tokens.add(newTypeReferenceToken(type, this.currentInternalTypeName));
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/StatementVisitor.class:(538)                 if (cc.getOtherTypes() != null)
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/StatementVisitor.class:(539)                     <b>for</b> (ObjectType otherType : cc.getOtherTypes()) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/StatementVisitor.class:(540)                         this.tokens.add(TextToken.VERTICALLINE);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/StatementVisitor.class:(541)                         this.tokens.add(newTypeReferenceToken(otherType, this.currentInternalTypeName));
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/StatementVisitor.class:(542)                     }  
--
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/TypeVisitor.class:(215)         int size = parameters.size();
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/TypeVisitor.class:(216)         if (size > 0) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/TypeVisitor.class:(217)             parameters.get(0).accept(this);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/TypeVisitor.class:(218)             <b>for</b> (int i = 1; i < size; i++) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/TypeVisitor.class:(219)                 this.tokens.add(TextToken.COMMA_SPACE);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/TypeVisitor.class:(220)                 parameters.get(i).accept(this);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/TypeVisitor.class:(221)             } 
--
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/TypeVisitor.class:(235)         int size = list.size();
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/TypeVisitor.class:(236)         if (size > 0) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/TypeVisitor.class:(237)             ((TypeArgumentVisitable)list.get(0)).accept(this);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/TypeVisitor.class:(238)             <b>for</b> (int i = 1; i < size; i++) {
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/TypeVisitor.class:(239)                 this.tokens.add(separator);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/TypeVisitor.class:(240)                 ((TypeArgumentVisitable)list.get(i)).accept(this);
org/jd/core/v1/service/fragmenter/javasyntaxtojavafragment/visitor/TypeVisitor.class:(241)             } 
--
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(22)         List<Fragment> fragments = message.<List<Fragment>>getBody();
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(23)         if (maxLineNumber != 0 && !containsByteCode && !showBridgeAndSynthetic && realignLineNumbers) {
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(24)             BuildSectionsVisitor buildSectionsVisitor = new BuildSectionsVisitor();
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(25)             <b>for</b> (Fragment fragment : fragments)
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(26)                 fragment.accept(buildSectionsVisitor); 
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(27)             List<Section> sections = buildSectionsVisitor.getSections();
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(28)             VisitorsHolder holder = new VisitorsHolder();
--
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(31)             int max = sections.size() * 2;
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(32)             if (max > 20)
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(33)                 max = 20; 
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(34)             <b>for</b> (int loop = 0; loop < max; ) {
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(35)                 visitor.reset();
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(36)                 <b>for</b> (Section section : sections) {
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(37)                     <b>for</b> (FlexibleFragment fragment : section.getFlexibleFragments())
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(38)                         fragment.accept(visitor); 
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(39)                     if (section.getFixedFragment() != null)
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(40)                         section.getFixedFragment().accept(visitor); 
--
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(41)                 } 
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(42)                 <b>for</b> (int redo = 0; redo < 10; redo++) {
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(43)                     boolean changed = false;
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(44)                     <b>for</b> (Section section : sections)
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(45)                         changed |= section.layout(false); 
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(46)                     if (!changed)
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(47)                         break; 
--
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(48)                 } 
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(49)                 int newSumOfRates = 0;
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(50)                 Section mostConstrainedSection = sections.get(0);
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(51)                 <b>for</b> (Section section : sections) {
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(52)                     section.updateRate();
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(53)                     if (mostConstrainedSection.getRate() < section.getRate())
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(54)                         mostConstrainedSection = section; 
--
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(63)                     loop++;
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(64)                 } 
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(65)             } 
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(66)             <b>for</b> (Section section : sections)
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(67)                 section.layout(true); 
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(68)             fragments.clear();
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(69)             <b>for</b> (Section section : sections) {
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(70)                 fragments.addAll((Collection)section.getFlexibleFragments());
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(71)                 FixedFragment fixedFragment = section.getFixedFragment();
org/jd/core/v1/service/layouter/LayoutFragmentProcessor.class:(72)                 if (fixedFragment != null)
--
org/jd/core/v1/service/layouter/model/Section.class:(68)     
org/jd/core/v1/service/layouter/model/Section.class:(69)     public void updateRate() {
org/jd/core/v1/service/layouter/model/Section.class:(70)         this.rate = 0;
org/jd/core/v1/service/layouter/model/Section.class:(71)         <b>for</b> (FlexibleFragment flexibleFragment : this.flexibleFragments) {
org/jd/core/v1/service/layouter/model/Section.class:(72)             if (flexibleFragment.getInitialLineCount() > flexibleFragment.getLineCount())
org/jd/core/v1/service/layouter/model/Section.class:(73)                 this.rate += flexibleFragment.getInitialLineCount() - flexibleFragment.getLineCount(); 
org/jd/core/v1/service/layouter/model/Section.class:(74)         } 
--
org/jd/core/v1/service/layouter/model/Section.class:(75)     }
org/jd/core/v1/service/layouter/model/Section.class:(76)     
org/jd/core/v1/service/layouter/model/Section.class:(77)     public boolean layout(boolean <b>for</b>ce) {
org/jd/core/v1/service/layouter/model/Section.class:(78)         if (this.fixedFragment != null) {
org/jd/core/v1/service/layouter/model/Section.class:(79)             int currentLineCount = 0;
org/jd/core/v1/service/layouter/model/Section.class:(80)             <b>for</b> (FlexibleFragment flexibleFragment : this.flexibleFragments)
org/jd/core/v1/service/layouter/model/Section.class:(81)                 currentLineCount += flexibleFragment.getLineCount(); 
org/jd/core/v1/service/layouter/model/Section.class:(82)             if (<b>for</b>ce || this.lastLineCount != currentLineCount) {
org/jd/core/v1/service/layouter/model/Section.class:(83)                 this.lastLineCount = currentLineCount;
org/jd/core/v1/service/layouter/model/Section.class:(84)                 if (this.targetLineCount != currentLineCount) {
org/jd/core/v1/service/layouter/model/Section.class:(85)                     AutoGrowthList filteredFlexibleFragments = new AutoGrowthList(this);
--
org/jd/core/v1/service/layouter/model/Section.class:(86)                     DefaultList<FlexibleFragment> constrainedFlexibleFragments = new DefaultList<>(this.flexibleFragments.size());
org/jd/core/v1/service/layouter/model/Section.class:(87)                     if (this.targetLineCount > currentLineCount) {
org/jd/core/v1/service/layouter/model/Section.class:(88)                         int i = this.delta = this.targetLineCount - currentLineCount;
org/jd/core/v1/service/layouter/model/Section.class:(89)                         <b>for</b> (FlexibleFragment flexibleFragment : this.flexibleFragments) {
org/jd/core/v1/service/layouter/model/Section.class:(90)                             if (flexibleFragment.getLineCount() < flexibleFragment.getMaximalLineCount())
org/jd/core/v1/service/layouter/model/Section.class:(91)                                 filteredFlexibleFragments.get(flexibleFragment.getWeight()).add(flexibleFragment); 
org/jd/core/v1/service/layouter/model/Section.class:(92)                         } 
--
org/jd/core/v1/service/layouter/model/Section.class:(93)                         <b>for</b> (DefaultList<FlexibleFragment> flexibleFragments : (Iterable<DefaultList<FlexibleFragment>>)filteredFlexibleFragments) {
org/jd/core/v1/service/layouter/model/Section.class:(94)                             constrainedFlexibleFragments.clear();
org/jd/core/v1/service/layouter/model/Section.class:(95)                             <b>for</b> (FlexibleFragment flexibleFragment : flexibleFragments) {
org/jd/core/v1/service/layouter/model/Section.class:(96)                                 if (flexibleFragment.getLineCount() < flexibleFragment.getInitialLineCount())
org/jd/core/v1/service/layouter/model/Section.class:(97)                                     constrainedFlexibleFragments.add(flexibleFragment); 
org/jd/core/v1/service/layouter/model/Section.class:(98)                             } 
--
org/jd/core/v1/service/layouter/model/Section.class:(99)                             expand(constrainedFlexibleFragments, <b>for</b>ce);
org/jd/core/v1/service/layouter/model/Section.class:(100)                             if (this.delta == 0)
org/jd/core/v1/service/layouter/model/Section.class:(101)                                 break; 
org/jd/core/v1/service/layouter/model/Section.class:(102)                         } 
--
org/jd/core/v1/service/layouter/model/Section.class:(103)                         if (this.delta > 0)
org/jd/core/v1/service/layouter/model/Section.class:(104)                             <b>for</b> (DefaultList<FlexibleFragment> flexibleFragments : (Iterable<DefaultList<FlexibleFragment>>)filteredFlexibleFragments) {
org/jd/core/v1/service/layouter/model/Section.class:(105)                                 expand(flexibleFragments, <b>for</b>ce);
org/jd/core/v1/service/layouter/model/Section.class:(106)                                 if (this.delta == 0)
org/jd/core/v1/service/layouter/model/Section.class:(107)                                     break; 
org/jd/core/v1/service/layouter/model/Section.class:(108)                             }  
--
org/jd/core/v1/service/layouter/model/Section.class:(109)                         return (i != this.delta);
org/jd/core/v1/service/layouter/model/Section.class:(110)                     } 
org/jd/core/v1/service/layouter/model/Section.class:(111)                     int oldDelta = this.delta = currentLineCount - this.targetLineCount;
org/jd/core/v1/service/layouter/model/Section.class:(112)                     <b>for</b> (FlexibleFragment flexibleFragment : this.flexibleFragments) {
org/jd/core/v1/service/layouter/model/Section.class:(113)                         if (flexibleFragment.getMinimalLineCount() < flexibleFragment.getLineCount())
org/jd/core/v1/service/layouter/model/Section.class:(114)                             filteredFlexibleFragments.get(flexibleFragment.getWeight()).add(flexibleFragment); 
org/jd/core/v1/service/layouter/model/Section.class:(115)                     } 
--
org/jd/core/v1/service/layouter/model/Section.class:(116)                     <b>for</b> (DefaultList<FlexibleFragment> flexibleFragments : (Iterable<DefaultList<FlexibleFragment>>)filteredFlexibleFragments) {
org/jd/core/v1/service/layouter/model/Section.class:(117)                         constrainedFlexibleFragments.clear();
org/jd/core/v1/service/layouter/model/Section.class:(118)                         <b>for</b> (FlexibleFragment flexibleFragment : flexibleFragments) {
org/jd/core/v1/service/layouter/model/Section.class:(119)                             if (flexibleFragment.getLineCount() > flexibleFragment.getInitialLineCount())
org/jd/core/v1/service/layouter/model/Section.class:(120)                                 constrainedFlexibleFragments.add(flexibleFragment); 
org/jd/core/v1/service/layouter/model/Section.class:(121)                         } 
--
org/jd/core/v1/service/layouter/model/Section.class:(122)                         compact(constrainedFlexibleFragments, <b>for</b>ce);
org/jd/core/v1/service/layouter/model/Section.class:(123)                         if (this.delta == 0)
org/jd/core/v1/service/layouter/model/Section.class:(124)                             break; 
org/jd/core/v1/service/layouter/model/Section.class:(125)                     } 
--
org/jd/core/v1/service/layouter/model/Section.class:(126)                     if (this.delta > 0)
org/jd/core/v1/service/layouter/model/Section.class:(127)                         <b>for</b> (DefaultList<FlexibleFragment> flexibleFragments : (Iterable<DefaultList<FlexibleFragment>>)filteredFlexibleFragments) {
org/jd/core/v1/service/layouter/model/Section.class:(128)                             compact(flexibleFragments, <b>for</b>ce);
org/jd/core/v1/service/layouter/model/Section.class:(129)                             if (this.delta == 0)
org/jd/core/v1/service/layouter/model/Section.class:(130)                                 break; 
org/jd/core/v1/service/layouter/model/Section.class:(131)                         }  
--
org/jd/core/v1/service/layouter/model/Section.class:(136)         return false;
org/jd/core/v1/service/layouter/model/Section.class:(137)     }
org/jd/core/v1/service/layouter/model/Section.class:(138)     
org/jd/core/v1/service/layouter/model/Section.class:(139)     protected void expand(DefaultList<FlexibleFragment> flexibleFragments, boolean <b>for</b>ce) {
org/jd/core/v1/service/layouter/model/Section.class:(140)         int oldDelta = Integer.MAX_VALUE;
org/jd/core/v1/service/layouter/model/Section.class:(141)         while (this.delta > 0 && this.delta < oldDelta) {
org/jd/core/v1/service/layouter/model/Section.class:(142)             oldDelta = this.delta;
--
org/jd/core/v1/service/layouter/model/Section.class:(143)             <b>for</b> (FlexibleFragment flexibleFragment : flexibleFragments) {
org/jd/core/v1/service/layouter/model/Section.class:(144)                 if (flexibleFragment.incLineCount(<b>for</b>ce) && 
org/jd/core/v1/service/layouter/model/Section.class:(145)                     --this.delta == 0)
org/jd/core/v1/service/layouter/model/Section.class:(146)                     break; 
org/jd/core/v1/service/layouter/model/Section.class:(147)             } 
--
org/jd/core/v1/service/layouter/model/Section.class:(148)         } 
org/jd/core/v1/service/layouter/model/Section.class:(149)     }
org/jd/core/v1/service/layouter/model/Section.class:(150)     
org/jd/core/v1/service/layouter/model/Section.class:(151)     protected void compact(DefaultList<FlexibleFragment> flexibleFragments, boolean <b>for</b>ce) {
org/jd/core/v1/service/layouter/model/Section.class:(152)         int oldDelta = Integer.MAX_VALUE;
org/jd/core/v1/service/layouter/model/Section.class:(153)         while (this.delta > 0 && this.delta < oldDelta) {
org/jd/core/v1/service/layouter/model/Section.class:(154)             oldDelta = this.delta;
--
org/jd/core/v1/service/layouter/model/Section.class:(155)             <b>for</b> (FlexibleFragment flexibleFragment : flexibleFragments) {
org/jd/core/v1/service/layouter/model/Section.class:(156)                 if (flexibleFragment.decLineCount(<b>for</b>ce) && 
org/jd/core/v1/service/layouter/model/Section.class:(157)                     --this.delta == 0)
org/jd/core/v1/service/layouter/model/Section.class:(158)                     break; 
org/jd/core/v1/service/layouter/model/Section.class:(159)             } 
--
org/jd/core/v1/service/layouter/model/Section.class:(163)     public boolean releaseConstraints(VisitorsHolder holder) {
org/jd/core/v1/service/layouter/model/Section.class:(164)         int flexibleCount = this.flexibleFragments.size();
org/jd/core/v1/service/layouter/model/Section.class:(165)         AbstractStoreMovableBlockFragmentIndexVisitorAbstract backwardSearchStartIndexesVisitor = holder.getBackwardSearchStartIndexesVisitor();
org/jd/core/v1/service/layouter/model/Section.class:(166)         AbstractStoreMovableBlockFragmentIndexVisitorAbstract <b>for</b>wardSearchEndIndexesVisitor = holder.getForwardSearchEndIndexesVisitor();
org/jd/core/v1/service/layouter/model/Section.class:(167)         AbstractSearchMovableBlockFragmentVisitor <b>for</b>wardSearchVisitor = holder.getForwardSearchVisitor();
org/jd/core/v1/service/layouter/model/Section.class:(168)         AbstractSearchMovableBlockFragmentVisitor backwardSearchVisitor = holder.getBackwardSearchVisitor();
org/jd/core/v1/service/layouter/model/Section.class:(169)         ListIterator<FlexibleFragment> iterator = this.flexibleFragments.listIterator(flexibleCount);
org/jd/core/v1/service/layouter/model/Section.class:(170)         backwardSearchStartIndexesVisitor.reset();
--
org/jd/core/v1/service/layouter/model/Section.class:(171)         <b>for</b>wardSearchEndIndexesVisitor.reset();
org/jd/core/v1/service/layouter/model/Section.class:(172)         while (iterator.hasPrevious() && backwardSearchStartIndexesVisitor.isEnabled())
org/jd/core/v1/service/layouter/model/Section.class:(173)             ((FlexibleFragment)iterator.previous()).accept(backwardSearchStartIndexesVisitor); 
org/jd/core/v1/service/layouter/model/Section.class:(174)         <b>for</b> (FlexibleFragment flexibleFragment : this.flexibleFragments) {
org/jd/core/v1/service/layouter/model/Section.class:(175)             flexibleFragment.accept(<b>for</b>wardSearchEndIndexesVisitor);
org/jd/core/v1/service/layouter/model/Section.class:(176)             if (!<b>for</b>wardSearchEndIndexesVisitor.isEnabled())
org/jd/core/v1/service/layouter/model/Section.class:(177)                 break; 
org/jd/core/v1/service/layouter/model/Section.class:(178)         } 
org/jd/core/v1/service/layouter/model/Section.class:(179)         int size = backwardSearchStartIndexesVisitor.getSize();
--
org/jd/core/v1/service/layouter/model/Section.class:(180)         Section nextSection = searchNextSection(<b>for</b>wardSearchVisitor);
org/jd/core/v1/service/layouter/model/Section.class:(181)         if (size > 1 && nextSection != null) {
org/jd/core/v1/service/layouter/model/Section.class:(182)             int index1 = flexibleCount - 1 - backwardSearchStartIndexesVisitor.getIndex(size / 2);
org/jd/core/v1/service/layouter/model/Section.class:(183)             int index2 = flexibleCount - 1 - backwardSearchStartIndexesVisitor.getIndex(0);
--
org/jd/core/v1/service/layouter/model/Section.class:(184)             int nextIndex = <b>for</b>wardSearchVisitor.getIndex();
org/jd/core/v1/service/layouter/model/Section.class:(185)             size = <b>for</b>wardSearchEndIndexesVisitor.getSize();
org/jd/core/v1/service/layouter/model/Section.class:(186)             if (size > 1) {
org/jd/core/v1/service/layouter/model/Section.class:(187)                 int index3 = <b>for</b>wardSearchEndIndexesVisitor.getIndex(0) + 1;
org/jd/core/v1/service/layouter/model/Section.class:(188)                 int index4 = <b>for</b>wardSearchEndIndexesVisitor.getIndex(size / 2) + 1;
org/jd/core/v1/service/layouter/model/Section.class:(189)                 Section previousSection = searchPreviousSection(backwardSearchVisitor);
org/jd/core/v1/service/layouter/model/Section.class:(190)                 if (nextSection.getRate() > previousSection.getRate()) {
org/jd/core/v1/service/layouter/model/Section.class:(191)                     int index = previousSection.getFlexibleFragments().size() - backwardSearchVisitor.getIndex();
--
org/jd/core/v1/service/layouter/model/Section.class:(198)             } 
org/jd/core/v1/service/layouter/model/Section.class:(199)             return true;
org/jd/core/v1/service/layouter/model/Section.class:(200)         } 
org/jd/core/v1/service/layouter/model/Section.class:(201)         size = <b>for</b>wardSearchEndIndexesVisitor.getSize();
org/jd/core/v1/service/layouter/model/Section.class:(202)         if (size > 1) {
org/jd/core/v1/service/layouter/model/Section.class:(203)             int index3 = <b>for</b>wardSearchEndIndexesVisitor.getIndex(0) + 1;
org/jd/core/v1/service/layouter/model/Section.class:(204)             int index4 = <b>for</b>wardSearchEndIndexesVisitor.getIndex(size / 2) + 1;
org/jd/core/v1/service/layouter/model/Section.class:(205)             Section previousSection = searchPreviousSection(backwardSearchVisitor);
org/jd/core/v1/service/layouter/model/Section.class:(206)             if (size > 1 && previousSection != null) {
org/jd/core/v1/service/layouter/model/Section.class:(207)                 int index = previousSection.getFlexibleFragments().size() - backwardSearchVisitor.getIndex();
--
org/jd/core/v1/service/layouter/model/Section.class:(217)         visitor.reset();
org/jd/core/v1/service/layouter/model/Section.class:(218)         while (section != null) {
org/jd/core/v1/service/layouter/model/Section.class:(219)             visitor.resetIndex();
org/jd/core/v1/service/layouter/model/Section.class:(220)             <b>for</b> (FlexibleFragment flexibleFragment : section.getFlexibleFragments()) {
org/jd/core/v1/service/layouter/model/Section.class:(221)                 flexibleFragment.accept(visitor);
org/jd/core/v1/service/layouter/model/Section.class:(222)                 if (visitor.getDepth() == 0)
org/jd/core/v1/service/layouter/model/Section.class:(223)                     return section; 
--
org/jd/core/v1/service/layouter/model/Section.class:(263)     protected void addFragmentsAtEnd(VisitorsHolder holder, int index, List<FlexibleFragment> flexibleFragments) {
org/jd/core/v1/service/layouter/model/Section.class:(264)         AbstractSearchMovableBlockFragmentVisitor visitor = holder.getForwardSearchVisitor();
org/jd/core/v1/service/layouter/model/Section.class:(265)         visitor.reset();
org/jd/core/v1/service/layouter/model/Section.class:(266)         <b>for</b> (FlexibleFragment flexibleFragment : flexibleFragments) {
org/jd/core/v1/service/layouter/model/Section.class:(267)             flexibleFragment.accept(visitor);
org/jd/core/v1/service/layouter/model/Section.class:(268)             if (visitor.getDepth() == 2)
org/jd/core/v1/service/layouter/model/Section.class:(269)                 break; 
--
org/jd/core/v1/service/layouter/model/Section.class:(284)     }
org/jd/core/v1/service/layouter/model/Section.class:(285)     
org/jd/core/v1/service/layouter/model/Section.class:(286)     protected void resetLineCount() {
org/jd/core/v1/service/layouter/model/Section.class:(287)         <b>for</b> (FlexibleFragment flexibleFragment : this.flexibleFragments)
org/jd/core/v1/service/layouter/model/Section.class:(288)             flexibleFragment.resetLineCount(); 
org/jd/core/v1/service/layouter/model/Section.class:(289)     }
org/jd/core/v1/service/layouter/model/Section.class:(290)     
--
org/jd/core/v1/service/tokenizer/javafragmenttotoken/JavaFragmentToTokenProcessor.class:(8)     public void process(Message message) throws Exception {
org/jd/core/v1/service/tokenizer/javafragmenttotoken/JavaFragmentToTokenProcessor.class:(9)         List<JavaFragment> fragments = message.<List<JavaFragment>>getBody();
org/jd/core/v1/service/tokenizer/javafragmenttotoken/JavaFragmentToTokenProcessor.class:(10)         TokenizeJavaFragmentVisitor visitor = new TokenizeJavaFragmentVisitor(fragments.size() * 3);
org/jd/core/v1/service/tokenizer/javafragmenttotoken/JavaFragmentToTokenProcessor.class:(11)         <b>for</b> (JavaFragment fragment : fragments)
org/jd/core/v1/service/tokenizer/javafragmenttotoken/JavaFragmentToTokenProcessor.class:(12)             fragment.accept(visitor); 
org/jd/core/v1/service/tokenizer/javafragmenttotoken/JavaFragmentToTokenProcessor.class:(13)         message.setBody(visitor.getTokens());
org/jd/core/v1/service/tokenizer/javafragmenttotoken/JavaFragmentToTokenProcessor.class:(14)     }
--
org/jd/core/v1/service/tokenizer/javafragmenttotoken/visitor/TokenizeJavaFragmentVisitor.class:(43)     
org/jd/core/v1/service/tokenizer/javafragmenttotoken/visitor/TokenizeJavaFragmentVisitor.class:(44)     protected static final KeywordToken IMPORT = new KeywordToken("import");
org/jd/core/v1/service/tokenizer/javafragmenttotoken/visitor/TokenizeJavaFragmentVisitor.class:(45)     
org/jd/core/v1/service/tokenizer/javafragmenttotoken/visitor/TokenizeJavaFragmentVisitor.class:(46)     protected static final KeywordToken FOR = new KeywordToken("<b>for</b>");
org/jd/core/v1/service/tokenizer/javafragmenttotoken/visitor/TokenizeJavaFragmentVisitor.class:(47)     
org/jd/core/v1/service/tokenizer/javafragmenttotoken/visitor/TokenizeJavaFragmentVisitor.class:(48)     protected static final KeywordToken TRUE = new KeywordToken("true");
org/jd/core/v1/service/tokenizer/javafragmenttotoken/visitor/TokenizeJavaFragmentVisitor.class:(49)     
--
org/jd/core/v1/service/tokenizer/javafragmenttotoken/visitor/TokenizeJavaFragmentVisitor.class:(274)         List<ImportsFragment.Import> imports = new DefaultList<>(fragment.getImports());
org/jd/core/v1/service/tokenizer/javafragmenttotoken/visitor/TokenizeJavaFragmentVisitor.class:(275)         imports.sort(NAME_COMPARATOR);
org/jd/core/v1/service/tokenizer/javafragmenttotoken/visitor/TokenizeJavaFragmentVisitor.class:(276)         this.tokens.add(StartMarkerToken.IMPORT_STATEMENTS);
org/jd/core/v1/service/tokenizer/javafragmenttotoken/visitor/TokenizeJavaFragmentVisitor.class:(277)         <b>for</b> (ImportsFragment.Import imp : imports) {
org/jd/core/v1/service/tokenizer/javafragmenttotoken/visitor/TokenizeJavaFragmentVisitor.class:(278)             this.tokens.add(IMPORT);
org/jd/core/v1/service/tokenizer/javafragmenttotoken/visitor/TokenizeJavaFragmentVisitor.class:(279)             this.tokens.add(TextToken.SPACE);
org/jd/core/v1/service/tokenizer/javafragmenttotoken/visitor/TokenizeJavaFragmentVisitor.class:(280)             this.tokens.add(new ReferenceToken(1, imp.getInternalName(), imp.getQualifiedName(), null, null));
--
org/jd/core/v1/service/tokenizer/javafragmenttotoken/visitor/TokenizeJavaFragmentVisitor.class:(286)     
org/jd/core/v1/service/tokenizer/javafragmenttotoken/visitor/TokenizeJavaFragmentVisitor.class:(287)     public void visit(LineNumberTokensFragment fragment) {
org/jd/core/v1/service/tokenizer/javafragmenttotoken/visitor/TokenizeJavaFragmentVisitor.class:(288)         this.knownLineNumberTokenVisitor.reset(fragment.getFirstLineNumber());
org/jd/core/v1/service/tokenizer/javafragmenttotoken/visitor/TokenizeJavaFragmentVisitor.class:(289)         <b>for</b> (Token token : fragment.getTokens())
org/jd/core/v1/service/tokenizer/javafragmenttotoken/visitor/TokenizeJavaFragmentVisitor.class:(290)             token.accept(this.knownLineNumberTokenVisitor); 
org/jd/core/v1/service/tokenizer/javafragmenttotoken/visitor/TokenizeJavaFragmentVisitor.class:(291)     }
org/jd/core/v1/service/tokenizer/javafragmenttotoken/visitor/TokenizeJavaFragmentVisitor.class:(292)     
--
org/jd/core/v1/service/tokenizer/javafragmenttotoken/visitor/TokenizeJavaFragmentVisitor.class:(475)     }
org/jd/core/v1/service/tokenizer/javafragmenttotoken/visitor/TokenizeJavaFragmentVisitor.class:(476)     
org/jd/core/v1/service/tokenizer/javafragmenttotoken/visitor/TokenizeJavaFragmentVisitor.class:(477)     public void visit(TokensFragment fragment) {
org/jd/core/v1/service/tokenizer/javafragmenttotoken/visitor/TokenizeJavaFragmentVisitor.class:(478)         <b>for</b> (Token token : fragment.getTokens())
org/jd/core/v1/service/tokenizer/javafragmenttotoken/visitor/TokenizeJavaFragmentVisitor.class:(479)             token.accept(this.unknownLineNumberTokenVisitor); 
org/jd/core/v1/service/tokenizer/javafragmenttotoken/visitor/TokenizeJavaFragmentVisitor.class:(480)     }
org/jd/core/v1/service/tokenizer/javafragmenttotoken/visitor/TokenizeJavaFragmentVisitor.class:(481)     
--
org/jd/core/v1/service/writer/WriteTokenProcessor.class:(15)         int minorVersion = ((Integer)message.<Integer>getHeader("minorVersion")).intValue();
org/jd/core/v1/service/writer/WriteTokenProcessor.class:(16)         printer.start(maxLineNumber, majorVersion, minorVersion);
org/jd/core/v1/service/writer/WriteTokenProcessor.class:(17)         visitor.start(printer, tokens);
org/jd/core/v1/service/writer/WriteTokenProcessor.class:(18)         <b>for</b> (Token token : tokens)
org/jd/core/v1/service/writer/WriteTokenProcessor.class:(19)             token.accept(visitor); 
org/jd/core/v1/service/writer/WriteTokenProcessor.class:(20)         visitor.end();
org/jd/core/v1/service/writer/WriteTokenProcessor.class:(21)         printer.end();
--
org/jd/core/v1/service/writer/visitor/PrintTokenVisitor.class:(148)     
org/jd/core/v1/service/writer/visitor/PrintTokenVisitor.class:(149)     protected int searchLineNumber() {
org/jd/core/v1/service/writer/visitor/PrintTokenVisitor.class:(150)         this.searchLineNumberVisitor.reset();
org/jd/core/v1/service/writer/visitor/PrintTokenVisitor.class:(151)         <b>for</b> (int i = this.index; i >= 0; i--) {
org/jd/core/v1/service/writer/visitor/PrintTokenVisitor.class:(152)             ((Token)this.tokens.get(i)).accept(this.searchLineNumberVisitor);
org/jd/core/v1/service/writer/visitor/PrintTokenVisitor.class:(153)             if (this.searchLineNumberVisitor.lineNumber != UNKNOWN_LINE_NUMBER)
org/jd/core/v1/service/writer/visitor/PrintTokenVisitor.class:(154)                 return this.searchLineNumberVisitor.lineNumber; 
--
org/jd/core/v1/service/writer/visitor/PrintTokenVisitor.class:(157)         } 
org/jd/core/v1/service/writer/visitor/PrintTokenVisitor.class:(158)         this.searchLineNumberVisitor.reset();
org/jd/core/v1/service/writer/visitor/PrintTokenVisitor.class:(159)         int size = this.tokens.size();
org/jd/core/v1/service/writer/visitor/PrintTokenVisitor.class:(160)         <b>for</b> (int j = this.index; j < size; j++) {
org/jd/core/v1/service/writer/visitor/PrintTokenVisitor.class:(161)             ((Token)this.tokens.get(j)).accept(this.searchLineNumberVisitor);
org/jd/core/v1/service/writer/visitor/PrintTokenVisitor.class:(162)             if (this.searchLineNumberVisitor.lineNumber != UNKNOWN_LINE_NUMBER)
org/jd/core/v1/service/writer/visitor/PrintTokenVisitor.class:(163)                 return this.searchLineNumberVisitor.lineNumber; 
--
org/jd/core/v1/util/DefaultList.class:(18)     public DefaultList(E element, E... elements) {
org/jd/core/v1/util/DefaultList.class:(19)         ensureCapacity(elements.length + 1);
org/jd/core/v1/util/DefaultList.class:(20)         add(element);
org/jd/core/v1/util/DefaultList.class:(21)         <b>for</b> (E e : elements)
org/jd/core/v1/util/DefaultList.class:(22)             add(e); 
org/jd/core/v1/util/DefaultList.class:(23)     }
org/jd/core/v1/util/DefaultList.class:(24)     
--
org/jd/core/v1/util/DefaultList.class:(25)     public DefaultList(E[] elements) {
org/jd/core/v1/util/DefaultList.class:(26)         if (elements != null && elements.length > 0) {
org/jd/core/v1/util/DefaultList.class:(27)             ensureCapacity(elements.length);
org/jd/core/v1/util/DefaultList.class:(28)             <b>for</b> (E e : elements)
org/jd/core/v1/util/DefaultList.class:(29)                 add(e); 
org/jd/core/v1/util/DefaultList.class:(30)         } 
org/jd/core/v1/util/DefaultList.class:(31)     }
--
org/jd/core/v1/util/DefaultStack.class:(65)         sb.append(", elements=[");
org/jd/core/v1/util/DefaultStack.class:(66)         if (this.head > 0) {
org/jd/core/v1/util/DefaultStack.class:(67)             sb.append(this.elements[0]);
org/jd/core/v1/util/DefaultStack.class:(68)             <b>for</b> (int i = 1; i < this.head; i++) {
org/jd/core/v1/util/DefaultStack.class:(69)                 sb.append(", ");
org/jd/core/v1/util/DefaultStack.class:(70)                 sb.append(this.elements[i]);
org/jd/core/v1/util/DefaultStack.class:(71)             } 
--
+ echo '*** eof test ***'
*** eof test ***
