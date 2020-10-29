### Simple tool for selectively copying files across volumes

Recursively copy matching files from one location to another given a set of file extensions to 
match against and paths that should be ignored. 
(todo: pass files that are corrupted and block read operation)

1. Adjust `config.json` file to your usecase/requirements:
   1. In the `lookFor` list specify file extensions you're interested in (`.jpg`? `.pdf`? `.docx`?)
   2. In the `filterOut` list specify files and directories you're not interested in (`\Windows`? `\Program Files`?)
2. Create a `clone` directory for the files to be copied over to, you can simply execute `mkdir clone` command
3. Execute the program `./hdd-files-recovery run /Path/To/Volume/Directory`

sample usage:
```bash
mkdir clone
./hdd-files-recovery /Volumes/Acer/Users/jack/Pictures
```

Feedback and improvements welcome, it's my first Go project so don't be too harsh ;)

#### A few words of the background

One of my family members recently asked me if I could help recover photos from a laptop that fell on the ground and is not booting anymore. 

They didn't want to give the HDD to a specialised lab (just yet, I guess?), but to just check if anything could be recovered with some low effort.

I plugged it to my laptop via an interface and saw the HDD come up - some structure was still discoverable, not all but some files browsable. That was a good start.

I decided it's a good enough reason to spend one evening writing a simple tool to scavenge what's still undamaged and only target specific files (photos in this case). 

#### usage tips and general information

You may need to run this tool a few times in case one or more files on the way are damaged and the tool freezes. No worries, it wont overwrite or double-copy the files, it will know which files have already been copied - **the tool is idempotent**. 
The way around this issue is to skip the problematic file by adding it to the `filterOut` config section and running the tool again. This only happens on damaged disks.

For now "clone" is the hardcoded directory name to where all the files will be copied to - on TODO to improve. 

You can run the tool against the whole drive, but if it's damaged, there's a chance the drive will reboot frequently not giving the tool enough time to go through all the files and copy everything. Only advised if necessary.

The tool **converts all extensions to lowercase automatically**, so a file named `photo.JPG` will be visible to the tool as `photo.jpg`

#### TODO:

- create the "clone" directory automatically or let the user specify a destination path
- implement timeouts for fetching files, atm the tool freezes when it tries to fetch a corrupted file
- implement more tests

#### Tests

You can run tests with `go test ./...`, for now only the file filtering functions are tested

#### disclaimer

There are plenty of other tools like this one, probably all of them better - I advise against using this tool on damaged disks with data that you **need** to recover. I don't take responsibility for loss of data and/or damage to the property.