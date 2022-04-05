

This is a small utility, it greps through the contents of an archive file, it also greps through any archive files that are embedded within an archive file.
They like to do that in java, where they have jars nested in jars.

The following archives are supported right now: zip|jar|war|ear|tar|tgz|taz|tar.gz|tbz2|tbz|tar.bz2|tar.bz
I am trusting this link with the definition of these extensions [link](https://www.gnu.org/software/tar/manual/tar.html#Compression)

The utility is written in golang, you can see how to use it from the test output:
The utility accepts a fixed search string, or a regular expression defined [here](https://github.com/google/re2/wiki/Syntax)


The usage of the utility is here:
<pre>
./rzgrep -h
Usage of ./rzgrep:
  -e string
    	regular expression to search for. Syntax defined here: https://github.com/google/re2/wiki/Syntax
  -in string
    	file or directory to scan
  -v	verbose output
</pre>

The test scripts runs its on a few examples, that's how it looks like:

<pre>
+ ./rzgrep -e Cl.se -in zip.jar
zip.jar|src/rzgrep.go:(95) 	defer archive.Close()
zip.jar|src/rzgrep.go:(105) 		fileReader.Close()
zip.jar|src/rzgrep.go:(134) 		fileReader.Close()
zip.jar|src/rzgrep.go:(172) 	defer file.Close()
zip.jar|src/rzgrep.go:(185) 	defer gzf.Close()
zip.jar|src/rzgrep.go:(203) 	defer file.Close()
zip.jar|src/rzgrep.go:(253) 	defer file.Close()
+ ./rzgrep -e Cl.se -in zip.jar
zip.jar|src/rzgrep.go:(95) 	defer archive.Close()
zip.jar|src/rzgrep.go:(105) 		fileReader.Close()
zip.jar|src/rzgrep.go:(134) 		fileReader.Close()
zip.jar|src/rzgrep.go:(172) 	defer file.Close()
zip.jar|src/rzgrep.go:(185) 	defer gzf.Close()
zip.jar|src/rzgrep.go:(203) 	defer file.Close()
zip.jar|src/rzgrep.go:(253) 	defer file.Close()
+ ./rzgrep -e Cl.se -in zip.ear
zip.ear|zip.jar|src/rzgrep.go:(95) 	defer archive.Close()
zip.ear|zip.jar|src/rzgrep.go:(105) 		fileReader.Close()
zip.ear|zip.jar|src/rzgrep.go:(134) 		fileReader.Close()
zip.ear|zip.jar|src/rzgrep.go:(172) 	defer file.Close()
zip.ear|zip.jar|src/rzgrep.go:(185) 	defer gzf.Close()
zip.ear|zip.jar|src/rzgrep.go:(203) 	defer file.Close()
zip.ear|zip.jar|src/rzgrep.go:(253) 	defer file.Close()
+ ./rzgrep -e Cl.se -in zip.tgz
zip.tgz|zip.jar|src/rzgrep.go:(95) 	defer archive.Close()
zip.tgz|zip.jar|src/rzgrep.go:(105) 		fileReader.Close()
zip.tgz|zip.jar|src/rzgrep.go:(134) 		fileReader.Close()
zip.tgz|zip.jar|src/rzgrep.go:(172) 	defer file.Close()
zip.tgz|zip.jar|src/rzgrep.go:(185) 	defer gzf.Close()
zip.tgz|zip.jar|src/rzgrep.go:(203) 	defer file.Close()
zip.tgz|zip.jar|src/rzgrep.go:(253) 	defer file.Close()
+
</pre>
