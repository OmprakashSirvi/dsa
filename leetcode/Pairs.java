package leetcode;

import java.util.HashMap;

public class Pairs {
    public static int[] sumPairs(int[] arr, int sum) {
        for (int i = 0; i < arr.length; i++) {
            for (int j = i + 1; j < arr.length; j++) {
                if (arr[i] + arr[j] == sum)
                    return new int[] { arr[i], arr[j] };
            }
        }
        return null;
    }

    // { 10, 5, 2, 3, -6, 9, 11 };

    // Map : -6 : 10, -1 : 5,

    public static int[] sumPairsImproved(int[] arr, int sum) {
        HashMap<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i < arr.length; i++) {
            if (map.containsKey(Integer.valueOf(arr[i])))
                return new int[] { arr[i], map.get(arr[i]) };
            map.put(Integer.valueOf(sum - arr[i]), Integer.valueOf(arr[i]));
        }

        return null;
    }
}
