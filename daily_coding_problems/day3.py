class Node:
    def __init__(self, val, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

    def serialize(self):
        leftSide = [None]
        rightSide = [None]
        if self.left:
            leftSide = self.left.serialize()
        if self.right:
            rightSide = self.right.serialize()
        return [self.val] + leftSide + rightSide

def deserialize(serialized_tree):
    print(serialized_tree)
    if len(serialized_tree) == 0:
        return

    value = serialized_tree[0]
    serialized_tree.pop(0)
    if value == None:
        return value
    
    tree = Node(value)
    tree.left = deserialize(serialized_tree)
    tree.right = deserialize(serialized_tree)
    return tree

def main():
    root = Node('10', Node('100'), Node('12'))
    serialized_tree = str(root.serialize())
    print(serialized_tree)
    new_tree = deserialize(serialized_tree.split(","))
    print(new_tree.left.left.val)
    
if __name__ == "__main__":
    main()

# Implementation copied from here https://learnersbucket.com/examples/algorithms/serialize-and-deserialize-binary-tree