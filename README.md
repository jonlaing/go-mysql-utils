# Go-Mysql-Utils

This project was born as a result of a Rails developer (me) being completely incapable of thinking about web apps in a non-Railsy way. I feared the old non-MVC style way I used to develop PHP apps, and I cringed. I really liked Go so far, and its syntax screamed to me that there had to be a better way. I looked around at ORM's and the likes, and I was pretty underwhelmed, so I decided to roll my own.

## Usage

### A Railsy Model

I'd like to demonstrate how one would go about making a railsy model using this technique. As you probably already know, Go isn't Object-Oriented in the way Ruby is. It's also a compiled language, not interpreted, like Ruby. This means the approach has to be a little different, but I think you'll find this method to be relatively painless.

```go
// I recommend putting your models in their own package, but that's just me
package post

import (
	"database/sql"
	"github.com/jonlaing/go-mysql-utils"
	"time"
	"strings"
)

// Lets pretend we're making a blog of sorts.
// The tags are essential as the utility uses reflection to turn them
// into MySQL statements. Primary Keys are denoted by writing `pk` 
// preceded by a comma. Otherwise, these statements must be valid MySQL.
type Post struct {
	Id 	int 		`mysql:"id INT NOT NULL AUTO_INCREMENT,pk"`
	Title 	string 		`mysql:"title VARCHAR(20) NOT NULL"`
	Body 	string 		`mysql:"body TEXT NOT NULL"`
	Created time.Time 	`mysql:"created DATETIME"`
	Updated time.Time 	`mysql:"updated DATETIME"`
}

var s MysqlUtil

func init() {
	// Get the configuration for the database
	s.Conn.GetConfigureation("config.gcfg", "development")	
	
	// Create the table based on the tags in Post
	err := s.CreateTable("posts", &Post{})
	if err != nil {
		// Deal with the error how you'd like
		panic(err)
	}
}

// Just a simple Find function, from this I assume You'd be a ble to surmise
// the rest. In your controller you'd call `post := post.Find(3)`. You'd also
// probably need some error handling, but this is just an example.
func Find(id int) (p Post) {
	fields := strings.Join(s.FieldList(&Post{}), ",")
	row := s.QueryRow("SELECT "+fields+" FROM posts WHERE id=?", id)
	row.Scan(&p.Id, &p.Title, &p.Body, &p.Created, &p.Updated)
}
```

### Using your Railsy model

Here is an example program that is next to useless. You'd probably want to use this in conjunction with a controller, but I'm too lazy to write one here. This should illustrate the point fairy well, though.

```go
import (
	"fmt"
	"post"
)

func main() {
	post := post.Find(3)

	fmt.Println("Title:", post.Title)
	fmt.Println("Body:", post.Body)
}
```

Pretty straight forward, right?

## TODO

I need to confirm that all of this works. This is really not stable, and if you're doing anything more than screwing around, I wouldn't rely on this just yet. I haven't even used it on a project yet. Do you see how young this repository is? There's likely to be a lot of revision. Seriously, don't use this for anything important yet.

I'm also on the fence about whether I want to have an automated model generator much like Rails' `rails generate model post`. I'm not sure if that's inline with the Go mentality or not. I guess time will tell.
