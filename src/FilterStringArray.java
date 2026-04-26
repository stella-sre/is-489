import java.util.Arrays;
import java.util.Scanner;

public class FilterStringArray {

    public static void main(String[] args) {

        try (Scanner sc = new Scanner(System.in)) {
            System.out.print("Please, enter any words separated by space: ");
            String userInput = sc.nextLine();

            if (userInput == null || userInput.trim().isEmpty()) {
                System.err.println("Error: no words were entered.");
                return;
            }

            System.out.print("Please, enter minimum word length to filter words: ");
            int minLength = sc.nextInt();

            if (minLength <= 0) {
                System.err.println("Error: minimum length must be a positive integer greater than 0.");
                return;
            }

            String[] words = userInput.split("\\s+");
            String[] filteredWords = filterWordsByLength(minLength, words);

            if (filteredWords.length == 0) {
                System.out.println("No words matched the minimum length of " + minLength + ".");
            } else {
                System.out.println(Arrays.toString(filteredWords));
            }

        } catch (java.util.InputMismatchException e) {
            System.err.println("Error: invalid input. Please enter an integer for the minimum length.");
        } catch (Exception e) {
            System.err.println("Error: unexpected problem → " + e.getMessage());
        }
    }

    public static String[] filterWordsByLength(int minLength, String[] words) {
        String[] filteredArray = new String[words.length];
        for (int i = 0; i < words.length; i++) {
            if (words[i].length() >= minLength) {
                filteredArray[i] = words[i];
            }
        }
        return filterNulls(filteredArray);
    }

    private static String[] filterNulls(String[] arr) {
        int newArraySize = 0;
        for (String word : arr) {
            if (word != null) {
                newArraySize++;
            }
        }
        String[] filteredArray = new String[newArraySize];
        int filteredArrayIndex = 0;
        for (String word : arr) {
            if (word != null) {
                filteredArray[filteredArrayIndex++] = word;
            }
        }
        return filteredArray;
    }
}