/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    var d []*ListNode
    for head!=nil{
        d=append(d,head)
        head = head.Next
    }
    return d[len(d)/2]
    
    
    
}