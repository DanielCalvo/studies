# 42. Understanding NPM scripts
- Uh-oh, node package manager!
```yaml
npm init myproject
```
- `package.json` contains NPM scripts, neat
```shell
npm start #Start is a reserved name though!
npm run start-server #You need to use `run` first -- start above is a special case
```

```shell
npm install nodemon --save-dev
# -g installs it on your entire machine
```

- `node_modules` will contain the source code of nodemon in this case

# 47. Understanding different Error Types
- Syntax error: Ya coded wrong son
- Runtime errors: You tried to execute some code that just breaks
- Logical errors: aka bugs