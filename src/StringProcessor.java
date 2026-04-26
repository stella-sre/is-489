public class StringProcessor {

    public static final String INPUT_DATA = "Login; Name; Email" + System.lineSeparator() +
            "peterson; Chris Peterson; peterson@outlook.com" + System.lineSeparator() +
            "james; Derek James; james@gmail.com" + System.lineSeparator() +
            "jackson; Walter Jackson; jackson@gmail.com" + System.lineSeparator() +
            "gregory; Mike Gregory; gregory@yahoo.com";

    public static void main(String[] args) {
        System.out.println("===== Convert 1 =====");
        System.out.println(convert1(INPUT_DATA));
        System.out.println("===== Convert 2 =====");
        System.out.println(convert2(INPUT_DATA));
    }

    public static String convert1(String input) {
        StringBuilder result = new StringBuilder();
        String[] lines = input.split(System.lineSeparator());
        for (int i = 1; i < lines.length; i++) {
            String[] wordsInLine = lines[i].split(";");
            result.append(wordsInLine[0].trim()).append(" ==> ").append(wordsInLine[2].trim()).append(System.lineSeparator());
        }
        return result.toString();
    }

    public static String convert2(String input) {
        StringBuilder result = new StringBuilder();
        String[] lines = input.split(System.lineSeparator());
        for (int i = 1; i < lines.length; i++) {
            String[] wordsInLine = lines[i].split(";");
            result.append(wordsInLine[1].trim()).append(" (email: ").append(wordsInLine[2].trim()).append(")").append(System.lineSeparator());
        }
        return result.toString();
    }
}