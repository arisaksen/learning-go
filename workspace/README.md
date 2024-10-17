# Run apps

## Command line should work fine

cd workspace
go run ./cmd/module1
go run ./cmd/module2

## Intellij

Step 1: Use "Edit Run Configurations"
Open IntelliJ IDEA.
Go to Run > Edit Configurations in the top menu.
Find or create your Go Run or Go Build configuration (depending on what you're doing).
Click the + button to add a new Go Build or Go Run configuration if it doesn't exist.
Step 2: Add the -buildvcs=false flag

# Docker build

Dockerfile must copy the code to build AND the library. When copy the library the build context must be root.

COPY cmd/module2 .
COPY common .

cd workspace
docker build -f cmd/module2/Dockerfile -t module2-img .