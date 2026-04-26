import java.util.Scanner;

public class SumDigitsInNumber {

    public static void main(String[] args) {

        try (Scanner sc = new Scanner(System.in)) {
            System.out.print("Please, enter integer: ");
            int number = sc.nextInt();

            int sumOfDigits = sumDigitsInNumber(number);
            System.out.println("Sum of digits: " + sumOfDigits);

        } catch (java.util.InputMismatchException e) {
            System.err.println("Error: invalid input. Please enter an integer number only.");
        } catch (Exception e) {
            System.err.println("Error: unexpected problem → " + e.getMessage());
        }
    }

    public static int sumDigitsInNumber(int number) {
        int result = 0;
        while (number != 0) {
            result += number % 10;
            number /= 10;
        }
        return Math.abs(result);
    }
}