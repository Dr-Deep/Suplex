package database

func (d *Database) Initialize() {}

/*
_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id TEXT PRIMARY KEY,
            name TEXT NOT NULL
        );
    `)
    if err != nil {
        return nil, err
    }
*/
