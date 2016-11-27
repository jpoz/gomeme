# gomeme

Take a gif, make a meme. Written in Go (Golang)

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

![output](https://cloud.githubusercontent.com/assets/12866/20644852/9dd55a7c-b3fb-11e6-8b88-4bcf0306afa7.gif)
