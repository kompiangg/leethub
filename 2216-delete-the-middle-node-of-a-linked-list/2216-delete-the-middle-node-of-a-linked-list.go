/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return nil
    }

    fast := head.Next.Next
    slow := head

    for fast != nil && fast.Next != nil{
        slow = slow.Next
        fast = fast.Next.Next
    }

    slow.Next = slow.Next.Next

    return head
}