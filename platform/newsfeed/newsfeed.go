package newsfeed

import "database/sql"

type Feed struct {
	DB *sql.DB
}

func (feed *Feed) Get() []Item {
	items := []Item{}
	rows, _ := feed.DB.Query(`
	SELECT * FROM newsfeed
	`)
	var id int
	var content string

	for rows.Next() {
		rows.Scan(&id, &content)
		item := Item{
			ID:      id,
			Content: content,
		}
		items = append(items, item)
	}
	return items
}

func (feed *Feed) Add(item Item) {
	stmt, _ := feed.DB.Prepare(`
		INSERT INTO newsfeed (content) values (?)
	`)

	stmt.Exec(item.Content)
}

func NewFeed(db *sql.DB) *Feed {
	stmt, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS "newsfeed" (
			"ID"	INTEGER NOT NULL UNIQUE,
			"content"	TEXT,
			PRIMARY KEY("ID" AUTOINCREMENT)
	);
	`)
	stmt.Exec()

	return &Feed{
		DB: db,
	}

}
