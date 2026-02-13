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
    - Query
      - ```DB.Query()``` is used for SELECT queries which return multiple rows.
      - ```DB.QueryRow()``` is used for SELECT queries which return a single row.
      - ```DB.Exec()``` is used for statements which don’t return rows (like INSERT and DELETE).
      - For specific commands, check the documentation of the driver.
      - Placeholder help avoid SQL injection attacks
      - PostgreSQL uses the ```$N``` notation instead of the ```?``` notation.
  - Executing SQL statements
  - Single-record
  - Multiple-record
  - Transactions

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
