package viewmodel

//Shop model
type Shop struct {
	Title string
	Active string
	Categories []Category 
}

//Category model
type Category struct {
	URL string
	ImageURL string
	Title string
	Description string	
}

//NewShop model
func NewShop() Shop{
	result := Shop{
		Title: "Lomenade Stand Supply - Shop",
		Active: "shop",
	}

	juiceCategory := Category{
		URL : "/shop_detals",
		ImageURL: "lemon.png",
		Title: "Juncie and Mixes",
		Description: "Explore our wide assortment of juices and mixes expected by today's lemonade stand clientelle. Now featuring a full line of organic juices that are guaranteed to be obtained from trees that have never been treated with pesticides or artificial fertilizers.",
	}

	supplyCategory := Category {
		URL : ".",
		ImageURL: "kiwi.png",
		Title: "Cups, Straws, and Other Supplies",
		Description: "From paper cups to bio-degradable plastic to straws and napkins, LSS is your source for the sundries that keep your stand running smoothly.",
	}

	advertiseCategory := Category{
		URL : ".",
		ImageURL : "pineapple.png",
		Title : "Sign and Advertising",
		Description: "Sure, you could just wait for people to find your stand along the side of the road, but if you want to take it to the next level, our premium line of advertising supplies.",
	}
	result.Categories = []Category{juiceCategory,supplyCategory,advertiseCategory}
	return result;
}