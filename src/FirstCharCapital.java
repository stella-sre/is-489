import java.util.Scanner;

public class FirstCharCapital {

    public static void main(String[] args) {

        try (Scanner sc = new Scanner(System.in)) {
            System.out.print("Please, enter any text: ");
            String userInput = sc.nextLine();

            if (userInput == null || userInput.trim().isEmpty()) {
                System.err.println("Error: no text was entered.");
                return;
            }

            System.out.println(firstCharToTitleCase(userInput));

        } catch (Exception e) {
            System.err.println("Error: unexpected problem → " + e.getMessage());
        }
    }

    public static String firstCharToTitleCase(String string) {
        char[] chars = string.toLowerCase().toCharArray();
        boolean found = false;
        for (int i = 0; i < chars.length; i++) {
            if (!found && Character.isLetter(chars[i])) {
                chars[i] = Character.toUpperCase(chars[i]);
                found = true;
            } else if (Character.isWhitespace(chars[i])) {
                found = false;
            }
        }
        return String.valueOf(chars);
    }
}