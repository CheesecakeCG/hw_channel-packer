# hare_ware's channnel packer
A tool for manipulating image channels

## Usage

-a --append
    Puts the second image as the last unfilled channel of the first image.
    `hwchanpack -a rgb.png a.png`
-c --channel
    Each image specifies the channel it should use. If not specified, the the channel will be black.
    `hwchanpack -c r.png b.png g.png a.png out.png`

-i --invert
    When used with -a, the alpha image will be inverted.
    `hwchanpack -a -i rgb.png a.png`

-h --help
    Prints this help message.
    `hwchanpack -h`