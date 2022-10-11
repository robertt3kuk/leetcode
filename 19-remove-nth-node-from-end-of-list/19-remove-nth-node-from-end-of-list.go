/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    var d []*ListNode
    h:=head
    for head!=nil{
        d=append(d,head)
        head = head.Next
    }
    
    
    head=h
    for i,c:= range d{
        if i!=0&&i!=len(d)-n{
            h.Next=c
            h=h.Next
        }
        if i==0&&i==len(d)-n{
            head = h.Next
        }
        if i+1==len(d)-n&&i+1==len(d)-1{
            h.Next=nil
            break
                
        }
        
    }
    return head
}