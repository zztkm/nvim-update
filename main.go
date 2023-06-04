package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

const version = "0.0.2"

const appimagefile = "nvim.appimage"

var revision = "HEAD"

func download(url string, savePath string) error {
	fmt.Printf("Downloding files... %s\n", savePath)

	err := os.MkdirAll(filepath.Dir(savePath), os.ModePerm)
	if err != nil {
		return err
	}

	// ファイルを作成して保存するためのオープンファイル
	file, err := os.Create(savePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// HTTP GETリクエストを送信
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// レスポンスのボディをファイルに書き込む
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	var showVersion bool
	var tag string

	flag.BoolVar(&showVersion, "version", false, "show version")
	flag.StringVar(&tag, "tag", "stable", "tag")
	flag.Parse()

	if showVersion {
		fmt.Printf("version: %s, revision: %s\n", version, revision)
		return
	}

	if flag.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <savePath>\n", os.Args[0])
		os.Exit(1)
	}

	base := flag.Arg(0)

	err := os.RemoveAll(base)
	if err != nil {
		log.Fatal(err)
	}

	url := "https://github.com/neovim/neovim/releases/download/" + tag + "/nvim.appimage"
	savePath := filepath.Join(base, appimagefile)
	err = download(url, savePath)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Chmod(savePath, 0755)
	if err != nil {
		log.Fatal(err)
	}

	absSavePath, err := filepath.Abs(savePath)
	if err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command(absSavePath, "--appimage-extract")
	cmd.Dir = base
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	path, err := filepath.Abs(filepath.Join(base, "squashfs-root/usr/bin"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nManually add the directory to your shell profile\n")
	fmt.Println(fmt.Sprintf("export PATH=\"%s:$PATH\"", path))
}
