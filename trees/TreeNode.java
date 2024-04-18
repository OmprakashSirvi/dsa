package trees;

import lombok.Getter;
import lombok.Setter;

@Getter
@Setter
public class TreeNode {
    private int data;
    private TreeNode leftChild;
    private TreeNode rightChild;

    public void insert(int value){
        if (value == data) return;

        if (value < data){
            if (leftChild == null){
                this.leftChild = new TreeNode(value);
            }else {
                insert(value);
            }
        }else {
            if (rightChild == null){
                this.rightChild = new TreeNode(value);
            } else {
                insert(value);
            }
        }
    }

    public TreeNode(int data) {
        this.data = data;
    }
}
