package main

import (
	"net/http"
	"strings"
)

type node struct {
	path     string
	children []*node

	// 如果这是叶子节点
	// 那么匹配上之后可以调用该方法
	handler handlerFunc
}

// 查找是否存在该子节点
func (n *node) findMatchChild(path string) (*node, bool) {
	var wildcardNode *node
	for _, child := range n.children {
		// 严格匹配
		if child.path == path && child.path != "*" {
			return child, true
		}
		// *模糊匹配
		if child.path == "*" {
			wildcardNode = child
		}
	}
	return wildcardNode, wildcardNode != nil
}

func (h *HandlerBasedOnTree) Route(method, pattern string, handlerFunc handlerFunc) {
	// 去掉左右两边的/
	pattern = strings.Trim(pattern, "/")
	paths := strings.Split(pattern, "/")

	cur := h.root
	for index, path := range paths {
		matchChild, ok := cur.findMatchChild(path)
		if ok {
			cur = matchChild
		} else {
			// 后面的都是找不到或者不存在的节点
			h.createSubTree(cur, paths[index:], handlerFunc)
			return
		}
	}
}

func (h *HandlerBasedOnTree) ServeHTTP(c *Context) {
	handler, found := h.findRouter(c.R.URL.Path)
	if !found {
		c.W.WriteHeader(http.StatusNotFound)
		_, _ = c.W.Write([]byte("Not Found"))
		return
	}
	handler(c) // 向下执行
}

func (h *HandlerBasedOnTree) createSubTree(root *node, paths []string, handlerFunc handlerFunc) {
	cur := root
	for _, path := range paths {
		newN := newNode(path)
		cur.children = append(cur.children, newN)
		cur = newN
	}
	cur.handler = handlerFunc
}

func (h *HandlerBasedOnTree) findRouter(path string) (handlerFunc, bool) {
	paths := strings.Split(strings.Trim(path, "/"), "/")
	cur := h.root
	for _, v := range paths {
		// 从子节点找一个对应的路由
		matchChild, ok := cur.findMatchChild(v)
		if !ok {
			return nil, false
		}
		cur = matchChild
	}
	if cur.handler == nil {
		return nil, false
	}
	return cur.handler, true
}

func newNode(path string) *node {
	return &node{
		path:     path,
		children: make([]*node, 0),
	}
}

type HandlerBasedOnTree struct {
	root *node
}
