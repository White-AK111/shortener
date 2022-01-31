package gin

import link "github.com/White-AK111/shortener/internal/pkg"

type delivery struct {
	links link.Usecase
}

func New(links link.Usecase) link.GinDelivery {
	ret := delivery{
		links: links,
	}
	return ret
}
