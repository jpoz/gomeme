# gomeme

Take a gif, make a meme. Written in Go (Golang)

## Installation

```
go get github.com/jpoz/gomeme/cmd/gomeme
```

## Usage

```shell
Usage: ./gomeme [options] input.gif output.gif

  -b string
        Bottom text of the meme.
  -fs float
        Font size of the text (default 42)
  -m int
        Margin around the text (default 10)
  -ss int
        Stroke size around the text (default 4)
  -t string
        Top text of the meme.
  -v    Displays more information.
```


## Example

```shell
gomeme -t "I am meme" -b "What are you?" input.gif output.gif
```

![output](https://cloud.githubusercontent.com/assets/12866/20644884/876f1740-b3fc-11e6-9718-15d7d69791a4.gif)
