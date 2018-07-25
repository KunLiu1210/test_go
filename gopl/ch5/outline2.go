package main

import "golang.org/x/net/html"

// forEachNode針對每個結點x,都會調用pre(x)和post(x)。
// pre和post都是可選的。
// 遍歷孩子結點之前,pre被調用
// 遍歷孩子結點之後，post被調用
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

