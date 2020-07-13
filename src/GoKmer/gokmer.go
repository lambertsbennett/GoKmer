package main

import (
	"github.com/lambertsbennett/gokmervec/src/KmerVec"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"sync"
	"time"
)

// gokmer -n 8 -file contig_file -k-mer-size 5 -o output file

func main(){
	var proc int
	flag.IntVar(&proc,"n",2,"Number of processors or threads to leverage.")

	var contigfile string
	flag.StringVar(&contigfile,"file","","Contig file in fasta format.")

	var out string
	flag.StringVar(&out,"o","./gokmerout.parquet","Output file.")

	var kmer int
	flag.IntVar(&kmer,"k-mer-size",0,"K-mer word size to use.")

	flag.Parse()

	if kmer == 0{
		log.Println("kmer size should be non-zero")
		os.Exit(10)
	}

	runtime.GOMAXPROCS(proc)

	ls := KmerVec.ReadFasta(contigfile)
	lsp := KmerVec.SequenceCollection{}

	fmt.Println("Processing sequences")
	start := time.Now()
	var wg sync.WaitGroup
	in := make(chan KmerVec.Sequence, len(ls))


	for i:=0; i< 100; i++ {
		wg.Add(1)
		go KmerVec.GetKmers(kmer,&wg,in,&lsp)

	}


	for _, s := range ls {
		in <- s
	}

	close(in)

	wg.Wait()

	lsp.ToParquet(out)

	t := time.Since(start)
	fmt.Printf("%v sequences analysed in %s \n",len(ls),t)

}
