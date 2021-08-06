# go365
learn a golang lesson a day from 001 to 365 

# how to use this course
Open terminal and clone this repo

```
git clone https://github.com/andrewarrow/go365.git
```

If it says git is not installed, [install it](https://www.atlassian.com/git/tutorials/install-git).

cd into the directory and then cd into the first 001 lesson

```
cd go365
cd 001
```

run the command `go build` to compile the `main.go` file in that directory.

If it says go is not installed, [install it](https://golang.org/doc/install).

Run the 001 program

```
./001
```

And then open the file `main.go` in a [text editor](https://www.techradar.com/best/best-text-editors).

Read all the words in the file even if they don't make sense to you.

Change some text that is in between two quotes.

Save the file.

Go back to terminal and run `go build` again.

Run the 001 program again

```
go build
./001
```

Notice the output of the program has your changes. 

Do this again.

Back to the text editor, make a change, save the file, run `go build` again. Run the program again.
Notice the change.

Do this again. Over and over, and get more and more bold with your changes.

When you type something that won't compile `go build` will tell you why with a specific line number where you have a mistake.

Keep going with the english words in 001/main.go and the few non-english words (the real golang programming stuff) over
and over until you see the cause and effect relationship between what you change in your text editor and what you see when you run the program.

When you are ready, move on to 002/main.go and keep going, one a day for 365 days.

# teaching style

Each day builds on the next. The first few days have a lot of `fmt.Println("text")` statements helping you
get used to a very common way to write messages to the user. And you yourself are the very first user
of your programs.

The first few days also have just one file `main.go`. After you get the hang of some core concepts
there will be days with more than one file, each one named something more descriptive that just the word main.

Each day has a few word title. Make sure when you finish a day and move on you can honestly read the title
of that day and nod your head affirmitively agreeing you `grok` it.

# table of contents

| 001 | a byte is a number from 0 to 255      |
| --- | ------------------------------------- |
| 002 | an uint16 is a number from 0 to 65535 | 
| 003 | know your ints and appends            |
| 004 | maps and Goroutines                   |
