package viewmodel

type Shop struct {
	Title      string
	Active     string
	Categories []Category
}

type Category struct {
	URL         string
	ImageURL    string
	Title       string
	Description string
}

func NewShop() Shop {
	result := Shop{
		Title:  "Penny Jewlery - Shop",
		Active: "shop",
	}
	juiceCategory := Category{
		URL:      "/shop_details",
		ImageURL: "lemon.png",
		Title:    "Juices and Mixes",
		Description: `Explore our wide variety of juices
						to temp the taste buds`,
	}
	supplyCategory := Category{
		URL:         ".",
		ImageURL:    "kiwi.png",
		Title:       "Cups, Straws and other supplies",
		Description: `From paper cups to plates, we have it all`,
	}
	advertiseCategory := Category{
		URL:         ".",
		ImageURL:    "pineapple.png",
		Title:       "Signs and Advertising",
		Description: `Wait for someone to find your business, this is not a good way to do it.`,
	}
	result.Categories = []Category{juiceCategory, supplyCategory, advertiseCategory}
	return result
}
