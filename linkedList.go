package linkedList

import (
    "fmt"
    "iter"

)

type Node[T any] struct {
    data    T
    prev    *Node[T]
    next    *Node[T]
}

type List[T any] struct {
    head    *Node[T]
    tail    *Node[T]
    size    int
}

func newNode[T any](data T, prev, next *Node[T]) *Node[T] {
    return &Node[T]{
        data:   data,
        prev:   prev,
        next:   next,
    }
}

func (n *Node[T]) Next() *Node[T] {
    return n.next
}

func (n *Node[T]) Prev() *Node[T] {
    return n.prev
}

func (n *Node[T]) Get() T {
    return n.data
}

func (n *Node[T]) Set(data T) {
    n.data = data
}

func New[T any]() *List[T] {
    return &List[T]{
        head:   nil,
        tail:   nil,
        size:   0,
    }
}

func FromSlice[T any](slice []T) *List[T] {
    l := New[T]()
    l.size = len(slice)
    if l.size == 0 {
        return l
    }
    node := newNode(slice[0], nil, nil)
    l.head = node
    l.tail = node
    for i := 1; i < l.size; i++ {
        node.next = newNode(slice[i], l.tail, nil)
        node = node.next
        l.tail = node
    }
    return l
}

func (l *List[T]) Append(data T) {
    if l.size == 0 {
        temp := newNode(data, nil, nil)
        l.head = temp
        l.tail = temp
    } else {
        temp := newNode(data, l.tail, nil)
        l.tail.next = temp
    }
    l.size += 1
}

func (l *List[T]) AppendAfter(data T, node *Node[T]) {
    if node.next == nil {
        node.next = newNode(data, node, nil)
        l.tail = node.next
    } else {
        node.next = newNode(data, node, node.next)
        node.next.prev = node.next
    }
    l.size += 1
}

func (l *List[T]) AppendBefore(data T, node *Node[T]) {
    if node.prev == nil {
        node.prev = newNode(data, nil, node)
        l.head = node.prev
    } else {
        node.prev = newNode(data, node.prev, node)
        node.prev.next = node.prev
    }
    l.size += 1
}

func (l *List[T]) AppendSlice(slice []T) {
    sliceLen := len(slice)
    if sliceLen == 0 {
        return
    }
    if l.size == 0 {
        l.size = sliceLen
        node := newNode(slice[0], nil, nil)
        l.head = node
        l.tail = node
        for i := 1; i < sliceLen; i++ {
            node.next = newNode(slice[i], node, nil)
            node = node.next
        }
        l.tail = node
    } else {
        l.size += sliceLen
        node := l.tail
        for i := 0; i < sliceLen; i++ {
            node.next = newNode(slice[i], node, nil)
            node = node.next
        }
        l.tail = node
    }
}

func (l *List[T]) AppendAt(data T, index int) error {
    if index < 0 || index > l.size {
        return fmt.Errorf("AppendAt: index out of bound")
    }
    if l.size == 0 {
        l.size += 1
        node := newNode(data, nil, nil)
        l.head = node
        l.tail = node
        return nil
    } else if index == l.size {
        l.Append(data)
        return nil
    }
    i := 0
    for node := l.head; node != nil; node = node.next {
        if i == index {
            if node.prev != nil {
                node.prev.next = newNode(data, node.prev, node)
                node.prev = node.prev.next
            } else {
                node.prev = newNode(data, nil, node)
                l.head = node.prev
            }
            break
        }
        i += 1
    }
    l.size += 1
    return nil
}

func (l *List[T]) AppendSliceAt(slice []T, index int) error {
    if index < 0 || index > l.size {
        return fmt.Errorf("AppendSliceAt: index out of bound")
    }
    sliceLen := len(slice)
    if l.size == 0 {
        l.size = sliceLen
        node := newNode(slice[0], nil, nil)
        l.head = node
        l.tail = node
        for i := 1; i < sliceLen; i++ {
            node.next = newNode(slice[i], node, nil)
            node = node.next
        }
        l.tail = node
        return nil
    } else if index == l.size {
        l.AppendSlice(slice)
        return nil
    }
    i := 0
    for node := l.head; node != nil; node = node.next {
        if i == index {
            if node.prev != nil {
                l.size += sliceLen
                next := node
                node = node.prev
                for i := 0; i < sliceLen; i++ {
                    node.next = newNode(slice[i], node, nil)
                    node = node.next
                }
                node.next = next
            } else {
                l.size += sliceLen
                next := l.head
                node := newNode(slice[0], nil, nil)
                l.head = node
                for i := 1; i < sliceLen; i++ {
                    node.next = newNode(slice[i], node, nil)
                    node = node.next
                }
                node.next = next
            }
        }
        i += 1
    }
    return nil
}

