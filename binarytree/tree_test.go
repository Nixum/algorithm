package binarytree

import (
	"fmt"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	root := CreateBinaryTree([]int{1, -1, 2, -1, -1, 3})
	fmt.Println(preorderTraversal(root))

	root = CreateBinaryTree([]int{1, 2, 3, 4, -1, 5, 6, 7})
	fmt.Println(TraversalByPreorderInIteration(root))

	root = CreateBinaryTree([]int{1, 2, 3, 4, -1, 5, 6, 7})
	fmt.Println(TraversalByPreorderInIteration2(root))
}

func TestPostorderTraversal(t *testing.T) {
	root := CreateBinaryTree([]int{1, -1, 2, -1, -1, 3})
	fmt.Println(postorderTraversal(root))

	root = CreateBinaryTree([]int{1, 2, 3, 4, -1, 5, 6, 7})
	fmt.Println(TraversalByPostorderInIteration(root))

	root = CreateBinaryTree([]int{1, 2, 3, 4, -1, 5, 6, 7})
	fmt.Println(TraversalByPostorderInIteration2(root))
}

func TestInorderTraversal(t *testing.T) {
	root := CreateBinaryTree([]int{1, -1, 2, -1, -1, 3})
	fmt.Println(inorderTraversal(root))

	root = CreateBinaryTree([]int{1, 2, 3, 4, -1, 5, 6, 7})
	fmt.Println(TraversalByInorderInIteration(root))

	root = CreateBinaryTree([]int{1, 2, 3, 4, -1, 5, 6, 7})
	fmt.Println(TraversalByInorderInIteration2(root))
}

func TestLevelOrder(t *testing.T) {
	root := CreateBinaryTree([]int{1, -1, 2, -1, -1, 3})
	fmt.Println(levelOrder(root))

	root = CreateBinaryTree([]int{1, 2, 3, 4, -1, 5, 6, 7})
	fmt.Println(levelOrder(root))

	root = CreateBinaryTree([]int{})
	fmt.Println(levelOrder(root))
}

func TestLevelOrderBottom(t *testing.T) {
	root := CreateBinaryTree([]int{1, -1, 2, -1, -1, 3})
	fmt.Println(levelOrderBottom(root))

	root = CreateBinaryTree([]int{1, 2, 3, 4, -1, 5, 6, 7})
	fmt.Println(levelOrderBottom(root))

	root = CreateBinaryTree([]int{})
	fmt.Println(levelOrderBottom(root))
}

func TestRightSideView(t *testing.T) {
	root := CreateBinaryTree([]int{1,2,3,-1,5,-1,4})
	fmt.Println(rightSideView(root))

	root = CreateBinaryTree([]int{1,-1,3})
	fmt.Println(rightSideView(root))

	root = CreateBinaryTree([]int{})
	fmt.Println(rightSideView(root))

	root = CreateBinaryTree([]int{1, 2})
	fmt.Println(rightSideView(root))
}

func TestAverageOfLevels(t *testing.T) {
	root := CreateBinaryTree([]int{3,9,20,-1,-1,15,7})
	fmt.Println(averageOfLevels(root))

	root = CreateBinaryTree([]int{3,9,20,15,7})
	fmt.Println(averageOfLevels(root))

	root = CreateBinaryTree([]int{})
	fmt.Println(averageOfLevels(root))

	root = CreateBinaryTree([]int{1, 2})
	fmt.Println(averageOfLevels(root))
}

func TestLevelOrder429(t *testing.T) {
	root := &Node{
		Val:      1,
		Children: []*Node{
			{
				Val:      4,
				Children: []*Node{
					{
						Val:      11,
						Children: nil,
					},
				},
			},
			{
				Val:      3,
				Children: []*Node{
					{
						Val:      31,
						Children: nil,
					},
					{
						Val:      21,
						Children: []*Node{
							{
								Val:      231,
								Children: nil,
							},
						},
					},
					{
						Val:      41,
						Children: nil,
					},
				},
			},
			{
				Val:      66,
				Children: nil,
			},
		},
	}
	fmt.Println(levelOrder429(root))
}

func TestLargestValues(t *testing.T) {
	root := CreateBinaryTree([]int{1,3,2,5,3,-1,9})
	fmt.Println(largestValues(root))

	root = CreateBinaryTree([]int{1,2,3})
	fmt.Println(largestValues(root))
}

func TestMaxDepth(t *testing.T) {
	root := CreateBinaryTree([]int{3,9,20,-1,-1,15,7})
	fmt.Println(maxDepth(root))

	root = CreateBinaryTree([]int{1,-1,2})
	fmt.Println(maxDepth(root))
}

