package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type TreeNode struct {
	Name     string
	Children map[string]*TreeNode
}

func NewTreeNode(name string) *TreeNode {
	return &TreeNode{
		Name:     name,
		Children: make(map[string]*TreeNode),
	}
}

func main() {
	// Verifica se um arquivo foi fornecido como argumento
	if len(os.Args) != 2 {
		fmt.Println("Uso: urltree arquivo_de_urls")
		os.Exit(1)
	}

	inputFile := os.Args[1]

	// Abre o arquivo de entrada
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		os.Exit(1)
	}
	defer file.Close()

	root := NewTreeNode("")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url := scanner.Text()
		parts := strings.Split(url, "/")
		current := root

		for _, part := range parts {
			if part != "" {
				child, exists := current.Children[part]
				if !exists {
					child = NewTreeNode(part)
					current.Children[part] = child
				}
				current = child
			}
		}
	}

	// Função recursiva para imprimir a árvore com formatação melhorada
	var printTree func(node *TreeNode, indent string, last bool)
	printTree = func(node *TreeNode, indent string, last bool) {
		if node.Name != "" {
			fmt.Print(indent)
			if last {
				fmt.Print("└─ ")
				indent += "    "
			} else {
				fmt.Print("├─ ")
				indent += "│   "
			}
			fmt.Println(node.Name)
		}

		keys := make([]string, 0, len(node.Children))
		for key := range node.Children {
			keys = append(keys, key)
		}
		for i, key := range keys {
			child := node.Children[key]
			printTree(child, indent, i == len(keys)-1)
		}
	}

	// Imprime a árvore com formatação melhorada, excluindo o nó raiz vazio
	printTree(root, "", true)

	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
		os.Exit(1)
	}
}

