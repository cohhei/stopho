package command

import (
	"strings"
	"testing"
)

func TestCmdSearch_WithoutSearchTerm(t *testing.T) {
	app, outStream := InitAppForTests()

	err := app.Run([]string{"stopho", "search"})
	if err != nil {
		t.Fatalf("expected no error from Run, got %s", err)
	}

	expected := "Specify a search term!"

	if !strings.Contains(outStream.String(), expected) {
		t.Errorf("\nexpected: \n%s\n got: \n%s", expected, outStream.String())
	}
}

func TestCmdSearch_All(t *testing.T) {
	app, outStream := InitAppForTests()

	err := app.Run([]string{"stopho", "search", "term"})
	if err != nil {
		t.Fatalf("expected no error from Run, got %s", err)
	}

	expected := `+----------------+--------------------------------------+
|      NAME      |                 URL                  |
+----------------+--------------------------------------+
| Unsplash       | https://unsplash.com/search/term     |
| Picjumbo       | https://picjumbo.com/?s=term         |
| StockSnap.io   | https://stocksnap.io/search/term     |
| Negative Space | https://www.negativespace.co/?s=term |
| Magdeleine     | http://magdeleine.co/?s=term         |
| Picography     | https://picography.co/?s=term        |
+----------------+--------------------------------------+`

	if !strings.Contains(outStream.String(), expected) {
		t.Errorf("\nexpected: \n%s\n got: \n%s", expected, outStream.String())
	}
}

func TestCmdSearch_Quiet(t *testing.T) {
	app, outStream := InitAppForTests()

	err := app.Run([]string{"stopho", "search", "-q", "term"})
	if err != nil {
		t.Fatalf("expected no error from Run, got %s", err)
	}

	expected := `https://unsplash.com/search/term
https://picjumbo.com/?s=term
https://stocksnap.io/search/term
https://www.negativespace.co/?s=term
http://magdeleine.co/?s=term
https://picography.co/?s=term`

	if !strings.Contains(outStream.String(), expected) {
		t.Errorf("\nexpected: \n%s\n got: \n%s", expected, outStream.String())
	}
}

func TestCmdSearch_All_Jp(t *testing.T) {
	app, outStream := InitAppForTests()

	err := app.Run([]string{"stopho", "-j", "search", "term"})
	if err != nil {
		t.Fatalf("expected no error from Run, got %s", err)
	}

	expected := `+----------------+--------------------------------------------------------------------+
|      NAME      |                                URL                                 |
+----------------+--------------------------------------------------------------------+
| Unsplash       | https://unsplash.com/search/term                                   |
| Picjumbo       | https://picjumbo.com/?s=term                                       |
| StockSnap.io   | https://stocksnap.io/search/term                                   |
| Negative Space | https://www.negativespace.co/?s=term                               |
| Magdeleine     | http://magdeleine.co/?s=term                                       |
| Picography     | https://picography.co/?s=term                                      |
| いらすとや     | http://www.irasutoya.com/search?q=term                             |
| マンガルー     | https://www.mangaloo.jp/m/comas?coma_search_form%5Bkeyword%5D=term |
| ぱくたそ       | https://www.pakutaso.com/search.html?search=term                   |
+----------------+--------------------------------------------------------------------+`

	if !strings.Contains(outStream.String(), expected) {
		t.Errorf("\nexpected: \n%s\n got: \n%s", expected, outStream.String())
	}
}

func TestCmdSearch_Quiet_Jp(t *testing.T) {
	app, outStream := InitAppForTests()

	err := app.Run([]string{"stopho", "search", "-j", "-q", "term"})
	if err != nil {
		t.Fatalf("expected no error from Run, got %s", err)
	}

	expected := `https://unsplash.com/search/term
https://picjumbo.com/?s=term
https://stocksnap.io/search/term
https://www.negativespace.co/?s=term
http://magdeleine.co/?s=term
https://picography.co/?s=term
http://www.irasutoya.com/search?q=term
https://www.mangaloo.jp/m/comas?coma_search_form%5Bkeyword%5D=term
https://www.pakutaso.com/search.html?search=term`

	if !strings.Contains(outStream.String(), expected) {
		t.Errorf("\nexpected: \n%s\n got: \n%s", expected, outStream.String())
	}
}

func TestCmdSearch_WithWrongSiteName(t *testing.T) {
	app, outStream := InitAppForTests()

	err := app.Run([]string{"stopho", "search", "-s", "wrongsitename", "term"})
	if err != nil {
		t.Fatalf("expected no error from Run, got %s", err)
	}

	expected := "No site matched `wrongsitename`"

	if !strings.Contains(outStream.String(), expected) {
		t.Errorf("\nexpected: \n%s\n got: \n%s", expected, outStream.String())
	}
}

func TestCmdSearch_WithSiteName(t *testing.T) {
	app, outStream := InitAppForTests()

	err := app.Run([]string{"stopho", "search", "-s", "Stock", "term"})
	if err != nil {
		t.Fatalf("expected no error from Run, got %s", err)
	}

	expected := `+--------------+----------------------------------+
|     NAME     |               URL                |
+--------------+----------------------------------+
| StockSnap.io | https://stocksnap.io/search/term |
+--------------+----------------------------------+`

	if !strings.Contains(outStream.String(), expected) {
		t.Errorf("\nexpected: \n%s\n got: \n%s", expected, outStream.String())
	}
}
