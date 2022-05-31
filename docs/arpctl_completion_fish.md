## arpctl completion fish

Generate the autocompletion script for fish

### Synopsis

Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	arpctl completion fish | source

To load completions for every new session, execute once:

	arpctl completion fish > ~/.config/fish/completions/arpctl.fish

You will need to start a new shell for this setup to take effect.


```
arpctl completion fish [flags]
```

### Options

```
  -h, --help              help for fish
      --no-descriptions   disable completion descriptions
```

### SEE ALSO

* [arpctl completion](arpctl_completion.md)	 - Generate the autocompletion script for the specified shell

###### Auto generated by spf13/cobra on 31-May-2022