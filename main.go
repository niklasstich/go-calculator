package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/niklasstich/calculator/evaluation"
	"github.com/niklasstich/calculator/parser"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	a := app.New()
	w := a.NewWindow("Calculator")
	go func(a *fyne.App) {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
		fmt.Println("Press CTRL + C to quit.")
		<-sig
		fmt.Println("Exiting.")
		(*a).Quit()
	}(&a)

	hello := widget.NewLabel("Hello world")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Click me", func() {
			hello.SetText("Thanks :)")
		}),
	))
	w.ShowAndRun()
}

func Evaluate(input string) string {
	tokens, err := parser.TokenizeString(input)
	if err != nil {
		log.Fatalf("Failed to tokenize input: %v\n", err)
	}

	tokens, err = parser.ReformToRPN(tokens)
	if err != nil {
		log.Fatalf("Failed to reform input: %v\n", err)
	}

	result, err := evaluation.EvaluateRPNExpression(tokens)
	if err != nil {
		log.Fatalf("Failed to evaluate input: %v\n", err)
	}

	return result.String()
}
