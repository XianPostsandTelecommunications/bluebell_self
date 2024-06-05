package badword

import (
	"bluebell/models"
	"bufio"
	"os"
	"sync"
)

// TrieV1Node 表示TrieV1树的节点
type TrieV1Node struct {
	children map[rune]*TrieV1Node // 子节点
	isEnd    bool
	Text     string
	Value    rune
	parent   *TrieV1Node // 父节点
}

// TrieV1 表示敏感词的TrieV1树
type TrieV1 struct {
	root *TrieV1Node
	lock sync.RWMutex
}

// NewTrieV1 创建一个新的TrieV1树
func NewTrieV1() *TrieV1 {
	return &TrieV1{
		root: &TrieV1Node{
			children: make(map[rune]*TrieV1Node),
			isEnd:    false,
		},
	}
}

// Insert 将一个敏感词插入到TrieV1树中
func (t *TrieV1) Insert(word string) {
	t.lock.Lock()
	defer t.lock.Unlock()

	node := t.root
	for _, char := range []rune(word) {
		if _, ok := node.children[char]; !ok {
			node.children[char] = &TrieV1Node{
				children: make(map[rune]*TrieV1Node),
				isEnd:    false,
				parent:   node,
				Value:    char,
			}
		}
		node = node.children[char]
	}

	node.Text = word
	node.isEnd = true
}

// Contains 检测文本中是否包含敏感词
func (t *TrieV1) Contains(text string) bool {
	t.lock.RLock()
	defer t.lock.RUnlock()
	node := t.root
	for _, char := range []rune(text) {
		if _, ok := node.children[char]; !ok {
			continue
		}
		node = node.children[char]
		if node.isEnd {
			return true
		}
	}
	return false
}

// Check 检测文本中是否包含敏感词，并返回第一个敏感词
func (t *TrieV1) Check(text string) string {
	t.lock.RLock()
	defer t.lock.RUnlock()

	node := t.root
	for _, char := range text {
		if _, ok := node.children[char]; !ok {
			continue
		}
		node = node.children[char]
		if node.isEnd {
			return node.Text
		}
	}

	return ""
}

// Rebuild 重新构建敏感词树
func (t *TrieV1) Rebuild(words []string) {
	t.lock.Lock()
	defer t.lock.Unlock()

	t.root = &TrieV1Node{}

	for _, word := range words {
		t.Insert(word)
	}
}

// Delete 删除一个敏感词
func (t *TrieV1) Delete(word string) {
	t.lock.Lock()
	defer t.lock.Unlock()

	node := t.root

	for _, char := range []rune(word) {
		if _, ok := node.children[char]; !ok {
			return
		}
		node = node.children[char]

		if node.isEnd {
			node.isEnd = false
			node.Text = ""

			if len(node.children) > 0 { // 有子节点，不能删除
				break
			}

			// 递归删除
			t.doDel(node)
		}

	}
}

func (t *TrieV1) doDel(node *TrieV1Node) {
	// 再次判断是否可以删除
	if node == nil || len(node.children) > 0 {
		return
	}

	// 从上级节点的children中删除本节点
	delete(node.parent.children, node.Value)

	// 判断上一层节点是否可以删除
	t.doDel(node.parent)
}

// LoadWordsFromFile 从文件加载敏感词列表
func (t *TrieV1) LoadWordsFromFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	t.Rebuild(words)
	return nil
}

// CheckPostForSensitiveWords 检查帖子内容是否包含敏感词
func CheckPostForSensitiveWords(post *models.Post, trie *TrieV1) bool {
	return trie.Contains(post.Content)
}
