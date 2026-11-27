# Records assets/demo.gif. Run from the repo root: vhs assets/demo.tape
# Uses a throwaway HOME, so your real ~/.claude, ~/.codex and ~/.cursor are untouched.

Output assets/demo.gif

Set Shell "bash"
Set FontSize 16
Set Width 1100
Set Height 720
Set Padding 20
Set Theme "Rose Pine"
Set TypingSpeed 60ms

Hide
Type `export GOCACHE=$(go env GOCACHE) GOMODCACHE=$(go env GOMODCACHE) GOPATH=$(go env GOPATH) && export HOME=$(mktemp -d) PS1='$ ' && mkdir -p $HOME/.claude $HOME/.codex $HOME/.cursor && go build -ldflags "-X main.version=$(git describe --tags --abbrev=0)" -o $HOME/bin/mallard . && export PATH=$HOME/bin:$PATH && clear`
Enter
Sleep 8s
Show

Type "mallard"
Sleep 400ms
Enter
Sleep 2s

# Install: pick agents
Enter
Sleep 2s

# Pick skills (all preselected)
Enter
Sleep 2s
Down 3
Sleep 1s

# Install
Enter
Sleep 4s
