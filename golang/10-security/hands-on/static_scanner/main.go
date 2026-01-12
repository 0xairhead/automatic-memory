package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: scanner <filename.go>")
		return
	}

	filename := os.Args[1]
	fset := token.NewFileSet()

	// Parse the file
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}

	fmt.Printf("Scanning %s for security issues...\n", filename)
	issuesFound := 0

	// Inspect the AST (Abstract Syntax Tree)
	ast.Inspect(node, func(n ast.Node) bool {
		// 1. Check imports for "unsafe"
		if importSpec, ok := n.(*ast.ImportSpec); ok {
			if importSpec.Path.Value == "\"unsafe\"" {
				fmt.Printf("⚠️  Line %d: Import of 'unsafe' package detected!\n", fset.Position(importSpec.Pos()).Line)
				issuesFound++
			}
		}

		// 2. Check for TODO comments (often technical debt/security risks)
		// (parser.ParseComments must be invalid for this to work, but comment map is separate)
		// Comments are attached to the file, let's scan them separately
		return true
	})

	// Scan comments
	for _, commentGroup := range node.Comments {
		for _, comment := range commentGroup.List {
			if containsTODO(comment.Text) {
				fmt.Printf("ℹ️  Line %d: TODO found: %s\n", fset.Position(comment.Pos()).Line, comment.Text)
				// issuesFound++ // Don't fail regarding TODOs, just inform.
			}
		}
	}

	if issuesFound == 0 {
		fmt.Println("✅ No critical issues found.")
	} else {
		fmt.Printf("❌ Found %d issues.\n", issuesFound)
	}
}

func containsTODO(text string) bool {
	// Simple check
	return len(text) > 4 && (text[0:4] == "TODO" || (len(text) > 5 && text[2:6] == "TODO"))
}
