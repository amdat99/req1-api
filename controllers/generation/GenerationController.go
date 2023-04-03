
package generationController

import (
	"req-api/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/PullRequestInc/go-gpt3"
	"context"
	"fmt"
	"log"
)


// var validate = validator.New()

//Get all requirements
func Gen(c *fiber.Ctx, session utils.Session) error {

	if(session.Org_id == ""){ return utils.ResError(c, "You must be logged in to use this feature", 401) }

	apiKey := ""

	ctx := context.Background()
	client := gpt3.NewClient(apiKey)

	resp, err := client.Completion(ctx, gpt3.CompletionRequest{
		Prompt:    []string{"Return fields for a health and safety form as a array  with maxlength,min, required,type,op"},
		MaxTokens: gpt3.IntPtr(600),
		Stop:      []string{"."},
		Echo:      true,
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp)
	fmt.Println(resp.Choices[0].Text)

	return utils.ResJSON(c, resp.Choices[0].Text)

}
