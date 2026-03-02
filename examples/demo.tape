Output examples/demo.gif

Require echo

Set Shell "bash"
Set FontSize 14
Set Width 1200
Set Height 600
Set Theme "Dracula"
Set Padding 20
Set PlaybackSpeed 1

# Setup shell prompt and alias (hidden from recording)
Hide
Type `export PS1='$ ' && alias isctl=./build/isctl`
Enter
Sleep 500ms
Type "clear"
Enter
Sleep 500ms
Show

# --- Basic query: get all NTP policies ---
Type `# Get all NTP policies`
Enter
Sleep 500ms
Type `isctl get ntp policy`
Sleep 500ms
Enter
Sleep 3s

# --- Query by name ---
Type `# Get a specific policy by name`
Enter
Sleep 500ms
Type `isctl get ntp policy --name isctl-bats-test-ntp-policy-2`
Sleep 500ms
Enter
Sleep 3s

# --- Filter results ---
Type `# Filter using the Intersight API query language`
Enter
Sleep 500ms
Type `isctl get ntp policy --filter "Name eq 'isctl-bats-test-ntp-policy-2'"`
Sleep 500ms
Enter
Sleep 3s

# --- JSON output piped to jq ---
Type `# Output as JSON and pipe to jq`
Enter
Sleep 500ms
Type `isctl get ntp policy -o json | jq -r '.[].Name'`
Sleep 500ms
Enter
Sleep 3s

# --- JSONPath output ---
Type `# JSONPath output - extract specific fields`
Enter
Sleep 500ms
Type `isctl get ntp policy -o jsonpath='[*].Name'`
Sleep 500ms
Enter
Sleep 3s

# --- Custom columns ---
Type `# Custom columns for tailored table output`
Enter
Sleep 500ms
Type `isctl get ntp policy --output custom-columns=NAME:.Name,ENABLED:.Enabled`
Sleep 500ms
Enter
Sleep 3s

Sleep 1s