func (l *List[T]) ToSlice() []T {
    res := make([]T, l.size)
    i := 0
    for node := l.head; node != nil; node = node.next {
        res[i] = node.data
        i += 1
    }
    return res
}

func (l *List[T]) Head() *Node[T] {
    return l.head
}

func (l *List[T]) Tail() *Node[T] {
    return l.tail
}

func (l *List[T]) Size() int {
    return l.size
}

func (l *List[T]) Get(index int) (T, error) {
    if index < 0 || index > l.size-1 {
        return *new(T), fmt.Errorf("Get: index out of bound")
    }
    i := 0
    for node := l.head; node != nil; node = node.next {
        if i == index {
            return node.data, nil
        }
        i += 1
    }
    return *new(T), fmt.Errorf("Get: unexpected error")
}

func (l *List[T]) IndexFunc(data T, eq func(T, T) bool) (int, error) {
    i := 0
    for node := l.head; node != nil; node = node.next {
        if eq(node.data, data) {
            return i, nil
        }
        i += 1
    }
    return -1, fmt.Errorf("IndexFunc: data not found")
} 

func (l *List[T]) DeleteFunc(data T, eq func(a, b T) bool) error {
    i := 0
    for node := l.head; node != nil; node = node.next {
        if eq(node.data, data) {
            if node.prev != nil {
                node.prev.next = node.next
            } else {
                l.head = node.next
            }
            if node.next != nil {
                node.next.prev = node.prev    
            } else {
                l.tail = node.prev
            }
            l.size -= 1
            return nil
        }
        i += 1
    }
    return fmt.Errorf("DeleteFunc: data not found")
}

func (l *List[T]) DeleteIndex(index int) error {
    if index < 0 || index > l.size-1 {
        return fmt.Errorf("DeleteIndex: index out of bound")
    }
    i := 0
    for node := l.head; node != nil; node = node.next {
        if i == index {
            if node.prev != nil {
                node.prev.next = node.next
            } else {
                l.head = node.next
            }
            if node.next != nil {
                node.next.prev = node.prev
            } else {
                l.tail = node.prev
            }
            l.size -= 1
            return nil
        }
        i += 1
    }
    return fmt.Errorf("DeleteIndex: unexpected error")
}

func (l *List[T]) DeleteNode(node *Node[T]) {
    if node.prev != nil {
        node.prev.next = node.next
    } else {
        l.head = node.next
    }
    if node.next != nil {
        node.next.prev = node.prev
    } else {
        l.tail = node.prev
    }
    l.size -= 1
}

func (l *List[T]) Set(data T, index int) error {
    if index < 0 || index > l.size-1 {
        return fmt.Errorf("Set: index out of bound")
    }
    i := 0
    for node := l.head; node != nil; node = node.next {
        if i == index {
            node.data = data
            return nil
        }
        i += 1
    }
    return fmt.Errorf("Set: unexpected error")
}

func (l *List[T]) All() iter.Seq2[int, T] {
    return func(yield func(int, T) bool) {
        i := 0
        for node := l.head; node != nil; node = node.next {
            if !yield(i, node.data) {
                return
            }
            i += 1
        }
    }
}

func (l *List[T]) Backward() iter.Seq2[int, T] {
    return func(yield func(int, T) bool) {
        i := l.size-1
        for node := l.tail; node != nil; node = node.prev {
            if !yield(i, node.data) {
                return
            }
            i -= 1
        }
    }
}

func swap[T any](n1, n2 *Node[T]) {
    n1.data, n2.data = n2.data, n1.data
}

func partition[T any](n1, n2 *Node[T], cmp func(a, b T) int) *Node[T] {
    pivot := n2.data
    i := n1.prev
    for j := n1; j != n2; j = j.next {
        if (cmp(j.data, pivot) <= 0) {
            if i == nil {
                i = n1
            } else {
                i = i.next
            }
            swap(i, j)
        }
    }
    if i == nil {
        i = n1
    } else {
        i = i.next
    }
    swap(i, n2)
    return i
}

func quickSort[T any](n1, n2 *Node[T], cmp func(a, b T) int) {
    if (n2 != nil && n1 != n2 && n1 != n2.next) {
        p := partition(n1, n2, cmp)
        quickSort(n1, p.prev, cmp)
        quickSort(p.next, n2, cmp)
    }
}

func (l *List[T]) SortFunc(cmp func(a, b T) int) {
    quickSort(l.head, l.tail, cmp)
}
