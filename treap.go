
package treap

import (
	"bytes"
	"math/rand"
	"fmt"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

/*
 Things of any kind can be values in the Treap.
*/
type Thing interface{}

/*
 TreapIterators iterate over the Treap.
*/
type TreapIterator func(key []byte, value Thing)

/*
 Treap is based on http://faculty.washington.edu/aragon/treaps.html like all other treaps.
*/
type Treap struct {
	size int
	root *node
}
/*
 Put will insert something in the Treap under a given key. 
 It will overwrite any previous value, and return the previous value and whether there was a previous value.
*/
func (self *Treap) Put(key []byte, value Thing) (old Thing, existed bool) {
	self.root, existed, old = self.root.insert(newNode(key, value))
	if !existed {
		self.size++
	}
	return
}
/*
 Get will return something from the Treap. 
 It will return the value found, and whether any value was found.
*/
func (self *Treap) Get(key []byte) (value Thing, existed bool) {
	return self.root.get(key)
}
/*
 Describe returns a string describing the tree structure of the Treap.
*/
func (self *Treap) Describe() string {
	buffer := bytes.NewBufferString(fmt.Sprintf("<Treap size:%v>\n", self.Size()))
	self.root.describe(0, buffer)
	return string(buffer.Bytes())
}
/*
 Del will remove a value from the Treap.
 It will return the value deleted and whether a value was deleted.
*/
func (self *Treap) Del(key []byte) (old Thing, existed bool) {
	self.root, existed, old = self.root.del(key)
	if existed {
		self.size--
	}
	return
}
/*
 Up will execute the given iterator on every key/value pair in the Treap where from <= key < below.
*/
func (self *Treap) Up(from, below []byte, f TreapIterator) {
	self.root.up(from, below, f)
}
/*
 Down will execute the given iterator on every key/value pair in the Treap where from >= key > above.
*/
func (self *Treap) Down(from, above []byte, f TreapIterator) {
	self.root.down(from, above, f)
}
/*
 Size will return the size of the Treap.
*/
func (self *Treap) Size() int {
	return self.size
}
/*
 String will return a simple string representation of the Treap.
*/
func (self *Treap) String() string {
	return fmt.Sprint(self.ToMap())
}
/*
 ToMap will return a map[string]Thing with the Treap keys converted to string.
*/
func (self *Treap) ToMap() map[string]Thing {
	rval := make(map[string]Thing)
	self.Up(nil, nil, func(key []byte, value Thing) {
		rval[string(key)] = value
	})
	return rval
}

func merge(left, right *node) (result *node) {
	if left == nil {
		result = right
	} else if right == nil {
		result = left
	} else if left.weight < right.weight {
		result, left.right = left, merge(left.right, right)
	} else {
		result, right.left = right, merge(right.left, left)
	}
	return
}

type node struct {
	weight int32
	left *node
	right *node
	key []byte
	value Thing
}
func newNode(key []byte, value Thing) (rval *node) {
	rval = &node{
		weight: rand.Int31(),
		key: key,
		value: value,
	}
	return
}
func (self *node) describe(indent int, buffer *bytes.Buffer) {
	if self != nil {
		self.left.describe(indent + 1, buffer)
		indentation := &bytes.Buffer{}
		for i := 0; i < indent; i++ {
			fmt.Fprint(indentation, " ")
		}
		fmt.Fprintf(buffer, "%v%v [%v] => %v\n", string(indentation.Bytes()), self.key, self.weight, self.value)
		self.right.describe(indent + 1, buffer)
	}
}
func (self *node) get(key []byte) (value Thing, existed bool) {
	if self != nil {
		switch bytes.Compare(key, self.key) {
		case -1:
			return self.left.get(key)
		case 1:
			return self.right.get(key)
		default:
			value, existed = self.value, true
		}
		
	}
	return
}
func (self *node) up(from, below []byte, f TreapIterator) {
	if self != nil {
		from_cmp := -1
		if from != nil {
			from_cmp = bytes.Compare(from, self.key)
		}
		below_cmp := 1
		if below != nil {
			below_cmp = bytes.Compare(below, self.key)
		}

		if from_cmp < 0 {
			self.left.up(from, below, f)
		}
		if from_cmp < 1 && below_cmp > 0 {
			f(self.key, self.value)
		}
		if below_cmp > 0 {
			self.right.up(from, below, f)
		}
	}
}
func (self *node) down(from, above []byte, f TreapIterator) {
	if self != nil {
		from_cmp := -1
		if from != nil {
			from_cmp = bytes.Compare(from, self.key)
		}
		above_cmp := 1
		if above != nil {
			above_cmp = bytes.Compare(above, self.key)
		}

		if from_cmp > 0 {
			self.right.down(from, above, f)
		}
		if from_cmp > -1 && above_cmp < 0 {
			f(self.key, self.value)
		}
		if above_cmp < 0 {
			self.left.down(from, above, f)
		}
	}
}
func (self *node) rotateLeft() (result *node) {
	result, self.left, self.left.right = self.left, self.left.right, self
	return
}
func (self *node) rotateRight() (result *node) {
	result, self.right, self.right.left = self.right, self.right.left, self
	return
}
func (self *node) del(key []byte) (result *node, existed bool, old Thing) {
	if self != nil {
		result = self
		switch bytes.Compare(key, self.key) {
		case -1:
			self.left, existed, old = self.left.del(key)
		case 1:
			self.right, existed, old = self.right.del(key)
		default:
			result, existed, old = merge(self.left, self.right), true, self.value
		}
	}
	return
}
func (self *node) insert(n *node) (result *node, existed bool, old Thing) {
	result = n
	if self != nil {
		result = self
		switch bytes.Compare(n.key, self.key) {
		case -1:
			self.left, existed, old = self.left.insert(n)
			if self.left.weight < self.weight {
				result = self.rotateLeft()
			}
		case 1:
			self.right, existed, old = self.right.insert(n)
			if self.right.weight < self.weight {
				result = self.rotateRight()
			}
		default:
			existed, old, self.value = true, self.value, n.value
		}
	}
	return
}