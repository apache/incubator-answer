package templaterender

import (
	"math"

	"github.com/answerdev/answer/internal/schema"
	"github.com/answerdev/answer/internal/service"
	"github.com/answerdev/answer/internal/service/tag"
	"github.com/google/wire"
)

// ProviderSetTemplateRenderController is template render controller providers.
var ProviderSetTemplateRenderController = wire.NewSet(
	NewTemplateRenderController,
)

type TemplateRenderController struct {
	questionService *service.QuestionService
	userService     *service.UserService
	tagService      *tag.TagService
}

func NewTemplateRenderController(
	questionService *service.QuestionService,
	userService *service.UserService,
	tagService *tag.TagService,

) *TemplateRenderController {
	return &TemplateRenderController{
		questionService: questionService,
		userService:     userService,
		tagService:      tagService,
	}
}

// Paginator page
// page : now page
// prepage : Number per page
// nums : Total
// Returns the contents of the page in the format of 1, 2, 3, 4, and 5. If the contents are less than 5 pages, the page number is returned
func Paginator(page, prepage int, nums int64) *schema.Paginator {
	if prepage == 0 {
		prepage = 10
	}

	var firstpage int //Previous page address
	var lastpage int  //Address on the last page
	//Generate the total number of pages based on the total number of nums and the number of prepage pages
	totalpages := int(math.Ceil(float64(nums) / float64(prepage))) //Total number of Pages
	if page > totalpages {
		page = totalpages
	}
	if page <= 0 {
		page = 1
	}
	var pages []int
	switch {
	case page >= totalpages-5 && totalpages > 5: //The last 5 pages
		start := totalpages - 5 + 1
		firstpage = page - 1
		lastpage = int(math.Min(float64(totalpages), float64(page+1)))
		pages = make([]int, 5)
		for i, _ := range pages {
			pages[i] = start + i
		}
	case page >= 3 && totalpages > 5:
		start := page - 3 + 1
		pages = make([]int, 5)
		firstpage = page - 3
		for i, _ := range pages {
			pages[i] = start + i
		}
		firstpage = page - 1
		lastpage = page + 1
	default:
		pages = make([]int, int(math.Min(5, float64(totalpages))))
		for i, _ := range pages {
			pages[i] = i + 1
		}
		firstpage = int(math.Max(float64(1), float64(page-1)))
		lastpage = page + 1
	}
	paginator := &schema.Paginator{}
	paginator.Pages = pages
	paginator.Totalpages = totalpages
	paginator.Firstpage = firstpage
	paginator.Lastpage = lastpage
	paginator.Currpage = page
	return paginator
}
