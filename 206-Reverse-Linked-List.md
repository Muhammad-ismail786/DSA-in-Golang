# 206. Reverse Linked List (Go)

## Intuition

Is problem mein humein Singly Linked List ko reverse karna hai.

Example:

Original:
`1 → 2 → 3 → nil`

Reversed:
`3 → 2 → 1 → nil`

Iske liye hum 3 pointers use karenge:

1. **prev:** Previous Node ka address store karta hai. Starting mein `nil` hota hai.
2. **current:** Current Node ka address store karta hai. Starting mein `head` hota hai.
3. **nextNode:** Agle Node ka address temporarily save karta hai, taake link reverse karte waqt agla Node lose na ho.

Go mein C++ ke `nullptr` ki jagah `nil` use hota hai.

## Approach

Hum `for` loop use karenge jo tab tak chalega jab tak `current != nil` ho.

Har iteration mein 4 steps perform honge:

**Step 1: Next Node ka address save karna**

`nextNode := current.Next`

Example:
- current = Node 1
- current.Next = Node 2
- nextNode = Node 2

**Step 2: Link reverse karna**

`current.Next = prev`

Node 1 ka Next ab nil ko point karega.

**Step 3: Previous pointer ko aage move karna**

`prev = current`

Ab prev Node 1 ko point karega.

**Step 4: Current pointer ko next Node par move karna**

`current = nextNode`

Ab current Node 2 ko point karega.

Ye 4 steps har Node ke liye repeat honge.

## Dry Run

Original Linked List:

`1 → 2 → 3 → nil`

| Iteration | current | nextNode | prev |
|---|---|---|---|
| Start | Node 1 | nil | nil |
| 1 | Node 1 | Node 2 | Node 1 |
| 2 | Node 2 | Node 3 | Node 2 |
| 3 | Node 3 | nil | Node 3 |
| End | nil | nil | Node 3 |

Final Linked List:

`3 → 2 → 1 → nil`

`return prev` use karte hain kyunke reverse hone ke baad prev naye first Node ko point karta hai.

## Complexity

- **Time Complexity:** O(n)

  Linked List ke har Node ko sirf aik baar visit karte hain.

- **Space Complexity:** O(1)

  Hum sirf 3 pointers (`prev`, `current`, `nextNode`) use karte hain. Extra array ya Linked List create nahi karte.

## Code

See `206-Reverse-Linked-List.go` for the Go solution.
