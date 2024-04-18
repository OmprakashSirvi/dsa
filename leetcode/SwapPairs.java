package leetcode;

import linkedlist.ListNode;

public class SwapPairs {
    public static ListNode swapPairs(ListNode head) {
        // use recursion to swap pairs of nodes
        if (head == null || head.next == null)
            return head;

        ListNode temp = head.next;
        head.next = swapPairs(head.next.next);
        temp.next = head;
        return temp;
    }
}
