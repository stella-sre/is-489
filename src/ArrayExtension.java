import java.util.Arrays;
import java.util.Random;
import java.util.Scanner;

public class ArrayExtension {

    public static final int MULTIPLIER = 2;

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);

        try {
            System.out.print("Please, enter length of initial array: ");
            int baseArrayLength = sc.nextInt();

            if (baseArrayLength <= 0) {
                System.err.println("Error: array length must be a positive integer greater than 0.");
                return;
            }

            int[] arr = generateRandomArray(baseArrayLength);
            int[] extendedArray = extendArray(arr);

            System.out.println("*** Initial array ***");
            System.out.println(Arrays.toString(arr));
            System.out.println("*** Extended array ***");
            System.out.println(Arrays.toString(extendedArray));

        } catch (java.util.InputMismatchException e) {
            System.err.println("Error: invalid input. Please enter an integer number only.");
        } catch (Exception e) {
            System.err.println("Error: unexpected problem → " + e.getMessage());
        } finally {
            sc.close();
        }
    }

    public static int[] extendArray(int[] arr) {
        int newArrayLength = arr.length * MULTIPLIER;
        int[] resultArray = Arrays.copyOf(arr, newArrayLength);
        for (int i = arr.length; i < newArrayLength; i++) {
            resultArray[i] = arr[i - arr.length] * MULTIPLIER;
        }
        return resultArray;
    }

    public static int[] generateRandomArray(int amountOfElements) {
        Random r = new Random();
        int[] resultArray = new int[amountOfElements];
        for (int i = 0; i < amountOfElements; i++) {
            resultArray[i] = r.nextInt(100) + 1;
        }
        return resultArray;
    }
}