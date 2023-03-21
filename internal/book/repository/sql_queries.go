package repository

const (
	qCreateBook = `
	INSERT INTO book (title, synopsis, cover_url, content, author_id)
    VALUES (?,?,?,?,?)
	RETURNING id
	`
	qGetBookByID = `
	SELECT b.id, b.title, b.synopsis, b.cover_url, b.content, a.full_name as author, b.created_at, b.updated_at
	FROM book b
	INNER JOIN author a on a.id = author_id
	WHERE b.id = ?
	`
	qCountBooks = `
	SELECT COUNT(*) AS total
	FROM book
	`
	qGetAllBook = `
	SELECT b.id, b.title, b.synopsis, b.cover_url, b.content, a.full_name as author, b.created_at, b.updated_at
	FROM book b
	INNER JOIN author a on a.id = author_id
	`
	qUpdateBook = `
	UPDATE book
	SET title = ?,
		synopsis = ?,
		cover_url = ?,
		content = ?,
		author_id = ?,
		updated_at = timezone('utc', now())
	WHERE id = ?
	`
	qDeleteBook = `
	DELETE FROM book
	WHERE id = ?
	`
)
