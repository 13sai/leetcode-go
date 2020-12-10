## morris遍历

> morris遍历是二叉树遍历算法的进阶算法，可以将非递归遍历中的空间复杂度降为O(1)。morris遍历利用的是树的叶节点左右孩子为空（树的大量空闲指针），实现空间开销的缩减。

morris遍历的实现原则

- 记作当前节点为cur。
- 如果cur无左子节点，cur向右移动（cur=cur.right）
- 如果cur有左子节点，找到cur左子树上最右的节点，记为mostright
- 如果mostright的right指针指向空，让其指向cur，cur向左移动（cur=cur.left）
- 如果mostright的right指针指向cur，让其指向空，cur向右移动（cur=cur.right）

实现以上的原则，即实现了morris遍历。

### 中序遍历

```
func inorderTraversal(root *TreeNode) (res []int) {
	for root != nil {
		if root.Left != nil {
			// predecessor 节点表示当前 root 节点向左走一步，然后一直向右走至无法走为止的节点
			predecessor := root.Left
			for predecessor.Right != nil && predecessor.Right != root {
				// 有右子树且没有设置过指向 root，则继续向右走
				predecessor = predecessor.Right
			}
			if predecessor.Right == nil {
				// 将 predecessor 的右指针指向 root，这样后面遍历完左子树 root.Left 后，就能通过这个指向回到 root
				predecessor.Right = root
				// 遍历左子树
				root = root.Left
			} else { // predecessor 的右指针已经指向了 root，则表示左子树 root.Left 已经访问完了
				res = append(res, root.Val)
				// 恢复原样
				predecessor.Right = nil
				// 遍历右子树
				root = root.Right
			}
		} else { // 没有左子树
			res = append(res, root.Val)
			// 若有右子树，则遍历右子树
			// 若没有右子树，则整颗左子树已遍历完，root 会通过之前设置的指向回到这颗子树的父节点
			root = root.Right
		}
	}
	return
}
```

Morris 遍历的核心思想是利用树的大量空闲指针，实现空间开销的极限缩减。其后序遍历规则总结如下：
1. 新建临时节点，令该节点为 root；
2. 如果当前节点的左子节点为空，则遍历当前节点的右子节点；
3. 如果当前节点的左子节点不为空，在当前节点的左子树中找到当前节点在中序遍历下的前驱节点；
- 如果前驱节点的右子节点为空，将前驱节点的右子节点设置为当前节点，当前节点更新为当前节点的左子节点。
- 如果前驱节点的右子节点为当前节点，将它的右子节点重新设为空。倒序输出从当前节点的左子节点到该前驱节点这条路径上的所有节点。当前节点更新为当前节点的右子节点。
4. 重复步骤 2 和步骤 3，直到遍历结束。





