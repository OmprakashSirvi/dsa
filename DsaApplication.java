import leetcode.Pairs;
import utils.Utils;

public class DsaApplication {
    public static void main(String[] args) {
        int[] arr = { 10, 5, 2, 3, -6, 9, 11 };

        int[] sumPairs = Pairs.sumPairs(arr, 4);
        Utils.PrintArray(sumPairs);

        Utils.PrintArray(Pairs.sumPairsImproved(arr, 5));
    }
}
