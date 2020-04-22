package plugin

import (
	"testing"

	md "github.com/arangodb-managed/html-to-markdown"
)

func TestStrikethroughDefault(t *testing.T) {
	conv := md.NewConverter("", true, nil)
	conv.Use(Strikethrough(""))

	input := `<p>Strikethrough uses two tildes. <del>Scratch this.</del></p>`
	expected := `Strikethrough uses two tildes. ~Scratch this.~`
	markdown, err := conv.ConvertString(input)
	if err != nil {
		t.Error(err)
	}

	if markdown != expected {
		t.Errorf("got '%s' but wanted '%s'", markdown, expected)
	}
}
func TestStrikethrough(t *testing.T) {
	conv := md.NewConverter("", true, nil)
	conv.Use(Strikethrough("~~"))

	input := `<p>Strikethrough uses two tildes. <del>Scratch this.</del></p>`
	expected := `Strikethrough uses two tildes. ~~Scratch this.~~`
	markdown, err := conv.ConvertString(input)
	if err != nil {
		t.Error(err)
	}

	if markdown != expected {
		t.Errorf("got '%s' but wanted '%s'", markdown, expected)
	}
}

func TestTaskListItems(t *testing.T) {
	conv := md.NewConverter("", true, nil)
	conv.Use(TaskListItems())

	input := `
	<ul>
		<li><input type=checkbox checked>Checked!</li>
		<li><input type=checkbox>Check Me!</li>
	</ul>
	`
	expected := `- [x] Checked!
- [ ] Check Me!`
	markdown, err := conv.ConvertString(input)
	if err != nil {
		t.Error(err)
	}

	if markdown != expected {
		t.Errorf("got '%s' but wanted '%s'", markdown, expected)
	}
}

func TestConfluenceCodeBlock(t *testing.T) {
	conv := md.NewConverter("", true, nil)
	conv.Use(ConfluenceCodeBlock())

	input := `<ac:structured-macro ac:name="code" ac:schema-version="1" ac:macro-id="150db472-e155-47c7-a195-c581bf891af5"><ac:plain-text-body><![CDATA[FOR stuff IN imdb_vertices
	FILTER LIKE(stuff.description, "%good%vs%evil%", true)
  RETURN stuff.description]]></ac:plain-text-body></ac:structured-macro>
some other stuff
<ac:structured-macro ac:name="code" ac:schema-version="1" ac:macro-id="150db472-e155-47c7-a195-c581bf891af5"><ac:parameter ac:name="language">sql</ac:parameter><ac:plain-text-body><![CDATA[FOR stuff IN imdb_vertices
	FILTER LIKE(stuff.description, "%good%vs%evil%", true)
  RETURN stuff.description]]></ac:plain-text-body></ac:structured-macro>`
	expected := "```" + `
FOR stuff IN imdb_vertices
	FILTER LIKE(stuff.description, "%good%vs%evil%", true)
  RETURN stuff.description
` + "```" + `
some other stuff
` + "```sql" + `
FOR stuff IN imdb_vertices
	FILTER LIKE(stuff.description, "%good%vs%evil%", true)
  RETURN stuff.description
` + "```"
	markdown, err := conv.ConvertString(input)
	if err != nil {
		t.Error(err)
	}

	if markdown != expected {
		t.Errorf("got '%s' but wanted '%s'", markdown, expected)
	}
}

func TestConfluenceAttachments(t *testing.T) {
	conv := md.NewConverter("", true, nil)
	conv.Use(ConfluenceAttachments())

	input := `<ri:attachment ri:filename="oasis.png" ri:version-at-save="1" />`
	expected := "![][oasis.png]"
	markdown, err := conv.ConvertString(input)
	if err != nil {
		t.Error(err)
	}

	if markdown != expected {
		t.Errorf("got '%s' but wanted '%s'", markdown, expected)
	}
}
