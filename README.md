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

The first few days have just one file `main.go`. After you get the hang of some core concepts
there will be days with more than one file, each one named something more descriptive that just the word main.

Each day has a title. Make sure when you finish a day and move on you can honestly read the title
of that day and nod your head affirmitively agreeing you `grok` it.

# table of contents

| day | title                                 |
| --- | ------------------------------------- |
| [001](https://github.com/andrewarrow/go365/blob/main/001/main.go) | a byte is a number from 0 to 255      |
| [002](https://github.com/andrewarrow/go365/blob/main/002/main.go) | an uint16 is a number from 0 to 65535 | 
| [003](https://github.com/andrewarrow/go365/blob/main/003/main.go) | know your ints and appends            |
| [004](https://github.com/andrewarrow/go365/blob/main/004/main.go) | maps and goroutines                   |
| [005](https://github.com/andrewarrow/go365/blob/main/005/main.go) | building something real               |
| [006](https://github.com/andrewarrow/go365/blob/main/006/main.go) | record in a map the stats             |
| [007](https://github.com/andrewarrow/go365/blob/main/007/main.go) |                |
| [008](https://github.com/andrewarrow/go365/blob/main/008/main.go) |                |
| [009](https://github.com/andrewarrow/go365/blob/main/009/main.go) |                |
| [010](https://github.com/andrewarrow/go365/blob/main/010/main.go) |                |
| [011](https://github.com/andrewarrow/go365/blob/main/011/main.go) |                |
| [012](https://github.com/andrewarrow/go365/blob/main/012/main.go) |                |
| [013](https://github.com/andrewarrow/go365/blob/main/013/main.go) |                |
| [014](https://github.com/andrewarrow/go365/blob/main/014/main.go) |                |
| [015](https://github.com/andrewarrow/go365/blob/main/015/main.go) |                |
| [016](https://github.com/andrewarrow/go365/blob/main/016/main.go) |                |
| [017](https://github.com/andrewarrow/go365/blob/main/017/main.go) |                |
| [018](https://github.com/andrewarrow/go365/blob/main/018/main.go) |                |
| [019](https://github.com/andrewarrow/go365/blob/main/019/main.go) |                |
| [020](https://github.com/andrewarrow/go365/blob/main/020/main.go) |                |
| [021](https://github.com/andrewarrow/go365/blob/main/021/main.go) |                |
| [022](https://github.com/andrewarrow/go365/blob/main/022/main.go) |                |
| [023](https://github.com/andrewarrow/go365/blob/main/023/main.go) |                |
| [024](https://github.com/andrewarrow/go365/blob/main/024/main.go) |                |
| [025](https://github.com/andrewarrow/go365/blob/main/025/main.go) |                |
| [026](https://github.com/andrewarrow/go365/blob/main/026/main.go) |                |
| [027](https://github.com/andrewarrow/go365/blob/main/027/main.go) |                |
| [028](https://github.com/andrewarrow/go365/blob/main/028/main.go) |                |
| [029](https://github.com/andrewarrow/go365/blob/main/029/main.go) |                |
| [030](https://github.com/andrewarrow/go365/blob/main/030/main.go) |                |
| [031](https://github.com/andrewarrow/go365/blob/main/031/main.go) |                |
| [032](https://github.com/andrewarrow/go365/blob/main/032/main.go) |                |
| [033](https://github.com/andrewarrow/go365/blob/main/033/main.go) |                |
| [034](https://github.com/andrewarrow/go365/blob/main/034/main.go) |                |
| [035](https://github.com/andrewarrow/go365/blob/main/035/main.go) |                |
| [036](https://github.com/andrewarrow/go365/blob/main/036/main.go) |                |
| [037](https://github.com/andrewarrow/go365/blob/main/037/main.go) |                |
| [038](https://github.com/andrewarrow/go365/blob/main/038/main.go) |                |
| [039](https://github.com/andrewarrow/go365/blob/main/039/main.go) |                |
| [040](https://github.com/andrewarrow/go365/blob/main/040/main.go) |                |
| [041](https://github.com/andrewarrow/go365/blob/main/041/main.go) |                |
| [042](https://github.com/andrewarrow/go365/blob/main/042/main.go) |                |
| [043](https://github.com/andrewarrow/go365/blob/main/043/main.go) |                |
| [044](https://github.com/andrewarrow/go365/blob/main/044/main.go) |                |
| [045](https://github.com/andrewarrow/go365/blob/main/045/main.go) |                |
| [046](https://github.com/andrewarrow/go365/blob/main/046/main.go) |                |
| [047](https://github.com/andrewarrow/go365/blob/main/047/main.go) |                |
| [048](https://github.com/andrewarrow/go365/blob/main/048/main.go) |                |
| [049](https://github.com/andrewarrow/go365/blob/main/049/main.go) |                |
| [050](https://github.com/andrewarrow/go365/blob/main/050/main.go) |                |
| [051](https://github.com/andrewarrow/go365/blob/main/051/main.go) |                |
| [052](https://github.com/andrewarrow/go365/blob/main/052/main.go) |                |
| [053](https://github.com/andrewarrow/go365/blob/main/053/main.go) |                |
| [054](https://github.com/andrewarrow/go365/blob/main/054/main.go) |                |
| [055](https://github.com/andrewarrow/go365/blob/main/055/main.go) |                |
| [056](https://github.com/andrewarrow/go365/blob/main/056/main.go) |                |
| [057](https://github.com/andrewarrow/go365/blob/main/057/main.go) |                |
| [058](https://github.com/andrewarrow/go365/blob/main/058/main.go) |                |
| [059](https://github.com/andrewarrow/go365/blob/main/059/main.go) |                |
| [060](https://github.com/andrewarrow/go365/blob/main/060/main.go) |                |
| [061](https://github.com/andrewarrow/go365/blob/main/061/main.go) |                |
| [062](https://github.com/andrewarrow/go365/blob/main/062/main.go) |                |
| [063](https://github.com/andrewarrow/go365/blob/main/063/main.go) |                |
| [064](https://github.com/andrewarrow/go365/blob/main/064/main.go) |                |
| [065](https://github.com/andrewarrow/go365/blob/main/065/main.go) |                |
| [066](https://github.com/andrewarrow/go365/blob/main/066/main.go) |                |
| [067](https://github.com/andrewarrow/go365/blob/main/067/main.go) |                |
| [068](https://github.com/andrewarrow/go365/blob/main/068/main.go) |                |
| [069](https://github.com/andrewarrow/go365/blob/main/069/main.go) |                |
| [070](https://github.com/andrewarrow/go365/blob/main/070/main.go) |                |
| [071](https://github.com/andrewarrow/go365/blob/main/071/main.go) |                |
| [072](https://github.com/andrewarrow/go365/blob/main/072/main.go) |                |
| [073](https://github.com/andrewarrow/go365/blob/main/073/main.go) |                |
| [074](https://github.com/andrewarrow/go365/blob/main/074/main.go) |                |
| [075](https://github.com/andrewarrow/go365/blob/main/075/main.go) |                |
| [076](https://github.com/andrewarrow/go365/blob/main/076/main.go) |                |
| [077](https://github.com/andrewarrow/go365/blob/main/077/main.go) |                |
| [078](https://github.com/andrewarrow/go365/blob/main/078/main.go) |                |
| [079](https://github.com/andrewarrow/go365/blob/main/079/main.go) |                |
| [080](https://github.com/andrewarrow/go365/blob/main/080/main.go) |                |
| [081](https://github.com/andrewarrow/go365/blob/main/081/main.go) |                |
| [082](https://github.com/andrewarrow/go365/blob/main/082/main.go) |                |
| [083](https://github.com/andrewarrow/go365/blob/main/083/main.go) |                |
| [084](https://github.com/andrewarrow/go365/blob/main/084/main.go) |                |
| [085](https://github.com/andrewarrow/go365/blob/main/085/main.go) |                |
| [086](https://github.com/andrewarrow/go365/blob/main/086/main.go) |                |
| [087](https://github.com/andrewarrow/go365/blob/main/087/main.go) |                |
| [088](https://github.com/andrewarrow/go365/blob/main/088/main.go) |                |
| [089](https://github.com/andrewarrow/go365/blob/main/089/main.go) |                |
| [090](https://github.com/andrewarrow/go365/blob/main/090/main.go) |                |
| [091](https://github.com/andrewarrow/go365/blob/main/091/main.go) |                |
| [092](https://github.com/andrewarrow/go365/blob/main/092/main.go) |                |
| [093](https://github.com/andrewarrow/go365/blob/main/093/main.go) |                |
| [094](https://github.com/andrewarrow/go365/blob/main/094/main.go) |                |
| [095](https://github.com/andrewarrow/go365/blob/main/095/main.go) |                |
| [096](https://github.com/andrewarrow/go365/blob/main/096/main.go) |                |
| [097](https://github.com/andrewarrow/go365/blob/main/097/main.go) |                |
| [098](https://github.com/andrewarrow/go365/blob/main/098/main.go) |                |
| [099](https://github.com/andrewarrow/go365/blob/main/099/main.go) |                |
| [100](https://github.com/andrewarrow/go365/blob/main/100/main.go) |                |
| [101](https://github.com/andrewarrow/go365/blob/main/101/main.go) |                |
| [102](https://github.com/andrewarrow/go365/blob/main/102/main.go) |                |
| [103](https://github.com/andrewarrow/go365/blob/main/103/main.go) |                |
| [104](https://github.com/andrewarrow/go365/blob/main/104/main.go) |                |
| [105](https://github.com/andrewarrow/go365/blob/main/105/main.go) |                |
| [106](https://github.com/andrewarrow/go365/blob/main/106/main.go) |                |
| [107](https://github.com/andrewarrow/go365/blob/main/107/main.go) |                |
| [108](https://github.com/andrewarrow/go365/blob/main/108/main.go) |                |
| [109](https://github.com/andrewarrow/go365/blob/main/109/main.go) |                |
| [110](https://github.com/andrewarrow/go365/blob/main/110/main.go) |                |
| [111](https://github.com/andrewarrow/go365/blob/main/111/main.go) |                |
| [112](https://github.com/andrewarrow/go365/blob/main/112/main.go) |                |
| [113](https://github.com/andrewarrow/go365/blob/main/113/main.go) |                |
| [114](https://github.com/andrewarrow/go365/blob/main/114/main.go) |                |
| [115](https://github.com/andrewarrow/go365/blob/main/115/main.go) |                |
| [116](https://github.com/andrewarrow/go365/blob/main/116/main.go) |                |
| [117](https://github.com/andrewarrow/go365/blob/main/117/main.go) |                |
| [118](https://github.com/andrewarrow/go365/blob/main/118/main.go) |                |
| [119](https://github.com/andrewarrow/go365/blob/main/119/main.go) |                |
| [120](https://github.com/andrewarrow/go365/blob/main/120/main.go) |                |
| [121](https://github.com/andrewarrow/go365/blob/main/121/main.go) |                |
| [122](https://github.com/andrewarrow/go365/blob/main/122/main.go) |                |
| [123](https://github.com/andrewarrow/go365/blob/main/123/main.go) |                |
| [124](https://github.com/andrewarrow/go365/blob/main/124/main.go) |                |
| [125](https://github.com/andrewarrow/go365/blob/main/125/main.go) |                |
| [126](https://github.com/andrewarrow/go365/blob/main/126/main.go) |                |
| [127](https://github.com/andrewarrow/go365/blob/main/127/main.go) |                |
| [128](https://github.com/andrewarrow/go365/blob/main/128/main.go) |                |
| [129](https://github.com/andrewarrow/go365/blob/main/129/main.go) |                |
| [130](https://github.com/andrewarrow/go365/blob/main/130/main.go) |                |
| [131](https://github.com/andrewarrow/go365/blob/main/131/main.go) |                |
| [132](https://github.com/andrewarrow/go365/blob/main/132/main.go) |                |
| [133](https://github.com/andrewarrow/go365/blob/main/133/main.go) |                |
| [134](https://github.com/andrewarrow/go365/blob/main/134/main.go) |                |
| [135](https://github.com/andrewarrow/go365/blob/main/135/main.go) |                |
| [136](https://github.com/andrewarrow/go365/blob/main/136/main.go) |                |
| [137](https://github.com/andrewarrow/go365/blob/main/137/main.go) |                |
| [138](https://github.com/andrewarrow/go365/blob/main/138/main.go) |                |
| [139](https://github.com/andrewarrow/go365/blob/main/139/main.go) |                |
| [140](https://github.com/andrewarrow/go365/blob/main/140/main.go) |                |
| [141](https://github.com/andrewarrow/go365/blob/main/141/main.go) |                |
| [142](https://github.com/andrewarrow/go365/blob/main/142/main.go) |                |
| [143](https://github.com/andrewarrow/go365/blob/main/143/main.go) |                |
| [144](https://github.com/andrewarrow/go365/blob/main/144/main.go) |                |
| [145](https://github.com/andrewarrow/go365/blob/main/145/main.go) |                |
| [146](https://github.com/andrewarrow/go365/blob/main/146/main.go) |                |
| [147](https://github.com/andrewarrow/go365/blob/main/147/main.go) |                |
| [148](https://github.com/andrewarrow/go365/blob/main/148/main.go) |                |
| [149](https://github.com/andrewarrow/go365/blob/main/149/main.go) |                |
| [150](https://github.com/andrewarrow/go365/blob/main/150/main.go) |                |
| [151](https://github.com/andrewarrow/go365/blob/main/151/main.go) |                |
| [152](https://github.com/andrewarrow/go365/blob/main/152/main.go) |                |
| [153](https://github.com/andrewarrow/go365/blob/main/153/main.go) |                |
| [154](https://github.com/andrewarrow/go365/blob/main/154/main.go) |                |
| [155](https://github.com/andrewarrow/go365/blob/main/155/main.go) |                |
| [156](https://github.com/andrewarrow/go365/blob/main/156/main.go) |                |
| [157](https://github.com/andrewarrow/go365/blob/main/157/main.go) |                |
| [158](https://github.com/andrewarrow/go365/blob/main/158/main.go) |                |
| [159](https://github.com/andrewarrow/go365/blob/main/159/main.go) |                |
| [160](https://github.com/andrewarrow/go365/blob/main/160/main.go) |                |
| [161](https://github.com/andrewarrow/go365/blob/main/161/main.go) |                |
| [162](https://github.com/andrewarrow/go365/blob/main/162/main.go) |                |
| [163](https://github.com/andrewarrow/go365/blob/main/163/main.go) |                |
| [164](https://github.com/andrewarrow/go365/blob/main/164/main.go) |                |
| [165](https://github.com/andrewarrow/go365/blob/main/165/main.go) |                |
| [166](https://github.com/andrewarrow/go365/blob/main/166/main.go) |                |
| [167](https://github.com/andrewarrow/go365/blob/main/167/main.go) |                |
| [168](https://github.com/andrewarrow/go365/blob/main/168/main.go) |                |
| [169](https://github.com/andrewarrow/go365/blob/main/169/main.go) |                |
| [170](https://github.com/andrewarrow/go365/blob/main/170/main.go) |                |
| [171](https://github.com/andrewarrow/go365/blob/main/171/main.go) |                |
| [172](https://github.com/andrewarrow/go365/blob/main/172/main.go) |                |
| [173](https://github.com/andrewarrow/go365/blob/main/173/main.go) |                |
| [174](https://github.com/andrewarrow/go365/blob/main/174/main.go) |                |
| [175](https://github.com/andrewarrow/go365/blob/main/175/main.go) |                |
| [176](https://github.com/andrewarrow/go365/blob/main/176/main.go) |                |
| [177](https://github.com/andrewarrow/go365/blob/main/177/main.go) |                |
| [178](https://github.com/andrewarrow/go365/blob/main/178/main.go) |                |
| [179](https://github.com/andrewarrow/go365/blob/main/179/main.go) |                |
| [180](https://github.com/andrewarrow/go365/blob/main/180/main.go) |                |
| [181](https://github.com/andrewarrow/go365/blob/main/181/main.go) |                |
| [182](https://github.com/andrewarrow/go365/blob/main/182/main.go) |                |
| [183](https://github.com/andrewarrow/go365/blob/main/183/main.go) |                |
| [184](https://github.com/andrewarrow/go365/blob/main/184/main.go) |                |
| [185](https://github.com/andrewarrow/go365/blob/main/185/main.go) |                |
| [186](https://github.com/andrewarrow/go365/blob/main/186/main.go) |                |
| [187](https://github.com/andrewarrow/go365/blob/main/187/main.go) |                |
| [188](https://github.com/andrewarrow/go365/blob/main/188/main.go) |                |
| [189](https://github.com/andrewarrow/go365/blob/main/189/main.go) |                |
| [190](https://github.com/andrewarrow/go365/blob/main/190/main.go) |                |
| [191](https://github.com/andrewarrow/go365/blob/main/191/main.go) |                |
| [192](https://github.com/andrewarrow/go365/blob/main/192/main.go) |                |
| [193](https://github.com/andrewarrow/go365/blob/main/193/main.go) |                |
| [194](https://github.com/andrewarrow/go365/blob/main/194/main.go) |                |
| [195](https://github.com/andrewarrow/go365/blob/main/195/main.go) |                |
| [196](https://github.com/andrewarrow/go365/blob/main/196/main.go) |                |
| [197](https://github.com/andrewarrow/go365/blob/main/197/main.go) |                |
| [198](https://github.com/andrewarrow/go365/blob/main/198/main.go) |                |
| [199](https://github.com/andrewarrow/go365/blob/main/199/main.go) |                |
| [200](https://github.com/andrewarrow/go365/blob/main/200/main.go) |                |
| [201](https://github.com/andrewarrow/go365/blob/main/201/main.go) |                |
| [202](https://github.com/andrewarrow/go365/blob/main/202/main.go) |                |
| [203](https://github.com/andrewarrow/go365/blob/main/203/main.go) |                |
| [204](https://github.com/andrewarrow/go365/blob/main/204/main.go) |                |
| [205](https://github.com/andrewarrow/go365/blob/main/205/main.go) |                |
| [206](https://github.com/andrewarrow/go365/blob/main/206/main.go) |                |
| [207](https://github.com/andrewarrow/go365/blob/main/207/main.go) |                |
| [208](https://github.com/andrewarrow/go365/blob/main/208/main.go) |                |
| [209](https://github.com/andrewarrow/go365/blob/main/209/main.go) |                |
| [210](https://github.com/andrewarrow/go365/blob/main/210/main.go) |                |
| [211](https://github.com/andrewarrow/go365/blob/main/211/main.go) |                |
| [212](https://github.com/andrewarrow/go365/blob/main/212/main.go) |                |
| [213](https://github.com/andrewarrow/go365/blob/main/213/main.go) |                |
| [214](https://github.com/andrewarrow/go365/blob/main/214/main.go) |                |
| [215](https://github.com/andrewarrow/go365/blob/main/215/main.go) |                |
| [216](https://github.com/andrewarrow/go365/blob/main/216/main.go) |                |
| [217](https://github.com/andrewarrow/go365/blob/main/217/main.go) |                |
| [218](https://github.com/andrewarrow/go365/blob/main/218/main.go) |                |
| [219](https://github.com/andrewarrow/go365/blob/main/219/main.go) |                |
| [220](https://github.com/andrewarrow/go365/blob/main/220/main.go) |                |
| [221](https://github.com/andrewarrow/go365/blob/main/221/main.go) |                |
| [222](https://github.com/andrewarrow/go365/blob/main/222/main.go) |                |
| [223](https://github.com/andrewarrow/go365/blob/main/223/main.go) |                |
| [224](https://github.com/andrewarrow/go365/blob/main/224/main.go) |                |
| [225](https://github.com/andrewarrow/go365/blob/main/225/main.go) |                |
| [226](https://github.com/andrewarrow/go365/blob/main/226/main.go) |                |
| [227](https://github.com/andrewarrow/go365/blob/main/227/main.go) |                |
| [228](https://github.com/andrewarrow/go365/blob/main/228/main.go) |                |
| [229](https://github.com/andrewarrow/go365/blob/main/229/main.go) |                |
| [230](https://github.com/andrewarrow/go365/blob/main/230/main.go) |                |
| [231](https://github.com/andrewarrow/go365/blob/main/231/main.go) |                |
| [232](https://github.com/andrewarrow/go365/blob/main/232/main.go) |                |
| [233](https://github.com/andrewarrow/go365/blob/main/233/main.go) |                |
| [234](https://github.com/andrewarrow/go365/blob/main/234/main.go) |                |
| [235](https://github.com/andrewarrow/go365/blob/main/235/main.go) |                |
| [236](https://github.com/andrewarrow/go365/blob/main/236/main.go) |                |
| [237](https://github.com/andrewarrow/go365/blob/main/237/main.go) |                |
| [238](https://github.com/andrewarrow/go365/blob/main/238/main.go) |                |
| [239](https://github.com/andrewarrow/go365/blob/main/239/main.go) |                |
| [240](https://github.com/andrewarrow/go365/blob/main/240/main.go) |                |
| [241](https://github.com/andrewarrow/go365/blob/main/241/main.go) |                |
| [242](https://github.com/andrewarrow/go365/blob/main/242/main.go) |                |
| [243](https://github.com/andrewarrow/go365/blob/main/243/main.go) |                |
| [244](https://github.com/andrewarrow/go365/blob/main/244/main.go) |                |
| [245](https://github.com/andrewarrow/go365/blob/main/245/main.go) |                |
| [246](https://github.com/andrewarrow/go365/blob/main/246/main.go) |                |
| [247](https://github.com/andrewarrow/go365/blob/main/247/main.go) |                |
| [248](https://github.com/andrewarrow/go365/blob/main/248/main.go) |                |
| [249](https://github.com/andrewarrow/go365/blob/main/249/main.go) |                |
| [250](https://github.com/andrewarrow/go365/blob/main/250/main.go) |                |
| [251](https://github.com/andrewarrow/go365/blob/main/251/main.go) |                |
| [252](https://github.com/andrewarrow/go365/blob/main/252/main.go) |                |
| [253](https://github.com/andrewarrow/go365/blob/main/253/main.go) |                |
| [254](https://github.com/andrewarrow/go365/blob/main/254/main.go) |                |
| [255](https://github.com/andrewarrow/go365/blob/main/255/main.go) |                |
| [256](https://github.com/andrewarrow/go365/blob/main/256/main.go) |                |
| [257](https://github.com/andrewarrow/go365/blob/main/257/main.go) |                |
| [258](https://github.com/andrewarrow/go365/blob/main/258/main.go) |                |
| [259](https://github.com/andrewarrow/go365/blob/main/259/main.go) |                |
| [260](https://github.com/andrewarrow/go365/blob/main/260/main.go) |                |
| [261](https://github.com/andrewarrow/go365/blob/main/261/main.go) |                |
| [262](https://github.com/andrewarrow/go365/blob/main/262/main.go) |                |
| [263](https://github.com/andrewarrow/go365/blob/main/263/main.go) |                |
| [264](https://github.com/andrewarrow/go365/blob/main/264/main.go) |                |
| [265](https://github.com/andrewarrow/go365/blob/main/265/main.go) |                |
| [266](https://github.com/andrewarrow/go365/blob/main/266/main.go) |                |
| [267](https://github.com/andrewarrow/go365/blob/main/267/main.go) |                |
| [268](https://github.com/andrewarrow/go365/blob/main/268/main.go) |                |
| [269](https://github.com/andrewarrow/go365/blob/main/269/main.go) |                |
| [270](https://github.com/andrewarrow/go365/blob/main/270/main.go) |                |
| [271](https://github.com/andrewarrow/go365/blob/main/271/main.go) |                |
| [272](https://github.com/andrewarrow/go365/blob/main/272/main.go) |                |
| [273](https://github.com/andrewarrow/go365/blob/main/273/main.go) |                |
| [274](https://github.com/andrewarrow/go365/blob/main/274/main.go) |                |
| [275](https://github.com/andrewarrow/go365/blob/main/275/main.go) |                |
| [276](https://github.com/andrewarrow/go365/blob/main/276/main.go) |                |
| [277](https://github.com/andrewarrow/go365/blob/main/277/main.go) |                |
| [278](https://github.com/andrewarrow/go365/blob/main/278/main.go) |                |
| [279](https://github.com/andrewarrow/go365/blob/main/279/main.go) |                |
| [280](https://github.com/andrewarrow/go365/blob/main/280/main.go) |                |
| [281](https://github.com/andrewarrow/go365/blob/main/281/main.go) |                |
| [282](https://github.com/andrewarrow/go365/blob/main/282/main.go) |                |
| [283](https://github.com/andrewarrow/go365/blob/main/283/main.go) |                |
| [284](https://github.com/andrewarrow/go365/blob/main/284/main.go) |                |
| [285](https://github.com/andrewarrow/go365/blob/main/285/main.go) |                |
| [286](https://github.com/andrewarrow/go365/blob/main/286/main.go) |                |
| [287](https://github.com/andrewarrow/go365/blob/main/287/main.go) |                |
| [288](https://github.com/andrewarrow/go365/blob/main/288/main.go) |                |
| [289](https://github.com/andrewarrow/go365/blob/main/289/main.go) |                |
| [290](https://github.com/andrewarrow/go365/blob/main/290/main.go) |                |
| [291](https://github.com/andrewarrow/go365/blob/main/291/main.go) |                |
| [292](https://github.com/andrewarrow/go365/blob/main/292/main.go) |                |
| [293](https://github.com/andrewarrow/go365/blob/main/293/main.go) |                |
| [294](https://github.com/andrewarrow/go365/blob/main/294/main.go) |                |
| [295](https://github.com/andrewarrow/go365/blob/main/295/main.go) |                |
| [296](https://github.com/andrewarrow/go365/blob/main/296/main.go) |                |
| [297](https://github.com/andrewarrow/go365/blob/main/297/main.go) |                |
| [298](https://github.com/andrewarrow/go365/blob/main/298/main.go) |                |
| [299](https://github.com/andrewarrow/go365/blob/main/299/main.go) |                |
| [300](https://github.com/andrewarrow/go365/blob/main/300/main.go) |                |
| [301](https://github.com/andrewarrow/go365/blob/main/301/main.go) |                |
| [302](https://github.com/andrewarrow/go365/blob/main/302/main.go) |                |
| [303](https://github.com/andrewarrow/go365/blob/main/303/main.go) |                |
| [304](https://github.com/andrewarrow/go365/blob/main/304/main.go) |                |
| [305](https://github.com/andrewarrow/go365/blob/main/305/main.go) |                |
| [306](https://github.com/andrewarrow/go365/blob/main/306/main.go) |                |
| [307](https://github.com/andrewarrow/go365/blob/main/307/main.go) |                |
| [308](https://github.com/andrewarrow/go365/blob/main/308/main.go) |                |
| [309](https://github.com/andrewarrow/go365/blob/main/309/main.go) |                |
| [310](https://github.com/andrewarrow/go365/blob/main/310/main.go) |                |
| [311](https://github.com/andrewarrow/go365/blob/main/311/main.go) |                |
| [312](https://github.com/andrewarrow/go365/blob/main/312/main.go) |                |
| [313](https://github.com/andrewarrow/go365/blob/main/313/main.go) |                |
| [314](https://github.com/andrewarrow/go365/blob/main/314/main.go) |                |
| [315](https://github.com/andrewarrow/go365/blob/main/315/main.go) |                |
| [316](https://github.com/andrewarrow/go365/blob/main/316/main.go) |                |
| [317](https://github.com/andrewarrow/go365/blob/main/317/main.go) |                |
| [318](https://github.com/andrewarrow/go365/blob/main/318/main.go) |                |
| [319](https://github.com/andrewarrow/go365/blob/main/319/main.go) |                |
| [320](https://github.com/andrewarrow/go365/blob/main/320/main.go) |                |
| [321](https://github.com/andrewarrow/go365/blob/main/321/main.go) |                |
| [322](https://github.com/andrewarrow/go365/blob/main/322/main.go) |                |
| [323](https://github.com/andrewarrow/go365/blob/main/323/main.go) |                |
| [324](https://github.com/andrewarrow/go365/blob/main/324/main.go) |                |
| [325](https://github.com/andrewarrow/go365/blob/main/325/main.go) |                |
| [326](https://github.com/andrewarrow/go365/blob/main/326/main.go) |                |
| [327](https://github.com/andrewarrow/go365/blob/main/327/main.go) |                |
| [328](https://github.com/andrewarrow/go365/blob/main/328/main.go) |                |
| [329](https://github.com/andrewarrow/go365/blob/main/329/main.go) |                |
| [330](https://github.com/andrewarrow/go365/blob/main/330/main.go) |                |
| [331](https://github.com/andrewarrow/go365/blob/main/331/main.go) |                |
| [332](https://github.com/andrewarrow/go365/blob/main/332/main.go) |                |
| [333](https://github.com/andrewarrow/go365/blob/main/333/main.go) |                |
| [334](https://github.com/andrewarrow/go365/blob/main/334/main.go) |                |
| [335](https://github.com/andrewarrow/go365/blob/main/335/main.go) |                |
| [336](https://github.com/andrewarrow/go365/blob/main/336/main.go) |                |
| [337](https://github.com/andrewarrow/go365/blob/main/337/main.go) |                |
| [338](https://github.com/andrewarrow/go365/blob/main/338/main.go) |                |
| [339](https://github.com/andrewarrow/go365/blob/main/339/main.go) |                |
| [340](https://github.com/andrewarrow/go365/blob/main/340/main.go) |                |
| [341](https://github.com/andrewarrow/go365/blob/main/341/main.go) |                |
| [342](https://github.com/andrewarrow/go365/blob/main/342/main.go) |                |
| [343](https://github.com/andrewarrow/go365/blob/main/343/main.go) |                |
| [344](https://github.com/andrewarrow/go365/blob/main/344/main.go) |                |
| [345](https://github.com/andrewarrow/go365/blob/main/345/main.go) |                |
| [346](https://github.com/andrewarrow/go365/blob/main/346/main.go) |                |
| [347](https://github.com/andrewarrow/go365/blob/main/347/main.go) |                |
| [348](https://github.com/andrewarrow/go365/blob/main/348/main.go) |                |
| [349](https://github.com/andrewarrow/go365/blob/main/349/main.go) |                |
| [350](https://github.com/andrewarrow/go365/blob/main/350/main.go) |                |
| [351](https://github.com/andrewarrow/go365/blob/main/351/main.go) |                |
| [352](https://github.com/andrewarrow/go365/blob/main/352/main.go) |                |
| [353](https://github.com/andrewarrow/go365/blob/main/353/main.go) |                |
| [354](https://github.com/andrewarrow/go365/blob/main/354/main.go) |                |
| [355](https://github.com/andrewarrow/go365/blob/main/355/main.go) |                |
| [356](https://github.com/andrewarrow/go365/blob/main/356/main.go) |                |
| [357](https://github.com/andrewarrow/go365/blob/main/357/main.go) |                |
| [358](https://github.com/andrewarrow/go365/blob/main/358/main.go) |                |
| [359](https://github.com/andrewarrow/go365/blob/main/359/main.go) |                |
| [360](https://github.com/andrewarrow/go365/blob/main/360/main.go) |                |
| [361](https://github.com/andrewarrow/go365/blob/main/361/main.go) |                |
| [362](https://github.com/andrewarrow/go365/blob/main/362/main.go) |                |
| [363](https://github.com/andrewarrow/go365/blob/main/363/main.go) |                |
| [364](https://github.com/andrewarrow/go365/blob/main/364/main.go) |                |
| [365](https://github.com/andrewarrow/go365/blob/main/365/main.go) |                |
