import data.Data;
import sort.InsertionSort;
import utils.Utils;

public class DsaApplication {
    public static void main(String[] args) {
        // Get an array from data
        int[] arr = Data.generateRandomArray(10);

        // Print array from the data
        Utils.PrintArray(arr);

        InsertionSort.sort(arr);
    }
}
