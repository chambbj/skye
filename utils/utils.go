package utils

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

func InferReaderFromExt(s string) (readerType, inputOpt string) {
	ext := filepath.Ext(s)
	if strings.EqualFold(ext, ".las") || strings.EqualFold(ext, ".laz") {
		inputOpt = "--readers.las.filename=" + s
		readerType = "readers.las"
	} else if strings.EqualFold(ext, ".bpf") {
		inputOpt = "--readers.bpf.filename=" + s
		readerType = "readers.bpf"
	}
	return
}

func InferWriterFromExt(s string) (writerType, outputOpt string) {
	ext := filepath.Ext(s)
	if strings.EqualFold(ext, ".las") || strings.EqualFold(ext, ".laz") {
		outputOpt = "--writers.las.filename=" + s
		writerType = "writers.las"
	} else if strings.EqualFold(ext, ".bpf") {
		outputOpt = "--writers.bpf.filename=" + s
		writerType = "writers.bpf"
	}
	return
}

func RunPdal(args ...string) {
	// cmd2 := exec.Command("pdal", args...)
	// stdout, err := cmd2.StdoutPipe()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// stderr, err := cmd2.StderrPipe()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// scannerOut := bufio.NewScanner(stdout)
	// go func() {
	// 	for scannerOut.Scan() {
	// 		fmt.Println(scannerOut.Text())
	// 	}
	// }()
	// scannerErr := bufio.NewScanner(stderr)
	// go func() {
	// 	for scannerErr.Scan() {
	// 		fmt.Println(scannerErr.Text())
	// 	}
	// }()
	// if err := cmd2.Start(); err != nil {
	// 	log.Fatal(err)
	// }
	// if err := cmd2.Wait(); err != nil {
	// 	log.Fatal(err)
	// }

	out, err := exec.Command("pdal", args...).CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
}

func RunPdalPipeline(args ...string) {
	args = append([]string{"pipeline"}, args...)
	RunPdal(args...)
}

func RunPdalTranslate(args ...string) {
	args = append([]string{"translate"}, args...)
	RunPdal(args...)
}

func OpenData(output string) {
	cmd2 := exec.Command("cmd", "/C", "start", output)
	// can we use Run here?
	if err := cmd2.Start(); err != nil {
		// need to double check that this catches case of no app being found
		log.Fatal(err)
	}
	if err := cmd2.Wait(); err != nil {
		log.Fatal(err)
	}
}
