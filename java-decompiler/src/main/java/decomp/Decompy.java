package decomp;
import  org.jd.core.v1.api.printer.Printer;
import org.jd.core.v1.api.loader.Loader;
import org.jd.core.v1.api.loader.LoaderException;
import org.jd.core.v1.ClassFileToJavaSourceDecompiler;

import java.io.ByteArrayOutputStream;
import java.io.IOException;
import java.io.InputStream;

public class Decompy {

    public static String decompile(Loader loader, String classPath) throws Exception {

        ClassFileToJavaSourceDecompiler decompiler = new ClassFileToJavaSourceDecompiler();

        if (loader == null) {
            loader = new DefaultLoader();
        }

        DecompilerPrinter printer = new DecompilerPrinter();
        decompiler.decompile(loader, printer, classPath);

        return printer.toString();
    }

    public static class DefaultLoader implements Loader {

        @Override
        public byte[] load(String internalName) throws LoaderException {

            InputStream is = this.getClass().getResourceAsStream("/" + internalName + ".class");

            if (is == null) {
                return null;
            } else {
                try (InputStream in = is; ByteArrayOutputStream out = new ByteArrayOutputStream()) {
                    byte[] buffer = new byte[1024];
                    int read = in.read(buffer);

                    while (read > 0) {
                        out.write(buffer, 0, read);
                        read = in.read(buffer);
                    }

                    byte [] result = out.toByteArray();
                    //System.err.println(">>> load " + internalName + " " + result.length);
                    return result;
                } catch (IOException e) {
                    throw new RuntimeException(e);
                }
            }
        }

        @Override
        public boolean canLoad(String internalName) {
            boolean result = this.getClass().getResource("/" + internalName + ".class") != null;
            //System.err.println(">>> canLoad " + internalName + " " + result);
            return result;
        }
    }

    public static class DecompilerPrinter implements Printer {

        private static final String INDENT = "    ";
        private static final String NEWLINE = "\n";

        private int indentationCount = 0;
        private final StringBuilder sb = new StringBuilder();

        @Override
        public String toString() {
            return sb.toString();
        }

        @Override
        public void start(int maxLineNumber, int majorVersion, int minorVersion) {
        }

        @Override
        public void end() {
        }

        @Override
        public void printText(String text) {
            sb.append(text);
        }

        @Override
        public void printNumericConstant(String constant) {
            sb.append(constant);
        }

        @Override
        public void printStringConstant(String constant, String ownerInternalName) {
            sb.append(constant);
        }

        @Override
        public void printKeyword(String keyword) {
            sb.append(keyword);
        }

        @Override
        public void printDeclaration(int type, String internalTypeName, String name, String descriptor) {
            sb.append(name);
        }

        @Override
        public void printReference(int type, String internalTypeName, String name, String descriptor,
                                   String ownerInternalName) {
            sb.append(name);
        }

        @Override
        public void indent() {
            this.indentationCount++;
        }

        @Override
        public void unindent() {
            this.indentationCount--;
        }

        @Override
        public void startLine(int lineNumber) {
            for (int i = 0; i < indentationCount; i++)
                sb.append(INDENT);
        }

        @Override
        public void endLine() {
            sb.append(NEWLINE);
        }

        @Override
        public void extraLine(int count) {
            while (count-- > 0)
                sb.append(NEWLINE);
        }

        @Override
        public void startMarker(int type) {
        }

        @Override
        public void endMarker(int type) {
        }
    }

}
