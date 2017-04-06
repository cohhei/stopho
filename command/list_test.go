package command

import (
	"strings"
	"testing"
)

func TestCmdList_All(t *testing.T) {
	app, outStream := InitAppForTests()

	err := app.Run([]string{"stopho", "list"})
	if err != nil {
		t.Fatalf("expected no error from Run, got %s", err)
	}

	expected := `+----------------+------------------------------+
|      NAME      |             URL              |
+----------------+------------------------------+
| Unsplash       | https://unsplash.com         |
| Picjumbo       | https://picjumbo.com         |
| StockSnap.io   | https://stocksnap.io         |
| Negative Space | https://www.negativespace.co |
| Magdeleine     | http://magdeleine.co/        |
| Picography     | https://picography.co        |
+----------------+------------------------------+`

	if !strings.Contains(outStream.String(), expected) {
		t.Errorf("\nexpected: \n%s\n got: \n%s", expected, outStream.String())
	}
}

func TestCmdList_Quiet(t *testing.T) {
	app, outStream := InitAppForTests()

	err := app.Run([]string{"stopho", "list", "-q"})
	if err != nil {
		t.Fatalf("expected no error from Run, got %s", err)
	}

	expected := `https://unsplash.com
https://picjumbo.com
https://stocksnap.io
https://www.negativespace.co
http://magdeleine.co/
https://picography.co`

	if !strings.Contains(outStream.String(), expected) {
		t.Errorf("\nexpected: \n%s\n got: \n%s", expected, outStream.String())
	}
}

func TestCmdList_All_Jp(t *testing.T) {
	app, outStream := InitAppForTests()

	err := app.Run([]string{"stopho", "-j", "list"})
	if err != nil {
		t.Fatalf("expected no error from Run, got %s", err)
	}

	expected := `+----------------+------------------------------+
|      NAME      |             URL              |
+----------------+------------------------------+
| Unsplash       | https://unsplash.com         |
| Picjumbo       | https://picjumbo.com         |
| StockSnap.io   | https://stocksnap.io         |
| Negative Space | https://www.negativespace.co |
| Magdeleine     | http://magdeleine.co/        |
| Picography     | https://picography.co        |
| いらすとや     | http://www.irasutoya.com     |
| マンガルー     | https://www.mangaloo.jp      |
| ぱくたそ       | https://www.pakutaso.com     |
+----------------+------------------------------+`

	if !strings.Contains(outStream.String(), expected) {
		t.Errorf("\nexpected: \n%s\n got: \n%s", expected, outStream.String())
	}
}

func TestCmdList_Quiet_JP(t *testing.T) {
	app, outStream := InitAppForTests()

	err := app.Run([]string{"stopho", "list", "-j", "-q"})
	if err != nil {
		t.Fatalf("expected no error from Run, got %s", err)
	}

	expected := `https://unsplash.com
https://picjumbo.com
https://stocksnap.io
https://www.negativespace.co
http://magdeleine.co/
https://picography.co
http://www.irasutoya.com
https://www.mangaloo.jp
https://www.pakutaso.com`

	if !strings.Contains(outStream.String(), expected) {
		t.Errorf("\nexpected: \n%s\n got: \n%s", expected, outStream.String())
	}
}
