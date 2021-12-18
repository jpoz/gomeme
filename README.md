# gomeme

Take a gif/jpeg/png/webp, make a meme. Written in Go (Golang)

## Installation

```
go get -u github.com/koalalorenzo/gomeme/cmd/gomeme
```

## Usage

```shell
Usage: gomeme [options] input.gif output.gif

  -b string
        Bottom text of the meme.
  -f string
        TrueType font path. Default is Hack-Bold.ttf
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
gomeme -t "Whatup internet?" -b "How you doing?" input.gif output.gif
```

![output](https://cloud.githubusercontent.com/assets/12866/20652316/e2555930-b4ab-11e6-9148-84e6bf0fc9d9.gif)

### Notes

`impact.ttf` From [here](https://github.com/neversmoke/lego/blob/817dce62321007d956c8e5823f09cfe7fab8b9ca/Lego/SiteBundle/Resources/public/css/inpact.ttf)
