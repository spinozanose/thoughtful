# thoughtful.ai
https://thoughtfulautomation.notion.site/Core-Engineering-Technical-Screen-b61b6f6980714c198dc49b91dd23d695

This is currently set up for a Macintosh, running on local. It could easily be adjusted for other environments -- look at the `scripts/build.sh` to see how the OS and architecture is specified. If this is to be run on Windows, shell or powershell scripts could be created, though it may be easier to just run the go build and go test commands manually, as done in the bash scripts.

After cloning to local, your best friends are the scripts in the `scripts` folder. If you want to verify the code, just run `scripts/test.sh` from the root of the repo. You can run the executable by running `scripts/run.sh` with some arguments from the command line.

Enjoy!