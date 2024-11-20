package cmd

import (
	"fmt"
	"os"

	"github.com/richardamare/countlines/internal/counter"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "countlines [extension]",
	Short: "Count lines of code in files with specific extension",
	Long: `Count lines of code in files with a specific extension.
Respects .gitignore rules if present in the current or parent directories.
Example: countlines go`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		extension := args[0]
		counts, err := counter.CountLines(extension)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		// Print results
		var total int
		for file, count := range counts {
			fmt.Printf("%d\t%s\n", count, file)
			total += count
		}
		fmt.Printf("\nTotal: %d lines in %d files\n", total, len(counts))
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
