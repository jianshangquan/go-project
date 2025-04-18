public class Main {
    public static void main(String[] args) {
        loop();
    }

    public static void loop() {
        long count = 0;
        long startTime = System.currentTimeMillis();

        while (System.currentTimeMillis() - startTime < (1000 * 10)) {  // loop for 1 second
            count++;
            // System.out.println("Java Looped: " + count);
        }

        System.out.println("Looped: " + count + " times in 1 second.");
    }
}
