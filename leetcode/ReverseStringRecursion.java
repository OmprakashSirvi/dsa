package leetcode;

public class ReverseStringRecursion {
    public void reverseString(char[] s) {
        helper(0, s.length - 1, s);
    }

    private void helper(int left, int right, char[] s) {
        if (left >= right)
            return;

        // swap elements
        char temp = s[left];
        s[left] = s[right];
        s[right] = temp;

        helper(left + 1, right - 1, s);
    }
}
