package command

type Site struct {
	Name      string `yaml:"name"`
	Url       string `yaml:"url"`
	SearchUrl string `yaml:"search_url"`
}

var Sites = []Site{
	{
		Name:      "Unsplash",
		Url:       "https://unsplash.com",
		SearchUrl: "https://unsplash.com/search/",
	},
	{
		Name:      "Picjumbo",
		Url:       "https://picjumbo.com",
		SearchUrl: "https://picjumbo.com/?s=",
	},
	{
		Name:      "StockSnap.io",
		Url:       "https://stocksnap.io",
		SearchUrl: "https://stocksnap.io/search/",
	},
	{
		Name:      "Negative Space",
		Url:       "https://www.negativespace.co",
		SearchUrl: "https://www.negativespace.co/?s=",
	},
	{
		Name:      "Magdeleine",
		Url:       "http://magdeleine.co/",
		SearchUrl: "http://magdeleine.co/?s=",
	},
	{
		Name:      "Picography",
		Url:       "https://picography.co",
		SearchUrl: "https://picography.co/?s=",
	},
}

var Sitesjp = []Site{
	{
		Name:      "いらすとや",
		Url:       "http://www.irasutoya.com",
		SearchUrl: "http://www.irasutoya.com/search?q=",
	},
	{
		Name:      "マンガルー",
		Url:       "https://www.mangaloo.jp",
		SearchUrl: "https://www.mangaloo.jp/m/comas?coma_search_form%5Bkeyword%5D=",
	},
	{
		Name:      "ぱくたそ",
		Url:       "https://www.pakutaso.com",
		SearchUrl: "https://www.pakutaso.com/search.html?search=",
	},
	{
		Name:      "ピクシブ",
		Url:       "https://www.pixiv.net",
		SearchUrl: "https://www.pixiv.net/search.php?word=",
	},
}
