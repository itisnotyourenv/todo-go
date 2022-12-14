CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  email VARCHAR(255) NOT NULL UNIQUE,
  username VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);


CREATE TABLE todo_lists (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  description TEXT,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE users_todo_lists (
  id SERIAL PRIMARY KEY,
--   FOREIGN KEY (user_id) REFERENCES users ON DELETE CASCADE
  user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
--   PRIMARY KEY (user_id, todo_list_id)
  list_id INTEGER NOT NULL REFERENCES todo_lists(id) ON DELETE CASCADE
);

CREATE TABLE todo_items (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  description TEXT,
  done BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE lists_items (
  id SERIAL PRIMARY KEY,
  list_id INTEGER NOT NULL REFERENCES todo_lists(id) ON DELETE CASCADE,
  item_id INTEGER NOT NULL REFERENCES todo_items(id) ON DELETE CASCADE
);