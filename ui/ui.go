package ui

import (
	"github.com/nraval1729/termnews/news"
	"fmt"
	"github.com/jroimartin/gocui"
	"github.com/pkg/browser"
	"math"
	"strings"
)

const articlesViewTitle = "articles"
const descriptionViewTitle = "description"
const pageSize = 25

var currPage = 0

func Run() error {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		return fmt.Errorf("ui.Run()::gocui.NewGui() threw %v\n", err)
	}
	defer g.Close()

	g.SetManagerFunc(func(gui *gocui.Gui) error {
		return layout(g)
	})

	if err := setKeybindings(g); err != nil {
		return fmt.Errorf("ui.Run()::ui.setKeybindings() threw %v\n", err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		return fmt.Errorf("ui.Run()::g.MainLopp() threw %v\n", err)
	}
	return nil
}

func layout(g *gocui.Gui) error {
	err := layoutArticles(g, news.GetNews())
	if err != nil {
		panic(err)
	}

	err = layoutDescription(g, news.GetNews())
	if err != nil {
		panic(err)
	}

	_, err = g.SetCurrentView(articlesViewTitle)
	if err != nil {
		panic(err)
	}
	return nil
}

func layoutDescription(g *gocui.Gui, nr *news.Resp) error {
	maxX, maxY := g.Size()

	v, err := g.SetView(descriptionViewTitle, 0, 5*maxY/6, maxX-1, maxY-1)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Title = getDescriptionViewTitle()
	v.Wrap = true
	articlesView, err := g.View(articlesViewTitle)
	if err != nil {
		return err
	}
	_, y := articlesView.Cursor()
	v.Clear()
	fmt.Fprintf(v, nr.Articles[getCurrPage()*pageSize+y].Description)

	return nil
}

func layoutArticles(g *gocui.Gui, nr *news.Resp) error {
	maxX, maxY := g.Size()
	v, err := g.SetView(articlesViewTitle, 0, 0, maxX-1, 5*maxY/6-1)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}
	v.Title = getArticlesViewTitle(len(nr.Articles))
	v.Highlight = true
	v.SelBgColor = gocui.ColorGreen
	v.SelFgColor = gocui.ColorBlack
	v.Clear()

	fmt.Fprintf(v, constructArticlesData(nr))
	fmt.Fprintf(v, getShortcutsHelpText())

	return nil
}

func constructArticlesData(nr *news.Resp) string {
	var str strings.Builder
	for _, article := range getArticlesOnCurrentPage(nr) {
		str.WriteString(getFormattedArticleItem(article.Title, article.Source.ID))
	}
	return str.String()
}

func getArticlesOnCurrentPage(nr *news.Resp) []news.Article {
	return nr.Articles[getCurrPage()*pageSize : min(len(nr.Articles), getCurrPage()*pageSize+pageSize)]
}

func setKeybindings(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	if err := g.SetKeybinding(articlesViewTitle, gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		return err
	}
	if err := g.SetKeybinding(articlesViewTitle, gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		return err
	}
	if err := g.SetKeybinding(articlesViewTitle, gocui.KeyArrowRight, gocui.ModNone, cursorRight); err != nil {
		return err
	}
	if err := g.SetKeybinding(articlesViewTitle, gocui.KeyCtrlD, gocui.ModNone, nextPage); err != nil {
		return err
	}
	if err := g.SetKeybinding(articlesViewTitle, gocui.KeyCtrlA, gocui.ModNone, prevPage); err != nil {
		return err
	}

	return nil
}

// Keybinding handlers
func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func cursorDown(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if cy < len(getArticlesOnCurrentPage(news.GetNews()))-1 {
			err := v.SetCursor(cx, cy+1)
			return err
		}
	}
	return nil
}

func cursorUp(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		cx, cy := v.Cursor()
		if cy > 0 {
			err := v.SetCursor(cx, cy-1)
			return err
		}
	}
	return nil
}

func cursorRight(g *gocui.Gui, v *gocui.View) error {
	if v != nil {
		_, currArticleIndex := v.Cursor()
		_ = browser.OpenURL(news.GetNews().Articles[getCurrPage()*pageSize+currArticleIndex].URL)
	}
	return nil
}

func nextPage(g *gocui.Gui, v *gocui.View) error {
	if v != nil && (getCurrPage()+1)*pageSize < len(news.GetNews().Articles) {
		incCurrPage()
		g.Update(func(gui *gocui.Gui) error {
			return layout(g)
		})
	}
	return nil
}

func prevPage(g *gocui.Gui, v *gocui.View) error {
	if v != nil && getCurrPage() >= 1 {
		decCurrPage()
		g.Update(func(gui *gocui.Gui) error {
			return layout(g)
		})
	}
	return nil
}

// Utilities
func getArticlesViewTitle(totalArticles int) string {
	return fmt.Sprintf(" %s (%d/%d) ", articlesViewTitle, getCurrPage()+1, getTotalPageCount(totalArticles))
}

func getShortcutsHelpText() string {
	return fmt.Sprintf("\n\n\nCtrlA/CtrlD [prev/next page], Cursor Up/Down [up/down one article], Cursor Right [open article in browser]")

}

func getTotalPageCount(totalArticles int) int {
	return int(math.Ceil(float64(totalArticles) / float64(pageSize)))
}

func getDescriptionViewTitle() string {
	return fmt.Sprintf(" %s ", descriptionViewTitle)
}

func getFormattedArticleItem(title, source string) string {
	return fmt.Sprintf("**\t\t%s => [%s]\n", title, source)
}

func getCurrPage() int {
	return currPage
}

func incCurrPage() {
	currPage++
}

func decCurrPage() {
	currPage--
}

// There's no inbuilt min function to compare ints :(
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
