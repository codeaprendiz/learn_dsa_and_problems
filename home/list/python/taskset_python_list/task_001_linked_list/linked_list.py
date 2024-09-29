# Create a Node class to create a node
class Node:
    def __init__(self, data):
        self.data = data
        self.next = None

# Create a LinkedList class
class LinkedList:
    def __init__(self):
        self.head = None

    # Method to add a node at begin of LL
    def insert_at_begin(self, data):
        new_node = Node(data)

        new_node.next = self.head
        self.head = new_node

    # Method to add a node at any index
    # Indexing starts from 0.
    def insert_at_index(self, data, index):
        new_node = Node(data)
        if index == 0:
            self.insert_at_begin(data)
            return

        current_position = 0
        current_node = self.head
        while current_node is not None and current_position != index-1: # 
            current_node = current_node.next
            current_position += 1

        if current_node is None:
            raise IndexError("Index out of bounds")

        new_node.next = current_node.next
        current_node.next = new_node
