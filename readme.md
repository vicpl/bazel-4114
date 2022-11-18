Sample repo to reproduce Intellij Bazel plugin [issue with Run configurations](https://github.com/bazelbuild/intellij/issues/4112). 

It has one service which prints all the command line arguments.

It also has one bazel Run configuration  (saved in .ijwb/.run) that launches `//service:service target` and has "External flags" set to `--test_path=$WorkspaceRoot$/testfile`

When run in debug mode, `$WorkspaceRoot$` is resolved and application prints the actual path 
```
2022/11/18 09:58:54 Received args:
2022/11/18 09:58:54 [--test_path=/repositories/bazel-4112/testfile]
```

When run whithout debugger, application print the literal string that was passed without resolving variables:
```
2022/11/18 09:59:50 Received args:
2022/11/18 09:59:50 [--test_path=$WorkspaceRoot$/testfile]
```

