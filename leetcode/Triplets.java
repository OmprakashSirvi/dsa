package leetcode;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

public class Triplets {
    public static List<List<Integer>> solve(int[] arr, int sum) {
        List<List<Integer>> returnArray = new ArrayList<>();
        Arrays.sort(arr);
        for (int i = 0; i < arr.length - 2; i++) {
            int currentElement = arr[i];
            // find pairs for arr[i] to arr[arr.length - 1]
            List<List<Integer>> pairs = Pairs.sortedSolve(Arrays.copyOfRange(arr, i + 1, arr.length),
                    sum - currentElement, currentElement);
            returnArray.addAll(pairs);
        }
        return returnArray;
    }
}
