# goMusicTagFixer

It is a simple command-line utility for to manage tags in mp3 files.
Its main task is to edit tags in directories with one command. 
For example, you have a catalog with a couple hundred different songs in mp3 
format. When you fill this in a physical player or add it to the library of a 
software player, you risk getting not just one folder with the name, say, 
"Collection", but a couple of hundred artists and albums in the library.
This utility allows you to solve this problem with a single command.

*It is also worth noting that for ease of reading, the utility's log is colored in colors other than white.*

Basic commands:

**scandir**  -   recursively scans the directory from which the utility was 
launched (where the terminal's "focus" was currently located) for mp3 music 
files. The contents of directories are displayed as a tree.

```
Starting program at 22:26:38 05-10-2020
Please enter the command:
scandir
Starting the folder scan at 22:26:41 05-10-2020
Temp
├─ 02 - Театр одного актера.mp3
├─ 03 - Последняя искра.mp3
├─ 04 - Шкатулка.mp3
├─ 05 - В последний раз.mp3
├─ 1983 - Rebel Yell
│  ├─ 01 - Rebel Yell.mp3
│  ├─ 02 - Daytime Drama.mp3
│  ├─ 03 - Eyes Without A Face.mp3
│  ├─ 04 - Blue Highway.mp3
│  ├─ 05 - Flesh For Fantasy.mp3
│  ├─ 06 - Catch My Fall.mp3
│  ├─ 07 - Crank Call.mp3
│  ├─ 08 - (Do Not) Stand In The Shadows.mp3
│  ├─ 09 - The Dead Next Door.mp3
│  ├─ Covers
├─ 1987 - Hot Track
│  ├─ 01 - Love Hurts.mp3
│  ├─ 02 - Shanghai'd in Shanghai.mp3
│  ├─ 03 - Carry Out Feelings.mp3
│  ├─ 04 - Razamanaz.mp3
│  ├─ 05 - Hair of the Dog.mp3
│  ├─ 06 - I Want to (Do Everything for You).mp3
│  ├─ 07 - This Flight Tonight.mp3
│  ├─ 08 - Broken Down Angel.mp3
│  ├─ 09 - Born to Love.mp3
│  ├─ 10 - My White Bicycle.mp3
│  ├─ 11 - Go Down Fighting.mp3
│  ├─ 12 - Vancouver Shakedown.mp3
│  ├─ Covers
├─ Космические Рейнджеры - Longway.mp3
├─ Космические Рейнджеры - Пробуждение.mp3
├─ Космические Рейнджеры 1 - Бой.mp3
├─ Космические рейнджеры 1 - Космический ветер.mp3
─────────────────────────────────────────
Scanning complete at  22:31:30 05-10-2020
Founded  49  mp3 files in  5  folders.

```

**deepscan**  -   its behavior is similar to the previous command, but there is 
additional functionality - it shows three main tags under each of the tracks:
artist, album and title.

```
Please enter the command:
deepscan
Deep scan started at  22:29:25 05-10-2020
Temp
├─ 1983 - Rebel Yell
│  ├─ 01 - Rebel Yell.mp3
│     Artist:   Rebel Yell 
│     Album:    Rebel Yell 
│     Title:   Rebel Yell 
│  ├─ 02 - Daytime Drama.mp3
│     Artist:   Rebel Yell 
│     Album:    Rebel Yell 
│     Title:    
│  ├─ 03 - Eyes Without A Face.mp3
│     Artist:   Rebel Yell 
│     Album:    Rebel Yell 
│     Title:    
│  ├─ 04 - Blue Highway.mp3
│     Artist:   Rebel Yell 
│     Album:    Rebel Yell 
│     Title:    
│  ├─ 05 - Flesh For Fantasy.mp3
│     Artist:   Rebel Yell 
│     Album:    Rebel Yell 
│     Title:    
│  ├─ 06 - Catch My Fall.mp3
│     Artist:   Rebel Yell 
│     Album:    Rebel Yell 
│     Title:    
│  ├─ 07 - Crank Call.mp3
│     Artist:   Rebel Yell 
│     Album:    Rebel Yell 
│     Title:   Crank Call 
│  ├─ 08 - (Do Not) Stand In The Shadows.mp3
│     Artist:   Rebel Yell 
│     Album:    Rebel Yell 
│     Title:   (Do Not) Stand In The Shadows 
│  ├─ 09 - The Dead Next Door.mp3
│     Artist:   Rebel Yell 
│     Album:    Rebel Yell 
│     Title:    
│  ├─ Covers
─────────────────────────────────────────
Scanning complete at  22:31:49 05-10-2020
Founded  49  mp3 files in  5  folders.

```

**mp3settags** -   the main command that the utility was designed for. It sets 
the album tag value for all tracks equal to the name of the directory where 
this file is located.

```
Please enter the command:
mp3settags
Changing of mp3 tags started at  22:30:37 05-10-2020
Temp
 02 - Театр одного актера.mp3  setting  Temp  album and artist name
 03 - Последняя искра.mp3  setting  Temp  album and artist name
 04 - Шкатулка.mp3  setting  Temp  album and artist name
 05 - В последний раз.mp3  setting  Temp  album and artist name
 Rebel Yell
 01 - Rebel Yell.mp3  setting   Rebel Yell  album and artist name
 02 - Daytime Drama.mp3  setting   Rebel Yell  album and artist name
 03 - Eyes Without A Face.mp3  setting   Rebel Yell  album and artist name
 04 - Blue Highway.mp3  setting   Rebel Yell  album and artist name
 05 - Flesh For Fantasy.mp3  setting   Rebel Yell  album and artist name
 06 - Catch My Fall.mp3  setting   Rebel Yell  album and artist name
 07 - Crank Call.mp3  setting   Rebel Yell  album and artist name
 08 - (Do Not) Stand In The Shadows.mp3  setting   Rebel Yell  album and artist name
 09 - The Dead Next Door.mp3  setting   Rebel Yell  album and artist name
Covers
 Hot Track
 01 - Love Hurts.mp3  setting   Hot Track  album and artist name
 02 - Shanghai'd in Shanghai.mp3  setting   Hot Track  album and artist name
 03 - Carry Out Feelings.mp3  setting   Hot Track  album and artist name
 04 - Razamanaz.mp3  setting   Hot Track  album and artist name
 05 - Hair of the Dog.mp3  setting   Hot Track  album and artist name
 06 - I Want to (Do Everything for You).mp3  setting   Hot Track  album and artist name
 07 - This Flight Tonight.mp3  setting   Hot Track  album and artist name
 08 - Broken Down Angel.mp3  setting   Hot Track  album and artist name
 09 - Born to Love.mp3  setting   Hot Track  album and artist name
 10 - My White Bicycle.mp3  setting   Hot Track  album and artist name
 11 - Go Down Fighting.mp3  setting   Hot Track  album and artist name
 12 - Vancouver Shakedown.mp3  setting   Hot Track  album and artist name
Covers
 Космические Рейнджеры - Longway.mp3  setting  Temp  album and artist name
 Космические Рейнджеры - Пробуждение.mp3  setting  Temp  album and artist name
 Космические Рейнджеры 1 - Бой.mp3  setting  Temp  album and artist name
 Космические рейнджеры 1 - Космический ветер.mp3  setting  Temp  album and artist name
Разное
 01 - 2000 - Frozen Features - The solitude.mp3  setting  Разное  album and artist name
 01 - Мир стекла.mp3  setting  Разное  album and artist name
 02 - 2000 - Frozen Features - The cry of silence.mp3  setting  Разное  album and artist name
 02 - Театр одного актера.mp3  setting  Разное  album and artist name
 03 - Последняя искра.mp3  setting  Разное  album and artist name
 04 - Шкатулка.mp3  setting  Разное  album and artist name
 05 - В последний раз.mp3  setting  Разное  album and artist name
 05 - О ней.mp3  setting  Разное  album and artist name
 07 - Ангел.mp3  setting  Разное  album and artist name
 08 - Гредят новый век.mp3  setting  Разное  album and artist name
 Григорий Семенов - Музыка людей (Космические рейнджеры).mp3  setting  Разное  album and artist name
 Космические Рейнджеры - Longway.mp3  setting  Разное  album and artist name
 Космические Рейнджеры - Безграничный космос.mp3  setting  Разное  album and artist name
 Космические Рейнджеры - Гимн.mp3  setting  Разное  album and artist name
 Космические Рейнджеры - Пробуждение.mp3  setting  Разное  album and artist name
 Космические Рейнджеры 1 - Бой.mp3  setting  Разное  album and artist name
 Космические рейнджеры 1 - Космический ветер.mp3  setting  Разное  album and artist name
 Космические рейнджеры 1 - Полетная.mp3  setting  Разное  album and artist name
 Космические рейнджеры OST - Starflight(2).mp3  setting  Разное  album and artist name
 Спящий - Abyssphere.mp3  setting  Разное  album and artist name
─────────────────────────────────────────
Changing complete at  22:30:37 05-10-2020
Changed  49  mp3 files in  5  folders.

```

**help**      -   view this help page.

The library is used for working with tags - [mikkyang/id3-go](https://github.com/mikkyang/id3-go)

#Build program

A convenient script has been written to build the program - build.sh. It compiles the program and puts the binary in the folder ./builds. 
