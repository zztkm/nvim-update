# nvim-update

Neovim update from AppImage.

https://github.com/neovim/neovim/wiki/Installing-Neovim

## Usage

```
nvim-update [path]
```

default tag is `stable`.

example: update to nightly
```bash
zztkm@DESKTOP-MA7I77G:~$ nvim --version
NVIM v0.9.1
Build type: Release
LuaJIT 2.1.0-beta3

   system vimrc file: "$VIM/sysinit.vim"
  fall-back for $VIM: "/__w/neovim/neovim/build/nvim.AppDir/usr/share/nvim"

Run :checkhealth for more info
zztkm@DESKTOP-MA7I77G:~$ nvim-update --tag nightly ~/nvim-install/
ファイルをダウンロードしています... /home/zztkm/nvim-install/nvim.appimage
...

zztkm@DESKTOP-MA7I77G:~$ nvim --version
NVIM v0.10.0-dev-436+ga8ee4c7a8
Build type: RelWithDebInfo
LuaJIT 2.1.0-beta3
Compilation: /usr/bin/gcc-10 -O2 -g -Og -g -Wall -Wextra -pedantic -Wno-unused-parameter -Wstrict-prototypes -std=gnu99 -Wshadow -Wconversion -Wvla -Wdouble-promotion -Wmissing-noreturn -Wmissing-format-attribute -Wmissing-prototypes -fno-common -Wno-unused-result -Wimplicit-fallthrough -fdiagnostics-color=always -fstack-protector-strong -DUNIT_TESTING -DINCLUDE_GENERATED_DECLARATIONS -D_GNU_SOURCE -DNVIM_TS_HAS_SET_MAX_START_DEPTH -I/__w/neovim/neovim/.deps/usr/include/luajit-2.1 -I/usr/include -I/__w/neovim/neovim/.deps/usr/include -I/__w/neovim/neovim/build/src/nvim/auto -I/__w/neovim/neovim/build/include -I/__w/neovim/neovim/build/cmake.config -I/__w/neovim/neovim/src -I/usr/include -I/__w/neovim/neovim/.deps/usr/include -I/__w/neovim/neovim/.deps/usr/include -I/__w/neovim/neovim/.deps/usr/include -I/__w/neovim/neovim/.deps/usr/include -I/__w/neovim/neovim/.deps/usr/include -I/__w/neovim/neovim/.deps/usr/include

   system vimrc file: "$VIM/sysinit.vim"
  fall-back for $VIM: "/__w/neovim/neovim/build/nvim.AppDir/usr/share/nvim"

Run :checkhealth for more info
```

## Install

```
go install github.com/zztkm/nvim-update@latest
```

## License

MIT


## Author

zztkm

