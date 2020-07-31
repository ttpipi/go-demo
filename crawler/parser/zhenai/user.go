package zhenai

import (
	"crawler/comm"
	"regexp"
)

var ()

type User struct {
	Url    string
	Name   string
	Age    string
	Gender string
}

var (
	nameRe   = regexp.MustCompile(`<h1 class="ceiling-name ib fl fs24 lh32 blue">([^>]+)</h1>`)
	idRe     = regexp.MustCompile(`.*album\.zhenai\.com/u/([0-9]+)`)
	ageRe    = regexp.MustCompile(`<td><span class="label">年龄：</span>([^>]+)</td>`)
	genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^>]+)</span></td>`)
)

func ParseUser(html []byte, url string) comm.Result {
	nameMatch := nameRe.FindSubmatch(html)
	idMatch := idRe.FindSubmatch([]byte(url))
	ageMatch := ageRe.FindAllSubmatch(html, -1)
	genderMatch := genderRe.FindAllSubmatch(html, -1)

	result := comm.Result{
		Profile: comm.Profile{
			Id: string(idMatch[1]),
		},
	}

	user := User{
		Url:    url,
		Name:   string(nameMatch[1]),
		Age:    string(ageMatch[0][1]),
		Gender: string((genderMatch[0][1])),
	}

	result.Profile.Data = user
	return result
}
