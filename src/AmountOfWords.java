import java.util.Scanner;

public class AmountOfWords {

    public static void main(String[] args) {

        try (Scanner sc = new Scanner(System.in)) {
            System.out.print("Please, enter any text: ");
            String userInput = sc.nextLine();

            if (userInput == null || userInput.trim().isEmpty()) {
                System.err.println("Error: no text was entered.");
                return;
            }

            int amountOfWords = getWordsAmount(userInput);
            System.out.println("Amount of words in your text: " + amountOfWords);

        } catch (Exception e) {
            System.err.println("Error: unexpected problem → " + e.getMessage());
        }
    }

    public static int getWordsAmount(String text) {
        return text.trim().split("[\\p{P}\\s]+").length;
    }
}