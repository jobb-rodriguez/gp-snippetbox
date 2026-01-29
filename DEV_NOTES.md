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

# MySQL
```bash
# installation
brew install mysql

# scaffold
sudo mysql

# alternate + enter password during installation
mysql -u root -p
```

> [!NOTE]
> See ```internal/sql``` for SQL commands.
