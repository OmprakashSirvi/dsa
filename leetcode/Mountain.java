package leetcode;

public class Mountain {
    public static int solve(int[] arr) {
        boolean descending = false;
        int maxLen = 0;
        int peakIndex = 0;
        int width = 0;

        for (int i = 1; i < arr.length - 1; i++) {
            // If peak is reached
            if (arr[i] > arr[i + 1] && arr[i] > arr[i - 1]) {
                descending = true;
                peakIndex = i;
                // going backward
                while (descending) {
                    if (i > 0 && arr[i] > arr[i - 1]) {
                        i--;
                    } else {
                        width = peakIndex - i;
                        descending = false;
                        i = peakIndex;
                    }
                }

                // going ahead
                descending = true;
                while (descending) {
                    if (i < arr.length - 1 && arr[i] > arr[i + 1])
                        i++;
                    else {
                        width = width + (i - peakIndex) + 1;
                        descending = false;
                        maxLen = Math.max(width, maxLen);
                    }
                }

            }
        }
        return maxLen;
    }
}
