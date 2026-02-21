# Development review notes
1. Struct methods allow to define the application's handlers, routes, and helpers.
2. (1) allows to keep ```main.go``` lean.
3. ```tmpl``` uses base tempaltes and calls other templates as well.
4. Project structure
 - ```cmd``` - application-specific code, e.g., main, helpers, routes
 - ```internal``` - non-application specific code, e.g., validation helpers and db models
 - ```ui``` - user-interface assets, e.g., html, static
5. ```flag``` allows to set flag inputs when running the application.
6. Centralized Errors can also be a model to other solutions.
7. Dependency injection makes your code more explicit, less error-prone, and easier to unit test than if you use global variables.
  - Do not use DI for (a) simple applications, (b) high-performance applications, (c) strict applications.
8. Database specific notes
  - Install drivers: [Go Drivers](https://go.dev/wiki/SQLDrivers)
    - Use ```@v1``` to refer to major version 1 and use the latest minor version.
  - Modules
    - ```go mod verify``` verifies the checksums of ```go.mod``` referencing ```go.sum```.
    - ```go.sum``` should not be edited.
    - ```go mod download``` allows uesrs to download the needed packages.
    - ```go get -u <link>```: ```-u``` allows to get the minor or patch release.
    - ```go mod tidy``` automatically removes any unused packages.
  - Creating a database connection pool: ```sql.Open()```
    - Read [```sql.DB```](https://pkg.go.dev/database/sql#DB)
  - Designing database models
    - A database model can be used by other applications in the future like a CLI. Hence, it belongs to ```internal```.
  - Executing SQL statements
    - Query
      - ```DB.Query()``` is used for SELECT queries which return multiple rows.
      - ```DB.QueryRow()``` is used for SELECT queries which return a single row.
      - ```DB.Exec()``` is used for statements which don’t return rows (like INSERT and DELETE).
      - For specific commands, check the documentation of the driver.
      - Placeholder help avoid SQL injection attacks
      - PostgreSQL uses the ```$N``` notation instead of the ```?``` notation.
    - Update model -> Update handler
  - Single-record
    - Mapping Values
      - CHAR, VARCHAR and TEXT map to string.
      - BOOLEAN maps to bool.
      - INT maps to int; BIGINT maps to int64.
      - DECIMAL and NUMERIC map to float.
      - TIME, DATE and TIMESTAMP map to time.Time.
      - For Time and Date, it's driver specific.
    - Use ```errors.Is()``` to check whether error matches.
  - Multiple-record
    - ```defer``` defers the execution of a function until the surrounding function returns.
  - Transactions
    - There is no guarantee that the db will use the same db connection. The solution is to use transactions.
    - ```Begin()``` creates a new ```sql.Tx``` object.
    - ```defer tx.Rollback()``` ensures it at the latest commit or most recent version if there is an error.
    - If there are no errors, call ```tx.Commit()```
    - Use transactions for commands with multiple statements.

> [!IMPORTANT]
> Closing a resultset with defer rows.Close() is critical in the code above. As long as a resultset is open it will keep the underlying database connection open

> [!NOTE]
> Go allows you to write portable code with ```database/sql``` package. However, read the driver documentation tot understand quirks and edge cases.

> [!DANGER]
> Go doesn't manage ```NULL``` values well. Avoid ```NULL``` values. SET ```NOT NULL``` cosntraints on all database columngs, along with sensible ```DEFAULT``` values as necessary.

> [!IMPORTANT]
> You must always call either Rollback() or Commit() before your function returns. If you don’t the connection will stay open and not be returned to the connection pool. This can lead to hitting your maximum connection limit/running out of resources. The simplest way to avoid this is to use defer tx.Rollback() like we are in the example above.

# MySQL
```bash
# installation
brew install mysql

# scaffold
sudo mysql

# alternate + enter password during installation
mysql -u root -p

# Change password of dolphin user
ALTER USER dolphin@localhost
IDENTIFIED BY 'qDvOD3@L10'; -- new password
```

> [!NOTE]
> See ```internal/sql``` for SQL commands.

9. Dynamic HTML templates
  - As a general rule, my advice is to get into the habit of always pipelining dot whenever you invoke a template with the {{template}} or {{block}} actions, unless you have a good reason not to.
  - ```html/template``` helps in avoiding XSS attacks.
  - Template actions aand functions
    - ```{{<keyword> .Foo}} C1 {{else}} C2 {{end}}```
    - ```<keyword>```: {if, with, range} (with and range change the value of dot)
    - [Read functions here](https://pkg.go.dev/text/template#hdr-Functions)
  - Caching
    - Why? Reduce code duplication
    - How?
      - Create an in-memory map in templates
      - Initialize the cache in main() and set it as a dependency via the application struct
      - Create a render helper method, which is called in handlers

# Git
```
# Deleting in remote branch (remote_name is usually main)
git push <remote_name> --delete <branch_name>
```

## Conventional Commits
[conventionalcommits.org](https://www.conventionalcommits.org/)

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]

e.g., Breaking Change example
feat(api)!: send an email to the customer when a product is shipped

e.g., Commit message with scope
feat(lang): add Polish language
```

1. ```fix```: a commit of the type fix patches a bug in your codebase (this correlates with PATCH in Semantic Versioning).
2. ```feat```: a commit of the type feat introduces a new feature to the codebase (this correlates with MINOR in Semantic Versioning).
3. BREAKING CHANGE: a commit that has a footer BREAKING CHANGE:, or appends a ! after the type/scope, introduces a breaking API change (correlating with MAJOR in Semantic Versioning). A BREAKING CHANGE can be part of commits of any type.
4. types other than fix: and feat: are allowed, for example @commitlint/config-conventional (based on the Angular convention) recommends build:, chore:, ci:, docs:, style:, refactor:, perf:, test:, and others.
5. footers other than BREAKING CHANGE: <description> may be provided and follow a convention similar to git trailer format.

**Types based on Angular Convention**
1. build: Changes that affect the build system or external dependencies (example scopes: gulp, broccoli, npm)
2. ci: Changes to our CI configuration files and scripts (example scopes: Travis, Circle, BrowserStack, SauceLabs)
3. docs: Documentation only changes
4. feat: A new feature
5. fix: A bug fix
6. perf: A code change that improves performance
7. refactor: A code change that neither fixes a bug nor adds a feature
8. style: Changes that do not affect the meaning of the code (white-space, formatting, missing semi-colons, etc)
9. test: Adding missing tests or correcting existing tests
