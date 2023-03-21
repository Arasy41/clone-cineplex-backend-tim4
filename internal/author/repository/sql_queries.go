package repository

const (
	qCreateAuthor = `
	INSERT INTO author (full_name)
    VALUES (?)
	RETURNING id
	`
	qGetAuthorByID = `
	SELECT id, full_name, created_at, updated_at
	FROM author
	WHERE id = ?
	`
	qGetAuthorByName = `
	SELECT id, full_name, created_at, updated_at
	FROM author
	WHERE full_name = ?
	`
	qGetAllAuthor = `
	SELECT id, full_name, created_at, updated_at
	FROM author
	`
	qUpdateAuthor = `
	UPDATE author
	SET full_name = ?,
		updated_at = timezone('utc', now())
	WHERE id = ?
	`
	qDeleteAuthor = `
	DELETE FROM author
	WHERE id = ?
	`
)
