# byteknot-api
Backend of byteknot.com that exposes API to create/update/serve articles and also comment, like others articles.

# This project follows Go Project Layout standards:
https://github.com/golang-standards/project-layout

# Prerequisite 
brew install mysql@8.4

by default brew installs MySQL database without a root password. To secure it run:
    `mysql_secure_installation`

```cmd
echo 'export PATH="/opt/homebrew/opt/mysql@8.4/bin:$PATH"' >> ~/.zshrc
export LDFLAGS="-L/opt/homebrew/opt/mysql@8.4/lib"
export CPPFLAGS="-I/opt/homebrew/opt/mysql@8.4/include"
```

# Start server
```cmd
brew services start mysql@8.4
```