func TestMinDepth(t *testing.T) {
	root := CreateBinaryTree([]int{1,2,3,4,5})
	fmt.Println(minDepth(root))

	root = CreateBinaryTree([]int{3,9,20,-1,-1,15,7})
	fmt.Println(minDepth(root))

	root = CreateBinaryTree([]int{1,-1,2})
	fmt.Println(minDepth(root))

	root = CreateBinaryTree([]int{2,-1,3,-1,4,-1,5,-1,6})
	fmt.Println(minDepth(root))
}

func TestBinaryTreePaths(t *testing.T) {
	root := CreateBinaryTree([]int{1,2,3,4,5})
	fmt.Println(binaryTreePaths(root))

	root = CreateBinaryTree([]int{3,9,20,-1,-1,15,7})
	fmt.Println(binaryTreePaths(root))

	root = CreateBinaryTree([]int{1,-1,2})
	fmt.Println(binaryTreePaths(root))

	root = CreateBinaryTree([]int{2,-1,3,-1,4,-1,5,-1,6})
	fmt.Println(binaryTreePaths(root))
}

func TestIsSameTree(t *testing.T) {
	root1 := CreateBinaryTree([]int{1,2,3,4,5})
	root2 := CreateBinaryTree([]int{1,2,3,4,5})
	fmt.Println(isSameTree(root1, root2))

	root1 = CreateBinaryTree([]int{1,2})
	root2 = CreateBinaryTree([]int{1, -1, 2})
	fmt.Println(isSameTree(root1, root2))

	root1 = CreateBinaryTree([]int{1,2, 1})
	root2 = CreateBinaryTree([]int{1, 1, 2})
	fmt.Println(isSameTree(root1, root2))
}

func TestSumOfLeftLeaves(t *testing.T) {
	root := CreateBinaryTree([]int{3,9,20,-1,-1,15,7})
	fmt.Println(sumOfLeftLeaves2(root))

	root = CreateBinaryTree([]int{1})
	fmt.Println(sumOfLeftLeaves(root))

	root = CreateBinaryTree([]int{1,-1,2})
	fmt.Println(sumOfLeftLeaves(root))

	root = CreateBinaryTree([]int{2,-1,3,-1,4,-1,5,-1,6})
	fmt.Println(sumOfLeftLeaves(root))
}

func TestFindBottomLeftValue(t *testing.T) {
	root := CreateBinaryTree([]int{3,9,20,-1,-1,15,7})
	fmt.Println(findBottomLeftValue(root))

	root = CreateBinaryTree([]int{1})
	fmt.Println(findBottomLeftValue(root))

	root = CreateBinaryTree([]int{2,1,3})
	fmt.Println(findBottomLeftValue(root))

	root = CreateBinaryTree([]int{1,2,3,4,-1,5,6,-1,-1,7})
	fmt.Println(findBottomLeftValue(root))
}

func TestHasPathSum(t *testing.T) {
	root := CreateBinaryTree([]int{5,4,8,11,-1,13,4,7,2,-1,-1,-1,1})
	fmt.Println(hasPathSum2(root, 22))

	root = CreateBinaryTree([]int{1, 2})
	fmt.Println(hasPathSum2(root, 1))

	root = CreateBinaryTree([]int{})
	fmt.Println(hasPathSum2(root, 0))

	root = CreateBinaryTree([]int{1,2,3})
	fmt.Println(hasPathSum2(root, 5))
}

func TestPathSum(t *testing.T) {
	root := CreateBinaryTree([]int{5,4,8,11,-1,13,4,7,2,-1,-1, -1, -1,5,1})
	fmt.Println(pathSum(root, 22))

	root = CreateBinaryTree([]int{1, 2, 3})
	fmt.Println(pathSum(root, 5))

	root = CreateBinaryTree([]int{})
	fmt.Println(pathSum(root, 0))

	root = CreateBinaryTree([]int{1,2})
	fmt.Println(pathSum(root, 0))
}

func TestGetMinimumDifference(t *testing.T) {
	root := CreateBinaryTree([]int{4,2,6,1,3})
	fmt.Println(getMinimumDifference(root))

	root = CreateBinaryTree([]int{1,0,48,-1,-1,12,49})
	fmt.Println(getMinimumDifference(root))

	root = CreateBinaryTree([]int{})
	fmt.Println(getMinimumDifference(root))

	root = CreateBinaryTree([]int{1,2})
	fmt.Println(getMinimumDifference(root))
}
