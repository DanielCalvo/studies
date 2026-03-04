# Getting this going
- Get the CLI: `go install github.com/spf13/cobra-cli@latest`
```shell
go mod init example
cobra-cli init
```

This is a dummy cobra project that will add, remove and search for people's names on a text file.

## review from February 2026
Let me review the docs and add some information here

## what the functions do/are for
- init(): define flags and handle configuration in your init() function (this is described in the context of root.go)
    - looks like on subcommads this just calls rootCmd.AddCommand by default
- RunE(): what you use to run a command and return an error

What you do in your simple mdscanner project is that you just call RunE everywhere and thats it!

## PreRun and PostRun Hooks
https://github.com/spf13/cobra/blob/main/site/content/user_guide.md#prerun-and-postrun-hooks
Kinda nitty gritty, but:
Short notes with codex, pending validation: Whats the difference between a persistent pre run and a regular pre run hook in cobra?
- PreRun and PreRunE: Run only for that exact command, and is not inherited by children
- PersistentPreRun and PersistentPreRunE: Intended for a shared set up across a command tree, defined on parent, available to children

However, by default cobra.EnableTraverseRunHooks is false. This means only one persistent pre-run is executed. This is the closest in the command chain, the child if it has on, otherwise the parent's.

As a general rule, put global set up in PersistentPreRun, and command specific on PreRun. Neat!

## other
- dang there is an automated help generator a and shell autocompletion generator feature and you can have plugins!
- Note that the cobra cli will put everything in cmd/ but you could have everything on your main too if you wanted
