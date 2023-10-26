-- +goose Up
-- +goose StatementBegin
INSERT INTO languages (name)
VALUES
    ('Java'),
    ('Python'),
    ('JavaScript'),
    ('C#'),
    ('C++'),
    ('Ruby'),
    ('Swift'),
    ('Kotlin'),
    ('HTML'),
    ('CSS'),
    ('Plain Text'),
    ('Go'),
    ('Dart'),
    ('Bash'),
    ('F#'),
    ('SQL'),
    ('Rust'),
    ('PHP'),
    ('TypeScript'),
    ('Objective-C'),
    ('R'),
    ('Scala'),
    ('Perl'),
    ('Haskell');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM languages WHERE name IN (
     'Java',
     'Python',
     'JavaScript',
     'C#',
     'C++',
     'Ruby',
     'Swift',
     'Kotlin',
     'HTML',
     'CSS',
     'Plain Text',
     'Go',
     'Dart',
     'Bash',
     'F#',
     'SQL',
     'Rust',
     'PHP',
     'TypeScript',
     'Objective-C',
     'R',
     'Scala',
     'Perl',
     'Haskell'
    );
-- +goose StatementEnd
