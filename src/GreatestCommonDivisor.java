import java.util.Scanner;

public class GreatestCommonDivisor {

    public static void main(String[] args) {

        try (Scanner sc = new Scanner(System.in)) {
            System.out.print("Please, enter two numbers separated by space: ");
            String userInput = sc.nextLine();

            if (userInput == null || userInput.trim().isEmpty()) {
                System.err.println("Error: no input was entered.");
                return;
            }

            String[] inputArgumentsArray = userInput.trim().split("\\s+");

            if (inputArgumentsArray.length != 2) {
                System.err.println("Error: please enter exactly two numbers separated by a space.");
                return;
            }

            int number1 = Integer.parseInt(inputArgumentsArray[0]);
            int number2 = Integer.parseInt(inputArgumentsArray[1]);

            System.out.println("GCD of " + number1 + " and " + number2 + " is: " + gcdRecursive(number1, number2));

        } catch (NumberFormatException e) {
            System.err.println("Error: invalid value detected → " + e.getMessage());
            System.err.println("Please enter integer numbers only.");
        } catch (Exception e) {
            System.err.println("Error: unexpected problem → " + e.getMessage());
        }
    }

    public static int gcdRecursive(int firstNumber, int secondNumber) {
        if (secondNumber == 0) {
            return Math.abs(firstNumber);
        } else {
            return gcdRecursive(secondNumber, firstNumber % secondNumber);
        }
    }
}