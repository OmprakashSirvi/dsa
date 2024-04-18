package leetcode;

import linkedlist.ListNode;

public class ReverseList {
    public static ListNode reverseList(ListNode head) {
        // reverse a list using recursion
        if (head == null || head.next == null)
            return head;

        ListNode reversedList = reverseList(head.next);
        head.next.next = head;
        head.next = null;
        return reversedList;
    }
}
