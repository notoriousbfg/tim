package interpreter

type OrderedMap struct {
	items      map[interface{}]*Element
	linkedList list
}

func NewOrderedMap() *OrderedMap {
	return &OrderedMap{
		items: make(map[interface{}]*Element),
	}
}

func (m *OrderedMap) Get(key interface{}) (interface{}, bool) {
	element, ok := m.items[key]
	if ok {
		return element.Value, true
	}
	return nil, false
}

func (m *OrderedMap) Set(key, value interface{}) bool {
	_, ok := m.items[key]
	if ok {
		m.items[key].Value = value
		return false
	}

	element := m.linkedList.PushBack(key, value)
	m.items[key] = element
	return true
}

func (m *OrderedMap) Len() int {
	return len(m.items)
}

func (m *OrderedMap) Keys() (keys []interface{}) {
	keys = make([]interface{}, 0, m.Len())
	for el := m.Front(); el != nil; el = el.Next() {
		keys = append(keys, el.Key)
	}
	return keys
}

func (m *OrderedMap) Delete(key interface{}) (didDelete bool) {
	element, ok := m.items[key]
	if ok {
		m.linkedList.Remove(element)
		delete(m.items, key)
	}

	return ok
}

func (m *OrderedMap) Front() *Element {
	return m.linkedList.Front()
}

func (m *OrderedMap) Back() *Element {
	return m.linkedList.Back()
}

func (m *OrderedMap) Copy() *OrderedMap {
	m2 := NewOrderedMap()

	for el := m.Front(); el != nil; el = el.Next() {
		m2.Set(el.Key, el.Value)
	}

	return m2
}

type Element struct {
	prev, next *Element

	Key   interface{}
	Value interface{}
}

func (e *Element) Next() *Element {
	return e.next
}

func (e *Element) Prev() *Element {
	return e.prev
}

type list struct {
	root Element
}

func (l *list) IsEmpty() bool {
	return l.root.next == nil
}

func (l *list) Front() *Element {
	return l.root.next
}

func (l *list) Back() *Element {
	return l.root.prev
}

// Remove removes e from its list
func (l *list) Remove(e *Element) {
	if e.prev == nil {
		l.root.next = e.next
	} else {
		e.prev.next = e.next
	}
	if e.next == nil {
		l.root.prev = e.prev
	} else {
		e.next.prev = e.prev
	}
	e.next = nil // avoid memory leaks
	e.prev = nil // avoid memory leaks
}

// PushFront inserts a new element e with value v at the front of list and returns e.
func (l *list) PushFront(key interface{}, value interface{}) *Element {
	e := &Element{Key: key, Value: value}
	if l.root.next == nil {
		// It's the first element
		l.root.next = e
		l.root.prev = e
		return e
	}

	e.next = l.root.next
	l.root.next.prev = e
	l.root.next = e
	return e
}

// PushBack inserts a new element e with value v at the back of list and returns e.
func (l *list) PushBack(key interface{}, value interface{}) *Element {
	e := &Element{Key: key, Value: value}
	if l.root.prev == nil {
		// It's the first element
		l.root.next = e
		l.root.prev = e
		return e
	}

	e.prev = l.root.prev
	l.root.prev.next = e
	l.root.prev = e
	return e
}
