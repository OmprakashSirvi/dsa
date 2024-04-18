package linkedlist;

public class ListNode {
    public int val;
    public ListNode next;

    public ListNode() {
    }

    public ListNode(int val) {
        this.val = val;
    }

    ListNode(int val, ListNode next) {
        this.val = val;
        this.next = next;
    }

    public ListNode add(int val) {
        ListNode node = new ListNode(val);
        ListNode current = this;
        while (current.next != null)
            current = current.next;

        // set the last element to the latest value
        current.next = node;

        return this;
    }

    public void print() {
        ListNode current = this;
        while (current != null) {
            System.out.print(current.val + " ");
            current = current.next;
        }
        System.out.println();
    }
}
