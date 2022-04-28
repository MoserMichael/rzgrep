package decomp;
import org.junit.Assert;
import org.junit.Test;


public class DecompTest  {

    @Test
    public void testDecompiler() {

        try {
            String result = Decompy.decompile(null, "decomp/Decompy");

            Assert.assertTrue(result != null && !result.equals(""));

            System.out.println(result);
        } catch (Exception ex) {
            ex.printStackTrace();
            Assert.assertTrue(false);
        }
    }
}
