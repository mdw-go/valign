# valign

Vertical alignment tool.

Given a file called query.sql with the following contents:

```sql
SELECT 'a' FROM table;
SELECT 'hello' FROM table;
SELECT 'goodbye' FROM table;
SELECT 'really-super-long-name' FROM table;
```

When processed with `valign`:

```
$ cat query.sql | valign -match "FROM"
SELECT 'a'                      FROM table;
SELECT 'hello'                  FROM table;
SELECT 'goodbye'                FROM table;
SELECT 'really-super-long-name' FROM table;
```
