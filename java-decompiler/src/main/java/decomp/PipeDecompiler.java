package decomp;

import org.jd.core.v1.api.loader.Loader;
import org.jd.core.v1.api.loader.LoaderException;

import java.io.*;

public class PipeDecompiler {
    private DataInputStream inData;
    private DataOutputStream outData;
    private RootDirLoader loader;

    private FileWriter logger;

    public static int CMD_GO2JAVA_DECOMPILE_CLASS = 2;
    public static int CMD_JAVA2GO_SEND_DECOMPILE_RESULT = 3;

    public static class RootDirLoader implements Loader {

        private static int SUFFIX_LEN = ".class".length();
        private PipeDecompiler decompiler;
        private String rootDir;

        RootDirLoader(PipeDecompiler decompiler) {
            this.decompiler = decompiler;
        }

        public void setRootDir(String rootDir) {
            if (!rootDir.endsWith("/")) {
                rootDir += "/";
            }
            this.rootDir = rootDir;
        }

        private String makeStringName(String internalName) {
            if (internalName.endsWith(".class")) {
                internalName = internalName.substring(0, internalName.length() - SUFFIX_LEN);
                String fileName = this.rootDir + internalName.replace('.', '/') + ".class";
                this.decompiler.log( "fileName: " + fileName);
                return fileName;
            }

            String fileName = this.rootDir + internalName.replace('.', '/');
            this.decompiler.log( "fileName: " + fileName);
            return fileName;
        }

        @Override
        public byte[] load(String internalName) throws LoaderException {

            this.decompiler.log("load: " + internalName);

            InputStream is = this.getClass().getResourceAsStream("/" + internalName + ".class");
            if (is == null) {
                String fileName = makeStringName(internalName);
                try {
                    is = new FileInputStream(fileName);
                } catch(IOException ex) {
                    this.decompiler.log("file not there! " + ex.toString());
                    return null;
                }
            }

            try (InputStream in = is; ByteArrayOutputStream out = new ByteArrayOutputStream()) {
                byte[] buffer = new byte[1024];
                int read = in.read(buffer);

                while (read > 0) {
                    out.write(buffer, 0, read);
                    read = in.read(buffer);
                }

                byte [] result = out.toByteArray();

                this.decompiler.log("load ok  " + internalName + " " + result.length);

                return result;
            } catch (IOException e) {
                throw new RuntimeException(e);
            }
        }

        @Override
        public boolean canLoad(String internalName) {
            String fileName = makeStringName(internalName);
            File fileCheck = new File(fileName);

            this.decompiler.log( "load: " + fileName + " exists: " + fileCheck.exists());

            return fileCheck.exists();
        }
    }

    PipeDecompiler() {
        inData = new DataInputStream(System.in);
        outData = new DataOutputStream(System.out);
        loader = new RootDirLoader(this);
        openLogger();
    }

    private void openLogger() {
        try {
            this.logger = new FileWriter("log.log");
            this.logger.write("init logger\n");
            this.logger.flush();
        } catch(IOException ex) {
        }
    }

    public void run() {
        try {
            log("waiting for commands...");
            for(;;) {
                int cmd = inData.readInt();
                if (cmd == CMD_GO2JAVA_DECOMPILE_CLASS) {

                    log("got CMD_GO2JAVA_DECOMPILE_CLASS");

                    String rootDir = inData.readUTF();
                    loader.setRootDir(rootDir);

                    log("got rootDir " + rootDir);

                    String className = inData.readUTF();

                    log("got className " + className);

                    String output = Decompy.decompile(loader, className);

                    log("output: " + output);

                    //outData.writeUTF(output);
                    outData.writeInt(CMD_JAVA2GO_SEND_DECOMPILE_RESULT);
                    byte [] outputBytes = output.getBytes("utf-8");
                    outData.writeInt(outputBytes.length);
                    outData.write(outputBytes);
                } else {
                    log("Wrong request type, got " + cmd);
                    break;
                }
            }
        } catch(IOException ex) {
            ex.printStackTrace();
        } catch(Exception exc) {
            exc.printStackTrace();
        }
    }

    private void log(String toLog) {
        try {
            this.logger.write(toLog + "\n");
            this.logger.flush();
        } catch(IOException ex) {
        }
    }
}
