DROP TABLE IF EXISTS link; -- TEST
CREATE TABLE IF NOT EXISTS link (
  hash    VARCHAR(32) NOT NULL PRIMARY KEY,
  target  VARCHAR(255) NOT NULL
);

INSERT INTO link
  (hash, target)
VALUES
  ('google', 'https://google.com'),
  ('icanpe', 'https://icanpe.com'),
  ('reddit', 'https://reddit.com');
